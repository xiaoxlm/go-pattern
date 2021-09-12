package builder

import (
	"go-pattern/utils"
	"testing"
)

func Test(t *testing.T) {
	devNode := NewDevNode("dev", []string{"评审需求", "研发", "单元测试通过"})
	productNode := NewProductNode("product", []string{"写需求", "验证通过"})

	applyWorkOrderBuilder := new(ApplyWorkOrderBuilder)

	director := Director{builder: applyWorkOrderBuilder}
	model := director.Generate("描述", []*Milestone{
		{
			Node:  devNode,
			Order: 1,
		},
		{
			Node:  productNode,
			Order: 2,
		},
	})

	utils.LogJSON(model)
}

// model
type NodeModel struct {
	Name string
	Jobs []string
}

type MilestoneModel struct {
	Node *NodeModel
	Order int
}

type WorkOrderModel struct {
	Milestones []*MilestoneModel
	Content string
	Category string
}

// logic
type IWorkOrderBuilder interface {
	BuildMilestone(milestone *Milestone) error
	BuildContent(content string) error
	GetResult() *WorkOrderModel
}

type Milestone struct {
	Node INode `json:"node"`
	Order int `json:"order"`
}

type INode interface {
	GetName() string
	GetJobs() []string
}

type Director struct {
	builder IWorkOrderBuilder
}

func (d *Director) SetBuilder(builder IWorkOrderBuilder) {
	d.builder = builder
}

func (d *Director) Generate(content string, milestones []*Milestone) *WorkOrderModel {
	for _, m := range milestones {
		_ = d.builder.BuildMilestone(m)
	}

	_ = d.builder.BuildContent(content)

	return d.builder.GetResult()
}

type ApplyWorkOrderBuilder struct {
	milestones []*Milestone
	content string
}

func (a *ApplyWorkOrderBuilder) BuildMilestone(milestone *Milestone) error {
	a.milestones = append(a.milestones, milestone)
	return nil
}

func (a *ApplyWorkOrderBuilder) BuildContent(content string) error {
	a.content = content
	return nil
}

func (a *ApplyWorkOrderBuilder) GetResult() *WorkOrderModel {
	model := new(WorkOrderModel)
	model.Category = "apply"
	model.Content = a.content
	model.Milestones = make([]*MilestoneModel, 0)

	for _, m := range a.milestones {
		milestoneModel := new(MilestoneModel)
		milestoneModel.Order = m.Order
		milestoneModel.Node = &NodeModel{
			Name: m.Node.GetName(),
			Jobs: m.Node.GetJobs(),
		}

		model.Milestones = append(model.Milestones, milestoneModel)
	}

	return model
}


type DevNode struct {
	name string
	jobs []string
}

func NewDevNode(name string, jobs []string) *DevNode {
	return &DevNode{
		name: name,
		jobs: jobs,
	}
}

func (d *DevNode) GetName() string {
	return d.name
}

func (d *DevNode) GetJobs() []string {
	return d.jobs
}

type ProductNode struct {
	name string
	jobs []string
}

func NewProductNode(name string, jobs []string) *ProductNode {
	return &ProductNode{
		name: name,
		jobs: jobs,
	}
}

func (d *ProductNode) GetName() string {
	return d.name
}

func (d *ProductNode) GetJobs() []string {
	return d.jobs
}