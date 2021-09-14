package main

import (
	"fmt"
	"testing"
)

func DivisionerTest(t *testing.T) {
	// name of test
	// fmt.Println(t.Name())

	res, err := Divisioner(12, 36)
	if err != nil {
		t.Fail()
		t.Errorf(err.Error())
	}
	fmt.Println(res)
}
