package nodes

import (
	"../logger"
	"errors"
)

type DataNode struct {
	tuples []*Tuple
	prop   *NodeP
}

func (dn *DataNode) MinKey() (int64, error) {
	//nil check
	if dn.tuples != nil {
		//return the tuple
		return dn.tuples[0].Key, nil
	} else {
		return -1, errors.New("No tuples in the data node.")
	}
}

func (dn *DataNode) Insert(t *Tuple) Node {
	if t.Key <= 0 {
		logger.Error().Println("Invalid index %v needs to be > 0", t.Key)
		return nil
	}
	if len(dn.tuples) < dn.prop.GetMaxKeysDn() {
		//insert
		dn.insertTuple(t)
		return nil
	} else {
		//create a new data node for data to be split into
		rtn := new(DataNode)
		rtn.prop = dn.prop
		//split the nodes.
		rtn.tuples = dn.tuples[dn.prop.GetMinKeysDn():]
		dn.tuples = dn.tuples[:dn.prop.GetMinKeysDn()]

		var toInsert *DataNode = nil
		//find which node the tuple should be inserted into
		if t.Key < rtn.tuples[0].Key {
			toInsert = dn
		} else {
			toInsert = rtn
		}
		//insert
		toInsert.insertTuple(t)
		//log
		logger.Debug().Println("splitting data node %v  %v", dn.tuples, rtn.tuples)

		return rtn
	}
}

func (dn *DataNode) insertTuple(t *Tuple) {
	//loop over the slice
	for i := 0; i < len(dn.tuples); i++ {
		//find the index where the new tuple goes (sorted order)
		if t.Key < dn.tuples[i].Key {
			//insert using append
			dn.tuples = append(dn.tuples[:i], append([]*Tuple{t}, dn.tuples[i:]...)...)
			return //we are done
		}
	}
	//if it wasn't found stick it on the end
	dn.tuples = append(dn.tuples, t)
}
