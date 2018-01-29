import QtQuick 2.0
import QtQuick.Controls 2.2
import QtQuick.Layouts 1.3
import QtQuick.Controls.Material 2.0
import QtQuick.Controls.Universal 2.0

Popup {
    height: settingsColumn.implicitHeight + topPadding + bottomPadding
    modal: true
    focus: true

    onClosed: {
        comboxLocale.val = BridgeConfigure.getLocaleIndex();
        comboxStyle.val = BridgeConfigure.getStyle();
    }
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
                id: comboxLocale
                property var val: BridgeConfigure.getLocaleIndex()
                currentIndex: val
                model: BridgeConfigure.getLocales()
                Layout.fillWidth: true
            }
        }

        RowLayout {
            spacing: 10

            Label {
                text: qsTr("Style:")
            }

            ComboBox {
                id: comboxStyle
                property var val: BridgeConfigure.getStyle()
                currentIndex: val
                model: BridgeConfigure.getStyles()
                Layout.fillWidth: true
            }
        }

        Label {
            text: qsTr("Restart required")
            color: "#e41e25"
            opacity: ((comboxLocale.currentIndex == comboxLocale.val) && (comboxStyle.currentIndex == comboxStyle.val))?0.0:1.0
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
                    var locale = comboxLocale.currentIndex;
                    var style = comboxStyle.currentIndex;
                    if((locale == comboxLocale.val) && (style == comboxStyle.val)){
                        settingsPopup.close();
                        return;
                    }

                    BridgeConfigure.save2(locale,style);
                    settingsPopup.close();
                }
                Layout.fillWidth: true
            }

            Button {
                id: cancelButton
                text: qsTr("Cancel")
                onClicked: {
                    settingsPopup.close();
                }
                Layout.fillWidth: true
            }
        }
    }
}
