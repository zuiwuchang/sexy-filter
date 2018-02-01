#include "bridgeplugins.h"
#include <boost/thread.hpp>
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
    int id = _id++;
    if(!id)
    {
        id = _id++;
    }

    auto ctx = this;
    boost::thread([=](){
        boost::this_thread::sleep(boost::posix_time::seconds(2));

        if(js == "1")
        {
            emit ctx->testReply(id,"",R"(
                                [
                                      {
                                      "Title":"t1",
                                      "Url":"u1"
                                      },
                                      {
                                      "Title":"t2",
                                      "Url":"u2"
                                      }
                                ]
                                      )");
        }
        else
        {
            emit ctx->testReply(id,"BridgePlugins::testUrl test error","");
        }
    });
    return id;
}
int BridgePlugins::testFile(QString js,QString file)
{
    int id = _id++;
    if(!id)
    {
        id = _id++;
    }

    auto ctx = this;
    boost::thread([=](){
        boost::this_thread::sleep(boost::posix_time::seconds(2));

        if(js == "1")
        {
            emit ctx->testReply(id,"",R"(
                                [
                                      {
                                      "Title":"t1",
                                      "Url":"u1"
                                      },
                                      {
                                      "Title":"t2",
                                      "Url":"u2"
                                      }
                                ]
                                      )");
        }
        else
        {
            emit ctx->testReply(id,"BridgePlugins::testFile test error","");
        }
    });
    return id;
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
int BridgePlugins::start(int pos)
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
