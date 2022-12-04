package db

import (
	"database/sql"
	"fmt"

	"github.com/shplume/system/util"
)

const (
	adminInsert = iota
	adminLogin
)

var (
	adminSQLString = []string{
		`INSERT INTO admins (account,psw) VALUES (%d,%s)`,
		`SELECT id,psw FROM admins WHERE account = '%s'`,
	}
)

func AdminInsertUser(db *sql.DB, sId int, psw string) error {
	hash, err := util.Generate(&psw)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf(adminSQLString[adminInsert], sId, hash)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalid
	}

	return nil
}

func AdminLogin(db *sql.DB, account string, psw string) error {
	var (
		id  uint32
		pwd string
	)

	sql := fmt.Sprintf(adminSQLString[adminLogin], account)
	err := db.QueryRow(sql).Scan(&id, &pwd)
	if err != nil {
		return err
	}

	if !util.Compare([]byte(pwd), &psw) {
		return errLoginFailed
	}

	return nil
}
