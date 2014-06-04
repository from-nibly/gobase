package main

import (
	"./logger"
	"./nodes"
	"fmt"
	"io/ioutil"
	"testing"
)

func getSettings() *nodes.NodeP {
	logger.Init(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard)
	prop := new(nodes.NodeP)
	prop.SetMaxKeysDn(4)
	prop.SetMaxKeysIn(4)
	return prop
}

func TestInsertOneMillion(t *testing.T) {

	tree := new(BTree)
	tree.SetProp(getSettings())

	for i := 0; i < 1000000; i++ {
		if i%1000 == 0 {
			fmt.Println("Inserting", i)
		}
		tree.Insert(&nodes.Tuple{int64(i), nil})
		//for j := 0; j <= i; j++ {
		//	if tree.Find(int64(i)) == nil {
		//		fmt.Println("Could not find ", j)
		//		t.FailNow()
		//	}
		//}
	}
}
