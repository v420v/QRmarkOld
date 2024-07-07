package common

import (
	"context"
	"errors"
)

type UserKey struct{}

type TraceIDKey struct{}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, TraceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(TraceIDKey{})
	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

func GetCurrentUserID(ctx context.Context) (int, error) {
	userID, isAuthorized := ctx.Value(UserKey{}).(int)

	if !isAuthorized {
		return userID, errors.New("Unauthorized")
	}

	return userID, nil
}
