//开始处理websocket
var conn;
var msg = document.getElementById("msg");
var log = document.getElementById("log");

///留言板
function appendLog(item) {
    var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    log.appendChild(item);
    if (doScroll) {
        log.scrollTop = log.scrollHeight - log.clientHeight;
    }
}
//留言方法
document.getElementById("form").onsubmit = function () {
    if (!conn) {
        return false;
    }
    if (!msg.value) {
        return false;
    }
    console.log(msg)
    var messageObj = {message:msg.value}
    conn.send("message|"+JSON.stringify(messageObj));
    msg.value = "";
    return false;
};
//改名方法
document.getElementById("changeName").onclick = function () {
    if (!conn) {
        return false;
    }
    var name = document.getElementById("name");
    if (!name.value) {
        return false;
    }
    var messageObj = {newName:name.value,type:"changeName"}
    conn.send("changeName|"+JSON.stringify(messageObj));
    name.value = "";
    return false;
};
//落座方法
function seat(position){
    var messageObj = {position: position,}
    conn.send("seat|"+JSON.stringify(messageObj));
}
// document.getElementById("seat").onclick = function () {
//     if (!conn) {
//         return false;
//     }
//     var messageObj = {type:"seat"}
//     conn.send("seat|"+JSON.stringify(messageObj));
//     return false;
// };

//加入房间
function intoRom(num){

}
if (window["WebSocket"]) {
    conn = new WebSocket("ws://"+document.domain+":8000/ws");
    conn.onclose = function (evt) {
        var item = document.createElement("div");
        item.innerHTML = "<b>Connection closed.</b>";
        appendLog(item);
    };
    conn.onmessage = function (evt) {
        var data = JSON.parse(evt.data)
        console.log(data)
        if(data.type == "message"){ //聊天消息
            var item = document.createElement("li");
            if (data.user.Type == "system"){
                item.innerText =  "【"+data.user.Name+ "】"+data.data.message;
            }else{
                item.innerText = data.user.Name + "说："+data.data.message;
            }
            appendLog(item);
        }else if(data.type == "play"){ //落子
            drew(data.position.x,data.position.y,data.user.Type)
        }else if(data.type == "radio"){ //广播消息
            var item = document.createElement("li");
            item.innerText = data.data.message;
            appendLog(item);
        }else if(data.type == "sync"){ //同步盘面信息
            sync(data.data)
        }

    };
} else {
    var item = document.createElement("div");
    item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
    appendLog(item);
}



//下棋方法
var btn = document.getElementById('qipan').getElementsByTagName("td");
var isClick = true;
var player = "A";
var nowPlayer = "A"
for (var i=0; i<btn.length; i++) {
    btn[i].addEventListener('click',click1,false);
}
function click1() {
    if (!conn) {
        return false;
    }
    if(isClick) {
        isClick = false;
        //事件
        position = checkPosition(this)
        var messageObj = {"x":position.x,"y":position.y}
        console.log(messageObj)
        conn.send("play|"+JSON.stringify(messageObj));
        //定时器
        setTimeout(function() {
            isClick = true;
        }, 200);//一秒内不能重复点击
        this.removeEventListener('click',click,false) //移除点击绑定
    }
}
//检查点击点位
function checkPosition(e) {
    var tr =  e.parentNode;
    var hang = Array.from(tr.querySelectorAll('td'));
    var lie =  Array.from(tr.parentNode.querySelectorAll('tr'));
    return {"x":hang.indexOf(e),"y":lie.indexOf(tr)}
}
//下棋
function drew(x,y, player) {
    console.log(x,y, player)
    qipan = document.getElementById('qipan');
    hang = qipan.querySelectorAll('tr')[y];
    check = hang.querySelectorAll('td')[x];
    check.style.backgroundImage="url(./img/play"+player+".png)";
    changePlayer();
    check.removeEventListener('click',click,false) //移除点击绑定
}
//切换当前玩家
function changePlayer(now = null) {
    if (now){
        nowPlayer = now;
    }else{
        nowPlayer = nowPlayer=="A"?"B":"A";
    }
    return nowPlayer;
}

function sync(data){
    info = data.info
    pan = data.pan
    qipan = document.getElementById('qipan');
    pan.forEach(function (v,y){
        v.forEach(function (v1,x){
            hang = qipan.querySelectorAll('tr')[y];
            check = hang.querySelectorAll('td')[x];
            check.style.backgroundImage="url(./img/play"+v1+".png)";
        })
    })
    document.getElementById("palyerA").getElementsByTagName("span")[0].innerHTML=info.PlayerA.Name;
    document.getElementById("palyerB").getElementsByTagName("span")[0].innerHTML=info.PlayerB.Name;

    console.log(info,pan)
}