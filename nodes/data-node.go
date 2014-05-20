package nodes

import "errors"
import "fmt"

type DataNode struct {
	tuples []*Tuple
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

func (dn *DataNode) Insert(t *Tuple) (Node, error) {
	if t.Key <= 0 {
		return nil, errors.New("Invalid Index key " + string(t.Key))
	}
	if len(dn.tuples) < maxKeysDn() {
		//insert
		dn.insertTuple(t)
		return nil, nil
	} else {
		//create a new data node for data to be split into
		rtn := new(DataNode)
		//split the nodes.
		rtn.tuples = dn.tuples[minKeysDn():]
		dn.tuples = dn.tuples[:minKeysDn()]

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
		fmt.Print("splitting data node")
		fmt.Print(dn.tuples)
		fmt.Println(rtn.tuples)

		return rtn, nil
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
