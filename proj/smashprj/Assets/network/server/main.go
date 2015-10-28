package main
 
import (
    "websocket"
    "fmt"
    "log"
    "net/http"
	"time"
	"strconv"
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
		return
	}
	if m.player1 == nil {
		m.player1 = player
		return
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
}

func (m *Room) Update(){
	for {
		fmt.Println(m.id," run...")
		time.Sleep(3*time.Second)	
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
	fmt.Println("add new room ", room.id)
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
 
		
        msg := "Received:  " + strconv.Itoa(room.id) + " " + reply + string(0)		
        fmt.Println("Sending to client: " + msg + " ")
        if err = websocket.Message.Send(ws, msg); err != nil {
            fmt.Println("Can't send")
            break
        }		
		
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