package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

type Item struct {
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Invoice struct {
	CustomerName string  `json:"customerName"`
	Items        []Item  `json:"items"`
	Discount     float64 `json:"discount"`
	TaxRate      float64 `json:"taxRate"`
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generate", generateInvoiceHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func generateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var invoice Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Calculate totals in INR
	subtotalINR := 0.0
	for _, item := range invoice.Items {
		subtotalINR += float64(item.Quantity) * item.UnitPrice
	}
	discountAmountINR := subtotalINR * (invoice.Discount / 100)
	taxAmountINR := (subtotalINR - discountAmountINR) * (invoice.TaxRate / 100)
	totalINR := subtotalINR - discountAmountINR + taxAmountINR

	// Generate PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Invoice")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Customer: "+invoice.CustomerName)
	pdf.Ln(10)

	// Table header
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(80, 10, "Description")
	pdf.Cell(30, 10, "Quantity")
	pdf.Cell(30, 10, "Unit Price")
	pdf.Cell(30, 10, "Total")
	pdf.Ln(10)

	// Table content
	pdf.SetFont("Arial", "", 12)
	for _, item := range invoice.Items {
		itemTotalINR := float64(item.Quantity) * item.UnitPrice
		pdf.Cell(80, 10, item.Description)
		pdf.Cell(30, 10, strconv.Itoa(item.Quantity))
		pdf.Cell(30, 10, strconv.FormatFloat(item.UnitPrice, 'f', 2, 64))
		pdf.Cell(30, 10, strconv.FormatFloat(itemTotalINR, 'f', 2, 64))
		pdf.Ln(10)
	}

	pdf.Ln(10)
	pdf.Cell(140, 10, "Subtotal: "+strconv.FormatFloat(subtotalINR, 'f', 2, 64))
	pdf.Ln(10)
	pdf.Cell(140, 10, "Discount ("+strconv.FormatFloat(invoice.Discount, 'f', 0, 64)+"%) : "+strconv.FormatFloat(discountAmountINR, 'f', 2, 64))
	pdf.Ln(10)
	pdf.Cell(140, 10, "Tax ("+strconv.FormatFloat(invoice.TaxRate, 'f', 0, 64)+"%) : "+strconv.FormatFloat(taxAmountINR, 'f', 2, 64))
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(140, 10, "Total: "+strconv.FormatFloat(totalINR, 'f', 2, 64))

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=invoice.pdf")
	pdf.Output(w)
}
