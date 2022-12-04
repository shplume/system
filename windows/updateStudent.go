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

func (c *Controller) GetUpdateStudentPage() *fyne.Window {
	page := c.App.NewWindow("修改信息")
	page.Resize(fyne.NewSize(400, 200))

	infoCon := container.New(layout.NewCenterLayout(), widget.NewLabel("学生信息修改\n"))

	idEntry := widget.NewEntry()
	idEntry.SetPlaceHolder("输入修改学生的学号")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("输入姓名")

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

	success := dialog.NewInformation("成功", "学生信息修改成功", page)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "学号", Widget: idEntry},
			{Text: "姓名", Widget: nameEntry},
			{Text: "性别", Widget: sexRadio},
			{Text: "电话号码", Widget: phoneEntry},
			{Text: "专业", Widget: specialtySelect},
			{Text: "班级", Widget: classEntry},
		},
		SubmitText: "修改",
		OnSubmit: func() {
			id, _ := strconv.Atoi(idEntry.Text)

			if len(nameEntry.Text) != 0 {
				err := db.UpdateStudent(c.DB, id, "s_name", nameEntry.Text)
				if err != nil {
					fmt.Println(1, err)
					information := dialog.NewInformation("更新失败", err.Error(), page)
					information.Show()
					return
				}
			}

			if sex != unknown {
				err := db.UpdateSex(c.DB, id, sex)
				if err != nil {
					fmt.Println(2, err)
					information := dialog.NewInformation("更新失败", err.Error(), page)
					information.Show()
					return
				}
			}

			if len(phoneEntry.Text) != 0 {
				err := db.UpdateStudent(c.DB, id, "phone", phoneEntry.Text)
				if err != nil {
					fmt.Println(3, err)
					information := dialog.NewInformation("更新失败", err.Error(), page)
					information.Show()
					return
				}
			}

			if specialty != "" {
				err := db.UpdateStudent(c.DB, id, "specialty", specialty)
				if err != nil {
					fmt.Println(4, err)
					information := dialog.NewInformation("更新失败", err.Error(), page)
					information.Show()
					return
				}
			}

			if len(classEntry.Text) != 0 {
				err := db.UpdateStudent(c.DB, id, "class", classEntry.Text)
				if err != nil {
					fmt.Println(5, err)
					information := dialog.NewInformation("更新失败", err.Error(), page)
					information.Show()
					return
				}
			}

			success.Show()
		},
	}

	content := container.New(layout.NewVBoxLayout(), infoCon, form)
	page.SetContent(content)

	return &page
}
