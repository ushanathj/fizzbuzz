// run package tests | run file tests
package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/ushanathj/fizzbuzz"
)

func TestFizz(t *testing.T) {
	t.Log("starting the fizz test")
	_, ok := fizzbuzz.Fizzbuzz(1)
	if ok {
		t.Logf("The value %v should not have returned true", 1)
		t.FailNow()
	}

	result, ok := fizzbuzz.Fizzbuzz(3)
	if !ok {
		t.Logf("The value %v should not have returned true", 3)
		t.Fail()
	}

	if result != "fizz" {
		t.Log("reslut should have been fizz")
		t.Fail()
	}

}

func ExampleFizzbuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	// Output: fizz

	result, _ = fizzbuzz.Fizzbuzz(5)
	fmt.Println(result)
	// Output: buzz

}
