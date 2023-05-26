package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/high-creek-software/visuals/internal/plausible"
	"golang.org/x/exp/slog"
)

type Sites struct {
    widget.BaseWidget
    
    repo SiteTokenRepo
    pairs []plausible.SiteTokenPair
    
    list *widget.List
    
    siteIDEntry *widget.Entry
    tokenEntry *widget.Entry
    saveBtn *widget.Button
}

func NewSites(repo SiteTokenRepo) *Sites {
    s := &Sites{repo: repo}
    s.ExtendBaseWidget(s)
    
    s.list = widget.NewList(s.length, s.createItem, s.updateItem)
    s.siteIDEntry = widget.NewEntry()
    s.siteIDEntry.PlaceHolder = "Site ID"
    s.tokenEntry = widget.NewEntry()
    s.tokenEntry.PlaceHolder = "Token"
    
    s.saveBtn = widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), s.save)
    
    s.loadPairs()
    
    return s
}

func (s *Sites) loadPairs() {
    var err error
    s.pairs, err = s.repo.List()
    if err != nil {
        slog.Info("error loading site token pairs", "error", err)
        return
    }
    s.list.Refresh()
}

func (s *Sites) CreateRenderer() fyne.WidgetRenderer {
    entriesList := container.NewVBox(s.siteIDEntry, s.tokenEntry, s.saveBtn)
    split := container.NewHSplit(s.list, entriesList)
    split.Offset = 0.22
    
    return widget.NewSimpleRenderer(split)    
}

func (s *Sites) length() int {
    return len(s.pairs)
}

func (s *Sites) createItem() fyne.CanvasObject {
    return widget.NewLabel("template")
}

func (s *Sites) updateItem(item widget.ListItemID, co fyne.CanvasObject) {
    pair := s.pairs[item]
    lbl := co.(*widget.Label)
    
    lbl.SetText(pair.SiteID)
}

func (s *Sites) save() {
    siteID := s.siteIDEntry.Text
    token := s.tokenEntry.Text
    
    pair := plausible.NewSiteTokenPair(siteID, token)
    err := s.repo.Store(pair)
    if err != nil {
        slog.Error("error storing site token pair", "error", err)
        return
    }
    slog.Info("saving pair", "pair", pair)
    s.loadPairs()
}