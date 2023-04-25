package main

// ** Server events and meaning ** //

// ** - 'user_conn' -- User connecting to server
// ** - 'user_disconn' -- User disconnecting from server
// ** - 'user_new' -- New user account
// ** - 'user_delete' -- User account deleted
// ** - 'user_update' -- User info update

// ** - 'chat_create' -- New Chat created
// ** - 'chat_delete' -- Chat deleted
// ** - 'chat_message' -- Chat message received
// ** - 'chat_join' -- Chat joined by another user
// ** - 'chat_invite' -- Chat invite received by another user
// ** - 'chat_leave' -- Chat left by another user

import (
	"yalk-backend/cattp"
	"yalk-backend/logger"

	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"nhooyr.io/websocket"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	log.Print("\033[H\033[2J")
	var version string = "pre-alpha" // make it os.env
	logger.LogColor("CORE", "Booting..")
	logger.LogColor("CORE", fmt.Sprintf("Chat Server version: %s", version)) // make it os.env
}

func main() {
	var wg sync.WaitGroup

	netConf := cattp.Config{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT_PLAIN"),
		URL:  os.Getenv("HTTP_URL"),
	}

	_channels := channels{
		Msg:     make(chan payload, 1),
		Dm:      make(chan map[string]any, 1),
		Notify:  make(chan payload, 1),
		Cmd:     make(chan payload),
		Conn:    make(chan payload),
		Disconn: make(chan payload),
	}

	_websocket := newWebSocketServer(nil, _channels)

	chatServer := &server{
		channels:  _channels,
		websocket: _websocket,
	}

	err := startHTTPServer[any](netConf, chatServer, nil)
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go chatServer.router()
	wg.Wait()
}

var connectHandle = cattp.HandlerFunc[*server](func(w http.ResponseWriter, r *http.Request, context *server) {

	logger.LogColor("WEBSOCK", fmt.Sprintf("Requested WebSocket - %s", r.RemoteAddr))

	conn, err := upgradeHttpRequest(w, r)
	logger.LogColor("WEBSOCK", fmt.Sprintf("Can't start accepting - %s", r.RemoteAddr))

	defer conn.Close(websocket.StatusInternalError, "Client disconnected")

	notify := make(chan bool)

	client := &websocketClient{
		Msgs: make(chan []byte, context.websocket.SubscriberMessageBuffer),
		CloseSlow: func() {
			conn.Close(websocket.StatusPolicyViolation, "connection too slow to keep up with messages")
		},
	}
	context.websocket.SubscribersMu.Lock()
	// context.websocket.Clients[session.UserID] = client
	context.websocket.Clients["test"] = client
	context.websocket.SubscribersMu.Unlock()

	var wg sync.WaitGroup

	// TODO: Properly introduce ping detection
	ticker := time.NewTicker(time.Second * time.Duration(100000))

	// **	Sender - From CLI to SRV	**	//
	wg.Add(1)
	go func(ticker *time.Ticker) {
		defer func() {
			wg.Done()
			notify <- true
		}()
	Run:
		for {
			t, payload, err := conn.Read(r.Context())
			fmt.Printf("Payload len: %v\n", len(payload))
			if err != nil && err != io.EOF {
				statusCode := websocket.CloseStatus(err)
				if statusCode == websocket.StatusGoingAway {
					log.Println("Graceful sender shutdown")
					ticker.Stop()
					break Run
				} else {
					log.Println("Sender - Error in reading from websocket context, client closed? Check main.go")
					break Run
				}
			}
			if t.String() == "MessageText" && err == nil {
				fmt.Printf("Message received: %s", payload)
				// err = server.handlePayload(payload, session.UserID)
				// if err != nil {
				// log.Printf("Sender - errors in broadcast: %v", err)
				// wg.Done()
				// return
				// }
			}
		}
	}(ticker)

	// **		Receiver from SRV to CLI		**	//
	wg.Add(1)
	go func(ticker *time.Ticker) {
		defer func() {
			wg.Done()
		}()

	Run:
		for {
			select {
			case <-notify:
				log.Println("Receiver - got shutdown signal")
				break Run
			case payload := <-client.Msgs:
				err = writeTimeout(r.Context(), time.Second*5, conn, payload)
				if err != nil {
					break Run
				}
			}
		}
	}(ticker)

	_payload := payload{
		Success: true,
		Event:   "user_conn",
	}

	payload, err := json.Marshal(_payload)

	testPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error in initial payload: %v\n", err)
	}
	err = writeTimeout(r.Context(), time.Second*5, conn, testPayload)

	if err != nil {
		log.Printf("Timeout in initial payload: %v\n", err)
	}
	log.Printf("OK - Full data sent to ID: %v\n", "test")

	wg.Wait()

	context.websocket.SubscribersMu.Lock()
	delete(context.websocket.Clients, "test")
	context.websocket.SubscribersMu.Unlock()
	onlineTick := time.NewTicker(time.Second * 10)
	<-onlineTick.C

})

func startHTTPServer[T any](conf cattp.Config, chatServer *server, gorm *gorm.DB) *cattp.Router[T] {
	context := chatServer

	router := cattp.New(context)
	router.HandleFunc("/ws", connectHandle)

	err := router.Listen(&conf)
	if err != nil {
		panic("can't start webapp")
	}

	log.Println("HTTP Server succesfully started") // TODO: Move back in main func
	return nil
}
func upgradeHttpRequest(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	var defaultOptions = &websocket.AcceptOptions{CompressionMode: websocket.CompressionNoContextTakeover, InsecureSkipVerify: true}
	var defaultSize int64 = 2097152 // 2Mb in bytes

	conn, err := websocket.Accept(w, r, defaultOptions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		r.Body.Close()
		return nil, err
	}

	conn.SetReadLimit(defaultSize)
	return conn, nil
}
