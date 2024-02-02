package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)



// GUI
type GUI struct {
	App          fyne.App
	RootWin      fyne.Window
}

// Create a new GUI, given a MacroManager ptr
func NewGUI() *GUI {
	gui := &GUI{}

	gui.App = app.New()
	gui.RootWin = gui.App.NewWindow("Images to PDF")
	gui.RootWin.CenterOnScreen()

	return gui
}
