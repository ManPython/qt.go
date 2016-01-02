package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qthreadpool.h
// dst-file: /src/core/qthreadpool.go
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
  // proto:  void QThreadPool::~QThreadPool();
extern void _ZN11QThreadPoolD0Ev(void* qthis);
  // proto:  int QThreadPool::expiryTimeout();
extern void _ZNK11QThreadPool13expiryTimeoutEv(void* qthis);
  // proto:  bool QThreadPool::waitForDone(int msecs);
extern void _ZN11QThreadPool11waitForDoneEi(void* qthis, int arg0);
  // proto:  const QMetaObject * QThreadPool::metaObject();
extern void _ZNK11QThreadPool10metaObjectEv(void* qthis);
  // proto:  void QThreadPool::cancel(QRunnable * runnable);
extern void _ZN11QThreadPool6cancelEP9QRunnable(void* qthis, void* arg0);
  // proto:  bool QThreadPool::tryStart(QRunnable * runnable);
extern void _ZN11QThreadPool8tryStartEP9QRunnable(void* qthis, void* arg0);
  // proto: static QThreadPool * QThreadPool::globalInstance();
extern void _ZN11QThreadPool14globalInstanceEv();
  // proto:  void QThreadPool::setMaxThreadCount(int maxThreadCount);
extern void _ZN11QThreadPool17setMaxThreadCountEi(void* qthis, int arg0);
  // proto:  void QThreadPool::setExpiryTimeout(int expiryTimeout);
extern void _ZN11QThreadPool16setExpiryTimeoutEi(void* qthis, int arg0);
  // proto:  void QThreadPool::reserveThread();
extern void _ZN11QThreadPool13reserveThreadEv(void* qthis);
  // proto:  void QThreadPool::clear();
extern void _ZN11QThreadPool5clearEv(void* qthis);
  // proto:  void QThreadPool::QThreadPool(QObject * parent);
extern void* dector_ZN11QThreadPoolC1EP7QObject(void* arg0);
extern void _ZN11QThreadPoolC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QThreadPool::start(QRunnable * runnable, int priority);
extern void _ZN11QThreadPool5startEP9QRunnablei(void* qthis, void* arg0, int arg1);
  // proto:  int QThreadPool::maxThreadCount();
extern void _ZNK11QThreadPool14maxThreadCountEv(void* qthis);
  // proto:  void QThreadPool::releaseThread();
extern void _ZN11QThreadPool13releaseThreadEv(void* qthis);
  // proto:  int QThreadPool::activeThreadCount();
extern void _ZNK11QThreadPool17activeThreadCountEv(void* qthis);
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

// class sizeof(QThreadPool)=1
type QThreadPool struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QThreadPool::~QThreadPool();
func (this *QThreadPool) FreeQThreadPool(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QThreadPool", "~QThreadPool", args)
  }

}

  // proto:  int QThreadPool::expiryTimeout();
func (this *QThreadPool) expiryTimeout(args ...interface{}) () {
  // expiryTimeout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QThreadPool13expiryTimeoutEv
  default:
    qtrt.ErrorResolve("QThreadPool", "expiryTimeout", args)
  }

}

  // proto:  bool QThreadPool::waitForDone(int msecs);
func (this *QThreadPool) waitForDone(args ...interface{}) () {
  // waitForDone(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool11waitForDoneEi
  default:
    qtrt.ErrorResolve("QThreadPool", "waitForDone", args)
  }

}

  // proto:  const QMetaObject * QThreadPool::metaObject();
func (this *QThreadPool) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QThreadPool10metaObjectEv
  default:
    qtrt.ErrorResolve("QThreadPool", "metaObject", args)
  }

}

  // proto:  void QThreadPool::cancel(QRunnable * runnable);
func (this *QThreadPool) cancel(args ...interface{}) () {
  // cancel(class QRunnable *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRunnable{}) // "QRunnable *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool6cancelEP9QRunnable
  default:
    qtrt.ErrorResolve("QThreadPool", "cancel", args)
  }

}

  // proto:  bool QThreadPool::tryStart(QRunnable * runnable);
func (this *QThreadPool) tryStart(args ...interface{}) () {
  // tryStart(class QRunnable *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRunnable{}) // "QRunnable *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool8tryStartEP9QRunnable
  default:
    qtrt.ErrorResolve("QThreadPool", "tryStart", args)
  }

}

  // proto: static QThreadPool * QThreadPool::globalInstance();
func (this *QThreadPool) globalInstance_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QThreadPool", "globalInstance", args)
  }

}

  // proto:  void QThreadPool::setMaxThreadCount(int maxThreadCount);
func (this *QThreadPool) setMaxThreadCount(args ...interface{}) () {
  // setMaxThreadCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool17setMaxThreadCountEi
  default:
    qtrt.ErrorResolve("QThreadPool", "setMaxThreadCount", args)
  }

}

  // proto:  void QThreadPool::setExpiryTimeout(int expiryTimeout);
func (this *QThreadPool) setExpiryTimeout(args ...interface{}) () {
  // setExpiryTimeout(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool16setExpiryTimeoutEi
  default:
    qtrt.ErrorResolve("QThreadPool", "setExpiryTimeout", args)
  }

}

  // proto:  void QThreadPool::reserveThread();
func (this *QThreadPool) reserveThread(args ...interface{}) () {
  // reserveThread()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool13reserveThreadEv
  default:
    qtrt.ErrorResolve("QThreadPool", "reserveThread", args)
  }

}

  // proto:  void QThreadPool::clear();
func (this *QThreadPool) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool5clearEv
  default:
    qtrt.ErrorResolve("QThreadPool", "clear", args)
  }

}

  // proto:  void QThreadPool::QThreadPool(QObject * parent);
func NewQThreadPool(args ...interface{}) QThreadPool {
  return QThreadPool{}
}

  // proto:  void QThreadPool::start(QRunnable * runnable, int priority);
func (this *QThreadPool) start(args ...interface{}) () {
  // start(class QRunnable *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRunnable{}) // "QRunnable *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool5startEP9QRunnablei
  default:
    qtrt.ErrorResolve("QThreadPool", "start", args)
  }

}

  // proto:  int QThreadPool::maxThreadCount();
func (this *QThreadPool) maxThreadCount(args ...interface{}) () {
  // maxThreadCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QThreadPool14maxThreadCountEv
  default:
    qtrt.ErrorResolve("QThreadPool", "maxThreadCount", args)
  }

}

  // proto:  void QThreadPool::releaseThread();
func (this *QThreadPool) releaseThread(args ...interface{}) () {
  // releaseThread()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QThreadPool13releaseThreadEv
  default:
    qtrt.ErrorResolve("QThreadPool", "releaseThread", args)
  }

}

  // proto:  int QThreadPool::activeThreadCount();
func (this *QThreadPool) activeThreadCount(args ...interface{}) () {
  // activeThreadCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QThreadPool17activeThreadCountEv
  default:
    qtrt.ErrorResolve("QThreadPool", "activeThreadCount", args)
  }

}

// <= body block end
