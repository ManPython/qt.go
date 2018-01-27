package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicstransform.h
// #include <qgraphicstransform.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QGraphicsTransform struct {
	*qtcore.QObject
}

func (this *QGraphicsTransform) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGraphicsTransform) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGraphicsTransformFromPointer(cthis unsafe.Pointer) *QGraphicsTransform {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsTransform{bcthis0}
}
func (*QGraphicsTransform) NewFromPointer(cthis unsafe.Pointer) *QGraphicsTransform {
	return NewQGraphicsTransformFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsTransform) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsTransform10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:60
// index:0
// Public
// void QGraphicsTransform(QObject *)
func NewQGraphicsTransform(parent *qtcore.QObject /*777 QObject **/) *QGraphicsTransform {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsTransformC1EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsTransformFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:61
// index:0
// Public virtual
// void ~QGraphicsTransform()
func DeleteQGraphicsTransform(*QGraphicsTransform) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsTransformD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:63
// index:0
// Public pure virtual
// void applyTo(QMatrix4x4 *)
func (this *QGraphicsTransform) ApplyTo(matrix *qtgui.QMatrix4x4 /*777 QMatrix4x4 **/) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsTransform7applyToEP10QMatrix4x4", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicstransform.h:66
// index:0
// Protected
// void update()
func (this *QGraphicsTransform) Update() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsTransform6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end