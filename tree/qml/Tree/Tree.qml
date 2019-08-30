import QtQuick 2.6
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4
import Tree 1.0

TreeController {
width: 800
height: 400
  TreeView {
        id: treeView
        anchors.fill: parent
        model: treeModel
        rowDelegate: Rectangle {
            id: rowDelegateRectangle
            height: 40
        }
        itemDelegate: Rectangle {
            id: delegateRectangle
            height: 40
            Text {
                anchors.verticalCenter: parent.verticalCenter
                anchors.left: parent.left
                color: "hotpink"
                text: styleData.value 
            }
            MouseArea {
                anchors.fill: parent
                onClicked: {
                    if (styleData.hasChildren) {
                        console.log(styleData.value + " has children")
                    }
                }
            }
        }
        style: TreeViewStyle {
            branchDelegate: Rectangle {
                width: 15; height: 15
                color: "#00000000"
                // anchors.right: treeView.right

                Image {
                    id: expandArrow
                    source: styleData.isExpanded ? "./images/close-round.png" : "./images/chevron-right.png"
                    sourceSize.width: parent.width
                    sourceSize.height: parent.height
                }

                // ColorOverlay {
                //     anchors.fill: expandArrow
                //     source: expandArrow
                //     color: "#293147"
                // }
            }
        }
        TableViewColumn {
            role: "display"
        }
        // TableViewColumn {
        //     role: "name"
        // }
    }
}
