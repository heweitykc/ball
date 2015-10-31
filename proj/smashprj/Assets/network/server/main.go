package main
 
import (
    "websocket"
    "fmt"
    "log"
    "net/http"
	"time"
//	"strconv"
)

const (  
  END, HELLO0, HELLO1 = "6", "100", "101"
  UP, DOWN, RIGHT, LEFT = "2", "3", "4", "5"
  START = "80,0|1,2"	//8n,n,n,...
  TILE =  "9"	//9n
)
 
var GlobalRooms = make(map[string]*Room)
var roomIndex int = 0

type Room struct{
	id int
	player0 *websocket.Conn
	player1 *websocket.Conn
}

func (m *Room) Init(){
	m.player0 = nil
	m.player1 = nil
	go m.Update()
}

func (m *Room) AddPlayer(player *websocket.Conn){
	if m.player0 == nil {
		m.player0 = player
		m.broadcast(HELLO0, m.player0)
		return
	}
	if m.player1 == nil {
		m.player1 = player
		m.broadcast(HELLO1, m.player1)
	}
	
	m.broadcast(START, m.player0)
	m.broadcast(START, m.player1)
}

//1 开始; 2 上; 3 下; 4 左; 5 右; 6 结束;
func (m *Room) broadcast(msg string, player *websocket.Conn){
	if player == nil {
		return
	}
	
	if err := websocket.Message.Send(player, msg); err != nil {
        fmt.Println("Can't send", err)
    }
}

func (m *Room) IsFull() bool{
	if m.player0 != nil && m.player1 != nil {
		return true
	}
	return false
}

func (m *Room) Exec(cmd string, player *websocket.Conn){
	fmt.Println("cmd:", cmd)
	m.broadcast(cmd, m.player0)
	m.broadcast(cmd, m.player1)
}

func (m *Room) Update(){
	for {
		//fmt.Println(m.id," run...")
		time.Sleep(5*time.Second)	
	}
}

//查找房间
func SearchRoom(player *websocket.Conn) *Room{
	for _,room := range GlobalRooms{
		if room.IsFull() == false {
			room.AddPlayer(player)
			return room;
		}
	}
	room := new(Room)
	room.id = roomIndex
	room.Init()
	room.AddPlayer(player)
	GlobalRooms[string(roomIndex)] = room
	roomIndex++
	fmt.Println("add new room-", room.id)
	return room
}


func Echo(ws *websocket.Conn) {
    var err error
	room := SearchRoom(ws)
    for {
        var reply string
 
        if err = websocket.Message.Receive(ws, &reply); err != nil {
            fmt.Println("Can't receive")
            break
        }
		room.Exec(reply, ws)
    }
}
 
func main() {
    fmt.Println("begin")
    http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line
 
    http.Handle("/socket", websocket.Handler(Echo))
 
    if err := http.ListenAndServe(":1234", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
 
    fmt.Println("end")
}