package bridge

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"sexy-filter/js"
)

type bridgeNode struct {
	core.QObject
	_ string `property:"title"`
	_ string `property:"url"`
	_ string `property:"pluginsId"`
	_ string `property:"pluginsName"`
}
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
	_ func(id int, emsg string) `signal:"testReply"`
	//返回 測試 model
	_ func() []*core.QObject `slot:"getModelTest"`
	//測試 model 已經改變
	_ func() `signal:"modelTestChanged"`

	//返回 當前 單件 狀態
	_ func() int `slot:"getStatus"`
	//返回 插件名稱 數組
	_ func() []string `slot:"getPlugins"`
	//單件 開始 工作 並返回 當前狀態
	_ func(style int, addr, user, pwd string, pos int) int `slot:"start"`
	//單件 停止 工作 並返回 當前狀態
	_ func() int `slot:"stop"`
	//單件 已開始工作 或 已停止 通知
	_ func(val int) `signal:"statusChanged"`

	//返回 查找 model
	_ func() []*core.QObject `slot:"getModelSearch"`
	//查找 model 已經改變
	_ func() `signal:"modelSearchChanged"`
	//查找 數據
	_ func(title, name, id, limit string) `slot:"search"`

	_ func() `constructor:"init"`

	modelTest    *wrapperNode
	tmpModelTest []*js.Node
	_            func() `signal:"qtModelTestChanged"`

	modelSearch    *wrapperNode
	tmpModelSearch []*js.Node
	_              func() `signal:"qtModelSearchChanged"`
}

func (b *bridgePlugins) init() {
	b.modelTest = newWrapperNode()
	b.modelSearch = newWrapperNode()
	b.ConnectDestroyBridgePlugins(func() {
		b.clearModelTest()
		b.clearModelSearch()
	})
	b.ConnectQtModelTestChanged(func() {
		if len(b.tmpModelTest) < 0 {
			return
		}
		b.modelTest.Append(b.tmpModelTest)
		b.tmpModelTest = nil

		//emit
		b.ModelTestChanged()
	})
	b.ConnectQtModelSearchChanged(func() {
		if len(b.tmpModelSearch) < 0 {
			return
		}

		b.modelSearch.Append(b.tmpModelSearch)
		b.tmpModelSearch = nil

		//emit
		b.ModelSearchChanged()
	})
}
func (b *bridgePlugins) getQtModelTest() []*core.QObject {
	return b.modelTest.qtArrs
}
func (b *bridgePlugins) getQtModelSearch() []*core.QObject {
	return b.modelSearch.qtArrs
}

func (b *bridgePlugins) clearModelTest() {
	m := b.modelTest
	if len(m.arrs) == 0 {
		return
	}
	m.Clear()

	//emit
	b.ModelTestChanged()
}
func (b *bridgePlugins) clearModelSearch() {
	m := b.modelSearch
	if len(m.arrs) == 0 {
		return
	}
	m.Clear()

	//emit
	b.ModelSearchChanged()
}
func (b *bridgePlugins) appendModelTest(nodes []*js.Node) {
	if len(nodes) == 0 {
		return
	}

	b.tmpModelTest = nodes
	b.QtModelTestChanged()
}
func (b *bridgePlugins) appendModelSearch(nodes []*js.Node) {
	if len(nodes) == 0 {
		return
	}

	b.tmpModelSearch = nodes
	b.QtModelSearchChanged()
}

type wrapperNode struct {
	qtArrs []*core.QObject
	arrs   []*bridgeNode
}

func newWrapperNode() *wrapperNode {
	return &wrapperNode{
		qtArrs: make([]*core.QObject, 0, 100),
		arrs:   make([]*bridgeNode, 0, 100),
	}
}
func (w *wrapperNode) Clear() (changed bool) {
	if len(w.arrs) == 0 {
		return
	}
	changed = true

	for _, p := range w.arrs {
		p.DestroyBridgeNode()
	}
	w.qtArrs = w.qtArrs[:0]
	w.arrs = w.arrs[:0]
	return
}
func (w *wrapperNode) Append(nodes []*js.Node) {
	for _, node := range nodes {
		nw := NewBridgeNode(nil)
		nw.SetTitle(node.Title)
		nw.SetUrl(node.Url)
		nw.SetPluginsId(node.PluginsId)
		nw.SetPluginsName(node.PluginsName)

		w.arrs = append(w.arrs, nw)
		w.qtArrs = append(w.qtArrs, core.NewQObjectFromPointer(nw.Pointer()))
	}
}

func initPlugins(context *qml.QQmlContext) {
	bridge := NewBridgePlugins(nil)

	bridge.ConnectGetPluginsFiles(js.GetPluginsFiles)
	bridge.ConnectGetTestFiles(js.GetTestFiles)
	id := 0
	bridge.ConnectTestUrl(func(style int, addr, user, pwd, jsFile string) (requestID int) {
		bridge.clearModelTest()

		id++
		requestID = id
		if requestID == 0 {
			id++
			requestID = id

		}
		go func() {
			nodes, e := js.TestUrl(style,
				addr, user, pwd,
				jsFile,
			)
			if e == nil {
				bridge.TestReply(requestID, "")
				bridge.appendModelTest(nodes)
			} else {
				emsg := "error : " + e.Error()
				bridge.TestReply(requestID, emsg)
			}
		}()
		return
	})
	bridge.ConnectTestFile(func(jsFile, testFile string) (requestID int) {
		bridge.clearModelTest()

		id++
		requestID = id
		if requestID == 0 {
			id++
			requestID = id

		}
		go func() {
			nodes, e := js.TestFile(jsFile, testFile)
			if e == nil {
				bridge.TestReply(requestID, "")
				bridge.appendModelTest(nodes)
			} else {
				emsg := "error : " + e.Error()
				bridge.TestReply(requestID, emsg)
			}
		}()
		return
	})
	bridge.ConnectGetModelTest(bridge.getQtModelTest)

	single := js.GetSingle()
	bridge.ConnectGetStatus(single.Status)
	bridge.ConnectGetPlugins(single.GetPlugins)
	bridge.ConnectStart(single.Start)
	bridge.ConnectStop(single.Stop)
	single.OnStatusChanged = func(val int) {
		bridge.StatusChanged(val)
	}

	bridge.ConnectSearch(func(title, name, id, limit string) {
		bridge.clearModelSearch()

		nodes := js.Search(title, name, id, limit)
		bridge.appendModelSearch(nodes)
	})
	bridge.ConnectGetModelSearch(bridge.getQtModelSearch)

	context.SetContextProperty("BridgePlugins", bridge)
}
