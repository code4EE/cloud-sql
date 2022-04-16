package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// 定义全局变量
var (
	srvInstance *ServerImpl = &ServerImpl{handlers: make(map[string]http.Handler)}
)

// 定义Server接口,这一步是为了优雅关闭server,见https://www.jianshu.com/p/ee2d2c6eb40d
type Server interface {
	ListenAndServe() error
	ShutDown(context.Context) error
}

type ServerImpl struct {
	mux      *http.ServeMux
	server   *http.Server
	handlers map[string]http.Handler
}

// InitServer 表示初始化Server
func InitServer(addr string) error {
	// TODO：检查addr是否合法
	srvInstance.mux = http.NewServeMux()
	srvInstance.server = &http.Server{
		Addr: addr,
	}
	return nil
}

// RegisterHandler 可以将用户自定义的handler注册到server里面
func RegisterHandler(pattern string, handler http.Handler) error {
	if !strings.HasPrefix(pattern, "/") {
		return fmt.Errorf("[Error] pattern must start with '/'")
	}
	srvInstance.handlers[pattern] = handler
	return nil
}

func RunServer(ctx context.Context) {
	log.Printf("Server is Running...")
	// 拆开多个handler
	for k, v := range srvInstance.handlers {
		srvInstance.mux.Handle(k, v)
	}
	srvInstance.server.Handler = srvInstance.mux
	var err error
	go func() {
		if err = srvInstance.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe:%+s", err)
		}
	}()

	// 如果ctx.Done()这个channel有值了代表上下文被取消了, 需要退出Server
	<-ctx.Done()
	log.Printf("Server is stopping...")
	ctxShutDown, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer func() {
		cancel()
	}()
	if err = srvInstance.server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server shutdown: %+s", err)
	}
	log.Printf("server shutdown gracefully")
	if err == http.ErrServerClosed {
		err = nil
	}
}
