package model

func Migration() {
	DB.AutoMigrate(&User{}, &Task{})
}
