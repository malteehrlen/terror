package terror

import "errors"

type TryFunc[R any] func() R
type CatchFunc[R any] func(Catchable) R
type FinallyFunc[R any] func(R) R
type Catchable any
type TryBlock[R any] struct {
	Try     TryFunc[R]
	Catch   CatchFunc[R]
	Finally FinallyFunc[R]
}

func (b TryBlock[R]) Run() (r R, err error) {
	defer func() {
		if c := recover(); c != nil {
			cString, ok := c.(string)
			if ok {
				err = errors.New("Uncaught exception: " + cString)
			} else {
				err = errors.New("Uncaught, unparseable exception (wow!)")
			}
		}
	}()
	if b.Catch != nil {
		if b.Finally != nil {
			r = terror(b.Try, b.Catch, b.Finally)
		} else {
		r = tryCatch(b.Try, b.Catch)
        }
	} else if b.Finally != nil {
		r = tryFinally(b.Try, b.Finally)
	} else {
		r = b.Try()
	}
	return
}

func Throw(v Catchable) {
	panic(v)
}
func tryCatch[R any](try TryFunc[R], catch CatchFunc[R]) R {
	return terror[R](try, catch, func(r R) R { return r })
}
func tryFinally[R any](try TryFunc[R], finally FinallyFunc[R]) R {
	catchFunc := func(c Catchable) R {
		Throw(c)
		return *new(R)
	}
	return terror[R](try, catchFunc, finally)
}
func terror[R any](try TryFunc[R], catch CatchFunc[R], finally FinallyFunc[R]) (r R) {
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
