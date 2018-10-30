package Twopset

import "CRDTexperiments/Gset"

type Twopset struct{
	addGset *Gset.Gset
	remGset *Gset.Gset
}

func Newtwopset() *Twopset{
	return &Twopset{
		addGset: Gset.NewGset(),
		remGset: Gset.NewGset(),
	}
}

func (p *Twopset) Add(element interface{}) {
	p.addGset.Add(element)
}

//set an error type to handle if the element doesn't exist
func (p *Twopset) Remove(element interface{}) error{
	if p.Query(element) != false{
	    p.remGset.Add(element)
	    return nil
	}
	return nil
}

func (p *Twopset) Query(element interface{}) bool{
	return (p.addGset.Query(element) && !p.remGset.Query(element))
}

func Compare(s, t *Twopset)  (bool, error){
	return false, nil
}

func Merge(s, t *Twopset) (*Twopset, error){
	new2Pset := Newtwopset()
	new2Pset.addGset, _ = Gset.Merge(s.addGset, t.addGset)
	new2Pset.remGset, _ = Gset.Merge(s.remGset, t.remGset)
	return new2Pset, nil
}