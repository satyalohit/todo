package server

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"todo/model"
)

func Database() *gorm.DB {

	var err error
	db, err := gorm.Open("postgres", "user=postgres password=admin dbname=todo sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("db connected")	
	db.AutoMigrate(&model.Task{})
	return db
}
