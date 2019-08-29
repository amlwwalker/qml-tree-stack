package tree

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func init() {
	fmt.Println("init TreeController")
	treeController_QmlRegisterType2("Tree", 1, 0, "TreeController")
}

type treeController struct {
	quick.QQuickItem
	_ func() `constructor:"init"`

	_ *core.QAbstractItemModel `property:"treeModel"`
}

func (f *treeController) init() {
	treeModel := gui.NewQStandardItemModel(nil)
	rootNode := treeModel.InvisibleRootItem()

	americaItem := gui.NewQStandardItem2("America")
	mexicoItem := gui.NewQStandardItem2("Canada")
	usaItem := gui.NewQStandardItem2("USA")
	bostonItem := gui.NewQStandardItem2("Boston")
	europeItem := gui.NewQStandardItem2("Europe")
	italyItem := gui.NewQStandardItem2("Italy")
	romeItem := gui.NewQStandardItem2("Rome")
	veronaItem := gui.NewQStandardItem2("Verona")

	rootNode.AppendRow2(americaItem)
	rootNode.AppendRow2(europeItem)
	americaItem.AppendRow2(mexicoItem)
	americaItem.AppendRow2(usaItem)
	usaItem.AppendRow2(bostonItem)
	europeItem.AppendRow2(italyItem)
	italyItem.AppendRow2(romeItem)
	italyItem.AppendRow2(veronaItem)
	fmt.Println("setting model")
	f.SetTreeModel(treeModel)
}
