# handy <!-- omit in toc -->
Small but useful Go generic functions.

# Table of Contents <!-- omit in toc -->
- [Showcase](#showcase)
  - [Pointer initializer](#pointer-initializer)
  - [Slice initializer](#slice-initializer)
  - [Ternary operator](#ternary-operator)

# Showcase
## Pointer initializer

```go
p := handy.New(42)
fmt.Println(*p) // 42
```

## Slice initializer

```go
ss := handy.Slice("foo", "bar")
fmt.Println(ss) // [foo bar]
```

## Ternary operator

```go
c := handy.Tern(true, "foo", "bar")
fmt.Println(c) // foo
```
