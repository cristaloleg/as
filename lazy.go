package as

import (
	"reflect"
	"sync"
)

type lazy struct {
	fn       reflect.Value
	params   []reflect.Value
	once     sync.Once
	hasValue bool
	result   interface{}
}

// Lazy ...
func Lazy(op interface{}, params ...interface{}) Value {
	t := reflect.TypeOf(op)
	if t.Kind() != reflect.Func {
		return nil
	}

	fn := reflect.ValueOf(op)
	in := make([]reflect.Value, len(params))

	for i := 0; i < len(params); i++ {
		if fn.Type().In(i) != reflect.TypeOf(params[i]) {
			return nil
		}
		in[i] = reflect.ValueOf(params[i])
	}

	return &lazy{
		fn:     fn,
		params: in,
	}
}

func (l *lazy) Get() interface{} {
	l.once.Do(func() {
		l.result = l.fn.Call(l.params)[0].Interface()
		l.hasValue = true
	})
	return l.result
}

func (l *lazy) HasValue() bool {
	return l.hasValue
}
