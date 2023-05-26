package mysql

import (
	"fmt"
	"lpr_system/models"
	"time"
)

// GetALLChildren 查询所有孩子信息
func GetALLChildren() (data []*models.ChildrenInfo, err error) {
	sqlStr := `select
	child_id, name, manager_name, count_test
	from children
	ORDER BY create_time
	DESC`

	data = make([]*models.ChildrenInfo, 0, 100)
	err = db.Select(&data, sqlStr)
	return
}

// GetChildrenCountTest 查询孩子测试次数
func GetChildrenCountTest(child_id int) (count_test string, err error) {
	sqlStr := `select
	count_test
	from children
	where child_id=?
	`
	err = db.Get(&count_test, sqlStr, child_id)
	fmt.Println(count_test)
	return
}

// GetAgeByTime 查询孩子某次次对应年龄
func GetAgeByTime(child_id int, test_n string) (age int, err error) {
	sqlStr := `select
	age
	from survey_gzj_info
	where child_id=? and test_n=?
	`
	err = db.Get(&age, sqlStr, child_id, test_n)

	fmt.Println(age)
	return
}

// CheckChildrenExist 孩子是否存在
func CheckChildrenExist(id int) (count int, err error) {
	sqlStr := `select count_test from children where child_id = ?`

	if err := db.Get(&count, sqlStr, id); err != nil {
		return 0, err
	}
	return
}

// Update count_test
func UpdateChild(n, cid int) (err error) {
	sqlStr := `update children set count_test = ? where child_id=?`

	if _, err := db.Exec(sqlStr, n, cid); err != nil {
		return err
	}
	return
}

// InsertChild 向数据库插入一条新的孩子记录
func InsertChild(child *models.ChildrenInfo) (err error) {

	now := time.Now()
	timestamp := now.Format("2006-01-02 15:04:05")
	sqlStr := `insert into children(child_id, name , manager_name, create_time,count_test) value(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, child.ChildID, child.Name, child.ManagerName, timestamp, child.TestCount)
	return
}

// GetSuggestionWithTime 查询某孩子某次的康复建议
func GetSuggestionWithTime(cid int, ctime string) (data string, err error) {
	sqlStr := `select
	suggestion
	from suggestion
	where child_id=? and test_n=?`

	err = db.Get(&data, sqlStr, cid, ctime)
	return
}

// PlotWithTime 绘制某孩子某次的得分曲线图
func PlotWithTime(cid int, ctime string) (data []int, err error) {
	data = make([]int, 0, 8)
	var pn [7]int
	sqlStr := `select  p_n
	from survey_gzj_info
	where child_id=? and test_n=?

	`
	err = db.Get(&pn[0], sqlStr, cid, ctime)
	data = append(data, pn[0])
	fmt.Println(data)

	sqlStr1 := `select  p_n
	from survey_cddz_info
	where child_id=? and test_n=?

	`
	err = db.Get(&pn[1], sqlStr1, cid, ctime)
	data = append(data, pn[1])
	fmt.Println(data)

	sqlStr2 := `select  p_n
	from survey_jxdz_info
	where child_id=? and test_n=?

	`
	err = db.Get(&pn[2], sqlStr2, cid, ctime)
	data = append(data, pn[2])
	fmt.Println(data)

	sqlStr3 := `select  p_n
	from survey_yyygt_info
	where child_id=? and test_n=?

	`
	err = db.Get(&pn[3], sqlStr3, cid, ctime)
	data = append(data, pn[3])
	fmt.Println(data)

	sqlStr4 := `select  p_n
	from survey_rzfz_info
	where child_id=? and test_n=?

	`
	err = db.Get(&pn[4], sqlStr4, cid, ctime)
	data = append(data, pn[4])
	fmt.Println(data)

	sqlStr5 := `select  p_n
	from survey_shjw_info
	where child_id=? and test_n=?

	`
	err = db.Get(&pn[5], sqlStr5, cid, ctime)
	data = append(data, pn[5])
	fmt.Println(data)

	sqlStr6 := `select  p_n
	from survey_shzl_info
	where child_id=? and test_n=?

	`
	err = db.Get(&pn[6], sqlStr6, cid, ctime)
	data = append(data, pn[6])
	fmt.Println(data)

	return
}
