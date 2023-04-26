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
	"yalk-backend/chat"
	"yalk-backend/logger"

	"encoding/json"
	"fmt"
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
	var version string = "pre-alpha"
	logger.Info("CORE", "Booting..")
	logger.Info("CORE", fmt.Sprintf("Chat Server version: %s", version)) // make it os.env
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	log.Print("\033[H\033[2J") // Clear console

}

func main() {
	var wg sync.WaitGroup

	netConf := cattp.Config{
		Host: os.Getenv("HTTP_HOST"),
		Port: os.Getenv("HTTP_PORT_PLAIN"),
		URL:  os.Getenv("HTTP_URL"),
	}

	chatServer := chat.NewServer(16)

	err := startHttpServer(netConf, chatServer, nil)
	if err != nil {
		panic(err)
	}
	wg.Add(1)
	go chatServer.Router()
	wg.Wait()
}

// TODO: Enum log events and colors 'info' 'warning 'error'
var connectHandle = cattp.HandlerFunc[*chat.Server](func(w http.ResponseWriter, r *http.Request, server *chat.Server) {

	logger.Info("WEBSOCK", fmt.Sprintf("Requested WebSocket - %s", r.RemoteAddr))

	// Upgrading HTTP request to Websocket
	conn, err := upgradeHttpRequest(w, r)
	if err != nil {
		logger.Err("WEBSOCK", fmt.Sprintf("Can't start accepting - %s", r.RemoteAddr))
	}

	// Defering normal closing if the function returns
	defer conn.Close(websocket.StatusNormalClosure, "Client disconnected")

	// Register and return client with Chat Server
	// Todo: Use profile instead of User ID?
	client := server.RegisterClient(conn, "test")

	notify := make(chan bool)

	var wg sync.WaitGroup
	var ticker = time.NewTicker(time.Second * time.Duration(100000))

	// TODO: Properly introduce ping detection
	channelsContext := &chat.MessageChannelsContext{
		NotifyChannel: notify,
		PingTicket:    ticker,
		WaitGroup:     &wg,
		ClientData:    client,
		Request:       r,
		Connection:    conn,
	}

	// **	Sender - From CLI to SRV	**	//
	wg.Add(1)
	go chat.Receiver(channelsContext)

	// **		Receiver from SRV to CLI		**	//
	wg.Add(1)
	go chat.Sender(channelsContext)

	_payload := chat.Payload{
		Success: true,
		Event:   "user_conn",
	}

	payload, err := json.Marshal(_payload)
	if err != nil {
		log.Printf("Error in initial payload: %v\n", err)
	}

	if chat.ClientWriteWithTimeout(r.Context(), time.Second*5, conn, payload); err != nil {
		log.Printf("Timeout in initial payload: %v\n", err)
	}

	log.Printf("OK - Full data sent to ID: %v\n", "test")

	wg.Wait()

	server.ClientsMu.Lock()
	delete(server.Clients, "test")
	server.ClientsMu.Unlock()
	onlineTick := time.NewTicker(time.Second * 10)
	<-onlineTick.C

})

func startHttpServer[T chat.Server](conf cattp.Config, chatServer *chat.Server, gorm *gorm.DB) *cattp.Router[T] {
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
