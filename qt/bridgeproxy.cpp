#include "bridgeproxy.h"
#include <boost/thread.hpp>
BridgeProxy::BridgeProxy(QObject *parent) : QObject(parent)
{

}
int BridgeProxy::testProxy(int style,QString addr,QString user,QString pwd,QString testUrl)
{
    int id = _id++;
    if(!id){
        id = _id++;
    }

    BridgeProxy* ctx = this;
    boost::thread([=](){
        boost::this_thread::sleep(boost::posix_time::seconds(2));

        if(style == 1)
        {
            emit ctx->testProxyReply(id,"");
        }
        else
        {
            emit ctx->testProxyReply(id,"BridgeProxy::testProxy test error");
        }
    });
    return id;
}
