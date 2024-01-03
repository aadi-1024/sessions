package session

import (
	"errors"
	"time"
)

type Session interface {
	Set(key, val string)
	Get(key string) (string, error)
	Delete(key string)
	Expired(exp time.Duration) bool
}

type MapSession struct {
	fields  map[string]string
	Sid     string
	created time.Time
}

func NewMapSession(sid string, created time.Time) *MapSession {
	return &MapSession{
		fields:  make(map[string]string),
		Sid:     sid,
		created: created,
	}
}

func (s *MapSession) Set(key, val string) {
	s.fields[key] = val
}

func (s *MapSession) Get(key string) (string, error) {
	val, ok := s.fields[key]
	if !ok {
		return "", errors.New("no matching field found")
	}
	return val, nil
}

func (s *MapSession) Delete(key string) {
	_, ok := s.fields[key]
	if ok {
		delete(s.fields, key)
	}
}

func (s *MapSession) Expired(exp time.Duration) bool {
	return s.created.Add(exp).Compare(time.Now()) < 1
}
