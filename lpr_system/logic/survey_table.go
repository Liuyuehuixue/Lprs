package logic

import (
	"bytes"
	"fmt"
	"lpr_system/dao/mysql"
	"lpr_system/models"
	"reflect"
	"strconv"
	"strings"
)

// GetALLSelectedSurveyItems 筛选逻辑的获取所有符合条件项目
func GetALLSelectedSurveyItems(id int, age int, testN string) (data models.SelectedSevenTableItems, err error) {
	// 查询status  0:存在记录可以插入 1：存在记录不能插入 2：不存在

	status, _ := mysql.GetStatusInfo(id, testN)
	fmt.Println(status)

	data, err = mysql.GetALLSelectedSurveyItems(id, age, testN)
	data.Status = status

	switch status {
	case 0:
		// 查询info表记录

		for sid := 0; sid < 7; sid++ {
			var seven_survey models.SevenSurvey
			switch sid {
			case 0:
				seven_survey.SurveyGzj, _ = mysql.GetOneInfoInT1(id, testN)
			case 1:
				seven_survey.SurveyCddz, _ = mysql.GetOneInfoInT2(id, testN)
				fmt.Println(seven_survey.SurveyCddz)
			case 2:
				seven_survey.SurveyJxdz, _ = mysql.GetOneInfoInT3(id, testN)
			case 3:
				seven_survey.SurveyYyygt, _ = mysql.GetOneInfoInT4(id, testN)
			case 4:
				seven_survey.SurveyRzfz, _ = mysql.GetOneInfoInT5(id, testN)
			case 5:
				seven_survey.SurveyShjw, _ = mysql.GetOneInfoInT6(id, testN)
			case 6:
				seven_survey.SurveyShzl, _ = mysql.GetOneInfoInT7(id, testN)

			}

			items := []string{}
			for i := 0; i < len(data.AllItems[sid].Items); i++ {
				items = append(items, data.AllItems[sid].Items[i].SurveyProjectEnname)
			}

			for i := 0; i < len(items); i++ {
				index, _ := strconv.Atoi(strings.Split(items[i], "_")[0])
				// 更新具体数值
				switch sid {
				case 0:
					v := reflect.ValueOf(seven_survey.SurveyGzj)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 1:
					v := reflect.ValueOf(seven_survey.SurveyCddz)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 2:
					v := reflect.ValueOf(seven_survey.SurveyJxdz)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 3:
					v := reflect.ValueOf(seven_survey.SurveyYyygt)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 4:
					v := reflect.ValueOf(seven_survey.SurveyRzfz)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 5:
					v := reflect.ValueOf(seven_survey.SurveyShjw)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 6:
					v := reflect.ValueOf(seven_survey.SurveyShzl)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				}
			}

		}
		return
	case 1:
		return models.SelectedSevenTableItems{}, err
	default:
		return
	}

}

// 查看逻辑的GetALLSelectedSurveyItems2 获取所有符合条件项目
func GetALLSelectedSurveyItems2(id int, age int, testN string) (data models.SelectedSevenTableItems, err error) {
	// 查询status  0:存在记录可以插入 1：存在记录不能插入 2：不存在

	status, _ := mysql.GetStatusInfo(id, testN)
	fmt.Println(status)

	data, err = mysql.GetALLSelectedSurveyItems(id, age, testN)
	data.Status = status

	switch status {
	case 0, 1:
		// 查询info表记录

		for sid := 0; sid < 7; sid++ {
			var seven_survey models.SevenSurvey
			switch sid {
			case 0:
				seven_survey.SurveyGzj, _ = mysql.GetOneInfoInT1(id, testN)
			case 1:
				seven_survey.SurveyCddz, _ = mysql.GetOneInfoInT2(id, testN)
				fmt.Println(seven_survey.SurveyCddz)
			case 2:
				seven_survey.SurveyJxdz, _ = mysql.GetOneInfoInT3(id, testN)
			case 3:
				seven_survey.SurveyYyygt, _ = mysql.GetOneInfoInT4(id, testN)
			case 4:
				seven_survey.SurveyRzfz, _ = mysql.GetOneInfoInT5(id, testN)
			case 5:
				seven_survey.SurveyShjw, _ = mysql.GetOneInfoInT6(id, testN)
			case 6:
				seven_survey.SurveyShzl, _ = mysql.GetOneInfoInT7(id, testN)

			}

			items := []string{}
			for i := 0; i < len(data.AllItems[sid].Items); i++ {
				items = append(items, data.AllItems[sid].Items[i].SurveyProjectEnname)
			}

			for i := 0; i < len(items); i++ {
				index, _ := strconv.Atoi(strings.Split(items[i], "_")[0])
				// 更新具体数值
				switch sid {
				case 0:
					v := reflect.ValueOf(seven_survey.SurveyGzj)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 1:
					v := reflect.ValueOf(seven_survey.SurveyCddz)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 2:
					v := reflect.ValueOf(seven_survey.SurveyJxdz)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 3:
					v := reflect.ValueOf(seven_survey.SurveyYyygt)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 4:
					v := reflect.ValueOf(seven_survey.SurveyRzfz)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 5:
					v := reflect.ValueOf(seven_survey.SurveyShjw)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				case 6:
					v := reflect.ValueOf(seven_survey.SurveyShzl)
					data.AllItems[sid].Items[i].Score = v.Field(index + 4).Interface().(string)
				}
			}

		}
		return
	default:
		return
	}

}

