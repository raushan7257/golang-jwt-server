package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func startRecording(conn *websocket.Conn) {
    fmt.Println("Start recording")
    // TODO: implement actual start recording logic here
    conn.WriteMessage(websocket.TextMessage, []byte("Recording started"))
}

// Stop recording logic
func stopRecording(conn *websocket.Conn) {
    fmt.Println("Stop recording")
    // TODO: implement actual stop recording logic here
    conn.WriteMessage(websocket.TextMessage, []byte("Recording stopped"))
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Upgrade error:", err)
        return
    }
    defer conn.Close()

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Read error:", err)
            break
        }

        command := string(msg)

        switch command{
            case "start":
                startRecording(conn)
            case "stop":
                stopRecording(conn)
            default:
                conn.WriteMessage(websocket.TextMessage, []byte("Unknown command"))
        }
    }
}
