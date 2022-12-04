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

const (
	student = iota
	teacher
)

func (c *Controller) GetLoginPage() *fyne.Window {
	page := c.App.NewWindow("登录")
	page.Resize(fyne.NewSize(400, 200))

	infoCon := container.New(layout.NewCenterLayout(), widget.NewLabel("学生信息管理系统"))

	identity := student
	identityRadio := widget.NewRadioGroup([]string{"学生", "老师"}, func(s string) {
		if s == "学生" {
			identity = student
		} else {
			identity = teacher
		}
	})

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("输入学号或工号")

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("输入密码")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "你的身份", Widget: identityRadio},
			{Text: "学号或工号", Widget: nameEntry},
			{Text: "密码", Widget: passEntry},
		},
		SubmitText: "登录",
		OnSubmit: func() {
			if len(nameEntry.Text) == 0 {
				information := dialog.NewInformation("请输入学号或工号", "学号或工号不允许为空", page)
				information.Show()
				return
			}

			if len(passEntry.Text) == 0 {
				information := dialog.NewInformation("请输入密码", "密码不允许为空", page)
				information.Show()
				return
			}

			if identity == student {
				id, _ := strconv.Atoi(nameEntry.Text)
				err := db.UserLogin(c.DB, id, passEntry.Text)
				if err != nil {
					information := dialog.NewInformation("登录失败", err.Error(), page)
					information.Show()
					return
				}

				newPage := *c.GetStudentPage(id, 0)
				newPage.Show()
			} else {
				err := db.AdminLogin(c.DB, nameEntry.Text, passEntry.Text)
				if err != nil {
					information := dialog.NewInformation("登录失败", err.Error(), page)
					information.Show()
					return
				}

				newPage := *c.GetTeacherPage()
				newPage.Show()
			}

			page.Close()
		},
	}

	formCon := container.New(layout.NewVBoxLayout(), form)

	labelCon := container.New(layout.NewCenterLayout(), widget.NewLabel("\n如忘记密码，请联系管理员找回"))

	content := container.New(layout.NewVBoxLayout(), infoCon, formCon, labelCon)
	page.SetContent(content)

	return &page
}