// GetItemsAllTimeById 显示某孩子某调查表的每次测试得分
func GetItemsAllTimeById(cid int, sid int, test_n string) (data models.OneSurveyInfo, err error) {

	data.Surveyname = id_surveyname[sid]
	data.AllTimesItems.TestN = test_n
	times, _ := mysql.GetTimesByID(cid)
	fmt.Println(times)
	if len(times) == 0 {
		return
	}
	max_t, _ := strconv.Atoi(times[len(times)-1])
	t, _ := strconv.Atoi(test_n)
	if max_t < t {
		return
	}
	age, err := mysql.GetAgeByTime(cid, test_n)

	alltabelitems, err := mysql.GetALLSelectedSurveyItems(cid, age, test_n)

	switch sid {
	case 1:
		allInfo := models.SurveyGzj{}
		allInfo, _ = mysql.GetOneInfoInT1(cid, test_n)
		fmt.Println(allInfo)

		data.AllTimesItems.PN = allInfo.PN
		data.AllTimesItems.EN = allInfo.EN
		data.AllTimesItems.FN = allInfo.FN

		fmt.Println(alltabelitems.AllItems[0].Items)

		for i := 0; i < len(alltabelitems.AllItems[0].Items); i++ {
			index, _ := strconv.Atoi(strings.Split(alltabelitems.AllItems[0].Items[i].SurveyProjectEnname, "_")[0])
			fmt.Println(index)
			// 更新具体数值
			v := reflect.ValueOf(allInfo)
			data.AllTimesItems.ItemsWithTest = append(data.AllTimesItems.ItemsWithTest, models.ItemOneSurvey{SurveyProjectCnname: alltabelitems.AllItems[0].Items[i].SurveyProjectCnname, Score: v.Field(index + 4).Interface().(string)})
		}
	case 2:
		allInfo := models.SurveyCddz{}
		allInfo, _ = mysql.GetOneInfoInT2(cid, test_n)
		fmt.Println(allInfo)

		data.AllTimesItems.PN = allInfo.PN
		data.AllTimesItems.EN = allInfo.EN
		data.AllTimesItems.FN = allInfo.FN

		for i := 0; i < len(alltabelitems.AllItems[1].Items); i++ {
			index, _ := strconv.Atoi(strings.Split(alltabelitems.AllItems[0].Items[i].SurveyProjectEnname, "_")[0])
			// 更新具体数值
			v := reflect.ValueOf(allInfo)
			data.AllTimesItems.ItemsWithTest = append(data.AllTimesItems.ItemsWithTest, models.ItemOneSurvey{SurveyProjectCnname: alltabelitems.AllItems[1].Items[i].SurveyProjectCnname, Score: v.Field(index + 4).Interface().(string)})
		}
	case 3:
		allInfo := models.SurveyJxdz{}
		allInfo, _ = mysql.GetOneInfoInT3(cid, test_n)
		fmt.Println(allInfo)

		data.AllTimesItems.PN = allInfo.PN
		data.AllTimesItems.EN = allInfo.EN
		data.AllTimesItems.FN = allInfo.FN

		for i := 0; i < len(alltabelitems.AllItems[2].Items); i++ {
			index, _ := strconv.Atoi(strings.Split(alltabelitems.AllItems[0].Items[i].SurveyProjectEnname, "_")[0])
			// 更新具体数值
			v := reflect.ValueOf(allInfo)
			data.AllTimesItems.ItemsWithTest = append(data.AllTimesItems.ItemsWithTest, models.ItemOneSurvey{SurveyProjectCnname: alltabelitems.AllItems[2].Items[i].SurveyProjectCnname, Score: v.Field(index + 4).Interface().(string)})
		}
	case 4:
		allInfo := models.SurveyYyygt{}
		allInfo, _ = mysql.GetOneInfoInT4(cid, test_n)
		fmt.Println(allInfo)

		data.AllTimesItems.PN = allInfo.PN
		data.AllTimesItems.EN = allInfo.EN
		data.AllTimesItems.FN = allInfo.FN

		for i := 0; i < len(alltabelitems.AllItems[3].Items); i++ {
			index, _ := strconv.Atoi(strings.Split(alltabelitems.AllItems[0].Items[i].SurveyProjectEnname, "_")[0])
			// 更新具体数值
			v := reflect.ValueOf(allInfo)
			data.AllTimesItems.ItemsWithTest = append(data.AllTimesItems.ItemsWithTest, models.ItemOneSurvey{SurveyProjectCnname: alltabelitems.AllItems[3].Items[i].SurveyProjectCnname, Score: v.Field(index + 4).Interface().(string)})
		}
	case 5:
		allInfo := models.SurveyRzfz{}
		allInfo, _ = mysql.GetOneInfoInT5(cid, test_n)
		fmt.Println(allInfo)

		data.AllTimesItems.PN = allInfo.PN
		data.AllTimesItems.EN = allInfo.EN
		data.AllTimesItems.FN = allInfo.FN

		for i := 0; i < len(alltabelitems.AllItems[4].Items); i++ {
			index, _ := strconv.Atoi(strings.Split(alltabelitems.AllItems[0].Items[i].SurveyProjectEnname, "_")[0])
			// 更新具体数值
			v := reflect.ValueOf(allInfo)
			data.AllTimesItems.ItemsWithTest = append(data.AllTimesItems.ItemsWithTest, models.ItemOneSurvey{SurveyProjectCnname: alltabelitems.AllItems[4].Items[i].SurveyProjectCnname, Score: v.Field(index + 4).Interface().(string)})
		}
	case 6:
		allInfo := models.SurveyShjw{}
		allInfo, _ = mysql.GetOneInfoInT6(cid, test_n)
		fmt.Println(allInfo)

		data.AllTimesItems.PN = allInfo.PN
		data.AllTimesItems.EN = allInfo.EN
		data.AllTimesItems.FN = allInfo.FN

		for i := 0; i < len(alltabelitems.AllItems[5].Items); i++ {
			index, _ := strconv.Atoi(strings.Split(alltabelitems.AllItems[0].Items[i].SurveyProjectEnname, "_")[0])
			// 更新具体数值
			v := reflect.ValueOf(allInfo)
			data.AllTimesItems.ItemsWithTest = append(data.AllTimesItems.ItemsWithTest, models.ItemOneSurvey{SurveyProjectCnname: alltabelitems.AllItems[5].Items[i].SurveyProjectCnname, Score: v.Field(index + 4).Interface().(string)})
		}
	case 7:
		allInfo := models.SurveyShzl{}
		allInfo, _ = mysql.GetOneInfoInT7(cid, test_n)
		fmt.Println(allInfo)

		data.AllTimesItems.PN = allInfo.PN
		data.AllTimesItems.EN = allInfo.EN
		data.AllTimesItems.FN = allInfo.FN

		for i := 0; i < len(alltabelitems.AllItems[6].Items); i++ {
			index, _ := strconv.Atoi(strings.Split(alltabelitems.AllItems[0].Items[i].SurveyProjectEnname, "_")[0])
			// 更新具体数值
			v := reflect.ValueOf(allInfo)
			data.AllTimesItems.ItemsWithTest = append(data.AllTimesItems.ItemsWithTest, models.ItemOneSurvey{SurveyProjectCnname: alltabelitems.AllItems[6].Items[i].SurveyProjectCnname, Score: v.Field(index + 4).Interface().(string)})
		}
	}

	if err != nil {
		return models.OneSurveyInfo{}, err
	}
	return
}

