**Invoice Generator using Go**
**Invoice Generator – Detailed Project Description**
The Invoice Generator is a full-stack web application built using Golang, designed to dynamically generate and download professional invoices in PDF format based on user input. This project showcases the integration of Go’s backend capabilities with a responsive, user-friendly frontend developed using HTML, CSS, and JavaScript, demonstrating how Go can power lightweight, real-world tools without relying on third-party databases or frameworks.

**Key Objectives:**
Create an intuitive web interface for users to input invoice data.

Dynamically handle multiple line items (products/services).

Perform calculations for subtotal, discount, tax, and final amount.

Generate a downloadable invoice in PDF format.

Ensure the entire application runs locally using Go’s net/http server.

**Technologies Used:**
Backend: Go (Golang), net/http, html/template, encoding/json

PDF Generation: gofpdf (open-source PDF generation library in Go)

Frontend: HTML, CSS (custom styling), JavaScript (DOM manipulation, fetch API)

Data Format: JSON (for client-server communication)

Deployment: Local server using Go standard library

**Functionality Breakdown
User Interface:**
Clean, responsive UI using semantic HTML and modern CSS.

Form-based structure where users input:

Customer name

Multiple line items: Description, Quantity, Unit Price

Optional: Discount percentage, Tax rate

JavaScript allows dynamic addition/removal of invoice items.

**Frontend Logic:**
JavaScript handles form validation and DOM manipulation.

On submission, gathers all form data into a structured JSON object.

Sends a POST request to the Go backend using fetch().

**Backend Logic (Golang):**
Accepts and decodes the JSON request.

Computes:

Subtotal by summing product of quantity and unit price.

Discount and tax amounts based on user input.

Final total = Subtotal − Discount + Tax.

Generates a PDF invoice using the gofpdf package:

Displays customer name

Lists all invoice items with calculated totals

Adds summary section: Subtotal, Discount, Tax, Total

Sends the PDF back as a downloadable response to the browser.

**PDF Output:**
Professionally formatted invoice with tables, fonts, and totals.

Automatically downloaded as invoice.pdf when generated.
