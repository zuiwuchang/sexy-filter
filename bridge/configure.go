package bridge

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"sexy-filter/configure"
	"sexy-filter/log"
)

type bridgeConfigure struct {
	core.QObject

	//返回 語言 列表
	_ func() []string `slot:"getLocales"`
	//返回 當前 選擇的 語言
	_ func() int `slot:"getLocaleIndex"`

	//返回 風格 列表
	_ func() []string `slot:"getStyles"`
	//返回 當前 選擇的 風格
	_ func() int `slot:"getStyle"`

	//保持 語言 風格 設置
	_ func(locale, style int) `slot:"save2"`
}

func initConfigure(context *qml.QQmlContext) {
	bridge := NewBridgeConfigure(nil)

	bridge.ConnectGetLocales(func() []string {
		langs := configure.GetLocales()
		arrs := make([]string, len(langs))
		for i, key := range langs {
			name := configure.GetLocaleName(key)
			if name == "" {
				arrs[i] = key
			} else {
				arrs[i] = key + "   " + name
			}
		}
		return arrs
	})
	bridge.ConnectGetLocaleIndex(func() int {
		return configure.GetLocaleIndex()
	})
	bridge.ConnectGetStyles(func() []string {
		return configure.GetStyles()
	})
	bridge.ConnectGetStyle(func() int {
		return configure.GetStyle()
	})
	bridge.ConnectSave2(func(locale, style int) {
		e := configure.Verification2(locale, style)
		if e != nil {
			if log.Warn != nil {
				log.Warn.Println(e)
			}
			return
		}
		configure.Set2(locale, style)
		s := core.NewQSettings5(nil)
		s.SetValue("locale", core.NewQVariant14(configure.GetLocale()))
		s.SetValue("style", core.NewQVariant7(style))
		s.Sync()

	})
	context.SetContextProperty("BridgeConfigure", bridge)
}
