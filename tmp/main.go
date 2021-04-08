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
	pdf.MultiCell(0, lineHt*1.0, "Hear is sum text, Ea voluptate reprehenderit enim esse amet tempor ut magna reprehenderit Non ullamco minim deserunt eiusmod id adipisicing nostrud sit aliquip incididunt quis aute.", gofpdf.BorderNone, gofpdf.AlignRight, false)

	// Basic Shapes
	pdf.SetFillColor(0, 255, 0)
	pdf.SetDrawColor(0, 0, 255)
	pdf.Rect(10, 100, 100, 100, "FD")
	pdf.SetFillColor(100, 200, 200)
	pdf.Polygon([]gofpdf.PointType{
		{110, 250},
		{160, 300},
		{110, 350},
		{60, 300},
	}, "F")

	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		panic(err)
	}
	//panic(err)
}
