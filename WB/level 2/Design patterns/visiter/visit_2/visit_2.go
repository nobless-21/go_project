/*
Задание: Расширение функциональности системы отчетов
Цели задания
Добавить новый тип отчета: Реализуйте новый тип отчета, например, CSVReport, который будет представлять отчет в формате CSV.
Добавить новую операцию: Реализуйте новую операцию для генерации отчетов, например, ExportToFile, которая будет экспортировать отчеты в файл.
Обновить интерфейс Visitor: Обновите интерфейс Visitor, чтобы он включал методы для нового типа отчета и новой операции.
Обновить класс ReportGenerator: Реализуйте логику для нового типа отчета и новой операции в классе ReportGenerator.
*/



package main

import "fmt"

type Visitor interface {
	VisitPDFReport(*PDFReport)
	VisitExcelReport(*ExcelReport)
	VisitHTMLReport(*HTMLReport)
	VisitCVCReport(*CVCReport)
}

type Report interface {
	Accept(Visitor)
}

type CVCReport struct{
	Content string
}

func (c *CVCReport) Accept(v Visitor){
	v.VisitCVCReport(c)
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

func (r *ReportGenerator) VisitCVCReport(report *CVCReport){
	fmt.Println("Генерация CVC отчета с содержимым:", report.Content)
}

func main() {
	reports := []Report{
		&PDFReport{Content: "Содержимое PDF отчета"},
		&ExcelReport{Content: "Содержимое Excel отчета"},
		&HTMLReport{Content: "Содержимое HTML отчета"},
		&CVCReport{Content: "Содержимое CVC отчета"},
	}

	generator := &ReportGenerator{}

	for _, report := range reports {
		report.Accept(generator)
	}
}