// InsertItemsGenSug
func InsertItemsGenSug(p *models.SelectedSevenTableItems, stat int) (err error) {
	//解析结构体对象p
	child_id := p.ID
	age := p.Age
	test_n := p.TestN
	max_ID, _ := mysql.GetMaxIdInfo1()

	// 生成建议，插入到suggestion表
	var bt bytes.Buffer
	var seven_survey models.SevenSurvey

	// 初始化结构体
	for sid := 0; sid < 7; sid++ {

		switch sid {
		case 0:
			seven_survey.SurveyGzj.ID = max_ID + 1
			seven_survey.SurveyGzj.ChildID = child_id
			seven_survey.SurveyGzj.TestN = test_n
			seven_survey.SurveyGzj.Age = age
			seven_survey.SurveyGzj.Status = stat
			//seven_survey.SurveyGzj.PN, seven_survey.SurveyGzj.EN, seven_survey.SurveyGzj.FN, seven_survey.SurveyGzj.Score = 0, 0, 0, 0
			flag := 0
			for i := 0; i < len(p.AllItems[0].Items); i++ {
				//得分
				s := p.AllItems[0].Items[i].Score
				if s == "P" {
					seven_survey.SurveyGzj.PN++
					seven_survey.SurveyGzj.Score++
				} else if s == "E" {
					seven_survey.SurveyGzj.EN++
				} else if s == "F" {
					println(s, flag)
					seven_survey.SurveyGzj.FN++
					if flag == 0 {
						str, _ := mysql.SearchSuggest(age, 1)
						bt.WriteString(str)
					}
					flag = 1
				}
				fmt.Println(seven_survey.SurveyGzj)

				index := strings.Split(p.AllItems[0].Items[i].SurveyProjectEnname, "_")[0]
				fmt.Println("case 0 ")
				fmt.Println(index)
				// 更新具体数值
				PValue := reflect.ValueOf(&seven_survey.SurveyGzj).Elem()
				fieldName := "P" + index
				fieldValuePtr := reflect.ValueOf(&s).Elem()
				structFieldValue := PValue.FieldByName(fieldName)
				if structFieldValue.IsValid() {
					if structFieldValue.CanSet() {
						structFieldValue.Set(fieldValuePtr)
					}
				}
			}
		case 1:
			seven_survey.SurveyCddz.ID = max_ID + 1
			seven_survey.SurveyCddz.ChildID = child_id
			seven_survey.SurveyCddz.TestN = test_n
			seven_survey.SurveyCddz.Age = age
			seven_survey.SurveyCddz.Status = stat
			//seven_survey.SurveyGzj.PN, seven_survey.SurveyGzj.EN, seven_survey.SurveyGzj.FN, seven_survey.SurveyGzj.Score = 0, 0, 0, 0
			flag := 0
			for i := 0; i < len(p.AllItems[1].Items); i++ {
				//得分
				s := p.AllItems[1].Items[i].Score
				if s == "P" {
					seven_survey.SurveyCddz.PN++
					seven_survey.SurveyCddz.Score++
				} else if s == "E" {
					seven_survey.SurveyCddz.EN++
				} else if s == "F" {
					println(s, flag)
					seven_survey.SurveyCddz.FN++
					if flag == 0 {
						str, _ := mysql.SearchSuggest(age, 2)
						bt.WriteString(str)
					}
					flag = 1
				}
				index := strings.Split(p.AllItems[1].Items[i].SurveyProjectEnname, "_")[0]
				fmt.Println("case 1 ")
				fmt.Println(index)
				// 更新具体数值
				PValue := reflect.ValueOf(&seven_survey.SurveyCddz).Elem()
				fieldName := "P" + index
				fieldValuePtr := reflect.ValueOf(&s).Elem()
				structFieldValue := PValue.FieldByName(fieldName)
				if structFieldValue.IsValid() {
					if structFieldValue.CanSet() {
						structFieldValue.Set(fieldValuePtr)
					}
				}
			}
		case 2:
			seven_survey.SurveyJxdz.ID = max_ID + 1
			seven_survey.SurveyJxdz.ChildID = child_id
			seven_survey.SurveyJxdz.TestN = test_n
			seven_survey.SurveyJxdz.Age = age
			seven_survey.SurveyJxdz.Status = stat
			//seven_survey.SurveyGzj.PN, seven_survey.SurveyGzj.EN, seven_survey.SurveyGzj.FN, seven_survey.SurveyGzj.Score = 0, 0, 0, 0
			flag := 0
			for i := 0; i < len(p.AllItems[2].Items); i++ {
				//得分
				s := p.AllItems[2].Items[i].Score
				if s == "P" {
					seven_survey.SurveyJxdz.PN++
					seven_survey.SurveyJxdz.Score++
				} else if s == "E" {
					seven_survey.SurveyJxdz.EN++
				} else if s == "F" {
					println(s, flag)
					seven_survey.SurveyJxdz.FN++
					if flag == 0 {
						str, _ := mysql.SearchSuggest(age, 3)
						bt.WriteString(str)
					}
					flag = 1
				}
				index := strings.Split(p.AllItems[2].Items[i].SurveyProjectEnname, "_")[0]
				fmt.Println("case 2 ")
				fmt.Println(index)
				// 更新具体数值
				PValue := reflect.ValueOf(&seven_survey.SurveyJxdz).Elem()
				fieldName := "P" + index
				fieldValuePtr := reflect.ValueOf(&s).Elem()
				structFieldValue := PValue.FieldByName(fieldName)
				if structFieldValue.IsValid() {
					if structFieldValue.CanSet() {
						structFieldValue.Set(fieldValuePtr)
					}
				}
			}
		case 3:
			seven_survey.SurveyYyygt.ID = max_ID + 1
			seven_survey.SurveyYyygt.ChildID = child_id
			seven_survey.SurveyYyygt.TestN = test_n
			seven_survey.SurveyYyygt.Age = age
			seven_survey.SurveyYyygt.Status = stat
			//seven_survey.SurveyGzj.PN, seven_survey.SurveyGzj.EN, seven_survey.SurveyGzj.FN, seven_survey.SurveyGzj.Score = 0, 0, 0, 0
			flag := 0
			for i := 0; i < len(p.AllItems[3].Items); i++ {
				//得分
				s := p.AllItems[3].Items[i].Score
				if s == "P" {
					seven_survey.SurveyYyygt.PN++
					seven_survey.SurveyYyygt.Score++
				} else if s == "E" {
					seven_survey.SurveyYyygt.EN++
				} else if s == "F" {
					println(s, flag)
					seven_survey.SurveyYyygt.FN++
					if flag == 0 {
						str, _ := mysql.SearchSuggest(age, 4)
						bt.WriteString(str)
					}
					flag = 1
				}
				index := strings.Split(p.AllItems[3].Items[i].SurveyProjectEnname, "_")[0]
				// 更新具体数值
				fmt.Println("case 3 ")
				fmt.Println(index)
				PValue := reflect.ValueOf(&seven_survey.SurveyYyygt).Elem()
				fieldName := "P" + index
				fieldValuePtr := reflect.ValueOf(&s).Elem()
				structFieldValue := PValue.FieldByName(fieldName)
				if structFieldValue.IsValid() {
					if structFieldValue.CanSet() {
						structFieldValue.Set(fieldValuePtr)
					}
				}
			}
		case 4:
			seven_survey.SurveyRzfz.ID = max_ID + 1
			seven_survey.SurveyRzfz.ChildID = child_id
			seven_survey.SurveyRzfz.TestN = test_n
			seven_survey.SurveyRzfz.Age = age
			seven_survey.SurveyRzfz.Status = stat
			//seven_survey.SurveyGzj.PN, seven_survey.SurveyGzj.EN, seven_survey.SurveyGzj.FN, seven_survey.SurveyGzj.Score = 0, 0, 0, 0
			flag := 0
			for i := 0; i < len(p.AllItems[4].Items); i++ {
				//得分
				s := p.AllItems[4].Items[i].Score
				if s == "P" {
					seven_survey.SurveyRzfz.PN++
					seven_survey.SurveyRzfz.Score++
				} else if s == "E" {
					seven_survey.SurveyRzfz.EN++
				} else if s == "F" {
					println(s, flag)
					seven_survey.SurveyRzfz.FN++
					if flag == 0 {
						str, _ := mysql.SearchSuggest(age, 5)
						bt.WriteString(str)
					}
					flag = 1
				}
				index := strings.Split(p.AllItems[4].Items[i].SurveyProjectEnname, "_")[0]
				fmt.Println("case 4 ")
				fmt.Println(index)
				// 更新具体数值
				PValue := reflect.ValueOf(&seven_survey.SurveyRzfz).Elem()
				fieldName := "P" + index
				fieldValuePtr := reflect.ValueOf(&s).Elem()
				structFieldValue := PValue.FieldByName(fieldName)
				if structFieldValue.IsValid() {
					if structFieldValue.CanSet() {
						structFieldValue.Set(fieldValuePtr)
					}
				}
			}
		case 5:
			seven_survey.SurveyShjw.ID = max_ID + 1
			seven_survey.SurveyShjw.ChildID = child_id
			seven_survey.SurveyShjw.TestN = test_n
			seven_survey.SurveyShjw.Age = age
			seven_survey.SurveyShjw.Status = stat
			//seven_survey.SurveyGzj.PN, seven_survey.SurveyGzj.EN, seven_survey.SurveyGzj.FN, seven_survey.SurveyGzj.Score = 0, 0, 0, 0
			flag := 0
			for i := 0; i < len(p.AllItems[5].Items); i++ {
				//得分
				s := p.AllItems[5].Items[i].Score
				if s == "P" {
					seven_survey.SurveyShjw.PN++
					seven_survey.SurveyShjw.Score++
				} else if s == "E" {
					seven_survey.SurveyShjw.EN++
				} else if s == "F" {
					println(s, flag)
					seven_survey.SurveyShjw.FN++
					if flag == 0 {
						str, _ := mysql.SearchSuggest(age, 6)
						bt.WriteString(str)
					}
					flag = 1
				}
				index := strings.Split(p.AllItems[5].Items[i].SurveyProjectEnname, "_")[0]
				fmt.Println("case 5 ")
				fmt.Println(index)
				// 更新具体数值
				PValue := reflect.ValueOf(&seven_survey.SurveyShjw).Elem()
				fieldName := "P" + index
				fieldValuePtr := reflect.ValueOf(&s).Elem()
				structFieldValue := PValue.FieldByName(fieldName)
				if structFieldValue.IsValid() {
					if structFieldValue.CanSet() {
						structFieldValue.Set(fieldValuePtr)
					}
				}
			}
		case 6:
			seven_survey.SurveyShzl.ID = max_ID + 1
			seven_survey.SurveyShzl.ChildID = child_id
			seven_survey.SurveyShzl.TestN = test_n
			seven_survey.SurveyShzl.Age = age
			seven_survey.SurveyShzl.Status = stat
			//seven_survey.SurveyGzj.PN, seven_survey.SurveyGzj.EN, seven_survey.SurveyGzj.FN, seven_survey.SurveyGzj.Score = 0, 0, 0, 0
			flag := 0
			for i := 0; i < len(p.AllItems[6].Items); i++ {
				//得分
				s := p.AllItems[6].Items[i].Score
				if s == "P" {
					seven_survey.SurveyShzl.PN++
					seven_survey.SurveyShzl.Score++
				} else if s == "E" {
					seven_survey.SurveyShzl.EN++
				} else if s == "F" {
					println(s, flag)
					seven_survey.SurveyShzl.FN++
					if flag == 0 {
						str, _ := mysql.SearchSuggest(age, 7)
						bt.WriteString(str)
					}
					flag = 1
				}
				index := strings.Split(p.AllItems[6].Items[i].SurveyProjectEnname, "_")[0]
				fmt.Println("case 6 ")
				fmt.Println(index)
				// 更新具体数值
				PValue := reflect.ValueOf(&seven_survey.SurveyShzl).Elem()
				fieldName := "P" + index
				fieldValuePtr := reflect.ValueOf(&s).Elem()
				structFieldValue := PValue.FieldByName(fieldName)
				if structFieldValue.IsValid() {
					if structFieldValue.CanSet() {
						structFieldValue.Set(fieldValuePtr)
					}
				}
			}
		}
	}

	suggestion := bt.String()
	fmt.Println(suggestion)

	id, _ := mysql.SelectSugID(child_id, test_n)
	fmt.Println("id", id)
	// suggest 不存在 插入
	if id == "-1" {
		// 往suggestion插入数据
		mysql.InsertSuggest(child_id, test_n, suggestion)
	} else {
		mysql.UpdateSuggest(child_id, test_n, suggestion)
	}

	// 根据结构体对象插入到info表
	gzj_exit, _ := mysql.GetOneInfoInT1(child_id, test_n)
	if reflect.DeepEqual(gzj_exit, reflect.Zero(reflect.TypeOf(gzj_exit)).Interface()) {
		mysql.InsertSurveyInfo(seven_survey)
	} else {
		mysql.DeleteSurveyInfo(child_id, test_n)
		mysql.InsertSurveyInfo(seven_survey)
	}

	if err != nil {
		return err
	}
	return
}
