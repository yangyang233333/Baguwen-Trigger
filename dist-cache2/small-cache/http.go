package small_cache

import (
	"net/http"
	"strings"
)

const defaultBasePath = "/small-cache/"

// HTTPPool 作为承载节点间 HTTP 通信的核心数据结构
type HTTPPool struct {
	// 本机的IP+PORT
	self string

	// 节点间进行通信的路由的前缀，也就是说domain.com/{basePath}/是用于节点间通信
	basePath string
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultBasePath,
	}
}

// 我们约定访问路径格式为 /<basepath>/<groupname>/<key>，
// 通过 groupname 得到 group 实例，再使用 group.Get(key) 获取缓存数据。
// 最终使用 w.Write() 将缓存值作为 httpResponse 的 body 返回。
func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/") // ["", basepath, groupname, key]
	if len(segs) != 4 {

	}

}