package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Shelter struct {
	gorm.Model
	name string
	email string
}
