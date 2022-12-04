package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/shplume/system/util"
)

const (
	userInsert = iota
	userLogin
	userDelete
	userUpdate
)

var (
	errInvalid     = errors.New("操作失败: 0 行受影响")
	errLoginFailed = errors.New("帐号或密码错误")

	userSQLString = []string{
		`INSERT INTO users (s_id,psw) VALUES (%d,'%s')`,
		`SELECT id,psw FROM users WHERE s_id = %d`,
		`DELETE FROM users WHERE s_id = %d`,
		`UPDATE users SET psw='%s' WHERE s_id = %d`,
	}
)

func InsertUser(db *sql.DB, sId int, psw string) error {
	hash, err := util.Generate(&psw)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf(userSQLString[userInsert], sId, hash)
	result, err := db.Exec(sql)
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return errInvalid
	}

	return nil
}

func UserLogin(db *sql.DB, sId int, psw string) error {
	var (
		id  uint32
		pwd string
	)

	sql := fmt.Sprintf(userSQLString[userLogin], sId)
	err := db.QueryRow(sql).Scan(&id, &pwd)
	if err != nil {
		return err
	}

	if !util.Compare([]byte(pwd), &psw) {
		return errLoginFailed
	}

	return nil
}

func DeleteUser(db *sql.DB, sId int) error {
	sql := fmt.Sprintf(userSQLString[userDelete], sId)
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(db *sql.DB, sId int, psw string) error {
	hash, err := util.Generate(&psw)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf(userSQLString[userUpdate], hash, sId)
	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}
