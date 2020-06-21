package main

import (
	"fmt"
	"log"

	"github.com/Piochat/GoPostgesql/db"
	"github.com/Piochat/GoPostgesql/models"
)

func exampleCreate() {
	e := models.Estudiante{
		Name:   "Alvaro",
		Active: true,
	}

	err := db.Create(e)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("CREADO EXITOSAMENTE", "!!!!!!!!!")
}

func exampleFind() {
	es, err := db.Find()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(es)
}

func main() {
	if db.PingToDB() {
		log.Fatalln("Error Unreachable database ")
	}

	// exampleCreate()
	exampleFind()

}
