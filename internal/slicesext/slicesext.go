package slicesext

import (
	"cmp"
	"slices"
)

func UniqueJoin[T cmp.Ordered](ss ...[]T) []T {
	var length int
	for _, s := range ss {
		length += len(s)
	}
	r := make([]T, length)
	var i int
	for _, s := range ss {
		i += copy(r[i:], s)
	}
	slices.Sort(r)
	return slices.Compact(r)
}

func FirstNonZero[T comparable](values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return zero
}
