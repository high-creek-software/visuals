package internal

import (
	"encoding/json"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/high-creek-software/visuals/internal/plausible"
	"github.com/high-creek-software/visuals/internal/ui"
	"golang.org/x/exp/slog"
)

type Visuals struct {
    app fyne.App
    mainWindow fyne.Window
    
    siteSelect *widget.Select
    siteTokenRepo ui.SiteTokenRepo
    docTabs *container.DocTabs
    
    availablePairs []plausible.SiteTokenPair
}

func NewVisualsApp(a fyne.App, win fyne.Window) *Visuals {
    v := &Visuals{app: a, mainWindow: win}
    v.initLayout()
    
    v.siteTokenRepo = NewSiteTokenFynePreferences(v.app.Preferences())
    
    return v
}

func (v *Visuals) initLayout() {
    
    settingsAction := widget.NewButtonWithIcon("", theme.SettingsIcon(), func() {
        content := ui.NewSettingsWindow(v.siteTokenRepo)
        win := v.app.NewWindow("Settings")
        win.SetContent(content)
        win.Resize(fyne.NewSize(800, 450))
        win.Show()
    })
    
    v.siteSelect = widget.NewSelect([]string{""}, v.siteSelected)
    v.docTabs = container.NewDocTabs()
    
    topBorder := container.NewBorder(nil, nil, nil, settingsAction, v.siteSelect)
    mainBorder := container.NewBorder(topBorder, nil, nil, nil, v.docTabs)
    
    v.mainWindow.SetContent(mainBorder)
}

func (v *Visuals) Start() {
    v.app.Preferences().AddChangeListener(func() {
        v.updateAvailableSites()
    })
    v.updateAvailableSites()
    v.mainWindow.ShowAndRun()
}

func (v *Visuals) updateAvailableSites() {
    pairs, err := v.siteTokenRepo.List()
    if err != nil {
        slog.Error("error loading site token paris", "error", err)
        return
    }
    
    v.availablePairs = pairs
    var ids []string
    for _, p := range pairs {
        ids = append(ids, p.SiteID)
    }
    
    v.siteSelect.Options = ids
    v.siteSelect.Refresh()
}

func (v *Visuals) siteSelected(selection string) {
    var pair plausible.SiteTokenPair
    for _, p := range v.availablePairs {
        if p.SiteID == selection {
            pair = p
        }
    }
    
    if pair.SiteID == "" {
        slog.Error("error finding pair", "selection", selection)
        return
    }
    
    statsRepo := plausible.NewStatsRepoResty(pair)
    tab := container.NewTabItem(selection, ui.NewSiteDisplay(v.mainWindow.Canvas(), statsRepo))
    v.docTabs.Append(tab)
}

const (
    keySites = "site_token_pairs"    
)

type SiteTokenRepoFynePreferences struct {
    prefs fyne.Preferences
}

func NewSiteTokenFynePreferences(prefs fyne.Preferences) ui.SiteTokenRepo {
    return &SiteTokenRepoFynePreferences{prefs: prefs}
}

func (str *SiteTokenRepoFynePreferences) List() ([]plausible.SiteTokenPair, error) {
    serialized := str.prefs.String(keySites)
    if serialized == "" {
        return nil, nil
    }
    
    var items []plausible.SiteTokenPair
    err := json.Unmarshal([]byte(serialized), &items)
    
    return items, err
}

func (str *SiteTokenRepoFynePreferences) Store(pair plausible.SiteTokenPair) error {
    items, err := str.List()
    if err != nil {
        return err
    }
    items = append(items, pair)
    return str.serialize(items)
}

func (str *SiteTokenRepoFynePreferences) Delete(pair plausible.SiteTokenPair) error {
    return nil
}

func (str *SiteTokenRepoFynePreferences) serialize(items []plausible.SiteTokenPair) error {
    serialized, err := json.Marshal(items)
    if err != nil {
        return err
    }
    
    str.prefs.SetString(keySites, string(serialized))
    return nil
}
