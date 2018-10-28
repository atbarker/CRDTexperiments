package addremove

import "2Pset"
import "Gset"

type AddRemove struct{
	V *2Pset
	E *Gset
}

func NewAddRemove() *AddRemove{
	return &AddRemove{
		V: 2Pset.New2Pset(),
		E: Gset.NewGset(),
	}
}

func (a *AddRemove) Lookup (element interface{}) bool{
	return false
}

func (a *AddRemove) QueryBefore(u, v interface{}) bool{
	return false
}

func (a *AddRemove) AddBetween(u, v, w interface{}) {

}

func (a *AddRemove) Remove(v interface{}){

}
