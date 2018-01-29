import QtQuick 2.0
import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3

Popup {
    modal: true
    focus: true
    contentHeight: aboutColumn.height
    Column {
        id: aboutColumn
        spacing: 10

        Label {
            text: qsTr("About")
            font.bold: true
        }

        Label {
            width: aboutDialog.availableWidth
            text: "這是一個類似 網絡爬蟲的 程式\n主要幫助司機在t66y等站點更方便的學習車技"
            wrapMode: Label.Wrap
            font.pixelSize: 12
        }
        Label {
            width: aboutDialog.availableWidth
            text: "git clone  https://github.com/zuiwuchang/sexy-filter.git"
            wrapMode: Label.Wrap
            font.pixelSize: 12
            MouseArea{
                anchors.fill: parent
                onClicked: Qt.openUrlExternally("https://github.com/zuiwuchang/sexy-filter")
            }
        }

        Label {
            width: aboutDialog.availableWidth
            text: "    作者     :   king"
            wrapMode: Label.Wrap
            font.pixelSize: 12
        }
        Label {
            width: aboutDialog.availableWidth
            text: "    博客     :   http://blog.king011.com/"
            wrapMode: Label.Wrap
            font.pixelSize: 12
            MouseArea{
                anchors.fill: parent
                onClicked: Qt.openUrlExternally("http://blog.king011.com/")
            }
        }
        Label {
            width: aboutDialog.availableWidth
            text: " google+ :   https://plus.google.com/u/0/陳宏杰"
            wrapMode: Label.Wrap
            font.pixelSize: 12
            MouseArea{
                anchors.fill: parent
                onClicked: Qt.openUrlExternally("https://plus.google.com/u/0/+%E5%AE%8F%E6%9D%B0%E9%99%B3")
            }
        }
    }
}
