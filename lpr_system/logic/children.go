package logic

import (
	"fmt"
	"lpr_system/dao/mysql"
	"lpr_system/models"
	"math"
	"strconv"
)

// GetALLChildren获取所有孩子信息
func GetALLChildren() (data []*models.ChildrenInfo, err error) {
	data, err = mysql.GetALLChildren()
	if err != nil {
		return nil, err
	}
	return
}

//
func UpdateChildInfo(p *models.ChildInfoTable) (err error) {
	// 1.判断孩子是否存在于表中,返回孩子次数
	count, err := mysql.CheckChildrenExist(p.ID)
	fmt.Println(count)

	if count == 0 {
		// 构造user实例
		child := &models.ChildrenInfo{
			ChildID:     p.ID,
			Age:         p.Age,
			Name:        p.Name,
			ManagerName: p.ManagerName,
			TestCount:   count + 1,
		}
		return mysql.InsertChild(child)
	}

	// 更新Child
	tn, _ := strconv.Atoi(p.TestN)
	if count != tn {
		return mysql.UpdateChild(tn, p.ID)
	}
	return
}

func GetSuggestionWithTime(cid int, ctime string) (data string, err error) {
	data, err = mysql.GetSuggestionWithTime(cid, ctime)
	if err != nil {
		return "", err
	}
	return
}

func PlotWithTime(cid int) (data []*models.PlotInfo, err error) {

	all_times, err := mysql.GetChildrenCountTest(cid)
	times, _ := strconv.Atoi(all_times)
	var age_score = make(map[int]int, 72)

	fmt.Println(all_times)

	// 先计算各表分数与总分
	allscore := [8]models.Score{}
	allscore[7].SurveyId = 8

	for i := 1; i <= times; i++ {
		age, _ := mysql.GetAgeByTime(cid, strconv.Itoa(i))
		allpn, _ := mysql.PlotWithTime(cid, strconv.Itoa(i))
		fmt.Println(allpn)
		allscore[7].PN = 0
		for j := 0; j <= 6; j++ {
			allscore[j].PN = allpn[j]
			allscore[j].SurveyId = j + 1
			allscore[7].PN += allpn[j]

		}
		data = append(data, &models.PlotInfo{TestN: i, Age: age, AllScore: allscore})
	}

	for i := 1; i <= times; i++ {
		for j := 0; j <= 7; j++ {

			base_age, _ := mysql.SelectBaseAge(j + 1)

			p_now := data[i-1].AllScore[j].PN
			fmt.Println("当前得分")
			fmt.Println(p_now)

			for m := 1; m <= 72; m++ {
				age_score[m] = base_age[m-1]
			}
			fmt.Println(age_score)
			fmt.Println(base_age)
			top, bot := 0, 0
			m1, m2 := 0, 0

			// 计算
		label:
			for m := 0; m < 72; m++ {
				if base_age[m] == p_now {
					for k, v := range age_score {
						if v == base_age[m] {
							data[i-1].AllScore[j].AgeEvaluate = k
							break label
						}
					}
				}
				fmt.Println(m, p_now)
				if base_age[m] > p_now {
					top = base_age[m]
					for {
						m--
						if base_age[m] != -1 {
							bot = base_age[m]
							for k, v := range age_score {
								if v == bot {
									m2 = k
								}
								if v == top {
									m1 = k
								}
							}
							fmt.Println(top, bot, m1, m2)
							p_nowf := float64(p_now)
							topf, botf, m1f, m2f := float64(top), float64(bot), float64(m1), float64(m2)
							x := ((p_nowf - botf) / (topf - botf)) * (m1f - m2f)

							fmt.Println(p_nowf)
							fmt.Println(x)
							data[i-1].AllScore[j].AgeEvaluate = int(math.Round(float64(x))) + m2
							break label
						}
					}
				}
			}
		}

	}

	if err != nil {
		return nil, err
	}
	return
}
