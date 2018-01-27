import QtQuick 2.0
import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import QtQuick.Controls.Material 2.0
import QtQuick.Controls.Universal 2.0

Popup {
    height: settingsColumn.implicitHeight + topPadding + bottomPadding
    modal: true
    focus: true

    contentItem: ColumnLayout {
        id: settingsColumn
        spacing: 20

        Label {
            text: qsTr("Settings")
            font.bold: true
        }

        RowLayout {
            spacing: 10

            Label {
                text: qsTr("Lang:")
            }

            ComboBox {
                property int styleIndex: -1
                model: ["zh_TW", "en_USA"]
                Component.onCompleted: {

                }
                Layout.fillWidth: true
            }
        }

        RowLayout {
            spacing: 10

            Label {
                text: qsTr("Style:")
            }

            ComboBox {
                id: styleBox
                property int styleIndex: -1
                model: ["Material Dark","Material Light", "Universal Dark","Universal Light"]
                Component.onCompleted: {

                }
                Layout.fillWidth: true
            }
        }

        Label {
            text: "Restart required"
            color: "#e41e25"
            opacity: 1.0
            horizontalAlignment: Label.AlignHCenter
            verticalAlignment: Label.AlignVCenter
            Layout.fillWidth: true
            Layout.fillHeight: true
        }

        RowLayout {
            spacing: 10

            Button {
                id: okButton
                text: qsTr("Sure")
                onClicked: {
                    settings.style = styleBox.displayText
                    settingsPopup.close()
                }

                Layout.fillWidth: true
            }

            Button {
                id: cancelButton
                text: qsTr("Cancel")
                onClicked: {
                    styleBox.currentIndex = styleBox.styleIndex
                    settingsPopup.close()
                }

                Layout.fillWidth: true
            }
        }
    }
}
