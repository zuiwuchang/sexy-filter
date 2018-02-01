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
	context.SetContextProperty("BridgePlugins", bridge)
}
