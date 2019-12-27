package proxy

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// Do clones an *http.Request and proxies it to the specified receiver, with timeout
func Do(r *http.Request, receiver string, timeout time.Duration) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	url := fmt.Sprintf("http://%s%s", receiver, r.RequestURI)
	newReq, _ := http.NewRequestWithContext(ctx, r.Method, url, r.Body)
	newReq.Header = r.Header
	if clientIP, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		if old, ok := r.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(old, ", ") + ", " + clientIP
		}
		newReq.Header.Set("X-Forwarded-For", clientIP)
	}
	return http.DefaultClient.Do(newReq)
}
