package Twopset

import "CRDTexperiments/Gset"
//import "labix.org/v1/vclock"

type IntMap map[int]interface{}

type Twopset struct{
	addGset         *Gset.Gset
	remGset         *Gset.Gset
	ACK             IntMap
	Deltas          IntMap
	DeltaCounter    int
	Interval        int
}

func Newtwopset() *Twopset{
	return &Twopset{
		addGset: Gset.NewGset(),
		remGset: Gset.NewGset(),
		DeltaCounter: 0,
	}
}

func (p *Twopset) Add(element, contents interface{}) {
	p.addGset.Add(element, contents)
}

func (p *Twopset) Fetch(element interface{}) interface{}{
	return p.addGset.Fetch(element)
}

//set an error type to handle if the element doesn't exist
func (p *Twopset) Remove(element interface{}) error{
	if p.Query(element) != false{
	    p.remGset.Add(element, struct{}{})
	    return nil
	}
	return nil
}

func (p *Twopset) List()  ([]interface{}){
	return p.addGset.List()
}

func (p *Twopset) Query(element interface{}) bool{
	return (p.addGset.Query(element) && !p.remGset.Query(element))
}

func Compare(s, t *Twopset)  (bool, error){
	return false, nil
}

//create deltas
func (p *Twopset) CreateDelta() *Twopset{

}

//join different deltas
func (p *Twopset) JoinDeltas(){

}


//Merge into the local state
func Merge(s, t *Twopset) (*Twopset, error){
	new2Pset := Newtwopset()
	new2Pset.addGset, _ = Gset.Merge(s.addGset, t.addGset)
	new2Pset.remGset, _ = Gset.Merge(s.remGset, t.remGset)
	return new2Pset, nil
}


func (p *Twopset) GarbageCollect(){
	for k, v := p.Deltas{

	}
}
