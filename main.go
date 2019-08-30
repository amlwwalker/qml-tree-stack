package main

import (
	"os"
	"sync"

	_ "github.com/amlwwalker/example-qml/TreeView/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

type Controller struct {
	core.QObject
	qApp *widgets.QApplication
}

var once sync.Once
var instance *Controller

func Instance() *Controller {
	once.Do(func() {
		instance = NewController(nil)
	})
	return instance
}

func main() {

	qApp := widgets.NewQApplication(len(os.Args), os.Args)
	Instance()
	Instance().qApp = qApp

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := qml.NewQQmlApplicationEngine(nil)
	app.AddImportPath("./tree/qml")
	app.Load(core.NewQUrl3("./main.qml", 0))
	widgets.QApplication_Exec()
}
