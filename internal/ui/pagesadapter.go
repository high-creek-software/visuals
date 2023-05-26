package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/high-creek-software/visuals/internal/plausible"
)

type PagesAdapter struct {
    pages []plausible.Page
    list *widget.List
}

func (pa *PagesAdapter) UpdateData(pages []plausible.Page) {
    pa.pages = pages
    pa.list.Refresh()
}

func (pa *PagesAdapter) Count() int {
    return len(pa.pages)
}

func (pa *PagesAdapter) Create() fyne.CanvasObject {
    return NewListItem()
}

func (pa *PagesAdapter) Update(id widget.ListItemID, co fyne.CanvasObject) {
    pg := pa.pages[id]
    co.(*ListItem).UpdateDataInt(pg.Page, pg.Visitors)
}