package nodes

import (
	"../logger"
)

type DataNode struct {
	tuples []*Tuple
	prop   *NodeP
}

func (this *DataNode) SetProp(prop *NodeP) {
	this.prop = prop
}

func (this *DataNode) MinKey() int64 {
	//nil check
	if this.tuples != nil {
		//return the tuple
		return this.tuples[0].Key
	} else {
		panic("There are no tuples in the data node")
	}
}

func (this *DataNode) Insert(t *Tuple) Node {
	if t.Key < 0 {
		logger.Error().Println("Invalid index %v needs to be >= 0", t.Key)
		return nil
	}
	if len(this.tuples) < this.prop.GetMaxKeysDn() {
		//insert
		this.insertTuple(t)
		return nil
	} else {
		//create a new data node for data to be split into
		rtn := new(DataNode)
		rtn.prop = this.prop
		//split the nodes.
		rtn.tuples = this.tuples[this.prop.GetMinKeysDn():]
		this.tuples = this.tuples[:this.prop.GetMinKeysDn()]

		var toInsert *DataNode = nil
		//find which node the tuple should be inserted into
		if t.Key < rtn.tuples[0].Key {
			toInsert = this
		} else {
			toInsert = rtn
		}
		//insert
		toInsert.insertTuple(t)
		//log
		logger.Debug().Println("splitting data node", this.tuples, rtn.tuples)

		return rtn
	}
}

func (this *DataNode) insertTuple(t *Tuple) {
	//loop over the slice
	for i := 0; i < len(this.tuples); i++ {
		//find the index where the new tuple goes (sorted order)
		if t.Key < this.tuples[i].Key {
			//insert using append
			this.tuples = append(this.tuples[:i], append([]*Tuple{t}, this.tuples[i:]...)...)
			return //we are done
		}
	}
	//if it wasn't found stick it on the end
	this.tuples = append(this.tuples, t)
}

func (this *DataNode) Dump(level int) {
	log := ""
	for i := 0; i < level; i++ {
		log += "\t"
	}

	for i := 0; i < len(this.tuples); i++ {
		log += string(this.tuples[i].Key) + " "
	}
	logger.Debug().Println(log)
}

func (this *DataNode) Find(key int64) *Tuple {
	for i := 0; i < len(this.tuples); i++ {
		if this.tuples[i].Key == key {
			return this.tuples[i]
		}
	}
	return nil
}
