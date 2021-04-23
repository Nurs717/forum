package model

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

// Limiter limit
type Limiter struct {
	ipCount map[string]int
	sync.Mutex
}

var limiter Limiter

func init() {
	limiter.ipCount = make(map[string]int)
}

// Limit rate limit
func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the IP address for the current user.
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Get the # of times the visitor has visited in the last 60 seconds
		limiter.Lock()
		count, ok := limiter.ipCount[ip]
		if !ok {
			limiter.ipCount[ip] = 0
		}
		if count > 30 {
			limiter.Unlock()
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}
		if count <= 30 {
			limiter.ipCount[ip]++
		}
		time.AfterFunc(time.Second*10, func() {
			limiter.Lock()
			limiter.ipCount[ip]--
			limiter.Unlock()
		})
		if limiter.ipCount[ip] == 30 {
			// set it to 20 so the decrement timers will only decrease it to
			// 10, and they stay blocked until the next timer resets it to 0
			limiter.ipCount[ip] = 40
			time.AfterFunc(time.Hour, func() {
				limiter.Lock()
				limiter.ipCount[ip] = 0
				limiter.Unlock()
			})
		}
		fmt.Println(limiter.ipCount)
		limiter.Unlock()
		next.ServeHTTP(w, r)
	})
}
