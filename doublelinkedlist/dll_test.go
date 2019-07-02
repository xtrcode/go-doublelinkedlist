
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
			for v := num-1; v > num-(num/3); v-- {
				list.Delete(v)
			}

			// delete start
			for v := 0; v < num/3; v++ {
				list.Delete(v)
			}
		})
	}
}
