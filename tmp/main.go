package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello Word")
	err := pdf.OutputFileAndClose("p1.pdf")
	if err != nil {
		fmt.Println(err.Error())
	}
	//panic(err)
}
