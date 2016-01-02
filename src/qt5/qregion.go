package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qregion.h
// dst-file: /src/gui/qregion.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QRegion::isNull();
extern void _ZNK7QRegion6isNullEv(void* qthis);
  // proto:  QRect QRegion::boundingRect();
extern void _ZNK7QRegion12boundingRectEv(void* qthis);
  // proto:  void QRegion::QRegion(const QRegion & region);
extern void* dector_ZN7QRegionC1ERKS_(void* arg0);
extern void _ZN7QRegionC1ERKS_(void* qthis, void* arg0);
  // proto:  int QRegion::rectCount();
extern void _ZNK7QRegion9rectCountEv(void* qthis);
  // proto:  void QRegion::translate(int dx, int dy);
extern void _ZN7QRegion9translateEii(void* qthis, int arg0, int arg1);
  // proto:  QRegion QRegion::united(const QRegion & r);
extern void _ZNK7QRegion6unitedERKS_(void* qthis, void* arg0);
  // proto:  QRegion QRegion::translated(const QPoint & p);
extern void demth_ZNK7QRegion10translatedERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRegion::swap(QRegion & other);
extern void demth_ZN7QRegion4swapERS_(void* qthis, void* arg0);
  // proto:  void QRegion::QRegion(const QBitmap & bitmap);
extern void* dector_ZN7QRegionC1ERK7QBitmap(void* arg0);
extern void _ZN7QRegionC1ERK7QBitmap(void* qthis, void* arg0);
  // proto:  void QRegion::~QRegion();
extern void _ZN7QRegionD0Ev(void* qthis);
  // proto:  void QRegion::translate(const QPoint & p);
extern void demth_ZN7QRegion9translateERK6QPoint(void* qthis, void* arg0);
  // proto:  void QRegion::QRegion();
extern void* dector_ZN7QRegionC1Ev();
extern void _ZN7QRegionC1Ev(void* qthis);
  // proto:  bool QRegion::contains(const QRect & r);
extern void _ZNK7QRegion8containsERK5QRect(void* qthis, void* arg0);
  // proto:  bool QRegion::isEmpty();
extern void _ZNK7QRegion7isEmptyEv(void* qthis);
  // proto:  QRegion QRegion::intersected(const QRect & r);
extern void _ZNK7QRegion11intersectedERK5QRect(void* qthis, void* arg0);
  // proto:  void QRegion::setRects(const QRect * rect, int num);
extern void _ZN7QRegion8setRectsEPK5QRecti(void* qthis, void* arg0, int arg1);
  // proto:  QVector<QRect> QRegion::rects();
extern void _ZNK7QRegion5rectsEv(void* qthis);
  // proto:  QRegion QRegion::subtracted(const QRegion & r);
extern void _ZNK7QRegion10subtractedERKS_(void* qthis, void* arg0);
  // proto:  bool QRegion::intersects(const QRect & r);
extern void _ZNK7QRegion10intersectsERK5QRect(void* qthis, void* arg0);
  // proto:  QRegion QRegion::translated(int dx, int dy);
extern void _ZNK7QRegion10translatedEii(void* qthis, int arg0, int arg1);
  // proto:  QRegion QRegion::intersected(const QRegion & r);
extern void _ZNK7QRegion11intersectedERKS_(void* qthis, void* arg0);
  // proto:  QRegion QRegion::united(const QRect & r);
extern void _ZNK7QRegion6unitedERK5QRect(void* qthis, void* arg0);
  // proto:  QRegion QRegion::xored(const QRegion & r);
extern void _ZNK7QRegion5xoredERKS_(void* qthis, void* arg0);
  // proto:  bool QRegion::intersects(const QRegion & r);
extern void _ZNK7QRegion10intersectsERKS_(void* qthis, void* arg0);
  // proto:  bool QRegion::contains(const QPoint & p);
extern void _ZNK7QRegion8containsERK6QPoint(void* qthis, void* arg0);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QRegion)=8
type QRegion struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QRegion::isNull();
func (this *QRegion) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion6isNullEv
  default:
    qtrt.ErrorResolve("QRegion", "isNull", args)
  }

}

  // proto:  QRect QRegion::boundingRect();
func (this *QRegion) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion12boundingRectEv
  default:
    qtrt.ErrorResolve("QRegion", "boundingRect", args)
  }

}

  // proto:  void QRegion::QRegion(const QRegion & region);
