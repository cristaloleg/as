package as

import "reflect"

type gArray struct {
	index int
	data  reflect.Value
}

type gChan struct {
	index int
	raw   reflect.Value
	next  reflect.Value
}

type gFunc struct {
	raw reflect.Value
}

type gMap struct {
	index int
	raw   reflect.Value
	data  []reflect.Value
}

// Generator ...
func Generator(obj ...interface{}) Value {
	if len(obj) == 0 {
		return nil
	}

	if len(obj) > 1 {
		// check obj isn't a func

		// treat obj as an array of objects
		return &gArray{
			data: reflect.ValueOf(obj),
		}
	}

	v := reflect.ValueOf(obj)
	switch v.Type().Kind() {
	case reflect.Array, reflect.Slice:
		return &gArray{
			data: v,
		}

	case reflect.Chan:
		return &gChan{
			raw: v,
		}

	case reflect.Func:
		return &gFunc{}

	case reflect.Map:
		return &gMap{
			raw:  v,
			data: v.MapKeys(),
		}

	default:
		panic("unsupported type, sorry")
	}
}

func (g *gArray) Get() interface{} {
	if !g.HasValue() {
		return nil
	}
	value := g.data.Index(g.index).Interface()
	g.index++
	return value
}

func (g *gArray) HasValue() bool {
	return g.data.Len() < g.index
}

func (g *gChan) Get() interface{} {
	if !g.HasValue() {
		return nil
	}
	return g.next
}

func (g *gChan) HasValue() bool {
	value, ok := g.raw.TryRecv()
	if !ok {
		return false
	}
	g.next = value
	return true
}

func (g *gFunc) Get() interface{} {
	if !g.HasValue() {
		return nil
	}
	return nil
}

func (g *gFunc) HasValue() bool {
	return true
}

func (g *gMap) Get() interface{} {
	if !g.HasValue() {
		return nil
	}
	key := g.data[g.index]
	value := g.raw.MapIndex(key).Interface()
	g.index++
	return value
}

func (g *gMap) HasValue() bool {
	return len(g.data) < g.index
}
