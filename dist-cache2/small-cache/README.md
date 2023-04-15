## 基于一致性哈希实现一个分布式缓存

### 基本架构如下
- common/ 中存放的是公共的代码，例如group_cache【单机】
- consistenhash/ 中存放的是一致性哈希的代码
- kvholder/ 中存放的是底层的SingleNode的代码  独立的程序
- proxy/ 中存放的是无状态路由层的代码    独立的程序
