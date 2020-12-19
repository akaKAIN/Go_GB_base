package benchtests

import (
	"fmt"
	"reflect"
	"testing"
)

// Сравнение производительности некоторых функций их пакета reflect
func BenchmarkTypeOf(b *testing.B) {
	name := "Ivan"
	for i := 0; i < b.N; i++ {

		_ = fmt.Sprintf("%v", reflect.TypeOf(name))
	}
}

func BenchmarkSprintf(b *testing.B) {
	name := "Ivan"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%v", name)
	}
}
