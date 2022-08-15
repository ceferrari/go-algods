package sorts

import (
	"golang.org/x/exp/constraints"
)

func findPivot[T constraints.Ordered](a []T, lo, hi int) int {
	mi := (lo + hi) / 2
	if a[lo] > a[mi] {
		a[lo], a[mi] = a[mi], a[lo]
	}
	if a[lo] > a[hi] {
		a[lo], a[hi] = a[hi], a[lo]
	}
	if a[mi] > a[hi] {
		a[mi], a[hi] = a[hi], a[mi]
	}
	a[mi], a[hi-1] = a[hi-1], a[mi]
	return hi - 1
}

func partition[T constraints.Ordered](a []T, lo, hi int) int {
	l, p := lo, findPivot(a, lo, hi)
	for r := lo; r < hi; r++ {
		if a[r] < a[p] {
			a[l], a[r] = a[r], a[l]
			l++
		}
	}
	a[l], a[p] = a[p], a[l]
	return l
}

func quickSort[T constraints.Ordered](a []T, lo, hi int) {
	if lo < hi {
		l := partition(a, lo, hi)
		quickSort(a, lo, l-1)
		quickSort(a, l+1, hi)
	}
}

func QuickSort[T constraints.Ordered](a []T) {
	if len(a) > 1 {
		quickSort(a, 0, len(a)-1)
	}
}
