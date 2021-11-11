//处理基础js效果
//弹窗效果
function g(o) { return document.getElementById(o); }

// function ShowBox() {
//     var romTable = "";
//     for (var i=0;i<5;i++){
//         romTable += "<tr>";
//         for(var j=0;j<5;j++)
//         {
//             romNum = i*5+j+1;
//             romTable += "<td onclick='intoRom("+romNum+")'>房间"+(romNum)+"<br>人数："+3+"人</td>";
//         }
//         romTable+="</tr>"
//     }
//     document.getElementById("rom").innerHTML=romTable;
//     g("backlayer").className = 'backLayer';
//     g("dialogbox").className = 'boxShow';
// }

// function CloseBox() {
//     if (rom == 0 ){
//         alert("请选择一个房间进入")
//         return false;
//     }
//     g("backlayer").className = 'hide';
//     g("dialogbox").className = 'hide';
// }

//html基础棋盘
//先画棋盘
var tableobj =document.getElementById("qipan");
for (var i=0;i<15;i++){
    var trobj =document.createElement("tr");
    for(var j=0;j<15;j++)
    {
        var tdobj = document.createElement("td");
        trobj.appendChild(tdobj);
    }
    tableobj.appendChild(trobj);
}