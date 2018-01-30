package bridge

import (
	"github.com/therecipe/qt/qml"
)

func Init(context *qml.QQmlContext) {
	initConfigure(context)
	initProxy(context)
}
