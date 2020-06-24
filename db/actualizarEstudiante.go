package db

import (
	"errors"

	"github.com/Piochat/GoPostgesql/models"
)

//Update funcion para modificar informacion
func Update(e models.Estudiante) error {
	query := `UPDATE estudiantes SET
				name = $1, age = $2, active = $3, updated_at = now()
				WHERE id = $4`

	stmt, err := DbCon.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active, e.ID)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error 1 affected row was expected ")
	}

	return nil
}
