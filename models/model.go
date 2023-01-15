package models

import "gorm.io/gorm"

type Blues struct {
	Id     int     `json:"id" gorm:"primary_key"`
	Number float32 `json:"number"`
	Title  string  `json:"title"`
	Level  string  `json:"level"`
	Page   int     `json:"page"`
	Path   string  `json:"path"`
}

type Reds struct {
	Id        int    `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Character string `json:"character"`
	Path      string `json:"path"`
	Level     string `json:"level"`
}

type PreExam struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Level string `json:"level"`
	Page  int    `json:"page"`
	Path  string `json:"path"`
}

type Masters struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Page  int    `json:"page"`
	Path  string `json:"path"`
}

type Dictionaries struct {
	Id    int    `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Path  string `json:"path"`
	Page  int    `json:"page"`
}

type Comics struct {
	gorm.Model
	Id     int
	Title  string
	Autor  string
	Status string
	Cover  string
}

type Contents struct {
	Id          int
	ComicId     int
	Comic       Comics
	Category    string
	Description string
	Chapters    int
}

type Chapters struct {
	ChapterNumber int
	SourceUrl     string
	Title         string
}
