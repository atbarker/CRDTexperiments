package Gset

//map interfaces (key) to structs (value) in our set
type baseSet map[interface{}]struct{}

//all our Gset has to contain is a single set that grows monotonically
type Gset struct {
	baseSet baseSet
}

func newGset() *Gset {
	return &Gset{baseSet: baseSet{}}
}


func (g *Gset) Add(element interface{}) error{
	g.baseSet[element] = struct{}{}
	return nil
}

func (g *Gset) Query(element interface{}) (bool, error){
	_, isThere := g.mainSet
	return isThere, nil
}

func (g *Gset) List()  ([]interface{}, error){
	elements := make([]interface{}, 0, len(g.baseSet))
	for element := range g.baseSet{
		elements = append(elements, element)
	}
	return elements
}

func (g *Gset) Length() (int, error){
	return len(g.baseSet), nil
}

func Compare() error{
	return nil
}

func Merge() (*Gset, error){
	return nil, nil
}
