import QtQuick 2.0
import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import Qt.labs.settings 1.0
Pane{
    id:thisView
    property real proxyNone: 0
    property real proxySocks5: 3
    property real requestID: 0
    Settings{
        id:settingsProxy
        category: "Proxy"
        property real pos: comboxType.currentIndex
        property string addr: textAddr.text
        property string user: textUser.text
        property string pwd: textPwd.text
        property string testUrl: textTestUrl.text
    }
    ColumnLayout {
        anchors.fill: parent
        Column{
            anchors.horizontalCenter: parent.horizontalCenter
            BusyIndicator{
                id:busyIndicator
                opacity: 0.0
            }
        }
        Column{
            anchors.horizontalCenter: parent.horizontalCenter

            GridLayout{
                width: thisView.width - 200
                Layout.fillHeight: true
                anchors.horizontalCenter: parent.horizontalCenter
                id:gridLayout
                columns: 2
                Label{
                    text:qsTr("type:")
                }
                ComboBox{
                    id:comboxType
                    Layout.fillWidth: true
                    model: [qsTr("none"),"http","https","socks5"]
                    Component.onCompleted: currentIndex = settingsProxy.pos
                }

                Label{
                    text: qsTr("addr:")
                }
                TextField{
                    enabled: comboxType.currentIndex != thisView.proxyNone
                    id:textAddr
                    Layout.fillWidth: true
                    placeholderText: qsTr("proxy address")
                    text: settingsProxy.addr
                }

                Label{
                    text: qsTr("user:")
                }
                TextField{
                    enabled: comboxType.currentIndex == thisView.proxySocks5
                    id:textUser
                    Layout.fillWidth: true
                    placeholderText: qsTr("proxy user name")
                    text: settingsProxy.user
                }

                Label{
                    text: qsTr("pwd:")
                }
                TextField{
                    enabled: comboxType.currentIndex == thisView.proxySocks5
                    id:textPwd
                    Layout.fillWidth: true
                    placeholderText: qsTr("proxy password")
                    text: settingsProxy.pwd
                }


                Label{
                    text: qsTr("test url:")
                }
                TextField{
                    enabled: comboxType.currentIndex != thisView.proxyNone
                    id:textTestUrl
                    Layout.fillWidth: true
                    placeholderText: qsTr("test proxy connect url")
                    text: settingsProxy.testUrl
                }
                Label{}
                Button{
                    enabled: comboxType.currentIndex != thisView.proxyNone
                    text: qsTr("test")
                    onClicked: {
                        thisView.requestID = BridgeProxy.testProxy(comboxType.currentIndex,
                                                         textAddr.text,
                                                         textUser.text,
                                                         textPwd.text,
                                                         textTestUrl.text
                                              );
                        gridLayout.enabled = false;
                        labelEmsg.visible = false;
                        busyIndicator.opacity = 1.0;
                    }
                }
                Connections{
                    target: BridgeProxy
                    onTestProxyReply: {
                        if(id != thisView.requestID){
                            return
                        }

                        thisView.requestID = 0;

                        gridLayout.enabled = true;
                        busyIndicator.opacity = 0.0;
                        if(emsg){
                            labelEmsg.visible = true;
                            labelEmsg.color="red";
                            labelEmsg.text = emsg;
                        }else{
                            labelEmsg.visible = true;
                            labelEmsg.color = "green";
                            labelEmsg.text = qsTr("connect success");
                        }
                    }
                }

            }

        }

        Column{
            anchors.horizontalCenter: parent.horizontalCenter
            Label{
                id:labelEmsg
                wrapMode: Label.Wrap
                visible: false
                color: "red"
            }
            Layout.fillHeight: true
        }
    }
}
