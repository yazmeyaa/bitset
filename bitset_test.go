package bitset_test

import (
	"testing"

	"github.com/yazmeyaa/bitset"
)

func TestIsSet(t *testing.T) {
	bs := bitset.NewBitSet(105)
	bs.Set(32)
	if !bs.IsSet(32) {
		t.Error("Value in BS must be 1")
	}
}

func TestSet(t *testing.T) {
	bs := bitset.NewBitSet(105)
	bs.Set(32)
	if !bs.IsSet(32) {
		t.Error("Value in BS must be 1")
	}
}

func TestClear(t *testing.T) {
	bs := bitset.NewBitSet(105)
	bs.Set(32)
	if !bs.IsSet(32) {
		t.Error("Value in BS must be 1")
	}

	bs.Clear(32)
	if bs.IsSet(32) {
		t.Error("Value is 1 after clear")
	}
}
