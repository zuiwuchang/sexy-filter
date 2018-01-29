import QtQuick 2.0
import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import Qt.labs.settings 1.0
Pane{
    Settings{
        id:settingsProxy
        category: "Proxy"
        property real pos: comboxType.currentIndex
        property string addr: textAddr.text
        property string user: textUser.text
        property string pwd: textPwd.text
    }
    ColumnLayout {
        anchors.centerIn: parent

        GridLayout{
            columns: 2
            Label{
                text:qsTr("type:")
            }
            ComboBox{
                id:comboxType
                Layout.fillWidth: true
                model: ["http","https","socks5"]
                Component.onCompleted: currentIndex = settingsProxy.pos
            }

            Label{
                text: qsTr("addr:")
            }
            TextField{
                id:textAddr
                Layout.fillWidth: true
                placeholderText: qsTr("proxy address")
                text: settingsProxy.addr
            }

            Label{
                text: qsTr("user:")
            }
            TextField{
                enabled: comboxType.currentIndex == 2
                id:textUser
                Layout.fillWidth: true
                placeholderText: qsTr("proxy user name")
                text: settingsProxy.user
            }

            Label{
                text: qsTr("pwd:")
            }
            TextField{
                enabled: comboxType.currentIndex == 2
                id:textPwd
                Layout.fillWidth: true
                placeholderText: qsTr("proxy password")
                text: settingsProxy.pwd
            }

            Label{}

            Button{
                text: qsTr("test")
                onClicked: {

                }
            }
        }
    }
}
