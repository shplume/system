package windows

import (
	"database/sql"

	"fyne.io/fyne/v2"
)

type Controller struct {
	App fyne.App
	DB  *sql.DB
}

func NewController(App *fyne.App, db *sql.DB) Controller {
	return Controller{App: *App, DB: db}
}
