package main

import (
	model "Abinash393/digital-trons-test/model"
)

func init() {
	model.InitializeDB()
	model.DB.AutoMigrate(&model.User{}, &model.Folder{}, &model.File{})
	model.Seed()
}

func main() {
}
