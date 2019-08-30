import QtQuick.Controls 2.4 //Menu
import QtQuick.Dialogs 1.2  //MessageDialog
import QtQuick.Layouts 1.3  //GridLayout
import QtQuick 2.2

import Tree 1.0

ApplicationWindow {
  id: app
  visible: true
  title: "Tree View"
  minimumWidth: 1500; minimumHeight: 800
  Tree {
    width: 400
    // anchors.fill: parent
  }
    // StackView
  // {
  //     id: stackView
  //     Layout.fillHeight: true
  //     Layout.fillWidth: true
  //     initialItem: Tree {}
  // }
}
