package mediamtx

// import "errors"

// type HookCallback func(HookType, map[string]any)
// type HookType int

// const (
// 	HOOK_runOnConnect HookType = iota
// 	HOOK_runOnDisconnect
// 	HOOK_runOnInit
// 	HOOK_runOnDemand
// 	HOOK_runOnUnDemand
// 	HOOK_runOnReady
// 	HOOK_runOnNotReady
// 	HOOK_runOnRead
// 	HOOK_runOnUnread
// 	HOOK_runOnRecordSegmentCreate
// 	HOOK_runOnRecordSegmentComplete
// )

// var hookTypeToString = map[HookType]string{
// 	HOOK_runOnConnect:               "runOnConnect",
// 	HOOK_runOnDisconnect:            "runOnDisconnect",
// 	HOOK_runOnInit:                  "runOnInit",
// 	HOOK_runOnDemand:                "runOnDemand",
// 	HOOK_runOnUnDemand:              "runOnUnDemand",
// 	HOOK_runOnReady:                 "runOnReady",
// 	HOOK_runOnNotReady:              "runOnNotReady",
// 	HOOK_runOnRead:                  "runOnRead",
// 	HOOK_runOnUnread:                "runOnUnread",
// 	HOOK_runOnRecordSegmentCreate:   "runOnRecordSegmentCreate",
// 	HOOK_runOnRecordSegmentComplete: "runOnRecordSegmentComplete",
// }

// type hookQueueType struct {
// 	datas map[string]any
// 	t     HookType
// }

// func (ht HookType) String() string {
// 	if str, ok := hookTypeToString[ht]; ok {
// 		return str
// 	}
// 	return "unknown"
// }

// func (ht HookType) Enable(vars string, restart bool, mtx *Mediamtx) error {

// 	queries := ""
// 	if len(vars) > 0 {
// 		queries = "?"
// 		queries += vars
// 	}
// 	url := "curl " + mtx.HookBaseUrl + "/" + ht.String() + queries

// 	switch ht {
// 	case HOOK_runOnConnect:
// 		return mtx.PatchGlobalConfiguration(map[string]any{
// 			ht.String():             url,
// 			ht.String() + "Restart": restart,
// 		})
// 	case HOOK_runOnDisconnect:
// 		return mtx.PatchGlobalConfiguration(map[string]any{
// 			ht.String(): url,
// 		})
// 	case HOOK_runOnInit:
// 		return errors.New("runOnInit need path")
// 	case HOOK_runOnDemand, HOOK_runOnReady, HOOK_runOnRead:
// 		return mtx.PatchDefaultPathConfiguration(map[string]any{
// 			ht.String():             url,
// 			ht.String() + "Restart": restart,
// 		})
// 	case HOOK_runOnUnDemand, HOOK_runOnNotReady, HOOK_runOnUnread,
// 		HOOK_runOnRecordSegmentComplete, HOOK_runOnRecordSegmentCreate:
// 		return mtx.PatchDefaultPathConfiguration(map[string]any{
// 			ht.String(): url,
// 		})
// 	default:
// 		return errors.New("unsupported hook")

// 	}
// 	return nil
// }
