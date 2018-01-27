package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkproxy.h
// #include <qnetworkproxy.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QNetworkProxyFactory struct {
	*qtrt.CObject
}

func (this *QNetworkProxyFactory) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QNetworkProxyFactory) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQNetworkProxyFactoryFromPointer(cthis unsafe.Pointer) *QNetworkProxyFactory {
	return &QNetworkProxyFactory{&qtrt.CObject{cthis}}
}
func (*QNetworkProxyFactory) NewFromPointer(cthis unsafe.Pointer) *QNetworkProxyFactory {
	return NewQNetworkProxyFactoryFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:219
// index:0
// Public
// void QNetworkProxyFactory()
func NewQNetworkProxyFactory() *QNetworkProxyFactory {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkProxyFactoryC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQNetworkProxyFactoryFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:220
// index:0
// Public virtual
// void ~QNetworkProxyFactory()
func DeleteQNetworkProxyFactory(*QNetworkProxyFactory) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkProxyFactoryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:224
// index:0
// Public static
// bool usesSystemConfiguration()
func (this *QNetworkProxyFactory) UsesSystemConfiguration() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkProxyFactory23usesSystemConfigurationEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QNetworkProxyFactory_UsesSystemConfiguration() bool {
	var nilthis *QNetworkProxyFactory
	rv := nilthis.UsesSystemConfiguration()
	return rv
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:225
// index:0
// Public static
// void setUseSystemConfiguration(bool)
func (this *QNetworkProxyFactory) SetUseSystemConfiguration(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkProxyFactory25setUseSystemConfigurationEb", ffiqt.FFI_TYPE_POINTER, enable)
	gopp.ErrPrint(err, rv)
}
func QNetworkProxyFactory_SetUseSystemConfiguration(enable bool) {
	var nilthis *QNetworkProxyFactory
	nilthis.SetUseSystemConfiguration(enable)
}

// /usr/include/qt/QtNetwork/qnetworkproxy.h:226
// index:0
// Public static
// void setApplicationProxyFactory(QNetworkProxyFactory *)
func (this *QNetworkProxyFactory) SetApplicationProxyFactory(factory *QNetworkProxyFactory /*777 QNetworkProxyFactory **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QNetworkProxyFactory26setApplicationProxyFactoryEPS_", ffiqt.FFI_TYPE_POINTER, factory)
	gopp.ErrPrint(err, rv)
}
func QNetworkProxyFactory_SetApplicationProxyFactory(factory *QNetworkProxyFactory /*777 QNetworkProxyFactory **/) {
	var nilthis *QNetworkProxyFactory
	nilthis.SetApplicationProxyFactory(factory)
}

//  body block end