package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qcollator.h
// dst-file: /src/core/qcollator.go
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
  // proto:  bool QCollator::numericMode();
extern void _ZNK9QCollator11numericModeEv(void* qthis);
  // proto:  void QCollator::QCollator(const QLocale & locale);
extern void* dector_ZN9QCollatorC1ERK7QLocale(void* arg0);
extern void _ZN9QCollatorC1ERK7QLocale(void* qthis, void* arg0);
  // proto:  void QCollator::setLocale(const QLocale & locale);
extern void _ZN9QCollator9setLocaleERK7QLocale(void* qthis, void* arg0);
  // proto:  void QCollator::setNumericMode(bool on);
extern void _ZN9QCollator14setNumericModeEb(void* qthis, bool arg0);
  // proto:  int QCollator::compare(const QChar * s1, int len1, const QChar * s2, int len2);
extern void _ZNK9QCollator7compareEPK5QChariS2_i(void* qthis, void* arg0, int arg1, void* arg2, int arg3);
  // proto:  QCollatorSortKey QCollator::sortKey(const QString & string);
extern void _ZNK9QCollator7sortKeyERK7QString(void* qthis, void* arg0);
  // proto:  int QCollator::compare(const QString & s1, const QString & s2);
extern void _ZNK9QCollator7compareERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QCollator::~QCollator();
extern void _ZN9QCollatorD0Ev(void* qthis);
  // proto:  bool QCollator::ignorePunctuation();
extern void _ZNK9QCollator17ignorePunctuationEv(void* qthis);
  // proto:  void QCollator::QCollator(const QCollator & );
extern void* dector_ZN9QCollatorC1ERKS_(void* arg0);
extern void _ZN9QCollatorC1ERKS_(void* qthis, void* arg0);
  // proto:  QLocale QCollator::locale();
extern void _ZNK9QCollator6localeEv(void* qthis);
  // proto:  void QCollator::swap(QCollator & other);
extern void _ZN9QCollator4swapERS_(void* qthis, void* arg0);
  // proto:  void QCollator::setIgnorePunctuation(bool on);
extern void _ZN9QCollator20setIgnorePunctuationEb(void* qthis, bool arg0);
  // proto:  void QCollatorSortKey::~QCollatorSortKey();
extern void _ZN16QCollatorSortKeyD0Ev(void* qthis);
  // proto:  void QCollatorSortKey::swap(QCollatorSortKey & other);
extern void _ZN16QCollatorSortKey4swapERS_(void* qthis, void* arg0);
  // proto:  int QCollatorSortKey::compare(const QCollatorSortKey & key);
extern void _ZNK16QCollatorSortKey7compareERKS_(void* qthis, void* arg0);
  // proto:  void QCollatorSortKey::QCollatorSortKey(const QCollatorSortKey & other);
extern void* dector_ZN16QCollatorSortKeyC1ERKS_(void* arg0);
extern void _ZN16QCollatorSortKeyC1ERKS_(void* qthis, void* arg0);
  // proto:  void QCollatorSortKey::QCollatorSortKey();
extern void* dector_ZN16QCollatorSortKeyC1Ev();
extern void _ZN16QCollatorSortKeyC1Ev(void* qthis);
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

// class sizeof(QCollator)=8
type QCollator struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QCollatorSortKey)=1
type QCollatorSortKey struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QCollator::numericMode();
func (this *QCollator) numericMode(args ...interface{}) () {
  // numericMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator11numericModeEv
  default:
    qtrt.ErrorResolve("QCollator", "numericMode", args)
  }

}

  // proto:  void QCollator::QCollator(const QLocale & locale);
func NewQCollator(args ...interface{}) QCollator {
  return QCollator{}
}

  // proto:  void QCollator::setLocale(const QLocale & locale);
func (this *QCollator) setLocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator9setLocaleERK7QLocale
  default:
    qtrt.ErrorResolve("QCollator", "setLocale", args)
  }

}

  // proto:  void QCollator::setNumericMode(bool on);
func (this *QCollator) setNumericMode(args ...interface{}) () {
  // setNumericMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator14setNumericModeEb
  default:
    qtrt.ErrorResolve("QCollator", "setNumericMode", args)
  }

}

  // proto:  int QCollator::compare(const QChar * s1, int len1, const QChar * s2, int len2);
func (this *QCollator) compare(args ...interface{}) () {
  // compare(const class QStringRef &, const class QStringRef &)
  // compare(const class QChar *, int, const class QChar *, int)
  // compare(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator7compareERK10QStringRefS2_
  case 1:
    // invoke: _ZNK9QCollator7compareEPK5QChariS2_i
  case 2:
    // invoke: _ZNK9QCollator7compareERK7QStringS2_
  default:
    qtrt.ErrorResolve("QCollator", "compare", args)
  }

}

  // proto:  QCollatorSortKey QCollator::sortKey(const QString & string);
func (this *QCollator) sortKey(args ...interface{}) () {
  // sortKey(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator7sortKeyERK7QString
  default:
    qtrt.ErrorResolve("QCollator", "sortKey", args)
  }

}

  // proto:  void QCollator::~QCollator();
func (this *QCollator) FreeQCollator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCollator", "~QCollator", args)
  }

}

  // proto:  bool QCollator::ignorePunctuation();
func (this *QCollator) ignorePunctuation(args ...interface{}) () {
  // ignorePunctuation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator17ignorePunctuationEv
  default:
    qtrt.ErrorResolve("QCollator", "ignorePunctuation", args)
  }

}

  // proto:  QLocale QCollator::locale();
func (this *QCollator) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QCollator6localeEv
  default:
    qtrt.ErrorResolve("QCollator", "locale", args)
  }

}

  // proto:  void QCollator::swap(QCollator & other);
func (this *QCollator) swap(args ...interface{}) () {
  // swap(class QCollator &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollator{}) // "QCollator &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator4swapERS_
  default:
    qtrt.ErrorResolve("QCollator", "swap", args)
  }

}

  // proto:  void QCollator::setIgnorePunctuation(bool on);
func (this *QCollator) setIgnorePunctuation(args ...interface{}) () {
  // setIgnorePunctuation(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QCollator20setIgnorePunctuationEb
  default:
    qtrt.ErrorResolve("QCollator", "setIgnorePunctuation", args)
  }

}

  // proto:  void QCollatorSortKey::~QCollatorSortKey();
func (this *QCollatorSortKey) FreeQCollatorSortKey(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "~QCollatorSortKey", args)
  }

}

  // proto:  void QCollatorSortKey::swap(QCollatorSortKey & other);
func (this *QCollatorSortKey) swap(args ...interface{}) () {
  // swap(class QCollatorSortKey &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollatorSortKey{}) // "QCollatorSortKey &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCollatorSortKey4swapERS_
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "swap", args)
  }

}

  // proto:  int QCollatorSortKey::compare(const QCollatorSortKey & key);
func (this *QCollatorSortKey) compare(args ...interface{}) () {
  // compare(const class QCollatorSortKey &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCollatorSortKey{}) // "const QCollatorSortKey &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QCollatorSortKey7compareERKS_
  default:
    qtrt.ErrorResolve("QCollatorSortKey", "compare", args)
  }

}

  // proto:  void QCollatorSortKey::QCollatorSortKey(const QCollatorSortKey & other);
func NewQCollatorSortKey(args ...interface{}) QCollatorSortKey {
  return QCollatorSortKey{}
}

// <= body block end
