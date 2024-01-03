package sessions

import (
	"crypto/rand"
	"log"
	"time"
)

type SessionManager struct {
	//Name describes the cookie name
	Name   string
	store  Store
	Expiry time.Duration
}

// NewSessionManager returns a session manager and the expiry routine
func NewSessionManager(exp time.Duration, store Store) (*SessionManager, func(time.Duration)) {
	if store == nil {
		store = NewMemStore()
	}
	ses := &SessionManager{store: store}
	f := func(interval time.Duration) {
		for {
			time.Sleep(interval)
			log.Println("running expiry routine")
			store.Expire(ses.Expiry)
		}
	}
	return ses, f
}

func (s *SessionManager) NewSession() (string, error) {
	sid := s.newSid()
	err := s.store.Save(sid, time.Now())
	if err != nil {
		return "", err
	}
	return sid, err
}

func (s *SessionManager) newSid() string {
	sid := make([]byte, 32) //32 byte sid
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
