package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qglyphrun.h
// dst-file: /src/gui/qglyphrun.go
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
  // proto:  void QGlyphRun::~QGlyphRun();
extern void _ZN9QGlyphRunD0Ev(void* qthis);
  // proto:  void QGlyphRun::setBoundingRect(const QRectF & boundingRect);
extern void _ZN9QGlyphRun15setBoundingRectERK6QRectF(void* qthis, void* arg0);
  // proto:  bool QGlyphRun::overline();
extern void _ZNK9QGlyphRun8overlineEv(void* qthis);
  // proto:  void QGlyphRun::setRawData(const quint32 * glyphIndexArray, const QPointF * glyphPositionArray, int size);
extern void _ZN9QGlyphRun10setRawDataEPKjPK7QPointFi(void* qthis, unsigned int* arg0, void* arg1, int arg2);
  // proto:  void QGlyphRun::setOverline(bool overline);
extern void _ZN9QGlyphRun11setOverlineEb(void* qthis, bool arg0);
  // proto:  void QGlyphRun::swap(QGlyphRun & other);
extern void _ZN9QGlyphRun4swapERS_(void* qthis, void* arg0);
  // proto:  void QGlyphRun::setUnderline(bool underline);
extern void _ZN9QGlyphRun12setUnderlineEb(void* qthis, bool arg0);
  // proto:  QVector<QPointF> QGlyphRun::positions();
extern void _ZNK9QGlyphRun9positionsEv(void* qthis);
  // proto:  void QGlyphRun::clear();
extern void _ZN9QGlyphRun5clearEv(void* qthis);
  // proto:  bool QGlyphRun::strikeOut();
extern void _ZNK9QGlyphRun9strikeOutEv(void* qthis);
  // proto:  void QGlyphRun::QGlyphRun();
extern void* dector_ZN9QGlyphRunC1Ev();
extern void _ZN9QGlyphRunC1Ev(void* qthis);
  // proto:  QRawFont QGlyphRun::rawFont();
extern void _ZNK9QGlyphRun7rawFontEv(void* qthis);
  // proto:  void QGlyphRun::setRawFont(const QRawFont & rawFont);
extern void _ZN9QGlyphRun10setRawFontERK8QRawFont(void* qthis, void* arg0);
  // proto:  void QGlyphRun::QGlyphRun(const QGlyphRun & other);
extern void* dector_ZN9QGlyphRunC1ERKS_(void* arg0);
extern void _ZN9QGlyphRunC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QGlyphRun::isRightToLeft();
extern void _ZNK9QGlyphRun13isRightToLeftEv(void* qthis);
  // proto:  QVector<quint32> QGlyphRun::glyphIndexes();
extern void _ZNK9QGlyphRun12glyphIndexesEv(void* qthis);
  // proto:  QRectF QGlyphRun::boundingRect();
extern void _ZNK9QGlyphRun12boundingRectEv(void* qthis);
  // proto:  void QGlyphRun::setRightToLeft(bool on);
extern void _ZN9QGlyphRun14setRightToLeftEb(void* qthis, bool arg0);
  // proto:  bool QGlyphRun::underline();
extern void _ZNK9QGlyphRun9underlineEv(void* qthis);
  // proto:  void QGlyphRun::setStrikeOut(bool strikeOut);
extern void _ZN9QGlyphRun12setStrikeOutEb(void* qthis, bool arg0);
  // proto:  bool QGlyphRun::isEmpty();
extern void _ZNK9QGlyphRun7isEmptyEv(void* qthis);
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

// class sizeof(QGlyphRun)=1
type QGlyphRun struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGlyphRun::~QGlyphRun();
func (this *QGlyphRun) FreeQGlyphRun(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGlyphRun", "~QGlyphRun", args)
  }

}

  // proto:  void QGlyphRun::setBoundingRect(const QRectF & boundingRect);
func (this *QGlyphRun) setBoundingRect(args ...interface{}) () {
  // setBoundingRect(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun15setBoundingRectERK6QRectF
  default:
    qtrt.ErrorResolve("QGlyphRun", "setBoundingRect", args)
  }

}

  // proto:  bool QGlyphRun::overline();
func (this *QGlyphRun) overline(args ...interface{}) () {
  // overline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun8overlineEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "overline", args)
  }

}

  // proto:  void QGlyphRun::setRawData(const quint32 * glyphIndexArray, const QPointF * glyphPositionArray, int size);
