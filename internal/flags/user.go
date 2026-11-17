package flags

import (
	"ServeGin/internal/global"
	"ServeGin/internal/models"
	"ServeGin/internal/utils/pwd"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/term"
	"gorm.io/gorm"
)

type User struct {
}

/*
命令行创建用户
*/
func (u *User) Create() {
	var user models.UserModel
	fmt.Println("请选择角色： 1 管理员 2 普通用户")
	_, err := fmt.Scanln(&user.RoleID)
	if err != nil {
		fmt.Println("输入错误", err)
		return
	}
	if user.RoleID != 1 && user.RoleID != 2 {
		fmt.Println("用户角色输入错误", err)
		return
	}

	// 无限循环，直到用户名合理
	for {
		fmt.Println("请输入用户名")
		_, err := fmt.Scanln(&user.UserName)
		// 1，处理直接回车(空输入)的情况
		if err != nil || user.UserName == "" {
			fmt.Println("用户名不能为空，请重新输入")
			continue
		}

		// 校验用户名是否存在
		var userModel models.UserModel
		err = global.DB.Take(&userModel, "user_name = ?", user.UserName).Error
		if err == nil {
			fmt.Println("用户名已存在")
			continue
		}
		// 数据库查询过程出现错误
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("查询数据库发生错误: ", err)
			return
		}
		break
	}

	fmt.Println("请输入密码")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("读取密码出错:", err)
		return
	}

	fmt.Println("请再次输入密码")
	rePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("读取密码出错:", err)
		return
	}
	if string(password) != string(rePassword) {
		fmt.Println("两次密码不一致")
		return
	}

	hashPassword, err := pwd.GenerateHashPassword(string(password))
	if err != nil {
		fmt.Println("密码加密失败", err)
	}

	// 根据用户名随机生成一个昵称
	nickName := fmt.Sprintf("%s_%d", user.UserName, rand.IntN(9000)+1000)
	err = global.DB.Create(&models.UserModel{
		UserName: user.UserName,
		RoleID:   user.RoleID,
		NickName: nickName,
		Password: hashPassword,
	}).Error
	if err != nil {
		logrus.Errorf("用户创建失败 %s", err)
		return
	}
	logrus.Infof("用户创建成功")
}

/*
获取用户列表
*/
func (u *User) List() {
	var userList []models.UserModel
	global.DB.Order("created_time desc").Limit(10).Find(&userList)
	for _, user := range userList {
		fmt.Printf("用户id：%d  用户名：%s 用户昵称：%s 用户角色：%d 创建时间：%s\n",
			user.ID,
			user.UserName,
			user.NickName,
			user.RoleID,
			user.CreatedTime.Format("2006-01-02 15:04:05"),
		)
	}
}
