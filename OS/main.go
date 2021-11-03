package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWindow fyne.Window = myApp.NewWindow("OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget

var img fyne.CanvasObject

var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main() {

	img = canvas.NewImageFromFile("Desktop.jpg")

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		showWeatherApp()
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalc()
	})

	btn3 = widget.NewButtonWithIcon("Gallary App", theme.StorageIcon(), func() {
		showGallaryApp()
	})

	btn4 = widget.NewButtonWithIcon("Text Editor", theme.DocumentIcon(), func() {
		showTextEditor()
	})

	btn5 = widget.NewButtonWithIcon("Music Player", theme.MediaMusicIcon(), func() {
		showMusicPlyer()
	})

	btn6 = widget.NewButtonWithIcon("News", theme.InfoIcon(), func() {
		showNews()
	})

	DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox(container.NewGridWithColumns(5, DeskBtn, btn1, btn2, btn3, btn4, btn5, btn6))

	myWindow.Resize(fyne.NewSize(1280, 680))
	myWindow.CenterOnScreen()

	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)

	myWindow.ShowAndRun()
}
