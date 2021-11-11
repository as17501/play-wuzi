package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"play/play"
	"strconv"
	"time"
)


// 只配置了跨域请求，其他配置见文档
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	// 创建静态资源服务
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// 创建websocket路由
	http.HandleFunc("/ws", handleConnections)

	// 开始监听处理消息
	go handleMessages()

	//创建定时器
	timeTickerChan := time.Tick(time.Second * 20)
	reData := new(play.Message)
	log.Print("hehhhh")
	go func() {
		//创建棋盘
		play.InitGame()

		//定时发起摘果子任务
		for {
			<-timeTickerChan
			play.Quest_zhai = 1
			reData.Type = "message"
			reData.Data = map[string]string{"message":"【小游戏】果子成熟了，有人要摘果子吗？发送“摘果子”来抢吧~"}
			reData.User = play.User{Name:"任务小精灵",Type:"system"}
			// 发送到通道
			play.Broadcast <- *reData
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

		}
	}()


	// 开启服务端，注意，ListenAndServe是阻塞的
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}


}

//ws处理器
func handleConnections(w http.ResponseWriter, r *http.Request) {
	//升级get请求到ws请求
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// 注册客户端
	play.Clients[ws] = &play.User{Name:"用户"+strconv.Itoa(int(rand.Int31n(1000000))),Type:""}

	for {
		// Read in a new message as JSON and map it to a Message object
		messageType,msg,err := ws.ReadMessage()
		//log.Println("收到消息",string(msg),messageType)
		if err != nil {
			log.Printf("error: %v - %d", err,messageType)
			delete(play.Clients, ws)
			break
		}
		//只接受文字消息
		if messageType != 1 {
			continue
		}
		data := make([]string,2,2)
		for k,v := range msg{
			if string(v) == "|" {
				data[0] = string(msg[0:k])
				data[1] = string(msg[k+1:])
			}
		}
		//log.Println("解析消息",data)
		switch data[0] {
		case "message":
			play.UserMessage(data[1], *play.Clients[ws])
			play.CheckQuest(*play.Clients[ws],data[1])
		case "play": //下棋方法
			play.Play(ws,data[1])
		case "seat":
			play.Seat(ws,data[1])

		case "changeName":
			play.ChangeName(play.Clients[ws],data[1])
		default:continue
		}
		if err != nil {
			log.Printf("任务异常，忽略: %v ", err)
			continue
		}

	}
}

//处理消息
func handleMessages() {
	for {
		// 接受消息
		data := <-play.Broadcast
		// 广播消息
		for client := range play.Clients {
			re,_ := json.Marshal(data)
			//log.Printf("广播：%+v",data)
			err := client.WriteMessage(websocket.TextMessage, re)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(play.Clients, client)
			}
		}
	}
}

