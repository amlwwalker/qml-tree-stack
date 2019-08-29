package main

import (
	"os"

	_ "github.com/amlwwalker/example-qml/NewTreeView/tree"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

func main() {

	widgets.NewQApplication(len(os.Args), os.Args)

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := qml.NewQQmlApplicationEngine(nil)
	app.AddImportPath("./tree/qml")
	app.Load(core.NewQUrl3("./main.qml", 0))
	widgets.QApplication_Exec()
}
