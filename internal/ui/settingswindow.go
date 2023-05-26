package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SettingsWindow struct {
	widget.BaseWidget

	appTabs *container.AppTabs
	sites *Sites
}

func NewSettingsWindow(str SiteTokenRepo) *SettingsWindow {
	sw := &SettingsWindow{}
	sw.ExtendBaseWidget(sw)

	sw.sites = NewSites(str)
	sw.appTabs = container.NewAppTabs(container.NewTabItem("Sites", sw.sites))
	
	return sw
}

func (sw *SettingsWindow) CreateRenderer() fyne.WidgetRenderer {
    return widget.NewSimpleRenderer(sw.appTabs)    
}