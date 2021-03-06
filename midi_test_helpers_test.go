// Copyright 2012 Joe Wass. All rights reserved.
// Use of this source code is governed by the MIT license
// which can be found in the LICENSE file.

// MIDI package
// A package for reading Standard Midi Files, written in Go.
// Joe Wass 2012
// joe@afandian.com

/*
 * Test helper functions.
 * Bits and pieces that don't appear to be in the test framework.
 */

package midi

import (
	"testing"
)

// Helper assertions
func assertHasFlag(value int, flag int, test *testing.T) {
	if value&flag == 0 {
		test.Fatal("Expected to find ", flag, " in ", value)
	}
}

// assertBytesEqual asserts that the given byte arrays or slices are equal.
func assertBytesEqual(a []byte, b []byte, t *testing.T) {
	if len(a) != len(b) {
		t.Fatal("Two arrays not equal", a, b)
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Fatal("Two arrays not equal. At ", i, " was ", a[i], " vs ", b[i])
		}
	}
}

// Assert uint16s equal
func assertUint16Equal(a uint16, b uint16, test *testing.T) {
	if a != b {
		test.Fatal(a, " != ", b)
	}
}

// Assert uint16s equal
func assertInt16sEqual(a int16, b int16, test *testing.T) {
	if a != b {
		test.Fatal(a, " != ", b)
	}
}

// Assert uint32s equal
func assertUint32Equal(a uint32, b uint32, test *testing.T) {
	if a != b {
		test.Fatal(a, " != ", b)
	}
}

// Assert uint16s equal
func assertIntsEqual(a int, b int, test *testing.T) {
	if a != b {
		test.Fatal(a, " != ", b)
	}
}

// Assert uint8s equal
func assertUint8sEqual(a uint8, b uint8, test *testing.T) {
	if a != b {
		test.Fatal(a, " != ", b)
	}
}

func assertFalse(a bool, test *testing.T) {
	if a == true {
		test.Fatal("Should have been false")
	}
}

func assertTrue(a bool, test *testing.T) {
	if a != true {
		test.Fatal("Should have been true")
	}
}

func assertNoError(e error, test *testing.T) {
	if e != nil {
		test.Fatal("Should have been no error, was ", e)
	}
}

func assertError(e error, expected error, test *testing.T) {
	if e != expected {
		test.Fatal("Should have been ", expected, " was ", e)
	}
}

// Assert two strings are equal
func assertStringsEqual(a string, b string, test *testing.T) {
	if a != b {
		test.Fatal(a, " != ", b)
	}
}
