package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qabstractslider.h
// dst-file: /src/widgets/qabstractslider.go
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
  // proto:  void QAbstractSlider::setSliderPosition(int );
extern void _ZN15QAbstractSlider17setSliderPositionEi(void* qthis, int arg0);
  // proto:  bool QAbstractSlider::isSliderDown();
extern void _ZNK15QAbstractSlider12isSliderDownEv(void* qthis);
  // proto:  int QAbstractSlider::value();
extern void _ZNK15QAbstractSlider5valueEv(void* qthis);
  // proto:  void QAbstractSlider::setInvertedControls(bool );
extern void _ZN15QAbstractSlider19setInvertedControlsEb(void* qthis, bool arg0);
  // proto:  void QAbstractSlider::QAbstractSlider(const QAbstractSlider & );
extern void* dector_ZN15QAbstractSliderC1ERKS_(void* arg0);
extern void _ZN15QAbstractSliderC1ERKS_(void* qthis, void* arg0);
  // proto:  int QAbstractSlider::minimum();
extern void _ZNK15QAbstractSlider7minimumEv(void* qthis);
  // proto:  int QAbstractSlider::singleStep();
extern void _ZNK15QAbstractSlider10singleStepEv(void* qthis);
  // proto:  int QAbstractSlider::pageStep();
extern void _ZNK15QAbstractSlider8pageStepEv(void* qthis);
  // proto:  void QAbstractSlider::setMaximum(int );
extern void _ZN15QAbstractSlider10setMaximumEi(void* qthis, int arg0);
  // proto:  bool QAbstractSlider::invertedControls();
extern void _ZNK15QAbstractSlider16invertedControlsEv(void* qthis);
  // proto:  void QAbstractSlider::setValue(int );
extern void _ZN15QAbstractSlider8setValueEi(void* qthis, int arg0);
  // proto:  void QAbstractSlider::~QAbstractSlider();
extern void _ZN15QAbstractSliderD0Ev(void* qthis);
  // proto:  void QAbstractSlider::setPageStep(int );
extern void _ZN15QAbstractSlider11setPageStepEi(void* qthis, int arg0);
  // proto:  void QAbstractSlider::setSliderDown(bool );
extern void _ZN15QAbstractSlider13setSliderDownEb(void* qthis, bool arg0);
  // proto:  int QAbstractSlider::maximum();
extern void _ZNK15QAbstractSlider7maximumEv(void* qthis);
  // proto:  const QMetaObject * QAbstractSlider::metaObject();
extern void _ZNK15QAbstractSlider10metaObjectEv(void* qthis);
  // proto:  void QAbstractSlider::setSingleStep(int );
extern void _ZN15QAbstractSlider13setSingleStepEi(void* qthis, int arg0);
  // proto:  void QAbstractSlider::setInvertedAppearance(bool );
extern void _ZN15QAbstractSlider21setInvertedAppearanceEb(void* qthis, bool arg0);
  // proto:  bool QAbstractSlider::hasTracking();
extern void _ZNK15QAbstractSlider11hasTrackingEv(void* qthis);
  // proto:  bool QAbstractSlider::invertedAppearance();
extern void _ZNK15QAbstractSlider18invertedAppearanceEv(void* qthis);
  // proto:  int QAbstractSlider::sliderPosition();
extern void _ZNK15QAbstractSlider14sliderPositionEv(void* qthis);
  // proto:  void QAbstractSlider::setTracking(bool enable);
extern void _ZN15QAbstractSlider11setTrackingEb(void* qthis, bool arg0);
  // proto:  void QAbstractSlider::QAbstractSlider(QWidget * parent);
extern void* dector_ZN15QAbstractSliderC1EP7QWidget(void* arg0);
extern void _ZN15QAbstractSliderC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QAbstractSlider::setRange(int min, int max);
extern void _ZN15QAbstractSlider8setRangeEii(void* qthis, int arg0, int arg1);
  // proto:  void QAbstractSlider::setMinimum(int );
extern void _ZN15QAbstractSlider10setMinimumEi(void* qthis, int arg0);
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

// class sizeof(QAbstractSlider)=1
type QAbstractSlider struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _sliderReleased QAbstractSlider_sliderReleased_signal;
//  _rangeChanged QAbstractSlider_rangeChanged_signal;
//  _sliderPressed QAbstractSlider_sliderPressed_signal;
//  _actionTriggered QAbstractSlider_actionTriggered_signal;
//  _valueChanged QAbstractSlider_valueChanged_signal;
//  _sliderMoved QAbstractSlider_sliderMoved_signal;
}

  // proto:  void QAbstractSlider::setSliderPosition(int );
func (this *QAbstractSlider) setSliderPosition(args ...interface{}) () {
  // setSliderPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider17setSliderPositionEi
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSliderPosition", args)
  }

}

  // proto:  bool QAbstractSlider::isSliderDown();
