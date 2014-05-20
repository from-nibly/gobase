package nodes

import "testing"

func TestInsertOne_MinKeysIsSame(t *testing.T) {
	d := new(DataNode)
	tu := new(Tuple)
	tu.Key = 5
	d.Insert(tu)

	key, err := d.MinKey()
	if err != nil {
		t.FailNow()
	}
	if key != 5 {
		t.FailNow()
	}
}

func TestInsertOutOfOrder_MinKeyIsRight(t *testing.T) {
	d := new(DataNode)
	tu := new(Tuple)
	tu.Key = 5
	d.Insert(tu)
	tu = new(Tuple)
	tu.Key = 3
	d.Insert(tu)

	key, err := d.MinKey()
	if err != nil {
		t.FailNow()
	}
	if key != 3 {
		t.FailNow()
	}
}
