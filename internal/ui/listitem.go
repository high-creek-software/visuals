package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ListItem struct {
    widget.BaseWidget
    
    title string
    data string
    
    titleLbl *widget.Label
    dataLbl *widget.Label
}

func NewListItem() *ListItem {
    li := &ListItem{}
    li.ExtendBaseWidget(li)
    
    li.titleLbl = widget.NewLabel("")
    li.dataLbl = widget.NewLabel("")
    
    return li
}

func (li *ListItem) CreateRenderer() fyne.WidgetRenderer {
    border := container.NewBorder(nil, nil, nil, li.dataLbl, li.titleLbl)
    
    return widget.NewSimpleRenderer(border)
}

func (li *ListItem) UpdateDataString(title, data string) {
    li.title = title
    li.data = data
    
    li.titleLbl.SetText(li.title)
    li.dataLbl.SetText(li.data)
}

func (li *ListItem) UpdateDataInt(title string, data int) {
    li.title = title
    li.data = fmt.Sprintf("%d", data)
    
    li.titleLbl.SetText(li.title)
    li.dataLbl.SetText(li.data)
}