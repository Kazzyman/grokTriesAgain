package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

// @formatter:off

type ColoredButton struct {
	widget.Button
	BackgroundColor color.Color
}
func NewColoredButton(label string, backgroundColor color.Color, tapped func()) *ColoredButton {
	btn := &ColoredButton{BackgroundColor: backgroundColor}
	btn.Text = label
	btn.OnTapped = tapped
	btn.ExtendBaseWidget(btn)
	return btn
}


// Custom renderer methods for coloredButton
func (b *ColoredButton) CreateRenderer() fyne.WidgetRenderer {
	text := widget.NewLabel(b.Text)
	text.Wrapping = fyne.TextWrapWord // Enable word wrapping
	background := canvas.NewRectangle(b.BackgroundColor)
	border := canvas.NewRectangle(color.Transparent)
	border.StrokeColor = color.Gray{0x60} // was 80 // space between buttons ? 
	border.StrokeWidth = 3 // was 2
	return &coloredButtonRenderer{
		button:     b,
		text:       text,
		background: background,
		border:     border,
		objects:    []fyne.CanvasObject{background, border, text},
	}
}


type coloredButtonRenderer struct {
	button     *ColoredButton
	text       *widget.Label // Change to *widget.Label
	background *canvas.Rectangle
	border     *canvas.Rectangle
	objects    []fyne.CanvasObject
}


func (r *coloredButtonRenderer) Refresh() {
	bgColor := r.button.BackgroundColor
	if r.button.Disabled() {
		if rgba, ok := bgColor.(color.RGBA); ok {
			rgba.A = 145 // ::: Lower opacity for disabled state
			r.background.FillColor = rgba
		}
	} else {
		r.background.FillColor = bgColor
	}
	r.text.SetText(r.button.Text)
	r.background.Refresh()
	r.border.Refresh()
	r.text.Refresh()
}
/* old non-dimming refresh 
func (r *coloredButtonRenderer) Refresh() {
	r.background.FillColor = r.button.BackgroundColor
	r.text.SetText(r.button.Text) // Use SetText
	r.background.Refresh()
	r.border.Refresh()
	r.text.Refresh()
}
 */

func (r *coloredButtonRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.border.Resize(size)
	r.text.Resize(fyne.NewSize(size.Width-20, size.Height-20)) // were -20 -20   ::: ineffective   
	r.text.Move(fyne.NewPos(15, 10)) // 10  10  ::: x is left margin, y is top margin;  10  1000 moves text down off the buttons
}
func (r *coloredButtonRenderer) MinSize() fyne.Size {
	textSize := r.text.MinSize()
	return fyne.NewSize(fyne.Max(textSize.Width+20, 200), fyne.Max(textSize.Height+20, 50)) // +20,200  +20,50 ::: ineffective 
}
func (r *coloredButtonRenderer) BackgroundColor() color.Color {
	return r.button.BackgroundColor
}
func (r *coloredButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}
func (r *coloredButtonRenderer) Destroy() {
	// No-op
}
