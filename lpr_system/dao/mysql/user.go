package mysql

//  Login

func Login(wechat_id string) (count int, err error) {
	sqlStr := `select count(wechat_id) from login where wechat_id = ?`

	if err := db.Get(&count, sqlStr, wechat_id); err != nil {
		return count, err
	}
	return
}
