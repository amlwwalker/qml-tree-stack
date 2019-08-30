package tree

import (
	"fmt"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"
)

func init() {
	treeController_QmlRegisterType2("Tree", 1, 0, "TreeController")
	fmt.Println("general intialisation")
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
	canadaItem := gui.NewQStandardItem2("Canada")

	usaItem := gui.NewQStandardItem2("USA")
	bostonItem := gui.NewQStandardItem2("Boston")
	europeItem := gui.NewQStandardItem2("Europe")
	italyItem := gui.NewQStandardItem2("Italy")
	romeItem := gui.NewQStandardItem2("Rome")
	veronaItem := gui.NewQStandardItem2("Verona")

	type City struct {
		*gui.QStandardItem
		Name       string `property:"name"`
		Population int64
	}
	// toronto := &City{gui.NewQStandardItem(), "Toronto", 1200}
	//canadian cities
	toronto := gui.NewQStandardItem2("Toronto")
	montreal := gui.NewQStandardItem2("Montreal")
	quebeck := gui.NewQStandardItem2("Quebeck")

	canadaItem.AppendRow2(montreal)
	canadaItem.AppendRow2(quebeck)
	canadaItem.AppendRow2(toronto)

	rootNode.AppendRow2(americaItem)
	rootNode.AppendRow2(europeItem)
	americaItem.AppendRow2(canadaItem)
	americaItem.AppendRow2(usaItem)
	usaItem.AppendRow2(bostonItem)
	europeItem.AppendRow2(italyItem)
	italyItem.AppendRow2(romeItem)
	italyItem.AppendRow2(veronaItem)

	f.SetTreeModel(treeModel)
	fmt.Println("treeController initialised")
}
