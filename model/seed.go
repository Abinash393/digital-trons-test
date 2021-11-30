package model

import "Abinash393/digital-trons-test/db"

func Seed() {
	user := User{Name: "Jon Doe", Email: "jon@test.com", Password: "qwerty"}
	db.ORM.Create(&user)

	rootFolder := Folder{ID: 1, Name: "root", ParaintFolderID: 1}
	db.ORM.Create(&rootFolder)
}
