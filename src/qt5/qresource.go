package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qresource.h
// dst-file: /src/core/qresource.go
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
  // proto:  void QResource::QResource(const QString & file, const QLocale & locale);
extern void* dector_ZN9QResourceC1ERK7QStringRK7QLocale(void* arg0, void* arg1);
extern void _ZN9QResourceC1ERK7QStringRK7QLocale(void* qthis, void* arg0, void* arg1);
  // proto:  QLocale QResource::locale();
extern void _ZNK9QResource6localeEv(void* qthis);
  // proto:  void QResource::setLocale(const QLocale & locale);
extern void _ZN9QResource9setLocaleERK7QLocale(void* qthis, void* arg0);
  // proto: static bool QResource::registerResource(const uchar * rccData, const QString & resourceRoot);
extern void _ZN9QResource16registerResourceEPKhRK7QString(unsigned char* arg0, void* arg1);
  // proto:  const uchar * QResource::data();
extern void _ZNK9QResource4dataEv(void* qthis);
  // proto: static QStringList QResource::searchPaths();
extern void _ZN9QResource11searchPathsEv();
  // proto:  QString QResource::fileName();
extern void _ZNK9QResource8fileNameEv(void* qthis);
  // proto:  QString QResource::absoluteFilePath();
extern void _ZNK9QResource16absoluteFilePathEv(void* qthis);
  // proto: static bool QResource::unregisterResource(const uchar * rccData, const QString & resourceRoot);
extern void _ZN9QResource18unregisterResourceEPKhRK7QString(unsigned char* arg0, void* arg1);
  // proto: static bool QResource::registerResource(const QString & rccFilename, const QString & resourceRoot);
extern void _ZN9QResource16registerResourceERK7QStringS2_(void* arg0, void* arg1);
  // proto: static void QResource::addSearchPath(const QString & path);
extern void _ZN9QResource13addSearchPathERK7QString(void* arg0);
  // proto:  qint64 QResource::size();
extern void _ZNK9QResource4sizeEv(void* qthis);
  // proto:  void QResource::~QResource();
extern void _ZN9QResourceD0Ev(void* qthis);
  // proto:  bool QResource::isValid();
extern void _ZNK9QResource7isValidEv(void* qthis);
  // proto:  void QResource::setFileName(const QString & file);
extern void _ZN9QResource11setFileNameERK7QString(void* qthis, void* arg0);
  // proto: static bool QResource::unregisterResource(const QString & rccFilename, const QString & resourceRoot);
extern void _ZN9QResource18unregisterResourceERK7QStringS2_(void* arg0, void* arg1);
  // proto:  bool QResource::isCompressed();
extern void _ZNK9QResource12isCompressedEv(void* qthis);
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

// class sizeof(QResource)=1
type QResource struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QResource::QResource(const QString & file, const QLocale & locale);
func NewQResource(args ...interface{}) QResource {
  return QResource{}
}

  // proto:  QLocale QResource::locale();
func (this *QResource) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource6localeEv
  default:
    qtrt.ErrorResolve("QResource", "locale", args)
  }

}

  // proto:  void QResource::setLocale(const QLocale & locale);
func (this *QResource) setLocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResource9setLocaleERK7QLocale
  default:
    qtrt.ErrorResolve("QResource", "setLocale", args)
  }

}

  // proto: static bool QResource::registerResource(const uchar * rccData, const QString & resourceRoot);
func (this *QResource) registerResource_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QResource", "registerResource", args)
  }

}

  // proto:  const uchar * QResource::data();
func (this *QResource) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource4dataEv
  default:
    qtrt.ErrorResolve("QResource", "data", args)
  }

}

  // proto: static QStringList QResource::searchPaths();
func (this *QResource) searchPaths_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QResource", "searchPaths", args)
  }

}

  // proto:  QString QResource::fileName();
func (this *QResource) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource8fileNameEv
  default:
    qtrt.ErrorResolve("QResource", "fileName", args)
  }

}

  // proto:  QString QResource::absoluteFilePath();
func (this *QResource) absoluteFilePath(args ...interface{}) () {
  // absoluteFilePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource16absoluteFilePathEv
  default:
    qtrt.ErrorResolve("QResource", "absoluteFilePath", args)
  }

}

  // proto: static bool QResource::unregisterResource(const uchar * rccData, const QString & resourceRoot);
func (this *QResource) unregisterResource_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QResource", "unregisterResource", args)
  }

}

  // proto: static void QResource::addSearchPath(const QString & path);
func (this *QResource) addSearchPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QResource", "addSearchPath", args)
  }

}

  // proto:  qint64 QResource::size();
func (this *QResource) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource4sizeEv
  default:
    qtrt.ErrorResolve("QResource", "size", args)
  }

}

  // proto:  void QResource::~QResource();
func (this *QResource) FreeQResource(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QResource", "~QResource", args)
  }

}

  // proto:  bool QResource::isValid();
func (this *QResource) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource7isValidEv
  default:
    qtrt.ErrorResolve("QResource", "isValid", args)
  }

}

  // proto:  void QResource::setFileName(const QString & file);
func (this *QResource) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QResource11setFileNameERK7QString
  default:
    qtrt.ErrorResolve("QResource", "setFileName", args)
  }

}

  // proto:  bool QResource::isCompressed();
func (this *QResource) isCompressed(args ...interface{}) () {
  // isCompressed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QResource12isCompressedEv
  default:
    qtrt.ErrorResolve("QResource", "isCompressed", args)
  }

}

// <= body block end
