package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func TalkMainWindow(gregeoApp fyne.App) {

	//appWindow := gregeoApp.NewWindow("成功登录")
	//appWindow.Resize(fyne.NewSize(300, 200))

	//menuList := widget.NewList(
	//	func() int { return 3 },
	//	func() fyne.CanvasObject {
	//		return widget.NewLabel("")
	//	},
	//	func(id widget.ListItemID, item fyne.CanvasObject) {
	//		switch id {
	//		case 0:
	//			item.(*widget.Label).SetText("对话")
	//		case 1:
	//			item.(*widget.Label).SetText("联系人")
	//		case 2:
	//			item.(*widget.Label).SetText("设置")
	//		}
	//	})
	//
	//rightContent := widget.NewLabel("请选择左侧菜单")
	//// 当用户选择菜单项时，更新右侧内容
	//menuList.OnSelected = func(id int) {
	//	switch id {
	//	case 0:
	//		rightContent.SetText("对话内容")
	//	case 1:
	//		rightContent.SetText("联系人列表")
	//	case 2:
	//		rightContent.SetText("设置页面")
	//	}
	//}
	//
	//// 创建一个水平分割容器
	//splitContainer := container.NewHSplit(menuList, rightContent)
	//splitContainer.SetOffset(0.25) // 设置左侧菜单的宽度比例
	//
	//// 设置主窗口的内容
	//appWindow.SetContent(splitContainer)
	//appWindow.Resize(fyne.NewSize(600, 400))

	//successMessage := widget.NewLabel("登录成功！欢迎使用应用！")
	//
	//appWindow.SetContent(container.NewVBox(
	//	successMessage,
	//	widget.NewButton("退出", func() {
	//		gregeoApp.Quit()
	//	}),
	//))

	appWindow := gregeoApp.NewWindow("成功登录")

	// 左侧的头像
	avatar := canvas.NewCircle(theme.PrimaryColor())
	avatar.Resize(fyne.NewSize(60, 60))
	//avatar.SetMinSize(fyne.NewSize(60, 60)) // 设置头像大小

	// 左侧菜单
	dialogMenu := widget.NewButton("对话", func() {})
	contactsMenu := widget.NewButton("联系人", func() {})
	favoritesMenu := widget.NewButton("收藏", func() {})

	menu := container.NewVBox(
		dialogMenu,
		contactsMenu,
		favoritesMenu,
	)

	// 底部的设置按钮
	settings := widget.NewButtonWithIcon("设置", theme.SettingsIcon(), func() {})

	// 左侧容器 (头像 + 菜单 + 设置)
	leftPanel := container.NewBorder(avatar, settings, nil, nil, menu)

	// 左侧固定宽度
	leftPanelContainer := container.NewMax(container.NewBorder(nil, nil, nil, nil, leftPanel))
	leftPanelContainer.Resize(fyne.NewSize(150, 600)) // 固定宽度为 150

	// 右侧的两栏
	chatList := widget.NewLabel("聊天列表区域")   // 右侧第一栏
	chatDetail := widget.NewLabel("聊天内容区域") // 右侧第二栏
	rightPanel := container.NewGridWithColumns(2, chatList, chatDetail)

	// 整体布局：左侧是固定宽度的栏目，右侧是两栏的区域
	content := container.NewHSplit(leftPanelContainer, rightPanel)
	content.Offset = 0.2 // 设置左右栏比例，20% 宽度用于左侧

	appWindow.SetContent(content)
	appWindow.Resize(fyne.NewSize(800, 600)) // 窗口大小

	appWindow.Show()
}
