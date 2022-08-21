package model

// migration 设置表结构自动同步
func migration() {
	//设置 user 表和 task 表自动迁移
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}, &Task{})

	if err != nil {
		return
	}
}
