package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// запускаем перед основными функциями по разу чтобы файл остался в памяти в файловом кеше
// ioutil.Discard - это ioutil.Writer который никуда не пишет
func init() {
	SlowSearch(ioutil.Discard)
	FastSearch(ioutil.Discard)
}

// -----
// go test -v

func TestSearch(t *testing.T) {
	slowOut := new(bytes.Buffer)
	SlowSearch(slowOut)
	slowResult := slowOut.String()

	fastOut := new(bytes.Buffer)
	FastSearch(fastOut)
	fastResult := fastOut.String()

	if slowResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, slowResult)
	}
}

// -----
// go test -bench . -benchmem

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowSearch(ioutil.Discard)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastSearch(ioutil.Discard)
	}
}

//BenchmarkSolution-8        500       2782432 ns/op      559910 B/op      10422 allocs/op
//1 вариант (со всеми полями структуры)
//BenchmarkFast-4   	     286	   3863703 ns/op	  774644 B/op	   12453 allocs/op
//2 Вариант (с игнорированием некоторых полей)
//BenchmarkFast-4            325       3566037 ns/op      714185 B/op      8453 allocs/op
//3 Вариант (Слияние функции чтения с функцией отбора из json мапы)
//BenchmarkFast-4            362       3288452 ns/op      569578 B/op      7441 allocs/op
