#ifndef BRIDGECONFIGURE_H
#define BRIDGECONFIGURE_H

#include <QObject>
#include <QVector>
class BridgeConfigure : public QObject
{
    Q_OBJECT
public:
    explicit BridgeConfigure(QObject *parent = nullptr);

signals:

public slots:
    QString getLocales();
    int getLocaleIndex();
    QString getStyles();
    int getStyle();
    void save2(int,int);
};

#endif // BRIDGECONFIGURE_H
