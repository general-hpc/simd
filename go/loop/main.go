package main

/*
#cgo CFLAGS: -O3 -mavx -mavx2
#include "immintrin.h"
#include "xmmintrin.h"

typedef struct Expr {
	int type;

	int valueN;
	float *value;

	int exprN;
	struct Expr *expr;
} Expr;

__m128 evalConst(Expr *expr, int i);
__m128 evalAdd(Expr *expr, int i);
__m128 evalPlus(Expr *expr, int i);
__m128 eval(Expr *expr, int i);

__m128 evalConst(Expr *expr, int i) {
	return _mm_load_ps(expr->value + i);
}

__m128 evalAdd(Expr *expr, int i) {
	__m128 v1 = eval(expr->expr, i);
	__m128 v2 = eval(expr->expr + 1, i);
	return _mm_add_ps(v1, v2);
}

__m128 evalPlus(Expr *expr, int i) {
	__m128 v1 = eval(expr->expr, i);
	__m128 v2 = eval(expr->expr + 1, i);
	return _mm_mul_ps(v1, v2);
}

__m128 eval(Expr *expr, int i) {
	switch (expr->type) {
		case 0:
			return evalConst(expr, i);
		case 1:
			return evalAdd(expr, i);
		default:
			return evalPlus(expr, i);
	}
}

void eval1(Expr *expr, float *value, int N) {
	for (int i = 0; i < N; i += 4) {
		__m128 v = eval(expr, i);
		_mm_store_ps(value + i, v);
	}
}

void eval2(float *a, float *b, float *c, int N) {
	Expr expr1 = {
		.type=0,
		.valueN=N,
		.value=a,
		.exprN=0,
		.expr=NULL,
	};
	Expr expr2 = {
		.type=0,
		.valueN=N,
		.value=b,
		.exprN=0,
		.expr=NULL,
	};
	Expr exprs[2] = {expr1, expr2};
	Expr expr = {
		.type=1,
		.valueN=0,
		.value=NULL,
		.exprN=2,
		.expr=exprs,
	};
	eval1(&expr, c, N);
}

void vaddSimd128(float *a, float *b, float *c, int n) {
	for (int i = 0; i < n; i += 4) {
		__m128 v1 = _mm_load_ps(a + i);
		__m128 v2 = _mm_load_ps(b + i);
		__m128 v = _mm_add_ps(v1, v2);
		_mm_store_ps(c + i, v);
	}
}

void vaddSimd256(float *a, float *b, float *c, int n) {
	for (int i = 0; i < n; i += 8) {
		__m256 v1 = _mm256_load_ps(a + i);
		__m256 v2 = _mm256_load_ps(b + i);
		__m256 v = _mm256_add_ps(v1, v2);
		_mm256_store_ps(c + i, v);
	}
}

void vadd(float *a, float *b, float *c, int n) {
	for (int i = 0; i < n; i += 1) {
		c[i] = a[i] + b[i];
	}
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

const (
	N = 1000000000
)

type Vector C.__m128

type Numeric interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~float32 | ~float64 |
	~complex64 | ~complex128
}

func getPointer[T any](s []T) unsafe.Pointer {
	// return unsafe.Pointer(&s[0])  // len(s) > 0
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	return unsafe.Pointer(sh.Data)
}

func getSize[T Numeric]() uintptr {
	return unsafe.Sizeof(T(0))
}

func generate() ([]float32, []float32, []float32) {
	a := make([]float32, N)
	b := make([]float32, N)
	c := make([]float32, N)

	for i := range a {
		a[i] = float32(i)
	}
	for i := range b {
		b[i] = float32(i)
	}

	return a, b, c
}

func check(c []float32) {
	for i := 0; i < N; i += 1 {
		if c[i] != 2 * float32(i) {
			print(i)
			break
		}
	}
}

func loadV1(a, b, c, d float32) Vector {
	x := []float32{a, b, c, d}
	y := (*C.float)(getPointer(x))
	return Vector(C._mm_load_ps(y))
}

func loadV2(y unsafe.Pointer) Vector {
	return Vector(C._mm_load_ps((*C.float)(y)))
}

func add(v1 Vector, v2 Vector) Vector {
	return Vector(C._mm_add_ps(C.__m128(v1), C.__m128(v2)))
}

func storeV1(v Vector) (float32, float32, float32, float32) {
	x := []float32{0, 0, 0, 0}
	y := (*C.float)(getPointer(x))
	C._mm_store_ps(y, C.__m128(v))
	return x[0], x[1], x[2], x[3]
}

func storeV2(v unsafe.Pointer, a Vector) {
	C._mm_store_ps((*C.float)(v), C.__m128(a))
}

func testA(a, b, c []float32) {
	startTime := time.Now()

	for i := 0; i < N; i++ {
		c[i] = a[i] + b[i]
	}

	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)

	check(c)
}

func testB(a, b, c []float32) {
	startTime := time.Now()

	for i := 0; i < N; i += 4 {
		v1 := loadV1(a[i], a[i + 1], a[i + 2], a[i + 3])
		v2 := loadV1(b[i], b[i + 1], b[i + 2], b[i + 3])
		v := add(v1, v2)
		c[i], c[i + 1], c[i + 2], c[i + 3] = storeV1(v)
	}

	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)

	check(c)
}

func testC(a, b, c []float32) {
	startTime := time.Now()

	aPtr := getPointer(a)
	bPtr := getPointer(b)
	cPtr := getPointer(c)
	size := getSize[float32]()

	for i := 0; i < N; i += 4 {
		offset := uintptr(i) * size
		v1 := loadV2(unsafe.Add(aPtr, offset))
		v2 := loadV2(unsafe.Add(bPtr, offset))
		v := add(v1, v2)
		storeV2(unsafe.Add(cPtr, offset), v)
	}

	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)

	check(c)
}

func testD(a, b, c []float32) {
	startTime := time.Now()

	aPtr := (*C.float)(getPointer(a))
	bPtr := (*C.float)(getPointer(b))
	cPtr := (*C.float)(getPointer(c))
	C.vaddSimd128(aPtr, bPtr, cPtr, N)

	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)

	check(c)
}

func testE(a, b, c []float32) {
	startTime := time.Now()

	aPtr := (*C.float)(getPointer(a))
	bPtr := (*C.float)(getPointer(b))
	cPtr := (*C.float)(getPointer(c))
	C.vaddSimd256(aPtr, bPtr, cPtr, N)

	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)

	check(c)
}

func testF(a, b, c []float32) {
	startTime := time.Now()

	aPtr := (*C.float)(getPointer(a))
	bPtr := (*C.float)(getPointer(b))
	cPtr := (*C.float)(getPointer(c))
	C.vadd(aPtr, bPtr, cPtr, N)

	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)

	check(c)
}

func testG(a, b, c []float32) {
	startTime := time.Now()

	aPtr := (*C.float)(getPointer(a))
	bPtr := (*C.float)(getPointer(b))
	cPtr := (*C.float)(getPointer(c))
	C.eval2(aPtr, bPtr, cPtr, N)

	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)

	check(c)
}

func main() {
	a, b, c := generate()
	testA(a, b, c)  // 2.13s - 2.68s
	testD(a, b, c)  // 1.87s - 2.12s
	testE(a, b, c)  // 1.86s - 1.91s
	testF(a, b, c)  // 1.87s - 1.90s
	                // cgo not support pass by ast
	testG(a, b, c)  // recurrence based, 1.90s - 1.91s
	                // stack based
}
