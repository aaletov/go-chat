package chat

import (
	"sync"
)

type ChatManager struct {
	Chats []Chat
}

func NewManager() ChatManager {
	return ChatManager{
		Chats: make([]Chat, 0),
	}
}

func (m *ChatManager) Add(chat Chat) {
	var wg sync.WaitGroup
	wg.Add(2)

	m.Chats = append(m.Chats, chat)

	go func() {
		for {
			err := sendMessage(chat.First, chat.Second)
			if err != nil {
				log.Println(err)
				break
			}
		}
		wg.Done()
	}()

	go func () {
		for {
			err := sendMessage(chat.Second, chat.First)
			if err != nil {
				log.Println(err)
				break
			}
		}
		wg.Done()
	}

	wg.Wait()

	chat.First.Close()
	chat.Second.Close()
	// delete chats
}
