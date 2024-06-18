package ident

import (
	"bytes"
	"reflect"

	"github.com/samborkent/uuidv7"
)

// ID constrains to types of [16]byte
type ID interface {
	~[16]byte
}

// NewID creates a uuidv7 based identifier of the specified type
func NewID[T ID]() T {
	var id T
	uid := uuidv7.New()
	copy(id[:], uid[:])
	return id
}

// CopyID returns a new instance of the identifer
func CopyID[T ID](src T) T {
	var id T
	copy(id[:], src[:])
	return id
}

// EqualValueID returns true if the identifers have the same value
func EqualValueID[T ID](dst, src T) bool {
	return bytes.Equal(dst[:], src[:])
}

// DeepEqualID returns true if the identifiers are of the same type and value
func DeepEqualID[T, U ID](dst T, src U) bool {
	tSrc := reflect.TypeOf(src)
	tDst := reflect.TypeOf(dst)
	if tSrc.String() != tDst.String() {
		return false
	}
	var s = CreateFrom[U, T](src)
	return EqualValueID(dst, s)
}

// CreateFrom convers an identifier to a different type, with the same value
func CreateFrom[T, U ID](src T) U {
	var dst U
	for i := 0; i < len(dst); i++ {
		dst[i] = src[i]
	}
	return dst
}

// EmptyID returns true if the identifier is blank
func EmptyID[T ID](src T) bool {
	var id T
	return EqualValueID(id, src)
}
