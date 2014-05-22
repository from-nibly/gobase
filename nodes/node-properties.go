package nodes

type NodeP struct {
	maxKeysDn int
	maxKeysIn int
}

func (prop *NodeP) SetMaxKeysIn(val int) {
	if val%2 != 0 {
		panic("Max keys must be even")
	}
	prop.maxKeysIn = val
}

func (prop *NodeP) SetMaxKeysDn(val int) {
	if val%2 != 0 {
		panic("Max keys must be even")
	}
	prop.maxKeysDn = val
}

func (prop *NodeP) GetMaxKeysDn() int {
	return prop.maxKeysDn
}

func (prop *NodeP) GetMaxKeysIn() int {
	return prop.maxKeysIn
}

func (prop *NodeP) GetMinKeysDn() int {
	return prop.maxKeysDn / 2
}

func (prop *NodeP) GetMinKeysIn() int {
	return prop.maxKeysIn / 2
}
