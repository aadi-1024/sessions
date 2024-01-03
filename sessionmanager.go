package sessions

import (
	"crypto/rand"
	"log"
	"time"
)

type SessionManager struct {
	store Store
}

func NewSessionManager(exp time.Duration) *SessionManager {
	return &SessionManager{
		NewMemStore(exp),
	}
}

func (s *SessionManager) NewSession() error {
	sid := s.newSid()
	err := s.store.Save(sid, time.Now())
	return err
}

func (s *SessionManager) newSid() string {
	sid := make([]byte, 16) //16 byte sid
	_, err := rand.Read(sid)

	if err != nil {
		log.Println(err)
	}

	id := s.store.Load(string(sid))
	for id != nil {
		_, err = rand.Read(sid)
		if err != nil {
			log.Println(err)
		}
		id = s.store.Load(string(sid))
	}

	return string(sid)
}
