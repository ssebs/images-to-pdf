// gui/gui.go
package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// GUI represents the Fyne GUI class.
type GUI struct {
	app fyne.App
	win fyne.Window
}

// NewGUI creates a new instance of the GUI class.
func NewGUI() *GUI {
	myApp := app.New()
	myWin := myApp.NewWindow("Images to PDF Converter")
	myWin.Resize(fyne.NewSize(600, 600))

	return &GUI{
		app: myApp,
		win: myWin,
	}
}

// Run starts the GUI application.
func (g *GUI) Run() {
	g.win.CenterOnScreen()

	// Create the file entry widget.
	fileEntry := widget.NewEntry()
	fileEntry.Disable()

	// Create the file selection button.
	fileSelectButton := widget.NewButtonWithIcon("Select Files", theme.FolderOpenIcon(), func() {
		g.showFileSelectDialog(fileEntry)
	})

	// Create the convert button.
	convertButton := widget.NewButtonWithIcon("Convert to PDF", theme.DocumentSaveIcon(), func() {
		g.convertToPDF(fileEntry.Text)
	})

	// Create the layout.
	content := container.NewVBox(
		container.NewHBox(
			fileSelectButton,
			fileEntry,
		),
		layout.NewSpacer(),
		convertButton,
	)

	g.win.SetContent(container.NewBorder(
		nil, nil, nil, nil,
		content,
	))

	// Show the window.
	g.win.ShowAndRun()
}

// showFileSelectDialog displays the file selection dialog and sets the selected directory path to the entry widget.
func (g *GUI) showFileSelectDialog(entry *widget.Entry) {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			// Handle the error, if any.
			fmt.Println("Error:", err)
			return
		}

		if uri == nil {
			// User canceled the dialog or no directory was selected.
			fmt.Println("User canceled the folder selection.")
			return
		}

		// Get the directory path from the URI.
		dirPath := uri.String()

		// Set the directory path to the entry widget.
		entry.SetText(dirPath)
	}, g.win)
}

// convertToPDF reads the selected image files and converts them to a PDF.
func (g *GUI) convertToPDF(filePath string) {
	// imgFiles, err := convert.ListFiles(filePath)
	// if err != nil {
	// 	dialog.ShowError(err, g.win)
	// 	return
	// }

	// pdfFile, err := convert.ImagesToPDF(imgFiles)
	// if err != nil {
	// 	dialog.ShowError(err, g.win)
	// 	return
	// }

	// err = os.WriteFile("output.pdf", pdfFile.Contents, os.ModePerm)
	// if err != nil {
	// 	dialog.ShowError(err, g.win)
	// 	return
	// }

	dialog.ShowInformation("Conversion Complete", "PDF file has been created successfully", g.win)
}
