package startdoor

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"imapp/windows"
)

func StartRun() {
	gregeoApp := app.New()
	gregeoApp.Settings().SetTheme(theme.LightTheme())
	//gregeoApp.
	windows.LoginWindow(gregeoApp)
}
