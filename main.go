package main

import (
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	//use time.Now() to get current time and .Format to format it as HH:MM:SS
	formatted := time.Now().Format("Time: 03:04:05")
	//Set the clock widgets text to the formatted time
	clock.SetText(formatted)
}

func main() {
	//Create a new app instance
	a := app.New()
	//Create a new Window on the app instance
	w := a.NewWindow("Clock")

	//instantiate clock as a Label widget with no text
	clock := widget.NewLabel("")
	//Call updateTime to set the clock
	updateTime(clock)

	//update the window with the clock widget
	w.SetContent(clock)
	//create goroutine to run updateTime once per second in the background
	go func() {
		//Once per second call update time
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	//Set the first window visable
	w.Show()
	
	// create a second window
	w2 := a.NewWindow("Larger")
	//set the content of the second window to a label that says More Content
	w2.SetContent(widget.NewLabel("More content"))
	//Resize the second window to be 100x100px large(this doesn't include the window chrome)
	w2.Resize(fyne.NewSize(100, 100))
	//Set the second window visable
	w2.Show()

	//Run the above defined app instance
	a.Run()
}
