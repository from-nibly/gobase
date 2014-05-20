package nodes

type IndexNode struct {
	nodes []*Node
	keys  []int64
}

func (in *IndexNode) Insert(t *Tuple) (*Node, error) {
	if len(in.keys) < maxKeysIn() {
		//just insert

	} else {
		//split

	}
	return nil, nil
}
