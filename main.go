package main

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	gopdfBarcode "github.com/jung-kurt/gofpdf/contrib/barcode"
	"log"
	"math"
)


func main () {
	var (
		str string = "pippo" // @TODO: accept the string as input
		pdf          *gofpdf.Fpdf
		bcode        barcode.Barcode
		err error
	)
	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)

	bcode, err = code128.Encode(str)
	if err != nil {
		log.Fatal(err)
	}

	key := gopdfBarcode.Register(bcode)
	// @TODO: find a beter way to scale the barcode
	gopdfBarcode.Barcode(pdf, key, 10, 10, math.Sqrt(500*float64(len(str))), 28, false)

	pdf.Text(10, 50, str)

	err = pdf.OutputFileAndClose("test.pdf") //@TODO: accept the  filename as input
	if err != nil {
		log.Fatal(err)
	}
}
