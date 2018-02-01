#ifndef BRIDGEPLUGINS_H
#define BRIDGEPLUGINS_H

#include <QObject>
#include <QStringList>
class BridgePlugins : public QObject
{
    Q_OBJECT
public:
    explicit BridgePlugins(QObject *parent = nullptr);

    int _id = 0;
signals:
    void testReply(int id,QString emsg,QString msg);
public slots:
    QStringList getPluginsFiles();
    QStringList getTestFiles();

    int testUrl(int style,QString addr,QString user,QString pwd, QString js);
    int testFile(QString js,QString file);
};

#endif // BRIDGEPLUGINS_H
