package main

import (
	"fmt"
	"log"

	"github.com/Piochat/GoPostgesql/db"
	"github.com/Piochat/GoPostgesql/models"
)

func exampleCreate() {
	e := models.Estudiante{
		Name:   "Alejandro",
		Age:    39,
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

func exampleUpdate() {
	e := models.Estudiante{
		ID:     1,
		Name:   "Alvaro",
		Age:    40,
		Active: true,
	}

	err := db.Update(e)
	if err != nil {
		log.Println(err.Error())
	}
}

func exampleDelete() {
	err := db.Delete(3)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	if db.PingToDB() {
		log.Fatalln("Error Unreachable database ")
	}

	//exampleCreate()
	exampleUpdate()
	exampleFind()

}
