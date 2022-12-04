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

func (c *Controller) GetUpdatePage() *fyne.Window {
	page := c.App.NewWindow("修改密码")
	page.Resize(fyne.NewSize(400, 200))

	infoCon := container.New(layout.NewCenterLayout(), widget.NewLabel("学生密码修改\n"))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("输入帐号")

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("输入密码")

	confirmEntry := widget.NewPasswordEntry()
	confirmEntry.SetPlaceHolder("确认密码")

	success := dialog.NewInformation("成功", "学生密码修改成功", page)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "帐号", Widget: nameEntry},
			{Text: "密码", Widget: passEntry},
			{Text: "密码", Widget: confirmEntry},
		},
		SubmitText: "修改",
		OnSubmit: func() {
			if len(nameEntry.Text) == 0 {
				information := dialog.NewInformation("请输入帐号", "帐号不允许为空", page)
				information.Show()
				return
			}

			if len(passEntry.Text) == 0 {
				information := dialog.NewInformation("请输入密码", "密码不允许为空", page)
				information.Show()
				return
			}

			if passEntry.Text != confirmEntry.Text {
				information := dialog.NewInformation("修改失败", "两次密码输入不一致", page)
				information.Show()
				return
			}

			id, _ := strconv.Atoi(nameEntry.Text)
			err := db.UpdateUser(c.DB, id, passEntry.Text)
			if err != nil {
				information := dialog.NewInformation("修改失败", err.Error(), page)
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