func NewQRegion(args ...interface{}) QRegion {
  return QRegion{}
}

  // proto:  int QRegion::rectCount();
func (this *QRegion) rectCount(args ...interface{}) () {
  // rectCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion9rectCountEv
  default:
    qtrt.ErrorResolve("QRegion", "rectCount", args)
  }

}

  // proto:  void QRegion::translate(int dx, int dy);
func (this *QRegion) translate(args ...interface{}) () {
  // translate(int, int)
  // translate(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegion9translateEii
  case 1:
    // invoke: _ZN7QRegion9translateERK6QPoint
  default:
    qtrt.ErrorResolve("QRegion", "translate", args)
  }

}

  // proto:  QRegion QRegion::united(const QRegion & r);
func (this *QRegion) united(args ...interface{}) () {
  // united(const class QRegion &)
  // united(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion6unitedERKS_
  case 1:
    // invoke: _ZNK7QRegion6unitedERK5QRect
  default:
    qtrt.ErrorResolve("QRegion", "united", args)
  }

}

  // proto:  QRegion QRegion::translated(const QPoint & p);
func (this *QRegion) translated(args ...interface{}) () {
  // translated(const class QPoint &)
  // translated(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion10translatedERK6QPoint
  case 1:
    // invoke: _ZNK7QRegion10translatedEii
  default:
    qtrt.ErrorResolve("QRegion", "translated", args)
  }

}

  // proto:  void QRegion::swap(QRegion & other);
func (this *QRegion) swap(args ...interface{}) () {
  // swap(class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegion4swapERS_
  default:
    qtrt.ErrorResolve("QRegion", "swap", args)
  }

}

  // proto:  void QRegion::~QRegion();
func (this *QRegion) FreeQRegion(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QRegion", "~QRegion", args)
  }

}

  // proto:  bool QRegion::contains(const QRect & r);
func (this *QRegion) contains(args ...interface{}) () {
  // contains(const class QRect &)
  // contains(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion8containsERK5QRect
  case 1:
    // invoke: _ZNK7QRegion8containsERK6QPoint
  default:
    qtrt.ErrorResolve("QRegion", "contains", args)
  }

}

  // proto:  bool QRegion::isEmpty();
func (this *QRegion) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion7isEmptyEv
  default:
    qtrt.ErrorResolve("QRegion", "isEmpty", args)
  }

}

  // proto:  QRegion QRegion::intersected(const QRect & r);
func (this *QRegion) intersected(args ...interface{}) () {
  // intersected(const class QRect &)
  // intersected(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion11intersectedERK5QRect
  case 1:
    // invoke: _ZNK7QRegion11intersectedERKS_
  default:
    qtrt.ErrorResolve("QRegion", "intersected", args)
  }

}

  // proto:  void QRegion::setRects(const QRect * rect, int num);
func (this *QRegion) setRects(args ...interface{}) () {
  // setRects(const class QRect *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QRegion8setRectsEPK5QRecti
  default:
    qtrt.ErrorResolve("QRegion", "setRects", args)
  }

}

  // proto:  QVector<QRect> QRegion::rects();
func (this *QRegion) rects(args ...interface{}) () {
  // rects()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion5rectsEv
  default:
    qtrt.ErrorResolve("QRegion", "rects", args)
  }

}

  // proto:  QRegion QRegion::subtracted(const QRegion & r);
func (this *QRegion) subtracted(args ...interface{}) () {
  // subtracted(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion10subtractedERKS_
  default:
    qtrt.ErrorResolve("QRegion", "subtracted", args)
  }

}

  // proto:  bool QRegion::intersects(const QRect & r);
func (this *QRegion) intersects(args ...interface{}) () {
  // intersects(const class QRect &)
  // intersects(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion10intersectsERK5QRect
  case 1:
    // invoke: _ZNK7QRegion10intersectsERKS_
  default:
    qtrt.ErrorResolve("QRegion", "intersects", args)
  }

}

  // proto:  QRegion QRegion::xored(const QRegion & r);
func (this *QRegion) xored(args ...interface{}) () {
  // xored(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QRegion5xoredERKS_
  default:
    qtrt.ErrorResolve("QRegion", "xored", args)
  }

}

// <= body block end
