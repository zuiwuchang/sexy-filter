#ifndef BRIDGEPROXY_H
#define BRIDGEPROXY_H

#include <QObject>

class BridgeProxy : public QObject
{
    Q_OBJECT
public:
    explicit BridgeProxy(QObject *parent = nullptr);

    int _id = 0;
signals:
    void testProxyReply(int id,QString emsg);
public slots:
    int testProxy(int style,QString addr,QString user,QString pwd,QString testUrl);
};

#endif // BRIDGEPROXY_H
