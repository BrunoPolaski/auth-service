package tests

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("\nExpected: %v \nGot: %v", expected, actual)
	}
}

func AssertNotEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected == actual {
		t.Errorf("\nExpected: %v \nGot: %v", expected, actual)
	}
}

func AssertNil(t *testing.T, actual interface{}) {
	isInterfaceNil := func(i interface{}) bool {
		return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
	}

	if actual != nil && !isInterfaceNil(actual) {
		t.Errorf("\nExpected: nil \nGot: %v", actual)
	}
}

func AssertNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		t.Errorf("\nExpected: not nil \nGot: nil")
	}
}
