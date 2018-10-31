package addremove

import "CRDTexperiments/Twopset"
import "CRDTexperiments/Gset"
import "labix.org/v1/vclock"


type Node struct{
	ID interface{}
}

//an edge points from the origin to the destination
//so left to right
type Edge struct{
	left *Node
	right *Node
}

type AddRemove struct{
	vectorClock *vclock.VClock
	V *Twopset.Twopset
	E *Gset.Gset
}

func NewNode(id interface{}) *Node{
	return &Node{
		ID: id,
	}
}

func NewAddRemove() *AddRemove{
	AR := &AddRemove{
		V: Twopset.Newtwopset(),
		E: Gset.NewGset(),
	}
	leftSentinel := NewNode("leftSentinel")
	rightSentinel := NewNode("rightSentinel")
	AR.V.Add("leftSentinel", leftSentinel)
	AR.V.Add("rightSentinel", rightSentinel)
	AR.AddEdge("sentinelEdge", "leftSentinel", "rightSentinel")
	return AR
}

func (a *AddRemove) Lookup (element interface{}) bool{
	if a.V.Query(element){
		return true
	}
	return false
}


//depth first search will likely be necessary
func (a *AddRemove) QueryBefore(u, v interface{}) bool{
	if a.V.Query(u) && a.V.Query(v){

	}
	return false
}

func (a *AddRemove) FetchNode(v interface{}) *Node{
	node := a.V.Fetch(v).(*Node)
	return node
}

func (a *AddRemove) AddEdge(edgename, u, v interface{}){
	if a.V.Query(u) && a.V.Query(v){
		newEdge := &Edge{
			left: a.FetchNode(u),
			right: a.FetchNode(v),
		}
		a.E.Add(edgename, newEdge)
	}
}

//TODO add something to generate a new unique ID for edges
func (a *AddRemove) AddBetween(u, v, w interface{}) {
	if !a.Lookup(w) && a.QueryBefore(u, w){
		a.V.Add(w, NewNode(w))
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
