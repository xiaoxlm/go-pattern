package builder

import (
	"go-pattern/utils"
	"testing"
)

func Test(t *testing.T) {
	devNode := NewDevNode("dev", []string{"评审需求", "研发", "单元测试通过"})
	testNode := NewTestNode("test", []string{"测试通过"})
	productNode := NewProductNode("product", []string{"验证通过, 发布"})

	applyWorkOrderBuilder := new(ApplyWorkOrderBuilder)

	director := Director{builder: applyWorkOrderBuilder}
	model := director.Generate("描述", []*Milestone{
		{
			Node:  devNode,
			Order: 1,
		},
		{
			Node:  testNode,
			Order: 2,
		},
		{
			Node:  productNode,
			Order: 3,
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
	SetMilestone(milestone *Milestone) error
	SetContent(content string) error
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
		_ = d.builder.SetMilestone(m)
	}

	_ = d.builder.SetContent(content)

	return d.builder.GetResult()
}

type ApplyWorkOrderBuilder struct {
	milestones []*Milestone
	content string
}

func (a *ApplyWorkOrderBuilder) SetMilestone(milestone *Milestone) error {
	a.milestones = append(a.milestones, milestone)
	return nil
}

func (a *ApplyWorkOrderBuilder) SetContent(content string) error {
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

type TestNode struct {
	name string
	jobs []string
}

func NewTestNode(name string, jobs []string) *TestNode {
	return &TestNode{
		name: name,
		jobs: jobs,
	}
}

func (d *TestNode) GetName() string {
	return d.name
}

func (d *TestNode) GetJobs() []string {
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