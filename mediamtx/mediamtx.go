package mediamtx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type mediamtx struct {
	hookQueue   chan hookQueueType
	hookBaseUrl *url.URL
	baseAddress string
	mutex       sync.Mutex
	host        string
	callbacks   map[HookType]HookCallback
	server http.Server
}

func CreateMtxApi(mtx_addr string, hookBaseUrl *url.URL) *mediamtx {
	return &mediamtx{
		hookBaseUrl: hookBaseUrl,
		baseAddress: mtx_addr,
		callbacks:   make(map[HookType]HookCallback),
	}
}

func (mtx *mediamtx) RunServer(host string) error {
	mtx.host = host
	mtx.hookQueue = make(chan hookQueueType)

	mux := http.NewServeMux()

	for hook := range hookTypeToString {
		mux.HandleFunc("/"+hook.String(), func(w http.ResponseWriter, r *http.Request) {
			datas := make(map[string]any)
			for key, values := range r.URL.Query() {
				if len(values) > 0 {
					datas[key] = values[0]
				}
			}
			select {
			case mtx.hookQueue <- hookQueueType{
				datas: datas,
				t:     hook,
			}:
			default:
				// If the channel is full, we drop the request to avoid blocking
				// the server. This is a simple way to handle backpressure.
				return
			}
		})
	}

	go func() {
		for {
			data, ok := <-mtx.hookQueue
			if !ok {
				return
			}
			{
				mtx.mutex.Lock()
				cb := mtx.callbacks[data.t]
				mtx.mutex.Unlock()

				if cb != nil {
					cb(data.t, data.datas)

				}
			}
		}
	}()
	mtx.server.Addr = host
	mtx.server.Handler = mux
	return mtx.server.ListenAndServe()
}

func (mtx *mediamtx) StopServer() {
	close(mtx.hookQueue)
	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := mtx.server.Shutdown(ctx); err != nil {
		fmt.Println("Shutdown error:", err)
	} else {
		fmt.Println("Server stopped gracefully")
	}

}
func (mtx *mediamtx) RegisterHookCallback(hook HookType, restart bool, vars string, callback HookCallback) error {
	err := hook.Enable(vars, restart, mtx)
	if err != nil {
		return err
	}

	mtx.mutex.Lock()
	defer mtx.mutex.Unlock()
	mtx.callbacks[hook] = callback
	return nil
}
