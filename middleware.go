package sessions

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"time"
)

// SessionsMiddleware manages the session cookie and the session stores
// if the request contains a cookie with a valid session id, the session
// is loaded and added to the request's Context. Otherwise, a new session
// and cookie get generated.
func SessionsMiddleware(s *SessionManager) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			cookie, err := r.Cookie(s.Name)
			rWithSession := r

			if err == http.ErrNoCookie {
				sid, err := s.NewSession()
				sidb := []byte(sid)
				if err != nil {
					log.Println(err)
				} else {
					b := make([]byte, base64.RawStdEncoding.EncodedLen(len(sidb)))

					base64.RawStdEncoding.Encode(b, sidb)

					http.SetCookie(w, &http.Cookie{
						Name:    s.Name,
						Value:   string(b),
						Expires: time.Now().Add(s.Expiry),
					})
					rWithSession = r.WithContext(context.WithValue(r.Context(), s.Name, s.store.Load(sid)))
				}
			} else if err != nil {
				log.Println(err)
			} else {
				//add the session to context
				b := make([]byte, base64.RawStdEncoding.DecodedLen(len([]byte(cookie.Value))))
				_, err := base64.RawStdEncoding.Decode(b, []byte(cookie.Value))
				if err != nil {
					fmt.Println(err)
				}

				rWithSession = r.WithContext(context.WithValue(r.Context(), s.Name, s.store.Load(string(b))))
			}

			next.ServeHTTP(w, rWithSession)
		})
	}
}
