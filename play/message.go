package play

import (
	"encoding/json"
	"log"
)

func SystemMessage(msg string){
	reData := new(Message)
	reData.Type = "message"
	reData.Data = map[string]string{"message":msg}
	reData.User = User{Name:"系统消息",Type:"system"}
	// 发送到通道
	Broadcast <- *reData
}
func UserMessage(msg string,user User)  {
	data := map[string]string{"message":""}
	err := json.Unmarshal([]byte(msg),&data)
	if err != nil{
		log.Println("无法解析的报文：",msg,"error:",err)
		return
	}


	reData := new(Message)
	reData.Type = "message"
	reData.Data = data
	reData.User = user
	// 发送到通道
	Broadcast <- *reData
}

//小游戏状态检查
func CheckQuest(user User,message string)  {
	if message == "{\"message\":\"摘果子\"}" && Quest_zhai == 1{
		Quest_zhai = 0
		SystemMessage("恭喜"+user.Name+"第一个摘取果子~咔吧咔吧可好吃了~")
	}
}