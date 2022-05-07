# handy <!-- omit in toc -->
Small but useful Go generic functions.

# Table of Contents <!-- omit in toc -->
- [Showcase](#showcase)
  - [Pointer initializer](#pointer-initializer)
  - [Slice initializer](#slice-initializer)
  - [Ternary operator](#ternary-operator)
  - [Unpack error](#unpack-error)
  - [Single-line errors.As](#single-line-errorsas)
    - [error type](#error-type)

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

## Unpack error

```go
w := iomock.ErrOnCallWriter(1, iomock.ErrWrite)
err := handy.UnpackError(io.WriteString(w, "foobar"))
fmt.Println(err) // write error
```

## Single-line errors.As

### error type

```go
type TestError string

func (s TestError) Error() string {
	return string(s)
}
```

```go
in := fmt.Errorf("wrap: %w", TestError("error"))
fmt.Println(handy.As[TestError](in)) // error true
```
