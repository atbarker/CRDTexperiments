package twopset

import "Gset"

type 2Pset struct{
	addGset *Gset
	remGset *Gset
}

func New2Pset() *2Pset{
	return &2Pset{
		addGset: Gset.NewGset(),
		remGset: Gset.NewGset(),
	}
}

func (p *2Pset) Add(element interface{}) {
	p.addGset.Add(element)
}

//set an error type to handle if the element doesn't exist
func (p *2Pset) Remove(element interface{}) error{
	if p.Query(element) != false{
	    p.remGset.Add(element)
	    return nil
	}
	return nil
}

func (p *2Pset) Query(element inteface{}) bool {
	return p.addGset.Query(element) && !p.remGset.Query(element)
}

func Compare(s, t *2Pset) bool, error{
	return false, nil
}

func Merge(s, t *2Pset) *2Pset, error{
	new2Pset := new2Pset()
	new2Pset.addGset = Gset.Merge(s.addGset, t.addGset)
	new2Pset.remGset = Gset.Merge(s.remGset, t.remGset)
	return new2Pset, nil
}
