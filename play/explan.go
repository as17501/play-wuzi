package play

import "github.com/gorilla/websocket"

var Qipan = make([][]int,15) //盘面信息
var RomInfo = new(Info) //对局信息
var Clients = make(map[*websocket.Conn]*User) // 客户端，标识：1,2 玩家，3观众，0未连接

var Quest_zhai = 0
var Broadcast = make(chan Message)           // broadcast channel

type Info struct {
	PlayerA User
	PlayerB User
	Statue int
}
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type User struct {
	Name string
	Type string
	Table int
}
type Message struct {
	User `json:"user"`
	Type string `json:"type"`
	Data interface{} `json:"data"`
}