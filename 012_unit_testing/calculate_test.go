package main

import "testing"

func TestCalculate(test *testing.T) {
	if calculate(3,3) != 5 {
		test.Error("Is not 5")
		test.Fail()
		test.FailNow()
	}
	test.Log("Success")
}
