package windows

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (c *Controller) GetQueryPage() *fyne.Window {
	page := c.App.NewWindow("查询")
	page.Resize(fyne.NewSize(400, 200))

	infoCon := container.New(layout.NewCenterLayout(), widget.NewLabel("查询学生信息\n"))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("输入要查询学生学号")

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("输入管理员密码")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "学号", Widget: nameEntry},
			{Text: "密码", Widget: passEntry},
		},
		SubmitText: "查询",
		OnSubmit: func() {
			if len(nameEntry.Text) == 0 {
				information := dialog.NewInformation("请输入学号", "学号不允许为空", page)
				information.Show()
				return
			}

			if len(passEntry.Text) == 0 {
				information := dialog.NewInformation("请输入密码", "管理员密码不允许为空", page)
				information.Show()
				return
			}

			id, _ := strconv.Atoi(nameEntry.Text)
			newPage := *c.GetStudentPage(id, 1)
			newPage.Show()

			page.Close()
		},
	}

	content := container.New(layout.NewVBoxLayout(), infoCon, form)
	page.SetContent(content)
	return &page
}
