package creational

import "testing"

func TestGetInstance(t *testing.T) {
	singleton1 := GetInstance()

	if singleton1 == nil {
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	expectedSingleton := singleton1

	currentCount := singleton1.AddOne()
	if currentCount != 1 {
		t.Error("the count must be 1 but it is %i\n", currentCount)
	}

	singleton2 := GetInstance()
	if singleton2 != expectedSingleton {
		t.Error("expected same instance in singleton 2")
	}

	currentCount = singleton1.AddOne()
	if currentCount != 2 {
		t.Error("currentCunt must be 2 but was %i\n", currentCount)
	}

}
