package mediamtx

type Mediamtx struct {
	// hookQueue   chan hookQueueType
	baseAddress string
	// mutex       sync.Mutex
	// authMutex   sync.Mutex
	// authCb      AuthenticationCallback
	// host        string
	// callbacks   map[HookType]HookCallback
	// server http.Server
}

func (m Mediamtx) GetBaseAddress()string{
	return m.baseAddress
}
func CreateMtxApi(mtx_addr string) *Mediamtx {
	return &Mediamtx{
		baseAddress: mtx_addr,
		// callbacks:   make(map[HookType]HookCallback),
	}
}

// func (mtx *Mediamtx) Routes(router *gin.Engine) {
// 	// Register hook handlers
// 	for hook := range hookTypeToString {
// 		hookCopy := hook // capture loop variable
// 		router.GET("/"+hook.String(), func(c *gin.Context) {
// 			datas := make(map[string]any)
// 			for key, values := range c.Request.URL.Query() {
// 				if len(values) > 0 {
// 					datas[key] = values[0]
// 				}
// 			}
// 			select {
// 			case mtx.hookQueue <- hookQueueType{
// 				datas: datas,
// 				t:     hookCopy,
// 			}:
// 			default:
// 				// Drop request if channel full
// 				c.Status(204) // No Content
// 				return
// 			}
// 			c.Status(200)
// 		})
// 	}

// 	// Auth endpoint
// 	router.POST("/auth", func(c *gin.Context) {
// 		var authData AuthenticationData
// 		if err := c.ShouldBindJSON(&authData); err != nil {
// 			c.JSON(400, gin.H{"error": "Invalid request body"})
// 			return
// 		}

// 		mtx.authMutex.Lock()
// 		cb := mtx.authCb
// 		mtx.authMutex.Unlock()

// 		if cb != nil {
// 			if cb(&authData) {
// 				c.String(200, "Authorized")
// 			} else {
// 				c.String(401, "Unauthorized")
// 			}
// 		}
// 	})

// 	// Background goroutine to process hook queue
// 	go func() {
// 		for data := range mtx.hookQueue {
// 			mtx.mutex.Lock()
// 			cb := mtx.callbacks[data.t]
// 			mtx.mutex.Unlock()

// 			if cb != nil {
// 				cb(data.t, data.datas)
// 			}
// 		}
// 	}()

// }

// func (mtx *Mediamtx) RunServer(host string) error {
// 	mtx.host = host
// 	mtx.hookQueue = make(chan hookQueueType)

// 	// Create Gin router
// 	router := gin.Default()
// 	mtx.Routes(router)
// 	mtx.server.Addr = host
// 	mtx.server.Handler = router
// 	return mtx.server.ListenAndServe()
// }

// func (mtx *Mediamtx) StopServer() {
// 	close(mtx.hookQueue)

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := mtx.server.Shutdown(ctx); err != nil {
// 		fmt.Println("Shutdown error:", err)
// 	} else {
// 		fmt.Println("Server stopped gracefully")
// 	}
// }

// func (mtx *Mediamtx) RegisterHookCallback(hook HookType, restart bool, vars string, callback HookCallback) error {
// 	err := hook.Enable(vars, restart, mtx)
// 	if err != nil {
// 		return err
// 	}

// 	mtx.mutex.Lock()
// 	defer mtx.mutex.Unlock()
// 	mtx.callbacks[hook] = callback
// 	return nil
// }
