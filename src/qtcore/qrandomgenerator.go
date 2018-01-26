package qtcore

// /usr/include/qt/QtCore/qrandom.h
// #include <qrandom.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QRandomGenerator struct {
	*qtrt.CObject
}

func (this *QRandomGenerator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRandomGenerator) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQRandomGeneratorFromPointer(cthis unsafe.Pointer) *QRandomGenerator {
	return &QRandomGenerator{&qtrt.CObject{cthis}}
}
func (*QRandomGenerator) NewFromPointer(cthis unsafe.Pointer) *QRandomGenerator {
	return NewQRandomGeneratorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qrandom.h:55
// index:0
// Public inline
// void QRandomGenerator(quint32)
func NewQRandomGenerator(seedValue uint) *QRandomGenerator {
	cthis := qtrt.Calloc(1, 256) // 2512
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGeneratorC2Ej", ffiqt.FFI_TYPE_VOID, cthis, seedValue)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:61
// index:1
// Public inline
// void QRandomGenerator(const quint32 *, qsizetype)
func NewQRandomGenerator_1(seedBuffer unsafe.Pointer /*666*/, len int64) *QRandomGenerator {
	cthis := qtrt.Calloc(1, 256) // 2512
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGeneratorC2EPKjx", ffiqt.FFI_TYPE_VOID, cthis, &seedBuffer, len)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:65
// index:2
// Public
// void QRandomGenerator(const quint32 *, const quint32 *)
func NewQRandomGenerator_2(begin unsafe.Pointer /*666*/, end unsafe.Pointer /*666*/) *QRandomGenerator {
	cthis := qtrt.Calloc(1, 256) // 2512
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGeneratorC2EPKjS1_", ffiqt.FFI_TYPE_VOID, cthis, &begin, &end)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:172
// index:3
// Protected
// void QRandomGenerator(enum QRandomGenerator::System)
func NewQRandomGenerator_3(arg0 int) *QRandomGenerator {
	cthis := qtrt.Calloc(1, 256) // 2512
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGeneratorC2ENS_6SystemE", ffiqt.FFI_TYPE_VOID, cthis, arg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQRandomGeneratorFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qrandom.h:77
// index:0
// Public inline
// quint32 generate()
func (this *QRandomGenerator) Generate() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator8generateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:152
// index:1
// Public inline
// void generate(quint32 *, quint32 *)
func (this *QRandomGenerator) Generate_1(begin unsafe.Pointer /*666*/, end unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator8generateEPjS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &begin, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:84
// index:0
// Public inline
// quint64 generate64()
func (this *QRandomGenerator) Generate64() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator10generate64Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:91
// index:0
// Public inline
// double generateDouble()
func (this *QRandomGenerator) GenerateDouble() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator14generateDoubleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtCore/qrandom.h:105
// index:0
// Public inline
// double bounded(double)
func (this *QRandomGenerator) Bounded(highest float64) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), highest)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtCore/qrandom.h:110
// index:1
// Public inline
// quint32 bounded(quint32)
func (this *QRandomGenerator) Bounded_1(highest uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), highest)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:118
// index:2
// Public inline
// int bounded(int)
func (this *QRandomGenerator) Bounded_2(highest int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), highest)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrandom.h:123
// index:3
// Public inline
// quint32 bounded(quint32, quint32)
func (this *QRandomGenerator) Bounded_3(lowest uint, highest uint) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEjj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), lowest, highest)
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qrandom.h:128
// index:4
// Public inline
// int bounded(int, int)
func (this *QRandomGenerator) Bounded_4(lowest int, highest int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7boundedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), lowest, highest)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qrandom.h:160
// index:0
// Public inline
// void seed(quint32)
func (this *QRandomGenerator) Seed(s uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator4seedEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:162
// index:0
// Public
// void discard(unsigned long long)
func (this *QRandomGenerator) Discard(z uint64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator7discardEy", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrandom.h:163
// index:0
// Public static inline
// QRandomGenerator::result_type min()
func (this *QRandomGenerator) Min() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator3minEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QRandomGenerator_Min() uint {
	var nilthis *QRandomGenerator
	rv := nilthis.Min()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:164
// index:0
// Public static inline
// QRandomGenerator::result_type max()
func (this *QRandomGenerator) Max() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator3maxEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QRandomGenerator_Max() uint {
	var nilthis *QRandomGenerator
	rv := nilthis.Max()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:166
// index:0
// Public static inline
// QRandomGenerator * system()
func (this *QRandomGenerator) System() *QRandomGenerator /*777 QRandomGenerator **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator6systemEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QRandomGenerator_System() *QRandomGenerator /*777 QRandomGenerator **/ {
	var nilthis *QRandomGenerator
	rv := nilthis.System()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:167
// index:0
// Public static inline
// QRandomGenerator * global()
func (this *QRandomGenerator) Global() *QRandomGenerator /*777 QRandomGenerator **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator6globalEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QRandomGenerator_Global() *QRandomGenerator /*777 QRandomGenerator **/ {
	var nilthis *QRandomGenerator
	rv := nilthis.Global()
	return rv
}

// /usr/include/qt/QtCore/qrandom.h:168
// index:0
// Public static inline
// QRandomGenerator securelySeeded()
func (this *QRandomGenerator) SecurelySeeded() *QRandomGenerator /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QRandomGenerator14securelySeededEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRandomGeneratorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QRandomGenerator_SecurelySeeded() *QRandomGenerator /*123*/ {
	var nilthis *QRandomGenerator
	rv := nilthis.SecurelySeeded()
	return rv
}

type QRandomGenerator__System = int

//  body block end