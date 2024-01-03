package stores

import (
	"github.com/aadi-1024/sessions/session"
	"time"
)

type Store interface {
	//Save stores a session
	Save(sid string, created time.Time) error

	//Load loads a session from the store, returns nil if not present
	Load(sid string) session.Session

	Delete(sid string) error

	Expire(duration time.Duration)
}
