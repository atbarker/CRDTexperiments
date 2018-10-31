package Gset

import "container/list"
//import "labix.org/v1/vclock"

//map interfaces (key) to structs (value) in our set
type baseSet map[interface{}]interface{}

//all our Gset has to contain is a single set that grows monotonically
type Gset struct {
	BaseSet baseSet
}


//used to contain operations that are then sent to another 
type OpList struct {
	Operation string
	Element   interface{}
	contents  struct{}
}

func NewGset() *Gset {
	return &Gset{BaseSet: baseSet{}}
}


func (g *Gset) Add(element, contents interface{}) error{
	g.BaseSet[element] = contents
	return nil
}

func (g *Gset) Fetch(element interface{}) interface{}{
	contents := g.BaseSet[element]
	return contents
}

func (g *Gset) Query(element interface{}) bool{
	_, isThere := g.BaseSet[element]
	return isThere
}

func (g *Gset) List()  ([]interface{}, error){
	elements := make([]interface{}, 0, len(g.BaseSet))
	for element := range g.BaseSet{
		elements = append(elements, element)
	}
	return elements, nil
}

func (g *Gset) Length() (int, error){
	return len(g.BaseSet), nil
}

func (g *Gset) ApplyOps(opslist *list.List) error {
	for e := opslist.Front(); e != nil; e = e.Next() {
		oplistElement := e.Value.(*OpList)
		if oplistElement.Operation == "Add" {
			g.Add(oplistElement.Element, oplistElement.contents)
			g.BaseSet[oplistElement.Element] = oplistElement.contents
		}else{
			return nil
		}
	}
	return nil
}


//check if S.A is a subset of T.A
func Compare(s, t *Gset) error{
	return nil
}

//merge two sets
func Merge(s, t *Gset) (*Gset, error){
	newGset := NewGset()
	for k, v := range s.BaseSet{
		newGset.BaseSet[k] = v
	}
	for k, v := range t.BaseSet{
		newGset.BaseSet[k] = v
	}
	return newGset, nil
}
