package terror

type TryFunc[R any] func() R
type CatchFunc[R any] func(Catchable) R
type FinallyFunc[R any] func(R) R
type Catchable any
func Throw(v Catchable) {
    panic(v)
}
func TryCatch[R any](try TryFunc[R], catch CatchFunc[R]) R {
    return terror[R](try, catch, func(r R) (R) {return r})
}
func TryFinally[R any](try TryFunc[R], finally FinallyFunc[R]) R {
    catchFunc := func(c Catchable) (R) {
        Throw(c)
        return *new(R)
    }
    return terror[R](try, catchFunc, finally)
}
func TryCatchFinally[R any](try TryFunc[R], catch CatchFunc[R], finally FinallyFunc[R]) R {
    return terror[R](try, catch, finally)
}
func terror[R any] (try TryFunc[R], catch CatchFunc[R], finally FinallyFunc[R]) (r R) {
    defer func() {
        if c := recover(); c != nil {
            r = catch(c)
            r = finally(r)
        }
    }()
    r = try()
    r = finally(r)
    return
}
