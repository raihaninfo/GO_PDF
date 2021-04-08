package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeLetter, "")
	w, h := pdf.GetPageSize()
	fmt.Printf("width=%v, height=%v\n", w, h)
	pdf.AddPage()

	// Basic Text stuff
	pdf.MoveTo(0, 0)
	pdf.SetFont("arial", "B", 30)
	_, lineHt := pdf.GetFontSize()
	pdf.SetTextColor(255, 0, 0)
	pdf.Text(0, lineHt, "Hello, World!")
	pdf.MoveTo(0, lineHt*2.0)
	pdf.SetFont("times", "", 18)
	pdf.SetTextColor(100, 100, 100)
	_, lineHt = pdf.GetFontSize()
	pdf.MultiCell(0, lineHt*1.5, "Hear is sum text, Ea voluptate reprehenderit enim esse amet tempor ut magna reprehenderit labore qui duis. Pariatur commodo et culpa non qui deserunt. Mollit cillum id sit deserunt adipisicing in aliquip. Irure mollit sunt et ea duis amet elit laborum do tempor fugiat esse. Non ullamco minim deserunt eiusmod id adipisicing nostrud sit aliquip incididunt quis aute.", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Basic Shapes

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
	//panic(err)
}
