package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showGallaryApp() {

	w := myApp.NewWindow("Gallary")
	w.Resize(fyne.NewSize(1280, 720))
	root_src := "C:\\Users\\ranji\\Desktop"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	tabs := container.NewAppTabs()
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			extensions := strings.Split(file.Name(), ".")[1]
			if extensions == "png" || extensions == "jpeg" {
				image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
				image.FillMode = canvas.ImageFillOriginal
				tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(container.NewBorder(DeskBtn, nil, nil, nil, tabs))
	w.Show()

}
