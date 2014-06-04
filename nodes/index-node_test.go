package nodes

import (
	"../logger"
	"os"
	"testing"
)

func TestInsertForEvenSplit_CanFindAllEntries(t *testing.T) {
	logger.Init(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stdout)
	node := new(DataNode)
	node.prop = getSettings()
	for i := 0; i < 4; i++ {
		node.Insert(&Tuple{int64(i * 10), nil})
		logger.Debug().Println("DUMP")
		node.Dump(0)
	}
	split := node.Insert(&Tuple{int64(40), nil})
	logger.Debug().Println(split)

	index := &IndexNode{[]Node{node, split}, []int64{split.MinKey()}, getSettings()}
	for i := 5; i < 12; i++ {
		index.Insert(&Tuple{int64(i * 10), nil})
	}
	index.Insert(&Tuple{int64(2), nil})
	index.Insert(&Tuple{int64(3), nil})
	isplit := index.Insert(&Tuple{int64(4), nil})

	logger.Debug().Println(isplit)
	for i := 0; i < 4; i++ {
		if index.Find(int64(i*10)) == nil {
			logger.Debug().Println("Could not find", i*10)
			t.FailNow()
		}
	}
	for i := 2; i < 5; i++ {
		if index.Find(int64(i)) == nil {
			logger.Debug().Println("Could not find", i*10)
			t.FailNow()
		}
	}
	for i := 4; i < 12; i++ {
		if isplit.Find(int64(i*10)) == nil {
			logger.Debug().Println("Could not find", i*10)
			t.FailNow()
		}
	}
}
