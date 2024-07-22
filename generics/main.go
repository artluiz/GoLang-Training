package main

import (
	"fmt"
)

type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

type MyNumber int

func SomaGenerica[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

func teste[T Ordered](n1 T, n2 T) T {
	if n1 == n2 {
		return n1
	}
	return n2
}

type stringer interface {
	String() string
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		result += val.String()
	}
	return result
}
func main() {

	fmt.Println(concat([]MyString{"a", "b", "c"}))
	//fmt.Println(SomaGenerica(map[string]Number{"a": 1.78}))
}
