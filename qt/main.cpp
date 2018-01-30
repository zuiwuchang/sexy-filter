#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include <QTranslator>
#include <QDebug>
#include <QQmlContext>

#include "bridgeconfigure.h"
#include "bridgeproxy.h"
int main(int argc, char *argv[])
{
#if defined(Q_OS_WIN)
    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);
#endif

    QGuiApplication app(argc, argv);
    app.setOrganizationName("cerberus");
    app.setOrganizationDomain("doc.king011.com");
    app.setApplicationName("go-qt-sexy-filter");


    QTranslator translator;
    translator.load(":locale/zh_TW.qm");
    app.installTranslator(&translator);

    QQmlApplicationEngine engine;
    QQmlContext* content = engine.rootContext();

    BridgeConfigure bridgeConfigure;
    content->setContextProperty("BridgeConfigure",&bridgeConfigure);
    BridgeProxy bridgeProxy;
    content->setContextProperty("BridgeProxy",&bridgeProxy);

    engine.load(QUrl(QStringLiteral("qrc:/views/main.qml")));
    if (engine.rootObjects().isEmpty())
        return -1;

    qDebug()<<engine.offlineStoragePath();
    return app.exec();
}
