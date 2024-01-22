package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisRateLimiter struct {
	host                     string
	port                     string
	db                       int
	password                 string
	rateLimitTTL             int
	maxRequestIPSecond       int
	maxRequestTokenSecond    int
	requestBlockingTimeIP    int
	requestBlockingTimeToken int
	rateLimiterIPEnabled     bool
	rateLimiterTokenEnabled  bool
	redisClient              *redis.Client
}

func NewRedisRateLimiter(
	host string,
	port string,
	db int,
	password string,
	rateLimitTTL int,
	maxRequestIPSecond int,
	maxRequestTokenSecond int,
	RequestBlockingTimeIP int,
	RequestBlockingTimeToken int,
	RateLimiterIPEnabled bool,
	RateLimiterTokenEnabled bool,
) *redisRateLimiter {
	fmt.Printf("%s:%s /n", host, port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	return &redisRateLimiter{
		host:                     host,
		port:                     port,
		db:                       db,
		password:                 password,
		rateLimitTTL:             rateLimitTTL,
		maxRequestIPSecond:       maxRequestIPSecond,
		maxRequestTokenSecond:    maxRequestTokenSecond,
		requestBlockingTimeIP:    RequestBlockingTimeIP,
		requestBlockingTimeToken: RequestBlockingTimeToken,
		rateLimiterIPEnabled:     RateLimiterIPEnabled,
		rateLimiterTokenEnabled:  RateLimiterTokenEnabled,
		redisClient:              redisClient,
	}
}

func (r *redisRateLimiter) isBlocked(ip, token string) bool {

	if token != "" {
		return r.isBlockedToken(token)
	}

	return r.isBlockedIP(ip)
}

func (r *redisRateLimiter) isBlockedIP(ip string) bool {
	result, err := r.redisClient.Keys(context.Background(), fmt.Sprintf("ip:%v:blocked", ip)).Result()
	if err != nil {
		log.Fatalln(err)
		return false
	}

	if len(result) > 0 {
		return true
	}

	return false
}

func (r *redisRateLimiter) isBlockedToken(token string) bool {
	result, err := r.redisClient.Keys(context.Background(), fmt.Sprintf("token:%v:blocked", token)).Result()
	if err != nil {
		log.Fatalln(err)
		return false
	}

	if len(result) > 0 {
		return true
	}

	return false
}

func (r *redisRateLimiter) BlockAccess(ip, token string) error {
	if token != "" {
		r.redisClient.Set(context.Background(), fmt.Sprintf("token:%v:blocked", token), 1, time.Duration(r.rateLimitTTL)*time.Second)
	}
	return nil
}

func (r *redisRateLimiter) isIpRateLimitExceeded(ip string) bool {
	result, err := r.redisClient.Keys(context.Background(), fmt.Sprintf("ip:%v:*", ip)).Result()
	if err != nil {
		log.Fatalln(err)
		return false
	}

	if len(result) >= r.maxRequestIPSecond {
		return true
	}

	return false
}

func (r *redisRateLimiter) isTokenRateLimitExceeded(token string) bool {
	result, err := r.redisClient.Keys(context.Background(), fmt.Sprintf("token:%v:*", token)).Result()
	if err != nil {
		log.Fatalln(err)
		return false
	}

	if len(result) >= r.maxRequestTokenSecond {
		return true
	}

	return false
}

func (r *redisRateLimiter) IsRateLimitExceeded(ip string, token string) bool {

	var isRateLimitExceeded bool

	if r.isBlocked(ip, token) {
		return true
	}

	if token != "" && r.rateLimiterTokenEnabled {
		isRateLimitExceeded = r.isTokenRateLimitExceeded(token)
	}

	if !isRateLimitExceeded && r.rateLimiterIPEnabled {
		isRateLimitExceeded = r.isIpRateLimitExceeded(ip)
	}

	if isRateLimitExceeded {
		r.BlockAccess(ip, token)
	}

	return isRateLimitExceeded

}

func (r *redisRateLimiter) registerIpAcess(ip string) error {
	r.redisClient.Set(context.Background(), fmt.Sprintf("ip:%v:%v", ip, time.Now().Format("20060102150405.000")), 1, time.Duration(r.rateLimitTTL)*time.Second)
	return nil
}

func (r *redisRateLimiter) registerTokenAcess(token string) error {
	r.redisClient.Set(context.Background(), fmt.Sprintf("token:%v:%v", token, time.Now().Format("20060102150405.000")), 1, time.Duration(r.rateLimitTTL)*time.Second)
	return nil
}

func (r *redisRateLimiter) RegisterAccess(ip string, token string) error {

	if token != "" {
		return r.registerTokenAcess(token)
	}

	return r.registerIpAcess(ip)

}
