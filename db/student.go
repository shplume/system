package db

import (
	"database/sql"
	"fmt"
)

const (
	studentInsert = iota
	studentInfo
	studentDelete
	studentUpdate
	sexUpdate
)

var (
	studentSQLString = []string{
		`INSERT INTO students (id,s_name,sex,specialty,class,phone) VALUES (%d,'%s',%d,'%s','%s','%s')`,
		`SELECT s_name,sex,specialty,class,phone FROM students WHERE id = '%d'`,
		`DELETE FROM students WHERE id = %d`,
		`UPDATE students SET %s='%s' WHERE id = %d`,
		`UPDATE students SET sex= %d WHERE id = %d`,
	}
)

type Student struct {
	Id        int
	Name      string
	Sex       int
	Specialty string
	Class     string
	Phone     string
}

func InsertStudent(db *sql.DB, id int, name string, sex int, specialty string, class string, phone string) error {
	sql := fmt.Sprintf(studentSQLString[studentInsert], id, name, sex, specialty, class, phone)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalid
	}

	return nil
}

func GetStudent(db *sql.DB, id int) (*Student, error) {
	s := Student{Id: id}

	sql := fmt.Sprintf(studentSQLString[studentInfo], s.Id)
	err := db.QueryRow(sql).Scan(&s.Name, &s.Sex, &s.Specialty, &s.Class, &s.Phone)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func DeleteStudent(db *sql.DB, id int) error {
	err := DeleteUser(db, id)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf(studentSQLString[studentDelete], id)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalid
	}

	return nil
}

func UpdateStudent(db *sql.DB, sId int, query string, value string) error {
	sql := fmt.Sprintf(studentSQLString[studentUpdate], query, value, sId)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func UpdateSex(db *sql.DB, sId int, value int) error {
	sql := fmt.Sprintf(studentSQLString[sexUpdate], value, sId)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