func (this *QGlyphRun) setRawData(args ...interface{}) () {
  // setRawData(const quint32 *, const class QPointF *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "const quint32 *"
  vtys[0][1] = reflect.TypeOf(QPointF{}) // "const QPointF *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun10setRawDataEPKjPK7QPointFi
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRawData", args)
  }

}

  // proto:  void QGlyphRun::setOverline(bool overline);
func (this *QGlyphRun) setOverline(args ...interface{}) () {
  // setOverline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun11setOverlineEb
  default:
    qtrt.ErrorResolve("QGlyphRun", "setOverline", args)
  }

}

  // proto:  void QGlyphRun::swap(QGlyphRun & other);
func (this *QGlyphRun) swap(args ...interface{}) () {
  // swap(class QGlyphRun &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGlyphRun{}) // "QGlyphRun &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun4swapERS_
  default:
    qtrt.ErrorResolve("QGlyphRun", "swap", args)
  }

}

  // proto:  void QGlyphRun::setUnderline(bool underline);
func (this *QGlyphRun) setUnderline(args ...interface{}) () {
  // setUnderline(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun12setUnderlineEb
  default:
    qtrt.ErrorResolve("QGlyphRun", "setUnderline", args)
  }

}

  // proto:  QVector<QPointF> QGlyphRun::positions();
func (this *QGlyphRun) positions(args ...interface{}) () {
  // positions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun9positionsEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "positions", args)
  }

}

  // proto:  void QGlyphRun::clear();
func (this *QGlyphRun) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun5clearEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "clear", args)
  }

}

  // proto:  bool QGlyphRun::strikeOut();
func (this *QGlyphRun) strikeOut(args ...interface{}) () {
  // strikeOut()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun9strikeOutEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "strikeOut", args)
  }

}

  // proto:  void QGlyphRun::QGlyphRun();
func NewQGlyphRun(args ...interface{}) QGlyphRun {
  return QGlyphRun{}
}

  // proto:  QRawFont QGlyphRun::rawFont();
func (this *QGlyphRun) rawFont(args ...interface{}) () {
  // rawFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun7rawFontEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "rawFont", args)
  }

}

  // proto:  void QGlyphRun::setRawFont(const QRawFont & rawFont);
func (this *QGlyphRun) setRawFont(args ...interface{}) () {
  // setRawFont(const class QRawFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRawFont{}) // "const QRawFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun10setRawFontERK8QRawFont
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRawFont", args)
  }

}

  // proto:  bool QGlyphRun::isRightToLeft();
func (this *QGlyphRun) isRightToLeft(args ...interface{}) () {
  // isRightToLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun13isRightToLeftEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "isRightToLeft", args)
  }

}

  // proto:  QVector<quint32> QGlyphRun::glyphIndexes();
func (this *QGlyphRun) glyphIndexes(args ...interface{}) () {
  // glyphIndexes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun12glyphIndexesEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "glyphIndexes", args)
  }

}

  // proto:  QRectF QGlyphRun::boundingRect();
func (this *QGlyphRun) boundingRect(args ...interface{}) () {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun12boundingRectEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "boundingRect", args)
  }

}

  // proto:  void QGlyphRun::setRightToLeft(bool on);
func (this *QGlyphRun) setRightToLeft(args ...interface{}) () {
  // setRightToLeft(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun14setRightToLeftEb
  default:
    qtrt.ErrorResolve("QGlyphRun", "setRightToLeft", args)
  }

}

  // proto:  bool QGlyphRun::underline();
func (this *QGlyphRun) underline(args ...interface{}) () {
  // underline()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun9underlineEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "underline", args)
  }

}

  // proto:  void QGlyphRun::setStrikeOut(bool strikeOut);
func (this *QGlyphRun) setStrikeOut(args ...interface{}) () {
  // setStrikeOut(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QGlyphRun12setStrikeOutEb
  default:
    qtrt.ErrorResolve("QGlyphRun", "setStrikeOut", args)
  }

}

  // proto:  bool QGlyphRun::isEmpty();
func (this *QGlyphRun) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QGlyphRun7isEmptyEv
  default:
    qtrt.ErrorResolve("QGlyphRun", "isEmpty", args)
  }

}

// <= body block end
