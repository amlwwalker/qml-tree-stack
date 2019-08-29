import QtQuick.Controls 2.4 //Menu
import QtQuick.Dialogs 1.2  //MessageDialog
import QtQuick.Layouts 1.3  //GridLayout
import QtQuick 2.2

ListView {
    Layout.fillHeight: true
    Layout.fillWidth: true
    model: treeModel.rowCount(rootIndex)
    clip: true
    snapMode: ListView.SnapToItem

    property var rootIndex

    // I used a QFileSytemModel in my example, so I had to manually 
    // fetch data when the rootIndex changed. You may not need this though.
    onRootIndexChanged: {
        if(treeModel.canFetchMore(rootIndex))
            treeModel.fetchMore(rootIndex)
    }
    Connections {
        target: treeModel
        onRowsInserted: {
            rootIndexChanged()
        }
    }

    delegate: ItemDelegate {
        property var modelIndex: treeModel.index(index,0, rootIndex)
        property bool hasChildren: treeModel.hasChildren(modelIndex)
        width: parent.width
        text: treeModel.data(modelIndex)
        onClicked: {
            if(hasChildren)
            {
                // Recursively add another TreeSlide, with a new rootIndex
                stackView.push("TreeSlide.qml", {rootIndex: modelIndex})
            }
        }
    }
}