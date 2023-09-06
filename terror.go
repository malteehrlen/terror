package terror


type TryFunc[R any] func() R
type CatchFunc[R any] func(Catchable) R
type FinallyFunc[R any] func(R) R

type Catchable any

func Terror[R any] (try TryFunc[R], catch CatchFunc[R], finally FinallyFunc[R]) (r R) {
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
