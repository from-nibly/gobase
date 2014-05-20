package main

import "fmt"
import "./nodes"

func main() {
	stuff := new(nodes.DataNode)

	for i := 0; i < 4; i++ {
		t := new(nodes.Tuple)
		t.Key = int64(i)
		stuff.Insert(t)
	}

	t := new(nodes.Tuple)
	t.Key = 4
	stuff.Insert(t)

	fmt.Println(stuff.MinKey())

}
