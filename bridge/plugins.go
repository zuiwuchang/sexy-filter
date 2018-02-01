package bridge

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"sexy-filter/js"
)

type bridgePlugins struct {
	core.QObject

	//返回 插件 檔案 名
	_ func() []string `slot:"getPluginsFiles"`

	//返回 測試 檔案 名
	_ func() []string `slot:"getTestFiles"`

	//測試 url 請求
	_ func(style int, addr, user, pwd, jsFile string) int `slot:"testUrl"`
	//測試 檔案 分析
	_ func(jsFile, testFile string) int `slot:"testFile"`
	//返回 測試請求 結果
	_ func(id int, emsg, msg string) `signal:"testReply"`

	//返回 當前 單件 狀態
	_ func() int `slot:"getStatus"`
	//返回 插件名稱 數組
	_ func() []string `slot:"getPlugins"`
	//單件 開始 工作 並返回 當前狀態
	_ func(pos int) int `slot:"start"`
	//單件 停止 工作 並返回 當前狀態
	_ func() int `slot:"stop"`
	//單件 已開始工作 或 已停止 通知
	_ func(val int) `signal:"statusChanged"`
}

func initPlugins(context *qml.QQmlContext) {
	bridge := NewBridgePlugins(nil)
	bridge.ConnectGetPluginsFiles(js.GetPluginsFiles)
	bridge.ConnectGetTestFiles(js.GetTestFiles)
	id := 0
	bridge.ConnectTestUrl(func(style int, addr, user, pwd, jsFile string) (requestID int) {
		id++
		requestID = id
		if requestID == 0 {
			id++
			requestID = id

		}
		go func() {
			msg, e := js.TestUrl(style,
				addr, user, pwd,
				jsFile,
			)
			if e == nil {
				bridge.TestReply(requestID, "", msg)
			} else {
				emsg := "error : " + e.Error()
				bridge.TestReply(requestID, emsg, "")
			}
		}()
		return
	})
	bridge.ConnectTestFile(func(jsFile, testFile string) (requestID int) {
		id++
		requestID = id
		if requestID == 0 {
			id++
			requestID = id

		}
		go func() {
			msg, e := js.TestFile(jsFile, testFile)
			if e == nil {
				bridge.TestReply(requestID, "", msg)
			} else {
				emsg := "error : " + e.Error()
				bridge.TestReply(requestID, emsg, "")
			}
		}()
		return
	})

	single := js.GetSingle()
	bridge.ConnectGetStatus(single.Status)
	bridge.ConnectGetPlugins(single.GetPlugins)
	bridge.ConnectStart(single.Start)
	bridge.ConnectStop(single.Stop)
	single.OnStatusChanged = func(val int) {
		bridge.StatusChanged(val)
	}

	context.SetContextProperty("BridgePlugins", bridge)
}
