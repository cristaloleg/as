package as_test

import (
	"testing"

	"github.com/cristaloleg/as"
)

func Test(t *testing.T) {
	fn := func(a int) int { return a * 10 }
	value := as.Lazy(fn, 20)
	if value == nil {
		t.Error("cannot instantiate")
	}

	if value.HasValue() {
		t.Error("value wasn't requested yet")
	}

	if got := value.Get().(int); got != 200 {
		t.Error("expected 200")
	}

	if !value.HasValue() {
		t.Error("value wasn already requested")
	}
}
