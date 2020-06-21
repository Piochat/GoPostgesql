package db

import (
	"database/sql"

	"github.com/Piochat/GoPostgesql/models"
	"github.com/lib/pq"
)

//Find devulve los alumnos de la base de datos
func Find() (estudiantes []models.Estudiante, err error) {
	query := `SELECT id, name, age, active, created_at, updated_at
				FROM estudiantes`

	timeNull := pq.NullTime{}
	intNull := sql.NullInt32{}

	rows, err := DbCon.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		e := models.Estudiante{}
		err = rows.Scan(
			&e.ID,
			&e.Name,
			&intNull,
			&e.Active,
			&e.CreatedAt,
			&timeNull)

		if err != nil {
			return
		}

		e.Age = int16(intNull.Int32)
		e.UpdateAt = timeNull.Time

		estudiantes = append(estudiantes, e)
	}

	return estudiantes, nil
}
