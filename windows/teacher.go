package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (c *Controller) GetTeacherPage() *fyne.Window {
	page := c.App.NewWindow("功能选择")
	page.Resize(fyne.NewSize(400, 220))

	title := container.New(layout.NewCenterLayout(), widget.NewLabel("功能选择\n"))

	recordBtn := widget.NewButton("学生信息录入", func() {
		newPage := *c.GetRecordPage()
		newPage.Show()
	})

	queryBtn := widget.NewButton("学生信息查询", func() {
		newPage := *c.GetQueryPage()
		newPage.Show()
	})

	modifyBtn := widget.NewButton("学生信息修改", func() {
		newPage := *c.GetUpdateStudentPage()
		newPage.Show()
	})

	deleteBtn := widget.NewButton("学生信息删除", func() {
		newPage := *c.GetDeletePage()
		newPage.Show()
	})

	registerBtn := widget.NewButton("学生帐号注册", func() {
		newPage := *c.GetRegisterPage()
		newPage.Show()
	})

	logoutBtn := widget.NewButton("修改学生密码", func() {
		newPage := *c.GetUpdatePage()
		newPage.Show()
	})

	exitBtn := widget.NewButton("退出登录", func() {
		page.Close()

		newPage := *c.GetLoginPage()
		newPage.Show()
	})

	infoCon := container.New(layout.NewGridLayout(2), recordBtn, queryBtn, modifyBtn, deleteBtn)
	accountCon := container.New(layout.NewGridLayout(2), registerBtn, logoutBtn)

	label := widget.NewLabel("")

	content := container.New(layout.NewVBoxLayout(), title, infoCon, accountCon, label, exitBtn)
	page.SetContent(content)

	return &page
}
