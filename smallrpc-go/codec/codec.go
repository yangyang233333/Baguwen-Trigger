package codec

// 消息编解码

/*
一个典型的 RPC 调用如下：
err = client.Call("Arith.Multiply", args, &reply)

客户端发送的请求包括服务名 Arith，方法名 Multiply，参数 args 三个，服务端的响应包括错误 error，
返回值 reply 2 个。我们将请求和响应中的参数和返回值抽象为 body，剩余的信息放在 header 中，那么就
可以抽象出数据结构 Header

*/
