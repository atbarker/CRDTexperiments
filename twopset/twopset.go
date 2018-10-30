package twopset

import "CRDTexperiments/Gset"

type twopset struct{
	addGset *Gset.Gset
	remGset *Gset.Gset
}

func Newtwopset() *twopset{
	return &twopset{
		addGset: Gset.NewGset(),
		remGset: Gset.NewGset(),
	}
}

func (p *twopset) Add(element interface{}) {
	p.addGset.Add(element)
}

//set an error type to handle if the element doesn't exist
func (p *twopset) Remove(element interface{}) error{
	if p.Query(element) != false{
	    p.remGset.Add(element)
	    return nil
	}
	return nil
}

func (p *twopset) Query(element interface{}) bool{
	return (p.addGset.Query(element) && !p.remGset.Query(element))
}

func Compare(s, t *twopset)  (bool, error){
	return false, nil
}

func Merge(s, t *twopset) (*twopset, error){
	new2Pset := Newtwopset()
	new2Pset.addGset, _ = Gset.Merge(s.addGset, t.addGset)
	new2Pset.remGset, _ = Gset.Merge(s.remGset, t.remGset)
	return new2Pset, nil
}
