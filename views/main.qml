import QtQuick 2.7
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.0
import QtQuick.Controls.Universal 2.0
import QtQuick.Layouts 1.3


ApplicationWindow {
  id: window
  visible: true
  title: qsTr("sexy filter")
  width: 800
  height: 600
  minimumWidth: 640
  minimumHeight: 480

  Material.theme: BridgeConfigure.getStyle()==0?"Dark":"Light"
  Universal.theme: BridgeConfigure.getStyle()==2?"Dark":"Light"

  header: ToolBar {
      RowLayout {
          spacing: 20
          anchors.fill: parent

          ToolButton {
              contentItem: Image {
                  fillMode: Image.Pad
                  horizontalAlignment: Image.AlignHCenter
                  verticalAlignment: Image.AlignVCenter
                  source: stackView.depth > 1 ? "images/back.png" : "images/drawer.png"
              }
              onClicked: {
                  if (stackView.depth > 1) {
                      stackView.pop()
                      listView.currentIndex = -1
                  } else {
                      drawer.open()
                  }
              }
          }

          Label {
              id: titleLabel
              text: listView.currentItem ? listView.currentItem.text : qsTr("filter")
              font.pixelSize: 20
              elide: Label.ElideRight
              horizontalAlignment: Qt.AlignHCenter
              verticalAlignment: Qt.AlignVCenter
              Layout.fillWidth: true
          }

          ToolButton {
              contentItem: Image {
                  fillMode: Image.Pad
                  horizontalAlignment: Image.AlignHCenter
                  verticalAlignment: Image.AlignVCenter
                  source: "images/menu.png"
              }
              onClicked: optionsMenu.open()

              Menu {
                  id: optionsMenu
                  x: parent.width - width
                  transformOrigin: Menu.TopRight

                  MenuItem {
                      text: qsTr("Settings")
                      onTriggered: settingsPopup.open()
                  }
                  MenuItem {
                      text: qsTr("About")
                      onTriggered: aboutDialog.open()
                  }
              }
          }
      }
  }
  Drawer {
          id: drawer
          width: Math.min(window.width, window.height) / 3 * 2
          height: window.height
          dragMargin: stackView.depth > 1 ? 0 : undefined

          ListView {
              id: listView
              currentIndex: -1
              anchors.fill: parent

              delegate: ItemDelegate {
                  width: parent.width
                  text: model.title
                  highlighted: ListView.isCurrentItem
                  onClicked: {
                      if (listView.currentIndex != index) {
                          listView.currentIndex = index
                          stackView.push(model.source)
                      }
                      drawer.close()
                  }
              }

              model: ListModel {
                  ListElement { title: qsTr("proxy settings"); source: "pages/proxy.qml" }
                  ListElement { title: qsTr("plugins test"); source: "pages/plugins.qml" }
              }

              ScrollIndicator.vertical: ScrollIndicator { }
          }
      }

      StackView {
          id: stackView
          anchors.fill: parent

          initialItem:Loader{
              source:"pages/main.qml"
          }
      }
      Settings{
          id: settingsPopup
          x: (window.width - width) / 2
          y: window.height / 6
          width: Math.min(window.width, window.height) / 4 * 3
      }
      About{
          id: aboutDialog
          x: (window.width - width) / 2
          y: window.height / 6
          width: Math.min(window.width, window.height) / 4 * 3
      }
}
