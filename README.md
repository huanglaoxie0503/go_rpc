```markdown
Go 客户端与服务端 RPC 调用的简单实现
    在 Go 语言中，我们可以使用标准库提供的 net/rpc 包很方便地编写 RPC 服务端和客户端程序，因为这个包实现了 RPC 协议的相关细节，使得在 Go 语言中实现 RPC 编程非常简单。
net/rpc 包允许 RPC 客户端程序通过网络或是其他 I/O 连接调用一个服务端对象的公开方法（大小字母开头）。在 RPC 服务端，需要将这个对象注册为可访问的服务，之后该对象的公开方法就能够以远程的方式提供访问。
一个 RPC 服务端可以注册多个不同类型的对象，但不允许注册同一类型的多个对象。此外，一个对象只有满足以下这些条件的方法，才能被 RPC 服务端设置为可提供远程访问： 
1.必须是在对象外部可公开调用的方法（首字母大写）；
2.必须有两个参数，且参数的类型都必须是包外部可以访问的类型或者是 Go 内建支持的类型；
3.第二个参数必须是一个指针；
4.方法必须返回一个 error 类型的值。
以上 4 个条件，可以简单地用如下这行代码表示：
```
```go
func (t *T) MethodName(argType T1, replyType *T2) error
```

```markdown
RPC 编程：引入 jsonrpc 包通过 JSON 对 RPC 传输数据进行编解码
```