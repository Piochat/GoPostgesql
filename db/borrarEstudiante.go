package db

import "errors"

func Delete(id init) error {
	query := `DELETE FROM estudiantes WHERE id = $1`

	stmt, err := DbCon.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error 1 affected row was expected ")
	}

	return nil
}
