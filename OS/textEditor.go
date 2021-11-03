package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func showTextEditor() {
	w := myApp.NewWindow("Text Editor")
	w.Resize(fyne.NewSize(500, 280))

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Text Editor"),
		),
	)
	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++
	}))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	input.Resize(fyne.NewSize(400, 400))

	saveBtn := widget.NewButton("Save File", func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textdata := []byte(input.Text)

				uc.Write(textdata)
			}, w)

		saveFileDialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
		saveFileDialog.Show()

	})

	openBtn := widget.NewButton("Open File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w := fyne.CurrentApp().NewWindow(string(output.StaticName))

				w.SetContent(container.NewVBox(viewData))

				w.Resize(fyne.NewSize(400, 400))

				w.Show()
			}, w)

		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialog.Show()
	})

	editorContainer := container.NewVBox(
		container.NewVBox(
			content,
			input,
			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)

	w.SetContent(container.NewBorder(DeskBtn, nil, nil, nil, editorContainer))
	w.Show()
}
