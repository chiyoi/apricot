package test

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, exp, rev any) {
	if !reflect.DeepEqual(exp, rev) {
		t.Errorf("Error (expect: %v, got: %v).", exp, rev)
	}
}
