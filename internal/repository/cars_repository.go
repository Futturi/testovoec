package repository

import (
	"fmt"
	"github.com/Futturi/testovoe/internal/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

type Cars_Repo struct {
	db *sqlx.DB
}

func NewCars_Repo(db *sqlx.DB) *Cars_Repo {
	return &Cars_Repo{db: db}
}

func (r *Cars_Repo) GetNomers(page int, query map[string]string) ([]models.Car, error) {
	var result []models.Car
	querys := make([]string, 0)
	values := make([]interface{}, 0)
	argid := 1
	for key, value := range query {
		if len(querys) == 0 {
			querys = append(querys, fmt.Sprintf("%v = $1", key))
		} else {
			querys = append(querys, fmt.Sprintf("AND %v = $%d", key, argid))
		}
		values = append(values, value)
		argid++
	}
	query1 := ""
	if len(query) == 0 {
		query1 = "SELECT regnum, mark, model, name, surname FROM nomera"
	} else {
		query1 = fmt.Sprintf("SELECT regnum, mark, model, name, surname FROM nomera WHERE %s", strings.Join(querys, " "))
	}
	fmt.Println(query1)
	if err := r.db.Select(&result, query1, values...); err != nil {
		return []models.Car{}, err
	}
	return result, nil
}
func (r *Cars_Repo) Delete(nomer models.Car) error {
	query := fmt.Sprintf("DELETE FROM nomera WHERE regnum = $1")
	_, err := r.db.Exec(query, nomer.Nomer)
	if err != nil {
		return err
	}
	return nil
}
func (r *Cars_Repo) Put(nomer models.CarUpdate) error {
	args := make([]interface{}, 0)
	setVal := make([]string, 0)
	argid := 1
	if nomer.Name != nil {
		setVal = append(setVal, fmt.Sprintf("name=$%d", argid))
		args = append(args, *nomer.Name)
		argid++
	}
	if nomer.Mark != nil {
		setVal = append(setVal, fmt.Sprintf("mark=$%d", argid))
		args = append(args, *nomer.Mark)
		argid++
	}
	if nomer.Model != nil {
		setVal = append(setVal, fmt.Sprintf("models=$%d", argid))
		args = append(args, *nomer.Model)
		argid++
	}
	if nomer.Surname != nil {
		setVal = append(setVal, fmt.Sprintf("surname=$%d", argid))
		args = append(args, *nomer.Surname)
		argid++
	}
	setQuery := strings.Join(setVal, ",")
	query := fmt.Sprintf("UPDATE nomera SET %s WHERE regNum = $%d", setQuery, argid)
	args = append(args, *nomer.Nomer)
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Cars_Repo) Insert(nomers []models.Car) error {
	queryStr := ""
	for ind, value := range nomers {
		if ind != len(nomers)-1 {
			queryStr += fmt.Sprintf("('%s', '%s', '%s', '%s', '%s'),", value.Nomer, value.Mark, value.Model, value.Name, value.Surname)
		} else {
			queryStr += fmt.Sprintf("('%s', '%s', '%s', '%s', '%s');", value.Nomer, value.Mark, value.Model, value.Name, value.Surname)
		}
	}
	query := fmt.Sprintf("INSERT INTO nomera(regNum, mark, model, name, surname) VALUES %s", queryStr)
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
