package main

import (
	"./logger"
	"./nodes"
)

type BTree struct {
	root nodes.Node
	left nodes.Node
	prop *nodes.NodeP
}

func (this *BTree) SetProp(prop *nodes.NodeP) {
	this.prop = prop
}

func (this *BTree) Insert(t *nodes.Tuple) {
	if this.root == nil {
		this.root = new(nodes.DataNode)
		this.root.SetProp(this.prop)
		this.left = this.root
	}
	split := this.root.Insert(t)
	logger.Debug().Println("checking split", split)
	logger.Debug().Println("checking split 1?", split)
	if split != nil {
		logger.Debug().Println("checking split 2?", split)
		newRoot := new(nodes.IndexNode)
		newRoot.SetProp(this.prop)
		logger.Debug().Println("checking split 3?", split)
		newRoot.Keys = []int64{split.MinKey()}
		newRoot.Nodes = []nodes.Node{this.root, split}
		this.root = newRoot
	}
}

func (this *BTree) Find(key int64) *nodes.Tuple {
	return this.root.Find(key)
}
