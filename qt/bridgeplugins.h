#ifndef BRIDGEPLUGINS_H
#define BRIDGEPLUGINS_H

#include <QObject>
#include <QStringList>
#include <QVariant>
#include <QList>
#include "nodeobject.h"


#define STATUS_NONE     0
#define STATUS_RUN      1
#define STATUS_RUNING   2
#define STATUS_STOPING  3


class BridgePlugins : public QObject
{
    Q_OBJECT
public:
    explicit BridgePlugins(QObject *parent = nullptr);

    int _id = 0;
    int _status = 0;
private:
    QList<QObject*> modelTest;
    void clearModelTest();
signals:
    void testReply(int id,QString emsg);
    void statusChanged(int val);
    void modelTestChanged();
public slots:
    QStringList getPluginsFiles();
    QStringList getTestFiles();

    int testUrl(int style,QString addr,QString user,QString pwd, QString js);
    int testFile(QString js,QString file);
    QVariant getModelTest();

    int getStatus();
    QStringList getPlugins();
    int start(int style,QString addr,QString user,QString pwd,int pos);
    int stop();
};

#endif // BRIDGEPLUGINS_H
