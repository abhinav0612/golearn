package models

import (
	"fmt"

	"gorm.io/gorm"
)

// gorm.Model definition
// type Model struct {
// 	ID        uint           `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
//   }

type Person struct {
	gorm.Model
	Name  string
	Email string
}

type Engineer struct {
	Person Person `gorm:"embedded"`
	Field  string
	YOE    int
}

// Model above is same as below

// type Engineer struct {
// 	gorm.Model
// 	Name  string
// 	Email string
// 	Field string
// 	YOE int
// }

// *********** Hooks **************

func (e *Engineer) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("Called after engineer is added ", e.Person.Name)
	return
}
