package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/x/fyne/layout"
	"github.com/high-creek-software/fynecharts"
	"github.com/high-creek-software/visuals/internal/plausible"
	"golang.org/x/exp/slog"
)

type SiteDisplay struct {
    widget.BaseWidget
    
    statsRepo plausible.StatsRepo
    
    timeChart *fynecharts.TimeSeriesChart
    
    topSourcesLbl *widget.RichText
    topSourcesList *widget.List
    topSourcesAdapter *SourcesAdapter
    
    topPagesLbl *widget.RichText
    topPagesList *widget.List
    topPagesAdapter *PagesAdapter
}

func NewSiteDisplay(canvas fyne.Canvas, statsRepo plausible.StatsRepo) *SiteDisplay {
    sd := &SiteDisplay{statsRepo: statsRepo}
    sd.ExtendBaseWidget(sd)
    
    sd.timeChart = fynecharts.NewTimeSeriesChart(canvas, "", nil, nil)
    sd.timeChart.SetMinHeight(250)
    sd.timeChart.UpdateDotDiameter(8)
    
    sd.topSourcesLbl = widget.NewRichTextFromMarkdown("### Top Sources")
    sd.topSourcesAdapter = &SourcesAdapter{}
    sd.topSourcesList = widget.NewList(sd.topSourcesAdapter.Count, sd.topSourcesAdapter.Create, sd.topSourcesAdapter.Update)
    sd.topSourcesAdapter.list = sd.topSourcesList
    
    sd.topPagesLbl = widget.NewRichTextFromMarkdown("### Top Pages")
    sd.topPagesAdapter = &PagesAdapter{}
    sd.topPagesList = widget.NewList(sd.topPagesAdapter.Count, sd.topPagesAdapter.Create, sd.topPagesAdapter.Update)
    sd.topPagesAdapter.list = sd.topPagesList
    
    sd.reloadData()
    
    return sd
}

func (sd *SiteDisplay) CreateRenderer() fyne.WidgetRenderer {
    
    sourcesBorder := container.NewBorder(sd.topSourcesLbl, nil, nil, nil, sd.topSourcesList)
    pagesBorder := container.NewBorder(sd.topPagesLbl, nil, nil, nil, sd.topPagesList)
    
    
    resp := layout.NewResponsiveLayout(
        layout.Responsive(container.NewMax(NewSpacer(0, 200), sourcesBorder), 1, 1, 0.5),
        layout.Responsive(container.NewMax(NewSpacer(0, 200), pagesBorder), 1, 1, 0.5),
    )
    
    return widget.NewSimpleRenderer(container.NewPadded(container.NewBorder(container.NewPadded(sd.timeChart), nil, nil, nil, resp)))
}

func (sd *SiteDisplay) reloadData() {
    go func() {
        if srcs, err := sd.statsRepo.LoadTopSources(""); err == nil {
            sd.topSourcesAdapter.UpdateData(srcs)
        } else {
            slog.Error("error loading top sources", "error", err)
        }
    }()
    
    go func() {
        if pgs, err := sd.statsRepo.LoadTopPages(""); err == nil {
            sd.topPagesAdapter.UpdateData(pgs)
        } else {
            slog.Error("error loading top pages", "error", err)
        }
    }()
    
    go func() {
        if ts, err := sd.statsRepo.LoadTimeSeries(""); err == nil {
            
            var lbls []string
            var data []float64
            
            for _, t := range ts {
                at := time.Time(t.Date)
                lbls = append(lbls, at.Format("2"))
                data = append(data, float64(t.Visitors))
            }
            
            sd.timeChart.UpdateData(lbls, data)
        } else {
            slog.Error("error loading time series", "error", err)
        }
    }()
    
    go func() {
       if agg, err := sd.statsRepo.LoadAggregate(""); err == nil {
           slog.Info("Aggregate data", "aggregate", agg)
       } else {
           slog.Error("error loading aggregate", "error", err)
       }
    }()
}