import QtQuick 2.6
// import QtQuick.Controls 1.4
import QtQuick.Controls 2.4

import Tree 1.0

TreeController {
  ListView {
    // Layout.fillHeight: true
    // Layout.fillWidth: true
    model: treeModel
    clip: true
    snapMode: ListView.SnapToItem

    property var rootIndex

    delegate: ItemDelegate {
        // property var modelIndex: treeModel.index(index,0, rootIndex)
        // property bool hasChildren: treeModel.hasChildren(modelIndex)
        width: parent.width
        text: "hello" //treeModel.data(modelIndex)
        onClicked: {
            // if(hasChildren)
            // {
            //     // Recursively add another TreeSlide, with a new rootIndex
            //     // stackView.push("TreeSlide.qml", {rootIndex: modelIndex})
            // }
        }
    }
  }
}
