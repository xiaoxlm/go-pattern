package observer

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	subject := new(Subject)

	subject.AttachObserver(new(NameObserver), new(DescObserver))

	subject.Name = "name1"
	subject.Desc = "desc1"
	subject.Notify()

	subject.Name = "name2"
	subject.Desc = "desc2"
	subject.Notify()
}

type Subject struct {
	observer []IObserver
	Name     string
	Desc     string
}

func (s *Subject) AttachObserver(o ...IObserver) {
	s.observer = append(s.observer, o...)
}

func (s *Subject) Notify() {
	for _, o := range s.observer {
		o.update(s)
	}
}

type IObserver interface {
	update(*Subject)
}

type NameObserver struct{}

func (*NameObserver) update(s *Subject) {
	fmt.Println("changed name:", s.Name)
}

type DescObserver struct{}

func (*DescObserver) update(s *Subject) {
	fmt.Println("changed desc:", s.Desc)
}
