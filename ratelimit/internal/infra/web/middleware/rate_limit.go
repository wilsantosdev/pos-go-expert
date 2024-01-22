package middleware

import (
	"encoding/json"
	"net"
	"net/http"
	"ratelimit/internal/entity"
)

func NewRateLimiter(limiter entity.Limiter) *rateLimiter {
	return &rateLimiter{
		Limiter: limiter,
	}
}

type rateLimiter struct {
	Limiter entity.Limiter
}

type rateLimiterErrorResponse struct {
	Message string `json:"message"`
}

func (rl *rateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get token from header
		token := r.Header.Get("API_KEY")

		// get ip from
		ip := r.Header.Get("X-FORWARDED-FOR")
		if ip == "" {
			ip = r.RemoteAddr
		}

		ip, _, _ = net.SplitHostPort(ip)

		if rl.Limiter.IsRateLimitExceeded(ip, token) {
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(rateLimiterErrorResponse{
				Message: "you have reached the maximum number of requests or actions allowed within a certain time frame",
			})
			return
		}

		rl.Limiter.RegisterAccess(ip, token)

		next.ServeHTTP(w, r)
	})
}
