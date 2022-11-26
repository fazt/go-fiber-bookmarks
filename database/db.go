package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Bookmark struct {
	gorm.Model
	Name string `json:"name"`
	URL  string `json:"url"`
}

func InitDB() error {
	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})

	if err != nil {
		return err
	}

	db.AutoMigrate(&Bookmark{})

	return nil
}

func CreateBookmark(name string, url string) (Bookmark, error) {
	bookmark := Bookmark{Name: name, URL: url}

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})

	if err != nil {
		return bookmark, err
	}

	db.Create(&bookmark)

	return bookmark, nil
}

func GetAllBookmarks() ([]Bookmark, error) {
	var bookmarks []Bookmark

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})

	if err != nil {
		return bookmarks, err
	}

	db.Find(&bookmarks)

	return bookmarks, nil
}

func GetBookmark(id string) (Bookmark, error) {
	var bookmark Bookmark

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})

	if err != nil {
		return bookmark, err
	}

	db.First(&bookmark, id)

	return bookmark, nil
}

func UpdateBookmark(id string, name string, url string) (Bookmark, error) {
	var bookmark Bookmark

	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})

	if err != nil {
		return bookmark, err
	}

	db.First(&bookmark, id)

	bookmark.Name = name
	bookmark.URL = url

	db.Save(&bookmark)

	return bookmark, nil
}

func DeleteBookmark(id string) (int64, error) {
	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})

	if err != nil {
		return 0, err
	}

	db = db.Delete(&Bookmark{}, id)

	return db.RowsAffected, nil
}
