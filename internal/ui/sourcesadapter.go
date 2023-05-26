package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/high-creek-software/visuals/internal/plausible"
)

type SourcesAdapter struct {
    sources []plausible.Source
    list *widget.List
}

func (sa *SourcesAdapter) UpdateData(sources []plausible.Source) {
    sa.sources = sources
    sa.list.Refresh()
}

func (sa *SourcesAdapter) Count() int {
    return len(sa.sources)
}

func (sa *SourcesAdapter) Create() fyne.CanvasObject {
    return NewListItem()
}

func (sa *SourcesAdapter) Update(id widget.ListItemID, co fyne.CanvasObject) {
    src := sa.sources[id]
    co.(*ListItem).UpdateDataInt(src.Source, src.Visitors)
}