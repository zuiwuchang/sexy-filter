package bridge

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"sexy-filter/net"
)

type bridgeProxy struct {
	core.QObject

	//測試代理
	_ func(style int, addr, user, pwd, testUrl string) int `slot:"testProxy"`
	_ func(id int, emsg string)                            `signal:"testProxyReply"`
}

func initProxy(context *qml.QQmlContext) {
	bridge := NewBridgeProxy(nil)

	id := 0
	bridge.ConnectTestProxy(func(style int, addr, user, pwd, testUrl string) (requestID int) {
		id++
		if id == 0 {
			id++
		}
		requestID = id

		go func() {
			e := net.TestProxy(style, addr, user, pwd, testUrl)
			if e != nil {
				emsg := "emsg : " + e.Error()
				bridge.TestProxyReply(requestID, emsg)
			} else {
				bridge.TestProxyReply(requestID, "")
			}
		}()

		return
	})

	context.SetContextProperty("BridgeProxy", bridge)
}
