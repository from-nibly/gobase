package nodes

import "testing"

func getSettings() *NodeP {
	prop := new(NodeP)
	prop.SetMaxKeysDn(4)
	prop.SetMaxKeysIn(4)
	return prop
}

func TestInsertOne_MinKeysIsSame(t *testing.T) {
	d := new(DataNode)
	d.prop = getSettings()
	tu := new(Tuple)
	tu.Key = 5
	d.Insert(tu)

	key := d.MinKey()
	if key != 5 {
		t.FailNow()
	}
}

func TestInsertOutOfOrder_MinKeyIsRight(t *testing.T) {
	d := new(DataNode)
	d.prop = getSettings()
	tu := new(Tuple)
	tu.Key = 5
	d.Insert(tu)
	tu = new(Tuple)
	tu.Key = 3
	d.Insert(tu)

	key := d.MinKey()
	if key != 3 {
		t.FailNow()
	}
}
