package model

func Seed() {
	rootFolder := Folder{ID: 1, Name: "root", ParaintFolderID: 1}
	DB.Create(&rootFolder)
}
