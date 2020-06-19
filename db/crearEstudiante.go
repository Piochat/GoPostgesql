package db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/Piochat/GoPostgesql/models"
)

//Create insert data in db
func Create(e models.Estudiante) error {
	query := `INSERT INTO estudiantes (name, age, active)
				VALUES ($1, $2,$3)`

	intNull := sql.NullInt32{}
	//strNull := sql.NullString{}
	//boolNull := sql.NullBool{}

	stmtIns, err := DbCon.Prepare(query)
	if err != nil {
		log.Println("Error Stmt ins [file crearEstudiante]", err.Error())
		return err
	}
	defer stmtIns.Close()

	if e.Age == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int32 = int32(e.Age)
	}

	// if e.Name == "" {
	// 	strNull.Valid = false
	// } else {
	// 	strNull.Valid = true
	// 	strNull.String = e.Name
	// }

	if e.Active {

	}

	result, err := stmtIns.Exec(e.Name, intNull, e.Active)
	if err != nil {
		log.Println("Error Stmt ins exec [file crearEstudiante]", err.Error())
		return err
	}

	i, _ := result.RowsAffected()
	if i != 1 {
		return errors.New("Error 1 affected row was expected ")
	}

	return err
}
