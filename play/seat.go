package play

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
)

var MsgNewName = map[string]string{
	"new":"",
}

func Seat(conn *websocket.Conn,msg string)  {
	data := map[string]string{"position":""}
	err := json.Unmarshal([]byte(msg),&data)
	log.Println(data,msg)
	if err != nil{
		log.Println("无法解析的报文：",msg,"error:",err)
		return
	}

	if Clients[conn].Type == "A" || Clients[conn].Type == "B"{
		err = errors.New("无效点击，已经坐下了")
		return
	}
	if  (data["position"] == "A" && RomInfo.PlayerA.Name != "") || (data["position"] == "B" && RomInfo.PlayerB.Name != ""){
		err = errors.New("无效点击，位置有人")
		return
	}
	if data["position"] == "A" {
		Clients[conn].Type = "A"
		RomInfo.PlayerA = *Clients[conn]
		SystemMessage(RomInfo.PlayerA.Name + "执白子！")
	}
	if data["position"] == "B" {
		Clients[conn].Type = "B"
		RomInfo.PlayerB = *Clients[conn]
		SystemMessage(RomInfo.PlayerB.Name + "执黑子！")
	}

	if RomInfo.PlayerA.Name != "" && RomInfo.PlayerB.Name != "" {
		SystemMessage("棋手已准备，开始对弈！")
		RomInfo.Statue = 1 //A棋手开始下棋
	}
	SyncInfo()
	return
}
func ChangeName(user *User,msg string)  {
	data := map[string]string{"newName":""}
	err := json.Unmarshal([]byte(msg),&data)
	log.Println(data,msg)
	if err != nil{
		log.Println("无法解析的报文：",msg,"error:",err)
		return
	}
	if data["newName"] == "" {
		err = errors.New("新名字不能为空")
		return
	}
	SystemMessage(user.Name + "改名为" + data["newName"] + "了！")
	user.Name = data["newName"]

	return
}
func IntoRom(user *User,msg string) (re map[string]string, err error) {
	data := map[string]string{"newName":""}
	err = json.Unmarshal([]byte(msg),&data)
	log.Println(data,msg)
	if err != nil{
		log.Println("无法解析的报文：",msg,"error:",err)
		return
	}
	if data["newName"] == "" {
		err = errors.New("新名字不能为空")
		return
	}
	re = make(map[string]string)
	re["message"] = user.Name + "改名为" + data["newName"] + "了！"
	user.Name = data["newName"]
	return
}
