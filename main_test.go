package main

import (
	"fmt"
	"testing"
	"time"
)

// запуск отдельного теста
// $ go test -run ^TestOne$ p006
func TestOne(t *testing.T) {
	a := 1
	b := 1
	if a != b {
		t.Fail()
		t.Errorf("%d != %d", a, b)
	}
}

func TestDivisioner(t *testing.T) {
	// name of test
	// fmt.Println(t.Name())

	res, err := Divisioner(12, 36)
	if err != nil {
		t.Fail()
		t.Errorf(err.Error())
	}
	fmt.Println("result:", res)

}

func BenchmarkMyStruct(b *testing.B) {
	ms := NewStruct(0, 0, 0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := ms.CompareFields()
		if err != nil {
			b.Fail()
			b.Error(err)
		}
	}

	b.Cleanup(ms.Cleanup)
}

// функции, которые начинаются с "Example" - это тесты,
// которые результат выводят в Stdout.
// Последний коммент (если он есть) - это ожидаемый результат,
// как видно, он оформляется специальным образом.
// Если функция выведет другой результат, то такой "тест-пример" зафейлится
func ExampleDivisioner() {
	time.Sleep(time.Millisecond * 500)
	res, _ := Divisioner(12, 4)
	fmt.Println(res)
	// Output: 3
}

// Следующим образом можно сделать субтесты.
// Видимо тесткейсы реализуются именно так.
// Суббенчмарки также возможны.
func TestCase1(t *testing.T) {
	// func setup for test case here

	// это тест1 тесткейса TestCase1
	t.Run("A=1,B=1", func(t *testing.T) {
		a := 1
		b := 1
		res := a / b
		fmt.Println("result:", res)
	})
	// это тест2 тесткейса TestCase1
	t.Run("A=2,B=1", func(t *testing.T) {
		a := 2
		b := 1
		_ = a / b
		// fmt.Println("result:", res)
	})
	// это тест3 тесткейса TestCase1
	t.Run("A=1,B=3", func(t *testing.T) {
		a := 1
		b := 3
		_ = a / b
		// fmt.Println("result:", res)
	})

	// panic! Паника в этом случае не будет обработана.
	// Выполнение тестов прекратится
	// и все последующие тесты не будут выполнены
	// t.Run("A=100,B=0", func(t *testing.T) {
	// 	a := 100
	// 	b := 0
	// 	res := a / b
	// 	fmt.Println(res)
	// })

	// здесь сделано что-то в роде субтесткейса тесткейса TestCase1
	t.Run("A=12,B=4,C=x", func(t *testing.T) {
		t.Run("C=21", func(t *testing.T) {
			a := 12
			b := 4
			c := 21
			res := a / b / c
			fmt.Println("result:", res)
		})
		// пропускаем выполнение теста
		t.Run("C=3", func(t *testing.T) {
			t.Skip()
			if !t.Skipped() {
				a := 12
				b := 4
				c := 3
				res := a / b / c
				fmt.Println("result:", res)
			}
		})
		t.Run("C=33", func(t *testing.T) {
			a := 12
			b := 4
			c := 3
			res := a / b / c
			fmt.Println("result:", res)
		})
	})

	// teardown func for test case here
	t.Cleanup(func() {
		fmt.Println("End of test case")
	})
}
