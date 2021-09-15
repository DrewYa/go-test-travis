package main

import (
	"fmt"
	"time"
)

func Divisioner(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("ZeroDivisionError")
	}
	return float64(a) / float64(b), nil
}

type MyStruct struct {
	F1, F2, F3 int
}

func NewStruct(f1, f2, f3 int) *MyStruct {
	ms := &MyStruct{f1, f2, f3}
	ms.Init()
	return ms
}

// имитируем долгую инициализацию функции
func (ms *MyStruct) Init() {
	time.Sleep(time.Second * 2)
}

// имитируем некое корректное завершение работы со структурой,
// и это отнимает опр. время
func (ms *MyStruct) Cleanup() {
	time.Sleep(time.Second * 2)
}

func (ms *MyStruct) CompareFields() error {
	switch {
	case ms.F1 < ms.F2:
		return fmt.Errorf("f2 larger f1")
	case ms.F2 < ms.F3:
		return fmt.Errorf("f3 larger f2")
	}
	return nil
}

func main() {
	res, _ := Divisioner(12, 5)
	fmt.Println("result:", res)
}
