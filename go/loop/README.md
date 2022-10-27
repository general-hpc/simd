## overview

- for / op [x]
- func / no func [x]
- func based [x] / object based

|      | lazy                            | eager                        |
|------|---------------------------------|------------------------------|
| gcc  | library (slow / can't pass ast) | library / auto vectorization |
| llvm | library / auto vectorization    | library / auto vectorization |
