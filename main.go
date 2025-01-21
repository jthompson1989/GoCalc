package main

import (
	"fmt"
)

func main() {
	var calc = Calculator{
		displayText:        "",
		calculationHistory: []Calculation{},
	}
	CalculatorLoop(calc)
}

func CalculatorLoop(calc Calculator) {
	for {
		fmt.Print("Enter Calculation (+ - * /): ")
		var calculation string
		_, err := fmt.Scanln(&calculation)

		if err != nil {
			fmt.Println(err.Error())
		}

		if calculation == "exit" {
			return
		} else if calculation == "history" {
			calc.DisplayHistory()
		} else {
			calc.SimpleCalculation(calculation)
			fmt.Println(calc.displayText)
		}

		fmt.Println("")
	}
}

// import (
// 	"image/color"
// 	"log"
// 	"os"

// 	"gioui.org/app"
// 	"gioui.org/op"
// 	"gioui.org/text"
// 	"gioui.org/widget/material"
// )

// func main() {
// 	go func() {
// 		window := new(app.Window)
// 		err := run(window)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		os.Exit(0)
// 	}()
// 	app.Main()
// }

// func run(window *app.Window) error {
// 	theme := material.NewTheme()
// 	var ops op.Ops
// 	for {
// 		switch e := window.Event().(type) {
// 		case app.DestroyEvent:
// 			return e.Err
// 		case app.FrameEvent:
// 			// This graphics context is used for managing the rendering state.
// 			gtx := app.NewContext(&ops, e)

// 			// Define an large label with an appropriate text:
// 			title := material.H1(theme, "Hello, Gio")

// 			// Change the color of the label.
// 			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
// 			title.Color = maroon

// 			// Change the position of the label.
// 			title.Alignment = text.Middle

// 			// Draw the label to the graphics context.
// 			title.Layout(gtx)

// 			// Pass the drawing operations to the GPU.
// 			e.Frame(gtx.Ops)
// 		}
// 	}
// }
