import QtQuick 2.0
import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import Qt.labs.settings 1.0
import QtQml.Models 2.2
Pane{
    id:thisView

    property real requestID: 0
    Settings{
        id:settingsPlugins
        category: "Plugins"
        property bool testUrl: switchTestUrl.checked
        property string jsFile
        property string testFile
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
                 id:gridLayout
                 columns:2
                 width: thisView.width - 200
                 anchors.horizontalCenter: parent.horizontalCenter


                 Label{
                     text: qsTr("select plugins file :")
                 }
                 ComboBox{
                     id:comboxJsFile
                     Layout.fillWidth: true
                     model: BridgePlugins.getPluginsFiles()
                     onCurrentTextChanged: settingsPlugins.jsFile = currentText
                     Component.onCompleted: {
                         var m = this.model;
                         for(var i=0;i<m.length;i++){
                             if(m[i] == settingsPlugins.jsFile){
                                 this.currentIndex = i;
                                 break;
                             }
                         }
                     }
                 }

                 Label{}
                 Switch{
                     id:switchTestUrl
                     checked: settingsPlugins.testUrl
                     text: qsTr("test url")
                 }
                 Label{
                     text: qsTr("select test file:")
                 }
                 ComboBox{
                     id:comboxTestFile
                     enabled: !switchTestUrl.checked
                     Layout.fillWidth: true
                     model: BridgePlugins.getTestFiles()
                     onCurrentTextChanged: settingsPlugins.testFile = currentText
                     Component.onCompleted: {
                         var m = this.model;
                         for(var i=0;i<m.length;i++){
                             if(m[i] == settingsPlugins.testFile){
                                 this.currentIndex = i;
                                 break;
                             }
                         }
                     }
                 }

                 Label{}
                 Button{
                     text: qsTr("test")
                     onClicked: {
                         modelRs.clear();

                         if(switchTestUrl.checked){
                             //url 測試
                             gridLayout.enabled = false;
                             busyIndicator.opacity = 1.0;
                             thisView.requestID = BridgePlugins.testUrl(
                                         settingsProxy.pos,
                                         settingsProxy.addr,
                                         settingsProxy.user,
                                         settingsProxy.pwd,
                                         comboxJsFile.currentText
                                         );
                         }else{
                             //檔案 分析
                             gridLayout.enabled = false;
                             busyIndicator.opacity = 1.0;
                             thisView.requestID = BridgePlugins.testFile(
                                         comboxJsFile.currentText,
                                         comboxTestFile.currentText
                                         );
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
            Connections{
                target: BridgePlugins
                onTestReply:{
                    if(thisView.requestID != id){
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
                        labelEmsg.text = qsTr("test success");

                        if(!msg){
                            return;
                        }


                        var arrs = JSON.parse(msg);
                        if(arrs && arrs.length > 0){
                            modelRs.append(arrs);
                        }
                    }

                }
            }

        }


        Pane{
            clip: true
            id:viewRs
            Layout.fillWidth: true
            Layout.fillHeight: true
            ScrollView{
                anchors.fill: parent
                ListView {
                     model: modelRs
                     delegate: RowLayout{
                         ItemDelegate{
                             text: Title
                             onClicked:Qt.openUrlExternally(Url)
                         }
                     }

                }
                ListModel{
                    id:modelRs
                }
            }
        }

    }
}
