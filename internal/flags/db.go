package flags

import (
	"ServeGin/internal/global"
	"ServeGin/internal/models"
	"fmt"
)

/*
迁移表
*/
func MigrateDB() error {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		return fmt.Errorf("表结构迁移失败: %w", err)
	}
	return nil
}
