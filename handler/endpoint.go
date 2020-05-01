package handler
import (
	"github.com/idawud/server-monitor/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true },
}

type WSEndPoint struct {
	l *log.Logger
}

func NewWebSocketEndpoint(l *log.Logger)  *WSEndPoint {
	return &WSEndPoint{l: l}
}

func (e *WSEndPoint) MainEndpoint(w http.ResponseWriter, r *http.Request) {
	e.l.Println("Main WebSocket (Start)")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		e.l.Println("Error: ", err)
	}
	e.l.Println("Client successfully connected")

	readerAndWriter(ws, e.l)
	e.l.Println("Main WebSocket (End)")
}

func readerAndWriter(ws *websocket.Conn, l *log.Logger) {
	for  {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			l.Println(err)
			return
		}
		log.Println("Client Message: ", string(p))
		for  {
			availability, err := service.GetAllAvailability()
			if err != nil {
				l.Println("Error: ", err)
			} else {
				if err := ws.WriteMessage(messageType, availability); err != nil {
					l.Println("Error:<> ", err)
					ws.Close()
				} else {
					l.Println("New Data Published")
				}
			}

			time.Sleep(time.Second * 15)
		}

	}
}