func (this *QAbstractSlider) isSliderDown(args ...interface{}) () {
  // isSliderDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider12isSliderDownEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "isSliderDown", args)
  }

}

  // proto:  int QAbstractSlider::value();
func (this *QAbstractSlider) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider5valueEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "value", args)
  }

}

  // proto:  void QAbstractSlider::setInvertedControls(bool );
func (this *QAbstractSlider) setInvertedControls(args ...interface{}) () {
  // setInvertedControls(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider19setInvertedControlsEb
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setInvertedControls", args)
  }

}

  // proto:  void QAbstractSlider::QAbstractSlider(const QAbstractSlider & );
func NewQAbstractSlider(args ...interface{}) QAbstractSlider {
  return QAbstractSlider{}
}

  // proto:  int QAbstractSlider::minimum();
func (this *QAbstractSlider) minimum(args ...interface{}) () {
  // minimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider7minimumEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "minimum", args)
  }

}

  // proto:  int QAbstractSlider::singleStep();
func (this *QAbstractSlider) singleStep(args ...interface{}) () {
  // singleStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider10singleStepEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "singleStep", args)
  }

}

  // proto:  int QAbstractSlider::pageStep();
func (this *QAbstractSlider) pageStep(args ...interface{}) () {
  // pageStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider8pageStepEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "pageStep", args)
  }

}

  // proto:  void QAbstractSlider::setMaximum(int );
func (this *QAbstractSlider) setMaximum(args ...interface{}) () {
  // setMaximum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider10setMaximumEi
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setMaximum", args)
  }

}

  // proto:  bool QAbstractSlider::invertedControls();
func (this *QAbstractSlider) invertedControls(args ...interface{}) () {
  // invertedControls()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider16invertedControlsEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "invertedControls", args)
  }

}

  // proto:  void QAbstractSlider::setValue(int );
func (this *QAbstractSlider) setValue(args ...interface{}) () {
  // setValue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider8setValueEi
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setValue", args)
  }

}

  // proto:  void QAbstractSlider::~QAbstractSlider();
func (this *QAbstractSlider) FreeQAbstractSlider(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractSlider", "~QAbstractSlider", args)
  }

}

  // proto:  void QAbstractSlider::setPageStep(int );
func (this *QAbstractSlider) setPageStep(args ...interface{}) () {
  // setPageStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider11setPageStepEi
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setPageStep", args)
  }

}

  // proto:  void QAbstractSlider::setSliderDown(bool );
func (this *QAbstractSlider) setSliderDown(args ...interface{}) () {
  // setSliderDown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider13setSliderDownEb
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSliderDown", args)
  }

}

  // proto:  int QAbstractSlider::maximum();
func (this *QAbstractSlider) maximum(args ...interface{}) () {
  // maximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider7maximumEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "maximum", args)
  }

}

  // proto:  const QMetaObject * QAbstractSlider::metaObject();
func (this *QAbstractSlider) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "metaObject", args)
  }

}

  // proto:  void QAbstractSlider::setSingleStep(int );
func (this *QAbstractSlider) setSingleStep(args ...interface{}) () {
  // setSingleStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider13setSingleStepEi
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setSingleStep", args)
  }

}

  // proto:  void QAbstractSlider::setInvertedAppearance(bool );
func (this *QAbstractSlider) setInvertedAppearance(args ...interface{}) () {
  // setInvertedAppearance(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider21setInvertedAppearanceEb
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setInvertedAppearance", args)
  }

}

  // proto:  bool QAbstractSlider::hasTracking();
func (this *QAbstractSlider) hasTracking(args ...interface{}) () {
  // hasTracking()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider11hasTrackingEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "hasTracking", args)
  }

}

  // proto:  bool QAbstractSlider::invertedAppearance();
func (this *QAbstractSlider) invertedAppearance(args ...interface{}) () {
  // invertedAppearance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider18invertedAppearanceEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "invertedAppearance", args)
  }

}

  // proto:  int QAbstractSlider::sliderPosition();
func (this *QAbstractSlider) sliderPosition(args ...interface{}) () {
  // sliderPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QAbstractSlider14sliderPositionEv
  default:
    qtrt.ErrorResolve("QAbstractSlider", "sliderPosition", args)
  }

}

  // proto:  void QAbstractSlider::setTracking(bool enable);
func (this *QAbstractSlider) setTracking(args ...interface{}) () {
  // setTracking(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider11setTrackingEb
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setTracking", args)
  }

}

  // proto:  void QAbstractSlider::setRange(int min, int max);
func (this *QAbstractSlider) setRange(args ...interface{}) () {
  // setRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider8setRangeEii
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setRange", args)
  }

}

  // proto:  void QAbstractSlider::setMinimum(int );
func (this *QAbstractSlider) setMinimum(args ...interface{}) () {
  // setMinimum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QAbstractSlider10setMinimumEi
  default:
    qtrt.ErrorResolve("QAbstractSlider", "setMinimum", args)
  }

}

// <= body block end
