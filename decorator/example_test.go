package decorator

import (
	"go-pattern/utils"
	"testing"
)

func Test(t *testing.T) {
	p := DecoratePeople(NewPeople)()
	utils.LogJSON(p)
}

type People struct {
	Age int
	Name string
	Desc string
}

func NewPeople() *People {
	return &People{}
}

type NewPeopleFNC func() *People

func DecoratePeople(fnc NewPeopleFNC) NewPeopleFNC {
	return func() *People {
		p := fnc()
		p.Desc = "from decorate"
		return p
	}
}
