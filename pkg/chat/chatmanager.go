package chat

import (
	"sync"
	"log"
	"time"
	"github.com/gorilla/websocket"
)

// key = ChatUniqueIdentifier; value = *websocket.Conn
type ChatManager struct {
	WaitingConns sync.Map
	Chats chan Chat
}

func NewManager() *ChatManager {
	return &ChatManager {
		WaitingConns: sync.Map{},
		Chats: make(chan Chat),
	}
}

func (m *ChatManager) Add(c *websocket.Conn, seq *ChatInitSequence) {
	keyPair := NewKeyPair(*seq)
	idPtr := keyPair.GetChatIdentifier()
	conn, loaded := m.WaitingConns.LoadOrStore(idPtr.String(), c)
	
	if loaded {
		m.WaitingConns.Delete(idPtr.String())
		chat := Chat{c, conn.(*websocket.Conn)}
		m.Chats <- chat
		log.Println("Two clients connected, starting chat...")
	} else {
		log.Println("First client connected, added to queue")	
	}
}

func (m *ChatManager) Start() {
	for chat := range m.Chats {
		close := make(chan bool)
		go func(){
			for {
				err := sendMessage(chat.First, chat.Second)
				if err != nil {
					log.Println(err)
					break
				}
			}
			close <- true
		}()
	
		go func(){
			for {
				err := sendMessage(chat.Second, chat.First)
				if err != nil {
					log.Println(err)
					break
				}
			}
			close <- true
		}()

		go func(){
			<- close
			deadline := time.Now().Add(time.Duration(1000000))
			chat.First.WriteControl(websocket.CloseMessage, []byte{}, deadline)
			chat.First.Close()
			chat.Second.WriteControl(websocket.CloseMessage, []byte{}, deadline)
			chat.Second.Close()
		}()
	}	
}
