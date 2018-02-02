#include "nodeobject.h"
#include <QDebug>
NodeObject::NodeObject(QObject *parent) : QObject(parent)
{

}
NodeObject::~NodeObject()
{

}
QString NodeObject::title()
{
    return _title;
}
void NodeObject::setTitle(const QString& val)
{
    if(_title != val)
    {
        _title = val;
        emit titleChanged(val);
    }
}
QString NodeObject::url()
{
    return _url;
}
void NodeObject::setUrl(const QString& val)
{
    if(_url != val)
    {
        _url = val;
        emit urlChanged(val);
    }
}
QString NodeObject::pluginsId()
{
    return _pluginsId;
}
void NodeObject::setPluginsId(const QString& val)
{
    if(_pluginsId != val)
    {
        _pluginsId = val;
        emit pluginsIdChanged(val);
    }
}
QString NodeObject::pluginsName()
{
    return _pluginsName;
}
void NodeObject::setPluginsName(const QString& val)
{
    if(_pluginsName != val)
    {
        _pluginsName = val;
        emit pluginsNameChanged(val);
    }
}
