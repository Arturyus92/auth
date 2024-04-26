package interceptor

import (
	"context"

	"github.com/Arturyus92/auth/internal/rate_limiter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RateLimiterInterceptor ...
type RateLimiterInterceptor struct {
	rateLimiter *rate_limiter.TokenBucketLimiter
}

// NewLimiterInterceptor ...
func NewLimiterInterceptor(rateLimiter *rate_limiter.TokenBucketLimiter) *RateLimiterInterceptor {
	return &RateLimiterInterceptor{rateLimiter: rateLimiter}
}

// Unary ...
func (r *RateLimiterInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !r.rateLimiter.Allow() {
		return nil, status.Error(codes.ResourceExhausted, "too many requests")
	}

	return handler(ctx, req)
}
