package windows

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/shplume/system/db"
)

func (c *Controller) GetStudentPage(id int, flag int) *fyne.Window {
	page := c.App.NewWindow("学生信息")
	page.Resize(fyne.NewSize(400, 200))

	s, err := db.GetStudent(c.DB, id)
	if err != nil {
		information := dialog.NewInformation("异常", err.Error(), page)
		information.Show()
		return &page
	}

	sex := "男"
	if s.Sex == 1 {
		sex = "女"
	}

	nameLabel := widget.NewLabel("姓名: " + s.Name)
	idLabel := widget.NewLabel("学号: " + fmt.Sprint(s.Id))
	sexLabel := widget.NewLabel("性别: " + sex)
	specialtyLabel := widget.NewLabel("专业: " + s.Specialty)
	classLabel := widget.NewLabel("班级: " + s.Class)
	phoneLabel := widget.NewLabel("电话号码: " + s.Phone)

	info := container.New(layout.NewVBoxLayout(), nameLabel, idLabel, sexLabel, specialtyLabel, classLabel, phoneLabel)

	exitBtn := widget.NewButton("退出登录", func() {
		page.Close()

		newPage := *c.GetLoginPage()
		newPage.Show()
	})

	if flag == 1 {
		exitBtn = widget.NewButton("关闭", func() {
			page.Close()
		})
	}

	page.SetContent(container.New(layout.NewVBoxLayout(), info, exitBtn))

	return &page
}
