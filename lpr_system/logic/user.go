package logic

import (
	"lpr_system/dao/mysql"
)

// Login
func Login(wechat_id string) (int, error) {
	count, err := mysql.Login(wechat_id)
	if err != nil {
		return count, err
	}
	return count, err
}
