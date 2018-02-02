import QtQuick 2.0
import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import Qt.labs.settings 1.0

Pane{
    id:thisView
    property real status: BridgePlugins.getStatus()
    readonly property real statusNone: 0    //空閒
    readonly property real statusRun: 1     //啓動
    readonly property real statusRuning: 2  //正在運行
    readonly property real statusStoping: 3 //正在停止
    Connections{
        target: BridgePlugins
        onStatusChanged : status = val
    }
    Settings{
        id:settings
        property real pluginsPos: comboxPlugins.currentIndex
        property string textTitle: textTitle.text
        property string textPluginsName: textPluginsName.text
        property string textPluginsId: textPluginsId.text
        property string textLimit: textLimit.text
    }
    Settings{
        id:settingsProxy
        category: "Proxy"
        property real pos
        property string addr
        property string user
        property string pwd
    }

    ColumnLayout{
        anchors.fill: parent
        GroupBox{
            Layout.fillWidth: true
            RowLayout{
                anchors.fill: parent
                ComboBox{
                    Layout.fillWidth: true
                    id:comboxPlugins
                    enabled: thisView.statusNone == thisView.status
                    model: BridgePlugins.getPlugins()
                    Component.onCompleted: currentIndex = settings.pluginsPos
                }
                ProgressBar {
                    Layout.fillWidth: true
                    indeterminate: thisView.statusNone != thisView.status
                }
                Button{
                    enabled: (thisView.statusNone == thisView.status) || (thisView.statusRuning == thisView.status)
                    text:thisView.statusNone == thisView.status ? qsTr("start") : qsTr("stop")
                    onClicked: {
                        if(thisView.statusNone == thisView.status){
                            var i = comboxPlugins.currentIndex;
                            if(i!=-1){
                                thisView.status = BridgePlugins.start(
                                            settingsProxy.pos,
                                            settingsProxy.addr,
                                            settingsProxy.user,
                                            settingsProxy.pwd,
                                            comboxPlugins.currentIndex);
                            }
                        }else if(thisView.statusRuning == thisView.status){
                            thisView.status = BridgePlugins.stop();
                        }
                    }
                }
            }
        }
        GroupBox{
            Layout.fillWidth: true
            GridLayout{
                width: parent.width
                anchors.horizontalCenter: parent.horizontalCenter
                columns: 3
                Label{
                    text: qsTr("Title:")
                }
                TextField{
                    id:textTitle
                    Layout.fillWidth: true
                    placeholderText: "きぃちゃん"
                    text: settings.textTitle
                }
                Button{
                    text: qsTr("search")
                    onClicked: {
                        BridgePlugins.search(
                                    textTitle.text,
                                    textPluginsName.text,
                                    textPluginsId.text,
                                    textLimit.text
                                    );
                    }
                }

                Label{
                    text: qsTr("PluginsName:")
                }
                TextField{
                    id:textPluginsName
                    Layout.fillWidth: true
                    placeholderText: "中文"
                    text: settings.textPluginsName
                }
                Label{}

                Label{
                    text: qsTr("PluginsId:")
                }
                TextField{
                    id:textPluginsId
                    Layout.fillWidth: true
                    placeholderText: "t66y"
                    text: settings.textPluginsId
                }
                Label{}

                Label{
                    text: qsTr("limit:")
                }
                TextField{
                    id:textLimit
                    Layout.fillWidth: true
                    placeholderText: "rows[,start]  (max 100)"
                    text: settings.textLimit
                }
            }
        }

        GroupBox{
            Layout.fillHeight: true
            Layout.fillWidth: true
            clip: true
            ScrollView{
                anchors.fill: parent
                clip: true
                ListView{
                    id:listView
                    model:BridgePlugins.getModelSearch()
                    delegate: RowLayout{
                        ItemDelegate{
                            text: index
                            onClicked:Qt.openUrlExternally(index)
                        }
                        ItemDelegate{
                            text: modelData.title
                            onClicked:Qt.openUrlExternally(modelData.url)
                        }
                        ItemDelegate{
                            text: modelData.pluginsName
                            onClicked:Qt.openUrlExternally(modelData.url)
                        }
                        ItemDelegate{
                            text: modelData.pluginsId
                            onClicked:Qt.openUrlExternally(modelData.url)
                        }
                    }
                }
                Connections{
                    target: BridgePlugins
                    onModelSearchChanged:listView.model = BridgePlugins.getModelSearch()
                }
            }
        }
    }
}



