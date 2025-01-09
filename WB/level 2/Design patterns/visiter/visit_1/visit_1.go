package main

import "fmt"

type Visitor interface {
	VisitPDFReport(*PDFReport)
	VisitExcelReport(*ExcelReport)
	VisitHTMLReport(*HTMLReport)
}

type Report interface {
	Accept(Visitor)
}

type PDFReport struct {
	Content string
}

func (p *PDFReport) Accept(v Visitor) {
	v.VisitPDFReport(p)
}

type ExcelReport struct {
	Content string
}

func (e *ExcelReport) Accept(v Visitor) {
	v.VisitExcelReport(e)
}

type HTMLReport struct {
	Content string
}

func (h *HTMLReport) Accept(v Visitor) {
	v.VisitHTMLReport(h)
}

type ReportGenerator struct{}

func (r *ReportGenerator) VisitExcelReport(report *ExcelReport) {
	fmt.Println("Генерация Excel отчета с содержимым:", report.Content)
}

func (r *ReportGenerator) VisitHTMLReport(report *HTMLReport) {
	fmt.Println("Генерация HTML отчета с содержимым:", report.Content)
}

func (r *ReportGenerator) VisitPDFReport(report *PDFReport){
	fmt.Println("Генерация PDF отчета с содержимым:", report.Content)
}

func main() {
	reports := []Report{
		&PDFReport{Content: "Содержимое PDF отчета"},
		&ExcelReport{Content: "Содержимое Excel отчета"},
		&HTMLReport{Content: "Содержимое HTML отчета"},
	}

	generator := &ReportGenerator{}

	for _, report := range reports {
		report.Accept(generator)
	}
}