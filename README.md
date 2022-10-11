# simd

## api

- api
  - SSE
  - SSE2
  - SSE3
  - SSSE3
  - SSE4.1
  - SSE4.2
  - AVX
  - AVX2
  - AVX-512
- vector
  - SIMD128
  - SIMD256
  - SIMD512

## impelementation

|      | library         | auto vectorization |
|------|-----------------|--------------------|
| gcc  | api / vector    | simple loop        |
| llvm | vector          | complex loop       |

## math

- [Calling fsincos instruction in LLVM slower than calling libc sin/cos functions?](https://stackoverflow.com/questions/12485190/calling-fsincos-instruction-in-llvm-slower-than-calling-libc-sin-cos-functions)
- [Why is the gcc math library so inefficient?](https://stackoverflow.com/questions/13875540/why-is-the-gcc-math-library-so-inefficient)
- [How does C compute sin() and other math functions?](https://stackoverflow.com/questions/2284860/how-does-c-compute-sin-and-other-math-functions)

## cpp

- [std::valarray](https://en.cppreference.com/w/cpp/numeric/valarray)
- [Extensions for parallelism, version 2](https://en.cppreference.com/w/cpp/experimental/parallelism_2)
- [google/highway](https://github.com/google/highway)
- [xtensor-stack/xsimd](https://github.com/xtensor-stack/xsimd)
- [VcDevel/std-simd](https://github.com/VcDevel/std-simd)
- [VcDevel/Vc](https://github.com/VcDevel/Vc)
- [p12tic/libsimdpp](https://github.com/p12tic/libsimdpp)
- [mendlin/llvm-SIMD](https://github.com/mendlin/llvm-SIMD)

## rust

- [Portable packed SIMD vector types](https://github.com/rust-lang/rfcs/pull/2948)
- [Tracking Issue for RFC 2948: Portable SIMD](https://github.com/rust-lang/rust/issues/86656)
- [rust-lang/portable-simd](https://github.com/rust-lang/portable-simd)

## golang

- [proposal: add package for using SIMD instructions](https://github.com/golang/go/issues/53171)

### asm

- [minio](https://github.com/minio)
- [klauspost](https://github.com/klauspost)

### llvm

- [go/gollvm](https://go.googlesource.com/gollvm/)
- [llir/llvm](https://github.com/llir/llvm)
- [tinygo-org/tinygo](https://github.com/tinygo-org/tinygo)
- [tinygo-org/go-llvm](https://github.com/tinygo-org/go-llvm)