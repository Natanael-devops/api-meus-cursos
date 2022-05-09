package models

import "gorm.io/gorm"

type Curso struct {
	gorm.Model
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	Professor string `json:"professor"`
}
