package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Spacer struct {
    widget.BaseWidget
    
    minWidth, minHeight float32
}

func NewSpacer(width, height float32) *Spacer {
    sp := &Spacer{minWidth: width, minHeight: height}
    sp.ExtendBaseWidget(sp)
    
    return sp
}

func (sp *Spacer) CreateRenderer() fyne.WidgetRenderer {
    return newSpacerRenderer(sp)
}

type spacerRenderer struct {
    sp *Spacer
}

func newSpacerRenderer(sp *Spacer) *spacerRenderer {
    return &spacerRenderer{sp: sp}
}

func (sr *spacerRenderer) Destroy() { }

func (sr *spacerRenderer) Layout(size fyne.Size) {
    
}

func (sr *spacerRenderer) MinSize() fyne.Size {
    return fyne.NewSize(sr.sp.minWidth, sr.sp.minHeight)
}

func (sr *spacerRenderer) Objects() []fyne.CanvasObject {
    return nil
}

func (sr *spacerRenderer) Refresh() {
    
}