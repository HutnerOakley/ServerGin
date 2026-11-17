package flags

import (
	"ServeGin/internal/global"
	"ServeGin/internal/models"

	"github.com/sirupsen/logrus"
)

/*
迁移表
*/
func MigrateDB() error {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		logrus.Errorf("表结构迁移失败 %s", err)
		return err
	}
	logrus.Infof("表结构迁移成功")
	return nil
}
