package addremove

import "CRDTexperiments/Twopset"
import "labix.org/v1/vclock"
import "container/list"


type Node struct{
	ID interface{}
}

//here we represent the element being added as an array
//0 is the element to add or remove (v)
//1 is the first element in an addbetween (u)
//2 is the last element in an addbetween (w)
type OpList struct {
	Operation string
	Element   []interface{}
	contents  struct{}
}

//an edge points from the origin to the destination
//so left to right
type Edge struct{
	left *Node
	right *Node
}

type AddRemove struct{
	vectorClock *vclock.VClock
	externalVectorClocks []vclock.VClock
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
	AR.AddEdge("leftSentinel", "rightSentinel")
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
	isBefore := false
	if a.V.Query(u) && a.V.Query(v){
		if edgeExists := a.FetchEdge(u, v); edgeExists!= nil {
			return true
		}
		if u.(*Node).ID == "leftSentinel" && v.(*Node).ID == "rightSentinel" {
			return true
		}
		edges := a.GetEdges(u)
		for k := range edges{
			isBefore = a.QueryBeforeRecurse(edges[k].(*Edge).right, v)
		}
	}
	return isBefore
}

func (a *AddRemove) QueryBeforeRecurse(u, v interface{}) bool{
	edges := a.GetEdges(u)
	isBefore := false
	//if we have hit the sentinel then we are done
	if len(edges) == 1 && edges[0].(*Edge).right.ID == "rightSentinel" {
		isBefore = false
	}
	for k := range edges{
		if edgeExists := a.FetchEdge(u, v); edgeExists != nil {
			isBefore = true
		}else{
			isBefore = a.QueryBeforeRecurse(edges[k].(*Edge).right, v)
		}
	}
	return isBefore
}

func (a *AddRemove) FetchNode(v interface{}) *Node{
	node := a.V.Fetch(v).(*Node)
	return node
}

//will return all edges in the set that contain a given node
func (a *AddRemove) FetchEdge(u interface{}, v interface{}) *Edge{
	edgeList := a.E.List()
	for k := range edgeList {
		if edgeList[k].(*Edge).left == u.(*Node) && edgeList[k].(*Edge).right == v.(*Node) {
			return edgeList[k].(*Edge)
		}
	}
	return nil
}

func (a *AddRemove) GetEdges(u interface{}) []interface{}{
	edges := a.E.List()
	returnEdges := make([]interface{}, 0)
	for k := range edges{
		if edges[k].(*Edge).left == u.(*Node){
			returnEdges = append(returnEdges, edges[k])
		}
	}
	return returnEdges

}

func (a *AddRemove) AddEdge(u, v interface{}){
	if a.V.Query(u) && a.V.Query(v){
		newEdge := &Edge{
			left: a.FetchNode(u),
			right: a.FetchNode(v),
		}
		a.E.Add(newEdge, nil)
	}
}

func (a *AddRemove) AddBetween(u, v, w interface{}) {
	if !a.Lookup(v) && a.QueryBefore(u, w){
		a.V.Add(v, NewNode(v))
		a.AddEdge(u, v)
		a.AddEdge(v, w)
	}
}

//needs a check to remove dangling edges, could possibly be done during garbage collection
func (a *AddRemove) Remove(v interface{}){
	if a.Lookup(v) && (v != "left" || v != "right"){
		a.V.Remove(v)
	}
}

func (a *AddRemove) RemoveEdge(v interface{}){
	if a.LookupEdge(v){
		a.E.Remove(v)
	}
}

func (a *AddRemove) ApplyOps(opslist *list.List) error{
	for e := opslist.Front(); e != nil; e = e.Next(){
		oplistElement := e.Value.(*OpList)
		if oplistElement.Operation == "AddBetween"{
			a.AddBetween(oplistElement.Element[1], oplistElement.Element[0], oplistElement.Element[2])
		}else if oplistElement.Operation == "Remove"{
			a.Remove(oplistElement.Element[0])
		}else{
			return nil
		}
	}
	return nil
}
