package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qlist.h
// dst-file: /src/core/qlist.go
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
  // proto:  void ** QListData::prepend();
extern void _ZN9QListData7prependEv(void* qthis);
  // proto:  void QListData::realloc(int alloc);
extern void _ZN9QListData7reallocEi(void* qthis, int arg0);
  // proto:  void ** QListData::end();
extern void demth_ZNK9QListData3endEv(void* qthis);
  // proto:  void QListData::remove(int i, int n);
extern void _ZN9QListData6removeEii(void* qthis, int arg0, int arg1);
  // proto:  void ** QListData::append(const QListData & l);
extern void _ZN9QListData6appendERKS_(void* qthis, void* arg0);
  // proto:  void QListData::remove(int i);
extern void _ZN9QListData6removeEi(void* qthis, int arg0);
  // proto:  void ** QListData::append();
extern void _ZN9QListData6appendEv(void* qthis);
  // proto:  void QListData::dispose();
extern void demth_ZN9QListData7disposeEv(void* qthis);
  // proto:  int QListData::size();
extern void demth_ZNK9QListData4sizeEv(void* qthis);
  // proto:  bool QListData::isEmpty();
extern void demth_ZNK9QListData7isEmptyEv(void* qthis);
  // proto:  void ** QListData::at(int i);
extern void demth_ZNK9QListData2atEi(void* qthis, int arg0);
  // proto:  void ** QListData::erase(void ** xi);
extern void _ZN9QListData5eraseEPPv(void* qthis, void* arg0);
  // proto:  void ** QListData::append(int n);
extern void _ZN9QListData6appendEi(void* qthis, int arg0);
  // proto:  void QListData::move(int from, int to);
extern void _ZN9QListData4moveEii(void* qthis, int arg0, int arg1);
  // proto:  void ** QListData::begin();
extern void demth_ZNK9QListData5beginEv(void* qthis);
  // proto:  void ** QListData::insert(int i);
extern void _ZN9QListData6insertEi(void* qthis, int arg0);
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

// class sizeof(QListData)=8
type QListData struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void ** QListData::prepend();
func (this *QListData) prepend(args ...interface{}) () {
  // prepend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData7prependEv
  default:
    qtrt.ErrorResolve("QListData", "prepend", args)
  }

}

  // proto:  void QListData::realloc(int alloc);
func (this *QListData) realloc(args ...interface{}) () {
  // realloc(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData7reallocEi
  default:
    qtrt.ErrorResolve("QListData", "realloc", args)
  }

}

  // proto:  void ** QListData::end();
func (this *QListData) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData3endEv
  default:
    qtrt.ErrorResolve("QListData", "end", args)
  }

}

  // proto:  void QListData::remove(int i, int n);
func (this *QListData) remove(args ...interface{}) () {
  // remove(int, int)
  // remove(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData6removeEii
  case 1:
    // invoke: _ZN9QListData6removeEi
  default:
    qtrt.ErrorResolve("QListData", "remove", args)
  }

}

  // proto:  void ** QListData::append(const QListData & l);
func (this *QListData) append(args ...interface{}) () {
  // append(const struct QListData &)
  // append()
  // append(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListData{}) // "const QListData &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData6appendERKS_
  case 1:
    // invoke: _ZN9QListData6appendEv
  case 2:
    // invoke: _ZN9QListData6appendEi
  default:
    qtrt.ErrorResolve("QListData", "append", args)
  }

}

  // proto:  void QListData::dispose();
func (this *QListData) dispose(args ...interface{}) () {
  // dispose(struct QListData::Data *)
  // dispose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  // vtys[0][0] = reflect.TypeOf(QListData::Data{}) // "QListData::Data *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData7disposeEPNS_4DataE
  case 1:
    // invoke: _ZN9QListData7disposeEv
  default:
    qtrt.ErrorResolve("QListData", "dispose", args)
  }

}

  // proto:  int QListData::size();
func (this *QListData) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData4sizeEv
  default:
    qtrt.ErrorResolve("QListData", "size", args)
  }

}

  // proto:  bool QListData::isEmpty();
func (this *QListData) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData7isEmptyEv
  default:
    qtrt.ErrorResolve("QListData", "isEmpty", args)
  }

}

  // proto:  void ** QListData::at(int i);
func (this *QListData) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData2atEi
  default:
    qtrt.ErrorResolve("QListData", "at", args)
  }

}

  // proto:  void ** QListData::erase(void ** xi);
func (this *QListData) erase(args ...interface{}) () {
  // erase(void **)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "void **"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData5eraseEPPv
  default:
    qtrt.ErrorResolve("QListData", "erase", args)
  }

}

  // proto:  void QListData::move(int from, int to);
func (this *QListData) move_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QListData", "move", args)
  }

}

  // proto:  void ** QListData::begin();
func (this *QListData) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QListData5beginEv
  default:
    qtrt.ErrorResolve("QListData", "begin", args)
  }

}

  // proto:  void ** QListData::insert(int i);
func (this *QListData) insert(args ...interface{}) () {
  // insert(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QListData6insertEi
  default:
    qtrt.ErrorResolve("QListData", "insert", args)
  }

}

// <= body block end
