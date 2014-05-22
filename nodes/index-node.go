package nodes

type IndexNode struct {
	nodes []*Node
	keys  []int64
	prop  *NodeP
}

func (in *IndexNode) Insert(t *Tuple) Node {
	//insert
	index := in.findNodeIndex(t.Key)
	node := in.nodes[index]
	n := (*node).Insert(t)
	//if n is not null we need to insert it into ourselves
	if n != nil {
		//insert the node into nodes.

		//if we are passed capacity we need to split
		if !(len(in.keys) < in.prop.GetMaxKeysIn()) {
			rtn := new(IndexNode)

			//if its an even split
			if index > in.prop.GetMinKeysIn()-1 {
				//split 2 3
				rtn.keys = in.keys[in.prop.GetMinKeysIn():]
				in.keys = in.keys[:in.prop.GetMinKeysIn()]

			} else {
				//if its an odd split
				//split 3 2
			}

		}
	}
	return nil
}

func (in *IndexNode) insertNode(n Node) {
	for i := 0; i < len(in.keys); i++ {

	}
}

func (in *IndexNode) findNodeIndex(key int64) int {
	for i := 0; i < len(in.keys); i++ {
		if key < in.keys[i] {
			return i
		}
	}
	return len(in.nodes) - 1
}
