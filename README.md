# Try-error (Terror for short)
Implements try/catch/finally so you can idk, mess up in new ways. Production ready! Try it today:
```
go get github.com/malteehrlen/terror@v0.1.1
```

## Example
Populate a TryBlock with code blocks and then run it. Guaranteed results!

```golang
tb := &terror.TryBlock[string]{}
tb.Try = func() string {
    terror.Throw("oh jeez")
    return "you will never get this"
}
tb.Catch = func(c terror.Catchable) string {
    cStr, ok := c.(string)
    if ok {
        return cStr
    }
    return "uncaught error"
}
tb.Finally = func(r string) string {
    println(r)
    return "finally"
}

a, err := tb.Run() // prints "oh jeez"
println(a) // prints "finally"
```
