package utilities

import (
	"crudwithmysqlgorm/models"
	"fmt"

	"gorm.io/gorm"
)

func CreateExample(db *gorm.DB) {

	// creating single item
	person := models.Person{Name: "Abhinav", Email: "abhinav@gmail.com"}
	db.Create(&person)
	fmt.Println(person.ID, person.Name)

	engineer := models.Engineer{Person: person, Field: "Computer Science", YOE: 2}
	db.Create(&engineer)
	fmt.Println(engineer.Person.ID, engineer.Field, engineer.Person)

	// creating multiple items

	people := []models.Person{
		{Name: "John", Email: "john@gmail.com"},
		{Name: "Jane", Email: "jane@gmail.com"},
	}

	db.Create(&people)

	// creating records with selected fields

	db.Select("Name").Create(&person)

	// creating records without specific fields

	db.Omit("Name").Create(&person)

}

func QueryExample(db *gorm.DB) {

	var person models.Person
	db.First(&person)
	fmt.Println(person.ID, person.Name)

	var p models.Person
	db.Model(&models.Person{}).First(&p)
	fmt.Println(p.ID, p.Name)

	var p1 models.Person
	db.First(&p1, 2)
	fmt.Println(p1.ID, p1.Name)

	var p2 models.Person
	db.First(&p2, 3)
	fmt.Println(p2.ID, p2.Name)

	var p3 models.Person
	db.Select("Name").Omit("Email").Where("name = ? and email = ?", "Abhinav", "abhinav@gmail.com").First(&p3)
	fmt.Println(p3.ID, p3.Name, p3.Email)

}
