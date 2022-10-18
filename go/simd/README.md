## code structure

- simd128
  - i8x16
    - for
    - simd
      - math lib c implementation
      - math lib go wrapper
    - llvm
  - i16x8
  - i32x4
  - i64x2
  - f32x4
  - f64x2
- simd256
- simd512

## api

```go
v1 := f32x4.Vector(1, 2, 3, 4)
v2 := f32x4.Vector(1, 2, 3, 4)
v := f32x4.Add(v1, v2)
```

