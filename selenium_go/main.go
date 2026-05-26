package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Hello Go GUI")

	message := widget.NewLabel("Hello from Go + Fyne!")

	button := widget.NewButton("Click me", func() {
		message.SetText("You clicked the button.")
	})

	window.SetContent(container.NewVBox(
		message,
		button,
	))

	window.ShowAndRun()
}