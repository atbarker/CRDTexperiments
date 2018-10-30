package optwopset

import "CRDTexperiments/Gset"
import "container/list"

type optwopset struct{
        addGset *Gset.Gset
}

type OpList struct {
        Operation string
        Element   interface{}
        contents  struct{}
}


func New2Pset() *optwopset{
        return &optwopset{
                addGset: Gset.NewGset(),
        }
}

func (p *optwopset) Add(element interface{}) {
        p.addGset.Add(element)
}

//set an error type to handle if the element doesn't exist
func (p *optwopset) Remove(element interface{}) error{
        if p.Query(element) != false{
		delete(p.addGset.BaseSet, element)
	}
	return nil
}

func (p *optwopset) Query(element interface{}) bool {
        return p.addGset.Query(element)
}

func (p *optwopset) ApplyOps(opslist *list.List) error {
        for e := opslist.Front(); e != nil; e = e.Next() {
                oplistElement := e.Value.(*OpList)
                if oplistElement.Operation == "Add" {
                        p.Add(oplistElement.Element)
                        p.addGset.BaseSet[oplistElement.Element] = oplistElement.contents
                }else if oplistElement.Operation == "Remove"{
                        p.Remove(oplistElement.Element)
                }
        }
        return nil
}


