package sessions

import (
	"errors"
	"fmt"
	"time"
)

type Store interface {
	//Store save a session
	Save(sid string, created time.Time) error

	//Load loads a session from the store, returns nil if not present
	Load(sid string) Session

	Delete(sid string) error

	Expire()
}

// ideally the implementation should be in a separate file

type MemStore struct {
	expiry time.Duration
	data   map[string]Session
}

func NewMemStore(exp time.Duration) *MemStore {
	m := &MemStore{
		expiry: exp,
		data:   make(map[string]Session),
	}
	go func() {
		for {
			time.Sleep(30 * time.Second)
			m.Expire()
			for k, _ := range m.data {
				fmt.Println(k)
			}
		}
	}()
	return m
}

func (m *MemStore) Save(sid string, created time.Time) error {
	_, ok := m.data[sid]
	if !ok {
		m.data[sid] = NewMapSession(sid, created)
		return nil
	}
	return errors.New("session with given id already exists")
}

func (m *MemStore) Load(sid string) Session {
	v, ok := m.data[sid]
	if !ok {
		return nil
	}
	return v
}

func (m *MemStore) Delete(sid string) error {
	_, ok := m.data[sid]
	if !ok {
		return errors.New("no session with given id")
	}
	delete(m.data, sid)
	return nil
}

func (m *MemStore) Expire() {
	for k := range m.data {
		if m.data[k].Expired(m.expiry) {
			delete(m.data, k)
		}
	}
}
