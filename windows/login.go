package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func LoginWindow(gregeoApp fyne.App) {
	loginWindow := gregeoApp.NewWindow("睿安即时应用")
	loginWindow.Resize(fyne.NewSize(800, 500))

	entry := widget.NewEntry()
	textArea := widget.NewMultiLineEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "用户名", Widget: entry},
			{Text: "密码", Widget: textArea},
		},
		OnSubmit: func() {
			log.Println("form submitted: ", entry.Text)
			log.Println("multi line submitted: ", textArea.Text)
			loginWindow.Hide()
			TalkMainWindow(gregeoApp)
		},
	}

	hello := widget.NewLabel("请登录")
	loginWindow.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("欢迎你")
		}),
	))
	loginWindow.SetContent(container.NewVBox(form))

	loginWindow.ShowAndRun()
}
