package entity

type Limiter interface {
	IsRateLimitExceeded(ip string, token string) bool
	RegisterAccess(ip string, token string) error
	BlockAccess(ip string, token string) error
}
