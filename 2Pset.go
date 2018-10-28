package 2Pset

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

func (p *2Pset) Remove(element interface{}) {
	p.remGset.Add(element)
}

func (p *2Pset) Query(element inteface{}) bool {
	return p.addGset.Query(element) && !p.remGset.Query(element)
}

func Compare(s, t *2Pset) bool, error{
	return false, nil
}

func Merge(s, t *2Pset) *2Pset, error{
	return nil, nil
}
