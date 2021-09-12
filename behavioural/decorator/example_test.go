package decorator

import (
	"go-pattern/utils"
	"testing"
)

func Test(t *testing.T) {
	t.Run("#NewPeople", func(t *testing.T) {
		p := NewPeople(10, "test")
		utils.LogJSON(p)
	})

	t.Run("#NewPeopleWithDesc", func(t *testing.T) {
		p := NewPeopleWithDesc(NewPeople)(10, "test", "add desc")
		utils.LogJSON(p)
	})
}

type People struct {
	Age int
	Name string

	Desc string
}

func NewPeople(age int, name string) *People {
	return &People{
		Age:  age,
		Name: name,
	}
}

// decorate NewPeople

type NewPeopleFNC func(age int, name string) *People
type NewPeopleFNCWithDesc func(age int, name, desc string) *People

func NewPeopleWithDesc(fnc NewPeopleFNC) NewPeopleFNCWithDesc {
	return func(age int, name, desc string) *People {
		p := fnc(age, name)
		p.Desc = desc
		return p
	}
}
