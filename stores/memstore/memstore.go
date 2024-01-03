package memstore

import (
	"errors"
	"github.com/aadi-1024/sessions/session"
	"time"
)

type MemStore struct {
	data map[string]session.Session
}

func NewMemStore() *MemStore {
	m := &MemStore{
		data: make(map[string]session.Session),
	}
	return m
}

func (m *MemStore) Save(sid string, created time.Time) error {
	_, ok := m.data[sid]
	if !ok {
		m.data[sid] = session.NewMapSession(sid, created)
		return nil
	}
	return errors.New("session with given id already exists")
}

func (m *MemStore) Load(sid string) session.Session {
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

func (m *MemStore) Expire(dur time.Duration) {
	for k := range m.data {
		if m.data[k].Expired(dur) {
			delete(m.data, k)
		}
	}
}
