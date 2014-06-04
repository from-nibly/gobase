package nodes

import (
	"../logger"
)

type IndexNode struct {
	Nodes []Node
	Keys  []int64
	prop  *NodeP
}

func (this *IndexNode) SetProp(prop *NodeP) {
	this.prop = prop
}

func (this *IndexNode) Insert(t *Tuple) Node {
	//insert
	index := this.findNodeIndex(t.Key)
	node := this.Nodes[index]
	n := node.Insert(t)

	//if n is not null we need to insert it into ourselves
	if n != nil {
		//insert the node into nodes.
		this.Nodes = append(this.Nodes[:index+1], append([]Node{n}, this.Nodes[index+1:]...)...)
		this.Keys = append(this.Keys[:index], append([]int64{n.MinKey()}, this.Keys[index:]...)...)

		//if we are passed capacity we need to split
		logger.Debug().Println("this keys ", this.Keys)

		if len(this.Keys) > this.prop.GetMaxKeysIn() {
			rtn := new(IndexNode)
			rtn.prop = this.prop

			//split keys 2 3
			rtn.Keys = this.Keys[this.prop.GetMinKeysIn():]
			this.Keys = this.Keys[:this.prop.GetMinKeysIn()]
			//split nodes 3 3
			rtn.Nodes = this.Nodes[this.prop.GetMinKeysIn()+1:]
			this.Nodes = this.Nodes[:this.prop.GetMinKeysIn()+1]
			//remove useless key
			rtn.Keys = rtn.Keys[1:]
			logger.Trace().Println("this split keys ", this.Keys)
			logger.Trace().Println("rtn keys ", rtn.Keys)
			logger.Trace().Println("STEP-5: ", this, rtn)
			return rtn
		}
	}
	return nil
}

func (this *IndexNode) insertNode(n Node) {
	for i := 0; i < len(this.Keys); i++ {

	}
}

func (this *IndexNode) findNodeIndex(key int64) int {
	for i := 0; i < len(this.Keys); i++ {
		if key < this.Keys[i] {
			return i
		}
	}
	return len(this.Nodes) - 1
}

func (this *IndexNode) MinKey() int64 {
	logger.Trace().Println("checking nodes", this.Nodes)
	return this.Nodes[0].MinKey()
}

func (this *IndexNode) Dump(level int) {
	log := ""
	for i := 0; i < level; i++ {
		log += "\t"
	}

	for i := 0; i < len(this.Keys); i++ {
		log += string(this.Keys[i]) + " "
	}
	logger.Debug().Println(log)

	for i := 0; i < len(this.Nodes); i++ {
		this.Nodes[i].Dump(level + 1)
	}
}

func (this *IndexNode) Find(key int64) *Tuple {
	for i := 0; i < len(this.Keys); i++ {
		if this.Keys[i] > key {
			return this.Nodes[i].Find(key)
		}
	}
	return this.Nodes[len(this.Nodes)-1].Find(key)
}
