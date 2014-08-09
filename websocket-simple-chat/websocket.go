package main

import (
	"log"
	"net"
	"net/http"
	"reflect"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/mattn/go-scan"
)

var ActiveClients = make(map[ClientConn]int)
var ActiveClientsRWMutex sync.RWMutex

type ClientConn struct {
	websocket *websocket.Conn
	clientIP  net.Addr
}

func addClient(cc ClientConn) {

	ActiveClientsRWMutex.Lock()
	ActiveClients[cc] = 0
	ActiveClientsRWMutex.Unlock()
}

func deleteClient(cc ClientConn) {

	ActiveClientsRWMutex.Lock()
	delete(ActiveClients, cc)
	ActiveClientsRWMutex.Unlock()
}

func broadcastMessage(messageType int, message []byte) {

	ActiveClientsRWMutex.RLock()

	defer ActiveClientsRWMutex.RUnlock()

	for client, _ := range ActiveClients {
		if err := client.websocket.WriteMessage(messageType, message); err != nil {
			return
		}
	}
}

func WebSocket(w http.ResponseWriter, r *http.Request) {

	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if _, ok := err.(websocket.HandshakeError); ok {

		http.Error(w, "Not a websocket handshake", 400)
		return

	} else if err != nil {

		log.Println(err)
		return

	}

	client := ws.RemoteAddr()
	clientConn := ClientConn{ws, client}
	addClient(clientConn)

	for {
		messageType, p, err := ws.ReadMessage()

		if err != nil {
			deleteClient(clientConn)
			log.Println(err)
			return
		}

		broadcastMessage(messageType, p)
		log.Println(ByteToStr(p))

		var _type string
		scan.ScanJSON(strings.NewReader(ByteToStr(p)), "/type", &_type)

		log.Println("type")
		log.Println(_type)

		if reflect.DeepEqual(_type, "new") {
			var _data string
			scan.ScanJSON(strings.NewReader(ByteToStr(p)), "/data", &_data)
			InsertEntry(_data)
		}

		log.Println(GetEntryAll())

	}

}
