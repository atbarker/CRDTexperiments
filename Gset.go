package Gset

//map interfaces (key) to structs (value) in our set
type baseSet map[interface{}]struct{}

//all our Gset has to contain is a single set that grows monotonically
type Gset struct {
	baseSet baseSet
}

func NewGset() *Gset {
	return &Gset{baseSet: baseSet{}}
}


func (g *Gset) Add(element interface{}) error{
	g.baseSet[element] = struct{}{}
	return nil
}

func (g *Gset) Query(element interface{}) (bool, error){
	_, isThere := g.baseSet[element]
	return isThere, nil
}

func (g *Gset) List()  ([]interface{}, error){
	elements := make([]interface{}, 0, len(g.baseSet))
	for element := range g.baseSet{
		elements = append(elements, element)
	}
	return elements, nil
}

func (g *Gset) Length() (int, error){
	return len(g.baseSet), nil
}


//check if S.A is a subset of T.A
func Compare(s, t *Gset) error{
	return nil
}

//merge two sets
func Merge(s, t *Gset) (*Gset, error){
	newGset := NewGset()
	for k, v := range s.baseSet{
		newGset.baseSet[k] = v
	}
	for k, v := range t.baseSet{
		newGset.baseSet[k] = v
	}
	return newGset, nil
}
