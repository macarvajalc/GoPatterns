package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter := GetInstance()

	if counter == nil  {
		t.Error("Expected pointer to Singleton after calling GetInstance(), not nil")
	}
}

func TestAddOne(t *testing.T){
	counter := GetInstance()
	expectedCounter := counter
	currentCounter := counter.AddOne()
	if currentCounter != 1{
		t.Error("After Calling for the first time to count, the count must be 1 but it is \n",currentCounter)
	}
	counter1 := GetInstance()
	if counter1 != expectedCounter{
		t.Error("Expected same instance in counter 2 but it got a different instance")
	}
	currentCounter = counter1.AddOne()
	if currentCounter != 2{
		t.Errorf("After calling AddOne using the second counter, the current count must be 2 but was %d\n",currentCounter)
	}
}