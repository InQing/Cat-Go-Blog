package service

import (
	"Go-Blog/dao"
	"Go-Blog/models"
	"Go-Blog/utils"
	"errors"
)

func Login(userName,passwd string) (*models.LoginRes,error) {
	// 密码加密
	passwd = utils.Md5Crypt(passwd,"black cat")
	user := dao.GetUser(userName,passwd);
	if user == nil {
		return nil,errors.New("账号密码不正确")
	}
	uid := user.Uid
	// 根据uid生成唯一token
	token,err := utils.Award(&uid)
	if err != nil {
		return nil,errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		Token: token,
		UserInfo: userInfo,
	}
	return lr,nil
}