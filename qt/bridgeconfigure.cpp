#include "bridgeconfigure.h"

BridgeConfigure::BridgeConfigure(QObject *parent) : QObject(parent)
{

}
QString BridgeConfigure::getLocales()
{
    return "";
}
int BridgeConfigure::getLocaleIndex()
{
    return 0;
}
QString BridgeConfigure::getStyles()
{
    return "";
}
int BridgeConfigure::getStyle()
{
    return 0;
}
void BridgeConfigure::save2(int,int)
{

}
