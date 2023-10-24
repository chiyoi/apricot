package neko

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"

	"github.com/chiyoi/apricot/logs"
)

var (
	ShutdownTimeout = time.Second * 5

	ErrWildcardPatternNeeded = errors.New("wildcard pattern needed")

	ls = func() *logs.Loggers {
		ls := logs.NewLoggers()
		ls.Prefix("[neko]")
		return ls
	}()
)

func IsWildcard(pattern string) bool {
	return strings.HasSuffix(pattern, "/")
}

func TrimPattern(path string, pattern string) string {
	return strings.TrimPrefix(path, pattern)
}

func PathResolver(pattern string) (resolve func(pattern string) string) {
	return func(p string) string {
		if IsWildcard(p) && !(pattern == "/" && p == "/") {
			return path.Join(pattern, p) + "/"
		}
		return path.Join(pattern, p)
	}
}

func AllowCrossOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", "*")
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func StartServer(srv *http.Server, tls bool) {
	ls.Info(fmt.Sprintln("Start server.", "srv.Addr:", srv.Addr))
	switch err := func() error {
		if tls {
			return srv.ListenAndServeTLS("", "")
		}
		return srv.ListenAndServe()
	}(); err {
	case http.ErrServerClosed:
		ls.Info(fmt.Sprintln("Stop server.", "srv.Addr:", srv.Addr))
	default:
		ls.Panic(err)
	}
}

func StopServer(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		ls.Error(err)
		return
	}
}

func Block() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	ls.Info("Stop:", <-sig)
}
