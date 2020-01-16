package main

/*
引入 jsonrpc 包通过 JSON 对 RPC 传输数据进行编解码
*/

type Item struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type RpcResponse struct {
	Ok bool `json:"ok"`
	Id int `json:"id"`
	Message string `json:"message"`
}

