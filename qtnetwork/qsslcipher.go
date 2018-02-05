package qtnetwork

// /usr/include/qt/QtNetwork/qsslcipher.h
// #include <qsslcipher.h>
// #include <QtNetwork>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QSslCipher struct {
	*qtrt.CObject
}

func (this *QSslCipher) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSslCipher) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSslCipherFromPointer(cthis unsafe.Pointer) *QSslCipher {
	return &QSslCipher{&qtrt.CObject{cthis}}
}
func (*QSslCipher) NewFromPointer(cthis unsafe.Pointer) *QSslCipher {
	return NewQSslCipherFromPointer(cthis)
}

// /usr/include/qt/QtNetwork/qsslcipher.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSslCipher()
func NewQSslCipher() *QSslCipher {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipherC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCipher)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcipher.h:59
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSslCipher(const QString &)
func NewQSslCipher_1(name *qtcore.QString) *QSslCipher {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipherC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCipher)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcipher.h:60
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSslCipher(const QString &, QSsl::SslProtocol)
func NewQSslCipher_2(name *qtcore.QString, protocol int) *QSslCipher {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipherC2ERK7QStringN4QSsl11SslProtocolE", qtrt.FFI_TYPE_POINTER, convArg0, protocol)
	gopp.ErrPrint(err, rv)
	gothis := NewQSslCipherFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSslCipher)
	return gothis
}

// /usr/include/qt/QtNetwork/qsslcipher.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QSslCipher()
func DeleteQSslCipher(this *QSslCipher) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipherD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtNetwork/qsslcipher.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QSslCipher &)
func (this *QSslCipher) Swap(other *QSslCipher) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QSslCipher4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtNetwork/qsslcipher.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QSslCipher) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtNetwork/qsslcipher.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QSslCipher) Name() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcipher.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int supportedBits()
func (this *QSslCipher) SupportedBits() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher13supportedBitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslcipher.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int usedBits()
func (this *QSslCipher) UsedBits() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher8usedBitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtNetwork/qsslcipher.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QString keyExchangeMethod()
func (this *QSslCipher) KeyExchangeMethod() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher17keyExchangeMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcipher.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QString authenticationMethod()
func (this *QSslCipher) AuthenticationMethod() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher20authenticationMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcipher.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QString encryptionMethod()
func (this *QSslCipher) EncryptionMethod() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher16encryptionMethodEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcipher.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString protocolString()
func (this *QSslCipher) ProtocolString() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher14protocolStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtNetwork/qsslcipher.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] QSsl::SslProtocol protocol()
func (this *QSslCipher) Protocol() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QSslCipher8protocolEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
