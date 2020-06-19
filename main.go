package main

import (
	"fmt"
	"log"

	"github.com/Piochat/GoPostgesql/db"
	"github.com/Piochat/GoPostgesql/models"
)

func main() {
	e := models.Estudiante{
		Name:   "Alvaro",
		Active: true,
	}

	if db.PingToDB() {
		log.Fatalln("Error Unreachable database ")
	}
	err := db.Create(e)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("CREADO EXITOSAMENTE", "!!!!!!!!!")
}
