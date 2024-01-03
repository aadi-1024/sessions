package session

import (
	"errors"
	"testing"
	"time"
)

// doesn't really make sense to test Expired

// any type which implements Session will work
var sessionObject = NewMapSession("sample", time.Now())

func init() {
	sessionObject.Set("a", "b")
	sessionObject.Set("c", "d")
}

func TestGet(t *testing.T) {
	table := []struct {
		name     string
		key      string
		expected string
		error    error
	}{
		{"should return", "a", "b", nil},
		{"should be error", "d", "", errors.New("no matching field found")},
	}

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			s, e := sessionObject.Get(v.key)
			if s != v.expected {
				t.Errorf("Failed test with key %v and returns %v %v", v.key, s, e)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	//have to make sure element gets deleted and nothing happens if no such element
	sessionObject.Set("tbd", "deleteme")
	sessionObject.Delete("tbd")

	if _, err := sessionObject.Get("tbd"); err == nil {
		t.Errorf("Element wasnt deleted")
	}
}
