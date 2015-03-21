/*
The MIT License (MIT)

Copyright (c) 2015 James Garfield

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package sort // import "goast.net/x/sort"

import (
	"sort"
)

//Impl Types: Requires a slice of any type (interface{})
type I interface{}
type Slice []I

//Related Type: Enables sorting as a slice operation
type _Sorter struct {
	Slice
	LessFunc func(I, I) bool
}

func (s _Sorter) Less(i, j int) bool {
	return s.LessFunc(s.Slice[i], s.Slice[j])
}

func (s Slice) Len() int {
	return len(s)
}

func (s Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//Sorts the slice in-place using the related type sorter
func (s Slice) Sort(less func(I, I) bool) {
	sort.Sort(_Sorter{s, less})
}
