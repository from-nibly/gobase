package nodes

type Node interface {
	Insert(t *Tuple) Node
	MinKey() int64
	Dump(level int)
	Find(key int64) *Tuple
	SetProp(prop *NodeP)
}

type Tuple struct {
	Key  int64
	Body []byte
}
