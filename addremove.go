package addremove

import "2Pset"
import "Gset"

type edge struct{
	left interface{}
	right interface{}
}

type AddRemove struct{
	V *2Pset
	E *Gset
}

func NewAddRemove() *AddRemove{
	AR := &AddRemove{
		V: twopset.New2Pset(),
		E: Gset.NewGset(),
	}
	AR.V.Add("left")
	AR.V.Add("right")
	initEdge := &edge{
		left: "left"
		right: "right"
	}
	AR.E.Add("initEdge")
	AR.E.baseSet["initEdge"] = initEdge
	return AR
}

func (a *AddRemove) Lookup (element interface{}) bool{
	if a.V.Query(element){
		return true
	}
	return false
}

func (a *AddRemove) QueryBefore(u, v interface{}) bool{
	if a.V.Query(u) && a.V.Query(v){

	}
	return false
}

func (a *AddRemove) AddEdge(edgename, u, v interface{}){

}

//TODO add something to generate a new unique ID for edges
func (a *AddRemove) AddBetween(u, v, w interface{}) {
	if !a.Lookup(w) && a.QueryBefore(u, w){
		a.V.Add(v)
		a.AddEdge("something", u, v)
		a.AddEdge("something2", v, w)
	}
}

func (a *AddRemove) Remove(v interface{}){
	if a.Lookup(v) && (v != "left" || v != "right"){
		a.V.Remove(v)
	}
}

func (a *AddRemove) RemoveEdge(v interface{}){

}
