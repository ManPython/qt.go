package qtgui

// /usr/include/qt/QtGui/qwindow.h
// #include <qwindow.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QWindow struct {
	*qtcore.QObject
	*QSurface
}

func (this *QWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWindow) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QSurface = NewQSurfaceFromPointer(cthis)
}
func NewQWindowFromPointer(cthis unsafe.Pointer) *QWindow {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQSurfaceFromPointer(cthis)
	return &QWindow{bcthis0, bcthis1}
}
func (*QWindow) NewFromPointer(cthis unsafe.Pointer) *QWindow {
	return NewQWindowFromPointer(cthis)
}

// /usr/include/qt/QtGui/qwindow.h:97
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:144
// index:0
// Public
// void QWindow(QScreen *)
func NewQWindow(screen *QScreen /*777 QScreen **/) *QWindow {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindowC2EP7QScreen", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWindowFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qwindow.h:145
// index:1
// Public
// void QWindow(QWindow *)
func NewQWindow_1(parent *QWindow /*777 QWindow **/) *QWindow {
	cthis := qtrt.Calloc(1, 256) // 40
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindowC2EPS_", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWindowFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qwindow.h:146
// index:0
// Public virtual
// void ~QWindow()
func DeleteQWindow(*QWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:148
// index:0
// Public
// void setSurfaceType(QSurface::SurfaceType)
func (this *QWindow) SetSurfaceType(surfaceType int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14setSurfaceTypeEN8QSurface11SurfaceTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), surfaceType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:149
// index:0
// Public virtual
// QSurface::SurfaceType surfaceType()
func (this *QWindow) SurfaceType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11surfaceTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:151
// index:0
// Public
// bool isVisible()
func (this *QWindow) IsVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:153
// index:0
// Public
// QWindow::Visibility visibility()
func (this *QWindow) Visibility() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow10visibilityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:154
// index:0
// Public
// void setVisibility(QWindow::Visibility)
func (this *QWindow) SetVisibility(v int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13setVisibilityENS_10VisibilityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), v)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:156
// index:0
// Public
// void create()
func (this *QWindow) Create() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow6createEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:158
// index:0
// Public
// WId winId()
func (this *QWindow) WinId() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow5winIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtGui/qwindow.h:160
// index:0
// Public
// QWindow * parent(QWindow::AncestorMode)
func (this *QWindow) Parent(mode int) *QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6parentENS_12AncestorModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:161
// index:1
// Public
// QWindow * parent()
func (this *QWindow) Parent_1() *QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:162
// index:0
// Public
// void setParent(QWindow *)
func (this *QWindow) SetParent(parent *QWindow /*777 QWindow **/) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setParentEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:164
// index:0
// Public
// bool isTopLevel()
func (this *QWindow) IsTopLevel() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow10isTopLevelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:166
// index:0
// Public
// bool isModal()
func (this *QWindow) IsModal() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow7isModalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:167
// index:0
// Public
// Qt::WindowModality modality()
func (this *QWindow) Modality() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8modalityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:168
// index:0
// Public
// void setModality(Qt::WindowModality)
func (this *QWindow) SetModality(modality int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setModalityEN2Qt14WindowModalityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), modality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:170
// index:0
// Public
// void setFormat(const QSurfaceFormat &)
func (this *QWindow) SetFormat(format *QSurfaceFormat) {
	var convArg0 = format.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setFormatERK14QSurfaceFormat", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:171
// index:0
// Public virtual
// QSurfaceFormat format()
func (this *QWindow) Format() *QSurfaceFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6formatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:172
// index:0
// Public
// QSurfaceFormat requestedFormat()
func (this *QWindow) RequestedFormat() *QSurfaceFormat /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow15requestedFormatEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:174
// index:0
// Public
// void setFlags(Qt::WindowFlags)
func (this *QWindow) SetFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow8setFlagsE6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:175
// index:0
// Public
// Qt::WindowFlags flags()
func (this *QWindow) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:176
// index:0
// Public
// void setFlag(Qt::WindowType, bool)
func (this *QWindow) SetFlag(arg0 int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow7setFlagEN2Qt10WindowTypeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:177
// index:0
// Public
// Qt::WindowType type()
func (this *QWindow) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:179
// index:0
// Public
// QString title()
func (this *QWindow) Title() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow5titleEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:181
// index:0
// Public
// void setOpacity(qreal)
func (this *QWindow) SetOpacity(level float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow10setOpacityEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), level)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:182
// index:0
// Public
// qreal opacity()
func (this *QWindow) Opacity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow7opacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qwindow.h:184
// index:0
// Public
// void setMask(const QRegion &)
func (this *QWindow) SetMask(region *QRegion) {
	var convArg0 = region.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow7setMaskERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:185
// index:0
// Public
// QRegion mask()
func (this *QWindow) Mask() *QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow4maskEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:187
// index:0
// Public
// bool isActive()
func (this *QWindow) IsActive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8isActiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:189
// index:0
// Public
// void reportContentOrientationChange(Qt::ScreenOrientation)
func (this *QWindow) ReportContentOrientationChange(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow30reportContentOrientationChangeEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:190
// index:0
// Public
// Qt::ScreenOrientation contentOrientation()
func (this *QWindow) ContentOrientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow18contentOrientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:192
// index:0
// Public
// qreal devicePixelRatio()
func (this *QWindow) DevicePixelRatio() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow16devicePixelRatioEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qwindow.h:194
// index:0
// Public
// Qt::WindowState windowState()
func (this *QWindow) WindowState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11windowStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:195
// index:0
// Public
// Qt::WindowStates windowStates()
func (this *QWindow) WindowStates() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12windowStatesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:196
// index:0
// Public
// void setWindowState(Qt::WindowState)
func (this *QWindow) SetWindowState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14setWindowStateEN2Qt11WindowStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:199
// index:0
// Public
// void setTransientParent(QWindow *)
func (this *QWindow) SetTransientParent(parent *QWindow /*777 QWindow **/) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow18setTransientParentEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:200
// index:0
// Public
// QWindow * transientParent()
func (this *QWindow) TransientParent() *QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow15transientParentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:202
// index:0
// Public
// bool isAncestorOf(const QWindow *, QWindow::AncestorMode)
func (this *QWindow) IsAncestorOf(child *QWindow /*777 const QWindow **/, mode int) bool {
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12isAncestorOfEPKS_NS_12AncestorModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:204
// index:0
// Public
// bool isExposed()
func (this *QWindow) IsExposed() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow9isExposedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:206
// index:0
// Public inline
// int minimumWidth()
func (this *QWindow) MinimumWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12minimumWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qwindow.h:207
// index:0
// Public inline
// int minimumHeight()
func (this *QWindow) MinimumHeight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13minimumHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qwindow.h:208
// index:0
// Public inline
// int maximumWidth()
func (this *QWindow) MaximumWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12maximumWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qwindow.h:209
// index:0
// Public inline
// int maximumHeight()
func (this *QWindow) MaximumHeight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13maximumHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qwindow.h:211
// index:0
// Public
// QSize minimumSize()
func (this *QWindow) MinimumSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:212
// index:0
// Public
// QSize maximumSize()
func (this *QWindow) MaximumSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:213
// index:0
// Public
// QSize baseSize()
func (this *QWindow) BaseSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8baseSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:214
// index:0
// Public
// QSize sizeIncrement()
func (this *QWindow) SizeIncrement() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13sizeIncrementEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:216
// index:0
// Public
// void setMinimumSize(const QSize &)
func (this *QWindow) SetMinimumSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14setMinimumSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:217
// index:0
// Public
// void setMaximumSize(const QSize &)
func (this *QWindow) SetMaximumSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14setMaximumSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:218
// index:0
// Public
// void setBaseSize(const QSize &)
func (this *QWindow) SetBaseSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setBaseSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:219
// index:0
// Public
// void setSizeIncrement(const QSize &)
func (this *QWindow) SetSizeIncrement(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow16setSizeIncrementERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:221
// index:0
// Public
// QRect geometry()
func (this *QWindow) Geometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8geometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:223
// index:0
// Public
// QMargins frameMargins()
func (this *QWindow) FrameMargins() *qtcore.QMargins /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12frameMarginsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:224
// index:0
// Public
// QRect frameGeometry()
func (this *QWindow) FrameGeometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13frameGeometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:226
// index:0
// Public
// QPoint framePosition()
func (this *QWindow) FramePosition() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13framePositionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:227
// index:0
// Public
// void setFramePosition(const QPoint &)
func (this *QWindow) SetFramePosition(point *qtcore.QPoint) {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow16setFramePositionERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:229
// index:0
// Public inline
// int width()
func (this *QWindow) Width() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qwindow.h:230
// index:0
// Public inline
// int height()
func (this *QWindow) Height() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qwindow.h:231
// index:0
// Public inline
// int x()
func (this *QWindow) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qwindow.h:232
// index:0
// Public inline
// int y()
func (this *QWindow) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qwindow.h:234
// index:0
// Public inline virtual
// QSize size()
func (this *QWindow) Size() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:235
// index:0
// Public inline
// QPoint position()
func (this *QWindow) Position() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8positionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:237
// index:0
// Public
// void setPosition(const QPoint &)
func (this *QWindow) SetPosition(pt *qtcore.QPoint) {
	var convArg0 = pt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setPositionERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:238
// index:1
// Public
// void setPosition(int, int)
func (this *QWindow) SetPosition_1(posx int, posy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setPositionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), posx, posy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:240
// index:0
// Public
// void resize(const QSize &)
func (this *QWindow) Resize(newSize *qtcore.QSize) {
	var convArg0 = newSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow6resizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:241
// index:1
// Public
// void resize(int, int)
func (this *QWindow) Resize_1(w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow6resizeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:243
// index:0
// Public
// void setFilePath(const QString &)
func (this *QWindow) SetFilePath(filePath *qtcore.QString) {
	var convArg0 = filePath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setFilePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:244
// index:0
// Public
// QString filePath()
func (this *QWindow) FilePath() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8filePathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:246
// index:0
// Public
// void setIcon(const QIcon &)
func (this *QWindow) SetIcon(icon *QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow7setIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:247
// index:0
// Public
// QIcon icon()
func (this *QWindow) Icon() *QIcon /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow4iconEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:249
// index:0
// Public
// void destroy()
func (this *QWindow) Destroy() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow7destroyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:253
// index:0
// Public
// bool setKeyboardGrabEnabled(bool)
func (this *QWindow) SetKeyboardGrabEnabled(grab bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow22setKeyboardGrabEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), grab)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:254
// index:0
// Public
// bool setMouseGrabEnabled(bool)
func (this *QWindow) SetMouseGrabEnabled(grab bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow19setMouseGrabEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), grab)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:256
// index:0
// Public
// QScreen * screen()
func (this *QWindow) Screen() *QScreen /*777 QScreen **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6screenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:257
// index:0
// Public
// void setScreen(QScreen *)
func (this *QWindow) SetScreen(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setScreenEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:259
// index:0
// Public virtual
// QAccessibleInterface * accessibleRoot()
func (this *QWindow) AccessibleRoot() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow14accessibleRootEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:260
// index:0
// Public virtual
// QObject * focusObject()
func (this *QWindow) FocusObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11focusObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:262
// index:0
// Public
// QPoint mapToGlobal(const QPoint &)
func (this *QWindow) MapToGlobal(pos *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11mapToGlobalERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:263
// index:0
// Public
// QPoint mapFromGlobal(const QPoint &)
func (this *QWindow) MapFromGlobal(pos *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13mapFromGlobalERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:266
// index:0
// Public
// QCursor cursor()
func (this *QWindow) Cursor() *QCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6cursorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:267
// index:0
// Public
// void setCursor(const QCursor &)
func (this *QWindow) SetCursor(arg0 *QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setCursorERK7QCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:268
// index:0
// Public
// void unsetCursor()
func (this *QWindow) UnsetCursor() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11unsetCursorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:271
// index:0
// Public static
// QWindow * fromWinId(WId)
func (this *QWindow) FromWinId(id uint64) *QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9fromWinIdEy", ffiqt.FFI_TYPE_POINTER, id)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QWindow_FromWinId(id uint64) *QWindow /*777 QWindow **/ {
	var nilthis *QWindow
	rv := nilthis.FromWinId(id)
	return rv
}

// /usr/include/qt/QtGui/qwindow.h:279
// index:0
// Public
// void requestActivate()
func (this *QWindow) RequestActivate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15requestActivateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:281
// index:0
// Public
// void setVisible(bool)
func (this *QWindow) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:283
// index:0
// Public
// void show()
func (this *QWindow) Show() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow4showEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:284
// index:0
// Public
// void hide()
func (this *QWindow) Hide() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow4hideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:286
// index:0
// Public
// void showMinimized()
func (this *QWindow) ShowMinimized() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13showMinimizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:287
// index:0
// Public
// void showMaximized()
func (this *QWindow) ShowMaximized() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13showMaximizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:288
// index:0
// Public
// void showFullScreen()
func (this *QWindow) ShowFullScreen() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14showFullScreenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:289
// index:0
// Public
// void showNormal()
func (this *QWindow) ShowNormal() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow10showNormalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:291
// index:0
// Public
// bool close()
func (this *QWindow) Close() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:292
// index:0
// Public
// void raise()
func (this *QWindow) Raise() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow5raiseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:293
// index:0
// Public
// void lower()
func (this *QWindow) Lower() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow5lowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:295
// index:0
// Public
// void setTitle(const QString &)
func (this *QWindow) SetTitle(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow8setTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:297
// index:0
// Public
// void setX(int)
func (this *QWindow) SetX(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow4setXEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:298
// index:0
// Public
// void setY(int)
func (this *QWindow) SetY(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow4setYEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:299
// index:0
// Public
// void setWidth(int)
func (this *QWindow) SetWidth(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow8setWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:300
// index:0
// Public
// void setHeight(int)
func (this *QWindow) SetHeight(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:301
// index:0
// Public
// void setGeometry(int, int, int, int)
func (this *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setGeometryEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), posx, posy, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:302
// index:1
// Public
// void setGeometry(const QRect &)
func (this *QWindow) SetGeometry_1(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:304
// index:0
// Public
// void setMinimumWidth(int)
func (this *QWindow) SetMinimumWidth(w int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15setMinimumWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:305
// index:0
// Public
// void setMinimumHeight(int)
func (this *QWindow) SetMinimumHeight(h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow16setMinimumHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:306
// index:0
// Public
// void setMaximumWidth(int)
func (this *QWindow) SetMaximumWidth(w int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15setMaximumWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:307
// index:0
// Public
// void setMaximumHeight(int)
func (this *QWindow) SetMaximumHeight(h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow16setMaximumHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:309
// index:0
// Public
// void alert(int)
func (this *QWindow) Alert(msec int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow5alertEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:311
// index:0
// Public
// void requestUpdate()
func (this *QWindow) RequestUpdate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13requestUpdateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:314
// index:0
// Public
// void screenChanged(QScreen *)
func (this *QWindow) ScreenChanged(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13screenChangedEP7QScreen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:315
// index:0
// Public
// void modalityChanged(Qt::WindowModality)
func (this *QWindow) ModalityChanged(modality int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15modalityChangedEN2Qt14WindowModalityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), modality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:316
// index:0
// Public
// void windowStateChanged(Qt::WindowState)
func (this *QWindow) WindowStateChanged(windowState int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow18windowStateChangedEN2Qt11WindowStateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), windowState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:317
// index:0
// Public
// void windowTitleChanged(const QString &)
func (this *QWindow) WindowTitleChanged(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow18windowTitleChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:319
// index:0
// Public
// void xChanged(int)
func (this *QWindow) XChanged(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow8xChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:320
// index:0
// Public
// void yChanged(int)
func (this *QWindow) YChanged(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow8yChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:322
// index:0
// Public
// void widthChanged(int)
func (this *QWindow) WidthChanged(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow12widthChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:323
// index:0
// Public
// void heightChanged(int)
func (this *QWindow) HeightChanged(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13heightChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:325
// index:0
// Public
// void minimumWidthChanged(int)
func (this *QWindow) MinimumWidthChanged(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow19minimumWidthChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:326
// index:0
// Public
// void minimumHeightChanged(int)
func (this *QWindow) MinimumHeightChanged(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow20minimumHeightChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:327
// index:0
// Public
// void maximumWidthChanged(int)
func (this *QWindow) MaximumWidthChanged(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow19maximumWidthChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:328
// index:0
// Public
// void maximumHeightChanged(int)
func (this *QWindow) MaximumHeightChanged(arg int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow20maximumHeightChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:330
// index:0
// Public
// void visibleChanged(bool)
func (this *QWindow) VisibleChanged(arg bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14visibleChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:331
// index:0
// Public
// void visibilityChanged(QWindow::Visibility)
func (this *QWindow) VisibilityChanged(visibility int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow17visibilityChangedENS_10VisibilityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visibility)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:332
// index:0
// Public
// void activeChanged()
func (this *QWindow) ActiveChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13activeChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:333
// index:0
// Public
// void contentOrientationChanged(Qt::ScreenOrientation)
func (this *QWindow) ContentOrientationChanged(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow25contentOrientationChangedEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:335
// index:0
// Public
// void focusObjectChanged(QObject *)
func (this *QWindow) FocusObjectChanged(object *qtcore.QObject /*777 QObject **/) {
	var convArg0 = object.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow18focusObjectChangedEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:337
// index:0
// Public
// void opacityChanged(qreal)
func (this *QWindow) OpacityChanged(opacity float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14opacityChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:340
// index:0
// Protected virtual
// void exposeEvent(QExposeEvent *)
func (this *QWindow) ExposeEvent(arg0 *QExposeEvent /*777 QExposeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11exposeEventEP12QExposeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:341
// index:0
// Protected virtual
// void resizeEvent(QResizeEvent *)
func (this *QWindow) ResizeEvent(arg0 *QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:342
// index:0
// Protected virtual
// void moveEvent(QMoveEvent *)
func (this *QWindow) MoveEvent(arg0 *QMoveEvent /*777 QMoveEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9moveEventEP10QMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:343
// index:0
// Protected virtual
// void focusInEvent(QFocusEvent *)
func (this *QWindow) FocusInEvent(arg0 *QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:344
// index:0
// Protected virtual
// void focusOutEvent(QFocusEvent *)
func (this *QWindow) FocusOutEvent(arg0 *QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:346
// index:0
// Protected virtual
// void showEvent(QShowEvent *)
func (this *QWindow) ShowEvent(arg0 *QShowEvent /*777 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:347
// index:0
// Protected virtual
// void hideEvent(QHideEvent *)
func (this *QWindow) HideEvent(arg0 *QHideEvent /*777 QHideEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:350
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QWindow) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:351
// index:0
// Protected virtual
// void keyPressEvent(QKeyEvent *)
func (this *QWindow) KeyPressEvent(arg0 *QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:352
// index:0
// Protected virtual
// void keyReleaseEvent(QKeyEvent *)
func (this *QWindow) KeyReleaseEvent(arg0 *QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:353
// index:0
// Protected virtual
// void mousePressEvent(QMouseEvent *)
func (this *QWindow) MousePressEvent(arg0 *QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:354
// index:0
// Protected virtual
// void mouseReleaseEvent(QMouseEvent *)
func (this *QWindow) MouseReleaseEvent(arg0 *QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:355
// index:0
// Protected virtual
// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QWindow) MouseDoubleClickEvent(arg0 *QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:356
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QWindow) MouseMoveEvent(arg0 *QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:358
// index:0
// Protected virtual
// void wheelEvent(QWheelEvent *)
func (this *QWindow) WheelEvent(arg0 *QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:360
// index:0
// Protected virtual
// void touchEvent(QTouchEvent *)
func (this *QWindow) TouchEvent(arg0 *QTouchEvent /*777 QTouchEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow10touchEventEP11QTouchEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:362
// index:0
// Protected virtual
// void tabletEvent(QTabletEvent *)
func (this *QWindow) TabletEvent(arg0 *QTabletEvent /*777 QTabletEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11tabletEventEP12QTabletEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:364
// index:0
// Protected virtual
// bool nativeEvent(const QByteArray &, void *, long *)
func (this *QWindow) NativeEvent(eventType *qtcore.QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 = eventType.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11nativeEventERK10QByteArrayPvPl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, &result)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QWindow__Visibility = int

const QWindow__Hidden QWindow__Visibility = 0
const QWindow__AutomaticVisibility QWindow__Visibility = 1
const QWindow__Windowed QWindow__Visibility = 2
const QWindow__Minimized QWindow__Visibility = 3
const QWindow__Maximized QWindow__Visibility = 4
const QWindow__FullScreen QWindow__Visibility = 5

type QWindow__AncestorMode = int

const QWindow__ExcludeTransients QWindow__AncestorMode = 0
const QWindow__IncludeTransients QWindow__AncestorMode = 1

//  body block end