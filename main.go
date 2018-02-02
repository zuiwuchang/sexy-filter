package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
	"os"
	"path/filepath"
	"sexy-filter/bridge"
	"sexy-filter/configure"
	"sexy-filter/js"
	kLog "sexy-filter/log"
)

func main() {
	configure.Init(os.Args[0])
	js.InitSingle()
	if e := js.InitDB(); e != nil {
		os.Exit(-1)
		return
	}

	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.QGuiApplication_SetWindowIcon(gui.NewQIcon5(":/views/images/sexy.ico"))

	app.SetOrganizationName("cerberus")
	app.SetOrganizationDomain("doc.king011.com")
	app.SetApplicationName("go-qt-sexy-filter")

	s := core.NewQSettings5(nil)
	//初始 語言 配置
	local := s.Value("locale", core.NewQVariant14("zh_TW")).ToString()
	configure.InitLocale(local)
	local = configure.GetLocale()
	//初始 界面 配置
	style := s.Value("style", core.NewQVariant7(0)).ToInt(false)
	configure.InitStyle(style)
	style = configure.GetStyle()

	if local == "zh_TW" {
		translator := core.NewQTranslator(nil)
		if translator.Load2(core.NewQLocale(), "zh_TW", "_", ":/locale", ".qm") {
			core.QCoreApplication_InstallTranslator(translator)
		} else {
			if kLog.Warn != nil {
				kLog.Warn.Println("cannot load translator", core.QLocale_System().Name())
			}
		}
	} else {
		dir, e := filepath.Abs(os.Args[0])
		if e == nil {
			dir = filepath.Dir(dir) + "/locale"
		} else {
			dir = "lolcale"

			if kLog.Warn != nil {
				kLog.Warn.Println(e)
			}
		}

		translator := core.NewQTranslator(nil)
		if translator.Load2(core.NewQLocale(), configure.GetLocale(), "_", dir, ".qm") {
			core.QCoreApplication_InstallTranslator(translator)
		} else {
			if kLog.Warn != nil {
				kLog.Warn.Println(dir)
				kLog.Warn.Println("cannot load translator", core.QLocale_System().Name())
			}

			if translator.Load2(core.NewQLocale(), "zh_TW", "_", ":/locale", ".qm") {
				core.QCoreApplication_InstallTranslator(translator)
			} else {
				if kLog.Warn != nil {
					kLog.Warn.Println("cannot load translator", core.QLocale_System().Name())
				}
			}
		}
	}
	if style < 2 {
		quickcontrols2.QQuickStyle_SetStyle("Material")
	} else {
		quickcontrols2.QQuickStyle_SetStyle("Universal")
	}

	engine := qml.NewQQmlApplicationEngine(nil)
	bridge.Init(engine.RootContext())

	//engine.Load(core.NewQUrl3("qrc:/views/main.qml", 0))
	engine.Load(core.NewQUrl3("views/main.qml", 0))
	gui.QGuiApplication_Exec()
}
