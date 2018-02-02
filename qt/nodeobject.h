#ifndef NODEOBJECT_H
#define NODEOBJECT_H

#include <QObject>

class NodeObject : public QObject
{
    Q_OBJECT
    Q_PROPERTY(QString title READ title WRITE setTitle NOTIFY titleChanged)
    Q_PROPERTY(QString url READ url WRITE setUrl NOTIFY urlChanged)
    Q_PROPERTY(QString pluginsId READ pluginsId WRITE setPluginsId NOTIFY pluginsIdChanged)
    Q_PROPERTY(QString pluginsName READ pluginsName WRITE setPluginsName NOTIFY pluginsNameChanged)
public:
    explicit NodeObject(QObject *parent = nullptr);
    virtual ~NodeObject();
private:
    QString _title;
    QString _url;
    QString _pluginsId;
    QString _pluginsName;
public:
    QString title();
    void setTitle(const QString& val);
    QString url();
    void setUrl(const QString& val);
    QString pluginsId();
    void setPluginsId(const QString& val);
    QString pluginsName();
    void setPluginsName(const QString& val);
signals:
    void titleChanged(const QString&);
    void urlChanged(const QString&);
    void pluginsIdChanged(const QString&);
    void pluginsNameChanged(const QString&);
public slots:

};

#endif // NODEOBJECT_H
