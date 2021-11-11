package play

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func Sync(clients map[*websocket.Conn]*User, conn *websocket.Conn,msg string) (re map[string]string, err error) {
	data := map[string]string{"position":""}
	err = json.Unmarshal([]byte(msg),&data)
	log.Println(data,msg)
	if err != nil{
		log.Println("无法解析的报文：",msg,"error:",err)
		return
	}
	a,b := 0,0
	for k,v := range clients{
		switch v.Type {
		case "A": a = 1
		case "B": b = 1
		}
		if k == conn {
			err = errors.New("已经坐下了")
			return
		}
	}
	if a==1 && b==1 {
		err = errors.New("坐满了")
		return
	}
	if a == 0 {
		clients[conn].Type = "A"
		a = 1
		re["message"] = clients[conn].Name + "执白子！"
	}else if b == 0 {
		clients[conn].Type = "B"
		b = 1
		re["message"] = clients[conn].Name + "执黑子！"
	}
	if a==1 && b==1 { //坐下后坐满了，开局
		//begain()
		fmt.Println("开局")
	}
	return
}
