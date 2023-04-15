package main

import (
	"common"
	"flag"
	"net/http"
	"strings"
)

var (
	port = flag.String("port", "10050", "proxy的端口")
)

type HTTPHandler struct {
}

// 用户访问
// 写 localhost:10050/kvcache/{group_name}/{key}/{val}
// 读 localhost:10050/kvcache/{group_name}/{key}
func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ["", "kvcache", "{group_name}", "{key}", "{val}"]
	segs := strings.Split(r.URL.Path, "/")
	if len(segs) < 3 {
		w.Write([]byte("参数太少!"))
		return
	}
	if segs[1] != "kvcache" {
		w.Write([]byte("segs[1] != kvcache"))
		return
	}

	if len(segs) == 4 { // 此时为读
		val, ok := Get(segs[2], segs[3])
		if !ok {
			w.Write([]byte(segs[3] + "对应的val不存在"))
		} else {
			w.Write([]byte(val))
		}
	} else if len(segs) == 5 { // 此时为写
		ok := Add(segs[2], segs[3], segs[4])
		if !ok {
			w.Write([]byte(segs[3] + "写入失败"))
		} else {
			w.Write([]byte(segs[3] + "写入成功"))
		}
	} else {
		w.Write([]byte("参数有误!"))
	}
}

func main() {
	flag.Parse()

	// 建立链接
	MakeConn()

	// 向底层的holder节点发送心跳检测，如果在规定时间没有得到回复，应该删除对应的虚拟节点
	Ping() // todo 该函数暂时未实现

	svr := HTTPHandler{}
	common.LogInstance().Info("proxy启动成功")

	http.ListenAndServe("127.0.0.1:"+*port, &svr)
}
