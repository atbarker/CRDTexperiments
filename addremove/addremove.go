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
	E *Twopset.Twopset
}

func NewNode(id interface{}) *Node{
	return &Node{
		ID: id,
	}
}

func NewAddRemove() *AddRemove{
	AR := &AddRemove{
		vectorClock: vclock.New(),
		V: Twopset.Newtwopset(),
		E: Twopset.Newtwopset(),
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

func (a *AddRemove) LookupEdge (element interface{}) bool{
	if a.E.Query(element){
		return true
	}
	return false
}


//depth first search will likely be necessary
//with the crazy interface madness this is going to be REALLY slow.
func (a *AddRemove) QueryBefore(u, v interface{}) bool{
	if a.V.Query(u) && a.V.Query(v){

	}
	return false
}

func (a *AddRemove) FetchNode(v interface{}) *Node{
	node := a.V.Fetch(v).(*Node)
	return node
}

//will return all edges in the set that contain a given node
func (a *AddRemove) FetchEdge(u interface{}) interface{}{
	return nil
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

//really only check if the edge exists at all
func (a *AddRemove) RemoveEdge(v interface{}){
	if a.LookupEdge(v){
		a.E.Remove(v)
	}
}
