package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"
	"log"
	"os"
)

func main() {
	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	translator := core.NewQTranslator(nil)
	if translator.Load2(core.NewQLocale(), "zh_TW", "_", "locale", ".qm") {
		core.QCoreApplication_InstallTranslator(translator)
	} else {
		log.Fatalln("cannot load translator", core.QLocale_System().Name())
	}

	quickcontrols2.QQuickStyle_SetStyle("material")

	engine := qml.NewQQmlApplicationEngine(nil)

	engine.Load(core.NewQUrl3("qml/main.qml", 0))

	gui.QGuiApplication_Exec()
}
