package models

type Login struct {
	WechatId string `json:"wechat_id" db:"wechat_id"`
}

type ChildrenInfo struct {
	ChildID     int    `json:"code" db:"child_id"`
	Age         int    `json:"age" db:"age"`
	TestCount   int    `json:"count_test" db:"count_test"`
	Name        string `json:"name" db:"name"`
	ManagerName string `json:"manager_name" db:"manager_name"`
}

type IDTimeInfo1 struct {
	ChildID int    `json:"child_id" db:"child_id"`
	Time    string `json:"test_n" db:"test_n"`
}
type IDTimeInfo struct {
	ChildID int `json:"child_id" db:"child_id"`
	// Time    string `json:"test_n" db:"test_n"`
}
