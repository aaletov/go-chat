package httputil

import (
	"context"
	"net"
	"net/http"
)

const (
	connContextKey = "http-conn"
)

func SaveConnInContext(ctx context.Context, c net.Conn) (context.Context) {
	return context.WithValue(ctx, connContextKey, c)
}

func GetConn(r *http.Request) (net.Conn) {
	return r.Context().Value(connContextKey).(net.Conn)
}