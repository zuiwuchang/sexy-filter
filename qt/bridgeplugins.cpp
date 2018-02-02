#include "bridgeplugins.h"
#include <boost/thread.hpp>
#include <QDebug>
BridgePlugins::BridgePlugins(QObject *parent) : QObject(parent)
{

}
QStringList BridgePlugins::getPluginsFiles()
{
    QStringList arrs;

    arrs.append("1");
    arrs.append("2");
    arrs.append("3");

    return arrs;
}
QStringList BridgePlugins::getTestFiles()
{
    QStringList arrs;

    arrs.append("t1");
    arrs.append("t2");
    arrs.append("t3");

    return arrs;
}
int BridgePlugins::testUrl(int style,QString addr,QString user,QString pwd, QString js)
{
    clearModelTest();

    int id = _id++;
    if(!id)
    {
        id = _id++;
    }

    auto p0 = new NodeObject();
    p0->setTitle("測試 url 0");
    p0->setUrl("https://www.google.com.tw");
    modelTest.append(p0);
    p0 = new NodeObject();
    p0->setTitle("測試 url 1");
    p0->setUrl("https://www.google.co.jp");
    modelTest.append(p0);

    auto ctx = this;
    boost::thread([=](){
        boost::this_thread::sleep(boost::posix_time::seconds(2));

        if(js == "1")
        {
            emit ctx->testReply(id,"");
        }
        else
        {
            emit ctx->testReply(id,"BridgePlugins::testUrl test error");
        }

        emit modelTestChanged();
    });
    return id;
}
void BridgePlugins::clearModelTest()
{
    if(modelTest.empty())
    {
        return;
    }

    for(auto p:modelTest)
    {
        delete (NodeObject*)(p);
    }
    modelTest.clear();
    emit modelTestChanged();
}
int BridgePlugins::testFile(QString js,QString file)
{
    clearModelTest();

    int id = _id++;
    if(!id)
    {
        id = _id++;
    }

    auto p0 = new NodeObject();
    p0->setTitle("測試 file 0");
    p0->setUrl("https://www.google.com.tw");
    modelTest.append(p0);
    p0 = new NodeObject();
    p0->setTitle("測試 file 1");
    p0->setUrl("https://www.google.co.jp");
    modelTest.append(p0);

    auto ctx = this;
    boost::thread([=](){
        boost::this_thread::sleep(boost::posix_time::seconds(2));

        if(js == "1")
        {
            emit ctx->testReply(id,"");
        }
        else
        {
            emit ctx->testReply(id,"BridgePlugins::testFile test error");
        }

        emit modelTestChanged();
    });
    return id;
}
QVariant BridgePlugins::getModelTest()
{
    return QVariant::fromValue(modelTest);
}
int BridgePlugins::getStatus()
{
    return _status;
}
QStringList BridgePlugins::getPlugins()
{
    QStringList arrs;

    arrs.append("1");
    arrs.append("2");
    arrs.append("3");

    return arrs;
}
int BridgePlugins::start(int style,QString addr,QString user,QString pwd,int pos)
{
    if(_status != STATUS_NONE){
        return _status;
    }
    _status = STATUS_RUN;

    auto ctx = this;
    boost::thread([=](){
        boost::this_thread::sleep(boost::posix_time::seconds(1));
        ctx->_status = STATUS_RUNING;
        emit ctx->statusChanged(STATUS_RUNING);
    });
    return _status;
}
int BridgePlugins::stop()
{
    if(_status != STATUS_RUNING){
        return _status;
    }
    _status = STATUS_STOPING;

    auto ctx = this;
    boost::thread([=](){
        boost::this_thread::sleep(boost::posix_time::seconds(1));
        ctx->_status = STATUS_NONE;
        emit ctx->statusChanged(STATUS_NONE);
    });
    return _status;
}
