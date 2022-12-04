package windows

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/shplume/system/db"
)

func (c *Controller) GetDeletePage() *fyne.Window {
	page := c.App.NewWindow("删除")
	page.Resize(fyne.NewSize(400, 200))

	infoCon := container.New(layout.NewCenterLayout(), widget.NewLabel("删除学生信息\n"))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("输入要删除学生学号")

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("输入管理员密码")

	success := dialog.NewInformation("成功", "学生帐号删除成功", page)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "学号", Widget: nameEntry},
			{Text: "密码", Widget: passEntry},
		},
		SubmitText: "删除",
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
			err := db.DeleteStudent(c.DB, id)
			if err != nil {
				information := dialog.NewInformation("失败", err.Error(), page)
				information.Show()
				return
			}

			success.Show()
		},
	}

	content := container.New(layout.NewVBoxLayout(), infoCon, form)
	page.SetContent(content)
	return &page
}
