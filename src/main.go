package main

import (
	"github.com/omarsagoo/omar-ekko/server"
)

// var (
// 	clients   = make(map[*websocket.Conn]bool) // client connections map
// 	broadcast = make(chan Message)             // Broadcast Channel
// 	// configure the upgrader
// 	upgrader = websocket.Upgrader{
// 		CheckOrigin: func(r *http.Request) bool {
// 			return true
// 		},
// 	}
// )

// type Message struct {
// 	Username string `json:"username"`
// 	Message  string `json:"message"`
// 	Avatar   string `json:"avatar"`
// }

// func handleConnections(w http.ResponseWriter, r *http.Request) {
// 	// upgrade Get request to websocket
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal("ConnectionHandler:", err)
// 	}

// 	// close when this function ends
// 	defer ws.Close()

// 	// add client to global clients map
// 	clients[ws] = true

// 	for {
// 		var msg Message

// 		// read new msg in as JSON, and initialize is as a Message object
// 		err := ws.ReadJSON(&msg)
// 		if err != nil {
// 			log.Printf("ConnectionHandlerMsg: %v", err)
// 			delete(clients, ws)
// 			break
// 		}
// 		// send the msg to the broadcast channel, to be handled by handleMessage goroutine
// 		broadcast <- msg
// 	}
// }

// func handleMessages() {
// 	for {
// 		msg := <-broadcast

// 		for client := range clients {
// 			err := client.WriteJSON(msg)
// 			if err != nil {
// 				log.Printf("error: %v", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }

func main() {
	// // create file server
	// fs := http.FileServer(http.Dir("../public"))
	// http.Handle("/", fs)

	// // configure websocket route
	// http.HandleFunc("/ws", handleConnections)

	// // start listening for messages
	// go handleMessages()

	// log.Println("http server started on port :8000")
	// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	server.Start()
}
