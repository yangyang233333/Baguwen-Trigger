package main

import (
	"net/http"
	sc "small-cache"
	"small-cache/common"
	"small-cache/common/group_cache"
)

var db = map[string]string{
	"key_1": "val_1",
	"key_2": "val_2",
	"key_3": "val_3",
}

func main() {
	selfAddr := "localhost:10086"
	peers := sc.NewHTTPPool(selfAddr)

	// 插入数据
	curG := group_cache.NewGroup("gname", 100)
	for k, v := range db {
		curG.Add(k, v)
	}

	common.LogInstance().Info("small-cache is listening at " + selfAddr)
	http.ListenAndServe(selfAddr, peers)
	// http://localhost:10086/small-cache/gname/key_1
	// http://localhost:10086/small-cache/gname/key_2
	// http://localhost:10086/small-cache/gname/key_3
}
