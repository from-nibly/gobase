package nodes

func maxKeysDn() int {
	return 4
}
func maxKeysIn() int {
	return 4
}
func minKeysDn() int {
	return maxKeysDn() / 2
}
func minKeysIn() int {
	return maxKeysIn() / 2
}

type Node interface {
	Insert(t *Tuple) Node
	MinKey() (int64, error)
}

type Tuple struct {
	Key  int64
	Body []byte
}
