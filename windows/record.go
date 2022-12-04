package windows

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/shplume/system/db"
)

const (
	unknown = iota - 1
	man
	woman
)

func (c *Controller) GetRecordPage() *fyne.Window {
	page := c.App.NewWindow("录入信息")
	page.Resize(fyne.NewSize(400, 200))

	infoCon := container.New(layout.NewCenterLayout(), widget.NewLabel("学生信息录入\n"))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("输入姓名")

	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("输入学号")

	sex := unknown
	sexRadio := widget.NewRadioGroup([]string{"男", "女"}, func(s string) {
		if s == "男" {
			sex = man
		} else {
			sex = woman
		}
	})

	specialty := ""
	specialtySelect := widget.NewSelect([]string{"CS", "EE", "ME", "MS", "CE"}, func(value string) {
		specialty = value
	})

	classEntry := widget.NewEntry()
	classEntry.SetPlaceHolder("输入班级")

	phoneEntry := widget.NewEntry()
	phoneEntry.SetPlaceHolder("输入电话号码")

	success := dialog.NewInformation("成功", "学生信息录入成功", page)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "姓名", Widget: nameEntry},
			{Text: "学号", Widget: idEntry},
			{Text: "性别", Widget: sexRadio},
			{Text: "电话号码", Widget: phoneEntry},
			{Text: "专业", Widget: specialtySelect},
			{Text: "班级", Widget: classEntry},
		},
		SubmitText: "注册",
		OnSubmit: func() {
			if len(nameEntry.Text) == 0 {
				information := dialog.NewInformation("请输入姓名", "姓名不允许为空", page)
				information.Show()
				return
			}

			if len(idEntry.Text) == 0 {
				information := dialog.NewInformation("请输入学号", "学号不允许为空", page)
				information.Show()
				return
			}

			if sex == unknown {
				information := dialog.NewInformation("请选择性别", "性别不允许为空", page)
				information.Show()
				return
			}

			if len(phoneEntry.Text) == 0 {
				information := dialog.NewInformation("请输入电话号码", "电话号码不允许为空", page)
				information.Show()
				return
			}

			if len(phoneEntry.Text) == 11 {
				information := dialog.NewInformation("请输入正确的电话号码", "电话号码格式错误", page)
				information.Show()
				return
			}

			if specialty == "" {
				information := dialog.NewInformation("请选择专业", "专业不允许为空", page)
				information.Show()
				return
			}

			if len(classEntry.Text) == 0 {
				information := dialog.NewInformation("请选择班级", "班级不允许为空", page)
				information.Show()
				return
			}

			id, _ := strconv.Atoi(idEntry.Text)
			err := db.InsertStudent(c.DB, id, nameEntry.Text, sex, specialty, classEntry.Text, phoneEntry.Text)
			if err != nil {
				information := dialog.NewInformation("录入失败", err.Error(), page)
				information.Show()
				fmt.Println(err)
				return
			}

			success.Show()
		},
	}

	content := container.New(layout.NewVBoxLayout(), infoCon, form)
	page.SetContent(content)

	return &page
}
