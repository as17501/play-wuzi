package play

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)


func Play(conn *websocket.Conn,jsonData string){
	data := make(map[string]int)
	err := json.Unmarshal([]byte(jsonData),&data)
	//log.Println(data,jsonData)
	if err != nil{
		log.Println("无法解析的报文：",jsonData,"error:",err)
		return
	}
	//记录当前数据
	if Clients[conn].Type == "A" && RomInfo.Statue == 1 && Qipan[data["y"]][data["x"]] == 0{
		//log.Print("A",data)
		Qipan[data["y"]][data["x"]] = 1
		RomInfo.Statue = 2
		checkWin(conn,data["x"],data["y"])
	}
	if Clients[conn].Type == "B" && RomInfo.Statue == 2 && Qipan[data["y"]][data["x"]] == 0{
		//log.Print("A",data)
		Qipan[data["y"]][data["x"]] = 2
		RomInfo.Statue = 1
		checkWin(conn,data["x"],data["y"])
	}

	//判断输赢
	//if true {
	//	data.Type = "play"
	//}else{
	//	//修改user
	//	//todo
	//	data.Type = "win"
	//}
	return
}
func checkWin(conn *websocket.Conn, x,y int)  {
	SyncInfo()
	mark := 0
	if Clients[conn].Type == "A" {
		mark = 1
	}else if Clients[conn].Type == "B"  {
		mark = 2
	}else {
		return
	}
	log.Print("当前点", Qipan[y][x])
	lianxu1,lianxu2,lianxu3,lianxu4 := 0,0,0,0
	temp_x,temp_y := 0,0
	//判断竖
	for i := -4; i< 5; i++ {
		temp_x = x
		temp_y = y+i
		if temp_x < 0 || temp_x > 14  || temp_y <0 || temp_y>14 {
			continue
		}
		log.Print(temp_x,temp_y, Qipan[temp_y][temp_x])
		if Qipan[temp_y][temp_x] == mark {
			lianxu1++
			if lianxu1 == 5{
				break
			}
		}else {
			lianxu1 = 0
		}
	}
	//判断横
	for i := -4; i< 5; i++ {
		temp_x = x+i
		temp_y = y
		if temp_x < 0 || temp_x > 14  || temp_y <0 || temp_y>14{
			continue
		}
		if Qipan[temp_y][temp_x] == mark {
			lianxu2++
			if lianxu2 == 5{
				break
			}
		}else {
			lianxu2 = 0
		}
	}
	//判断斜
	for i := -4; i< 5; i++ {
		temp_x = x+i
		temp_y = y+i
		if temp_x < 0 || temp_x > 14  || temp_y <0 || temp_y>14{
			continue
		}
		if Qipan[temp_y][temp_x] == mark {
			lianxu3++
			if lianxu3 == 5{
				break
			}
		}else {
			lianxu3 = 0
		}
	}
	for i := -4; i< 5; i++ {
		temp_x = x+i
		temp_y = y-i
		if temp_x < 0 || temp_x > 14  || temp_y <0 || temp_y>14{
			continue
		}
		if Qipan[temp_y][temp_x] == mark {
			lianxu4++
			if lianxu4 == 5{
				break
			}
		}else {
			lianxu4 = 0
		}
	}
	//综合判断
	if lianxu1 > 4 || lianxu2 > 4 || lianxu3 > 4 || lianxu4 > 4  {
		//胜利
		log.Print("胜利",Clients[conn])
		SystemMessage(Clients[conn].Name + "胜利，恭喜！！")
		SystemMessage(Clients[conn].Name + "20秒后重新开始，请稍后。")
		RomInfo.Statue = 0
		go func() {
			time.Sleep(20 * time.Second)
			InitGame()
			SystemMessage("新的一局开始啦，请棋手落座！")
		}()
	}

}

//同步盘面信息
func SyncInfo()  {
	data := make(map[string]interface{})
	data["pan"] = Qipan
	data["info"] = RomInfo
	reData := new(Message)
	reData.Type = "sync"
	reData.Data = data
	// 发送到通道
	Broadcast <- *reData

}

//初始化盘面
func InitGame()  {
	pan:=make([][]int,15)
	for  k1,_ := range pan {
		pan[k1] = make([]int,15)
	}
	Qipan = pan
	log.Print(Qipan)
	RomInfo = new(Info)
	for _,user := range Clients {
		user.Type = ""
	}
	SyncInfo()
}