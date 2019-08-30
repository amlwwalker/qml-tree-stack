import QtQuick 2.6
// import QtQuick.Controls 1.4
import QtQuick.Controls 2.4

import Tree 1.0

TreeController {
  ListView {
    id: view
    // Layout.fillHeight: true
    // Layout.fillWidth: true
    anchors.fill: parent
    model: treeModel
    clip: true
    snapMode: ListView.SnapToItem

    // property var rootIndex
    delegate: Rectangle {
        width: 200; height: 25
        Text { text: "hello" }

        MouseArea {
            anchors.fill: parent
            onClicked: {
                if (model.hasModelChildren)
                    view.model.rootIndex = view.model.modelIndex(index)
            }
        }
    }
  }
}
