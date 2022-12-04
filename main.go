package main

import (
	"database/sql"

	"fyne.io/fyne/v2/app"
	"github.com/shplume/system/theme"
	"github.com/shplume/system/windows"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&theme.Theme{})

	dsn := "sqlserver://sa:123456@localhost?database=test&connection+timeout=30"
	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		panic(err)
	}

	controller := windows.NewController(&myApp, db)

	loginPage := *controller.GetLoginPage()

	loginPage.Show()

	myApp.Run()
}
