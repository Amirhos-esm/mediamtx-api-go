package main

import (
	"fmt"

	mtx "github.com/Amirhos-esm/mediamtx-api-go/mediamtx"
)

func main() {

	api := mtx.CreateMtxApi("http://localhost:9997")

	paths, err := api.GetAllPath(0, 100)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(paths)

	cfg, err := api.GetDefaultPathConfiguration()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(cfg)
	/*
		err := api.RegisterHookCallback(mtx.HOOK_runOnReady, false,
			"path=$MTX_PATH&source_id=$MTX_SOURCE_ID&source_type=$MTX_SOURCE_TYPE",
			func(hook mtx.HookType, data map[string]any) {
				fmt.Println("Hook triggered:", hook, "Data:", data)
			})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = api.RegisterHookCallback(mtx.HOOK_runOnConnect, false,
			"path=$MTX_PATH&source_id=$MTX_SOURCE_ID&source_type=$MTX_SOURCE_TYPE",
			func(hook mtx.HookType, data map[string]any) {
				fmt.Println("Hook triggered:", hook, "Data:", data)
			})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = api.RegisterHookCallback(mtx.HOOK_runOnRead, false,
			"path=$MTX_PATH&source_id=$MTX_SOURCE_ID&source_type=$MTX_SOURCE_TYPE",
			func(hook mtx.HookType, data map[string]any) {
				fmt.Println("Hook triggered:", hook, "Data:", data)
			})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = api.RegisterHookCallback(mtx.HOOK_runOnRecordSegmentCreate, false,
			"path=$MTX_PATH&segment_path=$MTX_SEGMENT_PATH",
			func(hook mtx.HookType, data map[string]any) {
				fmt.Println("Hook triggered:", hook, "Data:", data)
			})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		err = api.RegisterHookCallback(mtx.HOOK_runOnRecordSegmentComplete, false,
			"path=$MTX_PATH&segment_path=$MTX_SEGMENT_PATH",
			func(hook mtx.HookType, data map[string]any) {
				fmt.Println("Hook triggered:", hook, "Data:", data)
			})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		err = api.RegisterHookCallback(mtx.HOOK_runOnNotReady, false,
			"path=$MTX_PATH&source_id=$MTX_SOURCE_ID&source_type=$MTX_SOURCE_TYPE",
			func(hook mtx.HookType, data map[string]any) {
				fmt.Println("Hook triggered:", hook, "Data:", data)
			})
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		api.AddAuthenticationCallback(func(authData *mtx.AuthenticationData) bool {
			fmt.Println("Authentication Data:", authData)
			return true
		})

		err = api.RunServer("localhost:10000")
		if err != nil {
			panic(err)
		}
	*/
}
