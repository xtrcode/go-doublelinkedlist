/*
This is free and unencumbered software released into the public domain.

Anyone is free to copy, modify, publish, use, compile, sell, or
distribute this software, either in source code form or as a compiled
binary, for any purpose, commercial or non-commercial, and by any
means.

In jurisdictions that recognize copyright laws, the author or authors
of this software dedicate any and all copyright interest in the
software to the public domain. We make this dedication for the benefit
of the public at large and to the detriment of our heirs and
successors. We intend this dedication to be an overt act of
relinquishment in perpetuity of all present and future rights to this
software under copyright law.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

For more information, please refer to <http://unlicense.org/>
*/

package doublelinkedlist

import (
	"fmt"
	"math"
	"testing"
)

var min = 15.
var max = 20.

func BenchmarkDoubleLinkedList_Append(b *testing.B) {
	for k := min; k <= max; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			list := NewDoubleLinkedList()

			for v := 0; v < b.N; v++ {
				list.Append(v)
			}
		})
	}
}

func BenchmarkDoubleLinkedList_Prepend(b *testing.B) {
	for k := min; k <= max; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			list := NewDoubleLinkedList()

			for v := 0; v < b.N; v++ {
				list.Prepend(v)
			}
		})
	}
}

func BenchmarkDoubleLinkedList_Delete(b *testing.B) {
	for k := min; k <= max; k++ {
		n := int(math.Pow(2, k))
		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			list := NewDoubleLinkedList()

			num := b.N

			for v := 0; v < num; v++ {
				list.Append(v)
			}

			// delete mid
			for v := num / 3; v < (num/3)*2; v++ {
				list.Delete(v)
			}

			// delete end
			for v := num; v > num-(num/3); v-- {
				list.Delete(v)
			}

			// delete start
			for v := 0; v < num/3; v++ {
				list.Delete(v)
			}
		})
	}
}
