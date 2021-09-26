package visitor

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	report := GenerateReport()
	report.Show(new(ReportVisitor))
}

type Application struct {
	name string
}

func (a *Application) GetName() string {
	return a.name
}

func (a *Application) Accept(v IVisitor) {
	v.visitApplication(a)
}

type Apply struct {
	month string
	count int
}

func (a *Apply) GetMonth() string {
	return a.month
}

func (a *Apply) GetCount() int {
	return a.count
}

func (a *Apply) Accept(v IVisitor) {
	v.visitApply(a)
}

type IVisitor interface {
	visitApplication(a *Application)
	visitApply(a *Apply)
}

type ReportVisitor struct {}

func (r *ReportVisitor) visitApplication(a *Application) {
	fmt.Printf("应用:%s\n", a.GetName())
}

func (r *ReportVisitor) visitApply(a *Apply) {
	fmt.Printf("%s月发布:%d次\n", a.GetMonth(), a.GetCount())
}

func GenerateReport() *ApplicationApplyReport {
	return NewApplicationApplyReport([]*ApplicationApply{
		{
			app:   &Application{
				name:    "监控",
			},
			apply: &Apply{
				month: "8",
				count: 3,
			},
		},
		{
			app:   &Application{
				name:    "商城",
			},
			apply: &Apply{
				month: "8",
				count: 5,
			},
		},
	})
}

type ApplicationApply struct {
	app *Application
	apply *Apply
}

type ApplicationApplyReport struct {
	statistics []*ApplicationApply
}

func NewApplicationApplyReport(statistics []*ApplicationApply) *ApplicationApplyReport {
	return &ApplicationApplyReport{statistics: statistics}
}

func (aa *ApplicationApplyReport) Show(visitor IVisitor) {
	for _, s := range aa.statistics {
		s.app.Accept(visitor)
		s.apply.Accept(visitor)
		fmt.Println("========")
	}
}

