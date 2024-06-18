package ident

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewID(t *testing.T) {

	type myID [16]byte

	id1 := NewID[myID]()
	id2 := NewID[myID]()

	if EqualValueID(id1, id2) {
		t.Fatal("should be different")
	}
}

func TestEmptyID(t *testing.T) {

	type myID [16]byte

	var id myID

	if !EmptyID(id) {
		t.Fatal("should be uninitialised")
	}
}

func TestCopyID(t *testing.T) {

	type myID [16]byte

	var id1 = NewID[myID]()

	var id2 = CopyID(id1)

	if !EqualValueID(id1, id2) {
		t.Fatal("should be identical")
	}
}

func TestDeepEqualID(t *testing.T) {

	type myID1 [16]byte
	type myID2 [16]byte

	var id1 = NewID[myID1]()

	var id2 = CreateFrom[myID1, myID2](id1)

	if DeepEqualID(id1, id2) {
		t.Fatal("unexpected equality")
	}
}

func TestDeepEqualID_1(t *testing.T) {

	type myID1 [16]byte

	var id1 = NewID[myID1]()
	var id2 = CreateFrom[myID1, myID1](id1)

	if !DeepEqualID(id1, id2) {
		t.Fatal("unexpected equality")
	}
}

func ExampleNewID() {

	type myID [16]byte

	id := NewID[myID]()

	fmt.Println(reflect.TypeOf(id))
	// Output: id.myID
}
