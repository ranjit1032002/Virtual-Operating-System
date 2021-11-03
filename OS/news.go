package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var num_articles int = 1

func showNews() {
	w := myApp.NewWindow("News App")
	w.Resize(fyne.NewSize(400, 400))

	//Api Part
	URL := "https://gnews.io/api/v4/search?q=kolkata&token=6b02f0fb87f4aca9066a1385a1bc08ca"

	res, _ := http.Get(URL)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	news, _ := UnmarshalNews(body)

	label := widget.NewLabel(fmt.Sprintf("No. Of Articles:%d",
		news.TotalArticles))

	label2 := widget.NewLabel(fmt.Sprintf("%s", news.Articles[1].Title))
	label2.TextStyle = fyne.TextStyle{Bold: true}
	label2.Wrapping = fyne.TextWrapBreak

	entry1 := widget.NewLabel(fmt.Sprintf("%s", news.Articles[1].Description))
	entry1.Wrapping = fyne.TextWrapBreak

	btn := widget.NewButton("Next", func() {
		num_articles += 1
		label2.Text = news.Articles[num_articles].Title
		entry1.Text = news.Articles[num_articles].Description
		label2.Refresh()
		entry1.Refresh()
	})

	label3 := canvas.NewText("News App", color.White)
	label3.Alignment = fyne.TextAlignCenter
	label.TextStyle = fyne.TextStyle{Bold: true}

	img := canvas.NewImageFromFile("news.png")
	img.FillMode = canvas.ImageFillOriginal

	e := container.NewVBox(label3, label2, entry1, btn)

	e.Resize(fyne.NewSize(300, 300))

	c := container.NewBorder(img, label, nil, nil, e)

	w.SetContent(c)

	w.Show()

}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    news, err := UnmarshalNews(bytes)
//    bytes, err = news.Marshal()

func UnmarshalNews(data []byte) (News, error) {
	var r News
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *News) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type News struct {
	TotalArticles int64     `json:"totalArticles"`
	Articles      []Article `json:"articles"`
}

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	URL         string `json:"url"`
	Image       string `json:"image"`
	PublishedAt string `json:"publishedAt"`
	Source      Source `json:"source"`
}

type Source struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
