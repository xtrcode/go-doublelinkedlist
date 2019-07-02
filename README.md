# Straightforward, in-place double linked list implementation in Go
## Usage
```go
list := NewDoubleLinkedList()
for v := 0; v < 100; v++ {
    list.Append(v)
}

list.Delete(5)
list.Prepend(8)

node := list.Find(6)
if node == nil {
    panic("Node not found!")
}

node.SetNext(nil)
node.SetPrevious(nil)
node.SetValue(5)
fmt.Println(node.Compare(5))
fmt.Println(node.GetValue())

```

## Benchmark
Hardware: Dell XPS 13 (9360) - i7 8550u/16GB
OS: Fedora 29
<pre>
  goos: linux
  goarch: amd64
  pkg: go-doublelinkedlist/doublelinkedlist
  BenchmarkDoubleLinkedList_Append/32768-8                10000000               102 ns/op
  BenchmarkDoubleLinkedList_Append/65536-8                20000000               115 ns/op
  BenchmarkDoubleLinkedList_Append/131072-8               10000000               115 ns/op
  BenchmarkDoubleLinkedList_Append/262144-8               20000000               116 ns/op
  BenchmarkDoubleLinkedList_Append/524288-8               20000000               115 ns/op
  BenchmarkDoubleLinkedList_Append/1048576-8              20000000               119 ns/op
  BenchmarkDoubleLinkedList_Prepend/32768-8               10000000               121 ns/op
  BenchmarkDoubleLinkedList_Prepend/65536-8               20000000               122 ns/op
  BenchmarkDoubleLinkedList_Prepend/131072-8              10000000               122 ns/op
  BenchmarkDoubleLinkedList_Prepend/262144-8              20000000               119 ns/op
  BenchmarkDoubleLinkedList_Prepend/524288-8              20000000               121 ns/op
  BenchmarkDoubleLinkedList_Prepend/1048576-8             20000000               113 ns/op
  BenchmarkDoubleLinkedList_Delete/32768-8                  100000            142778 ns/op
  BenchmarkDoubleLinkedList_Delete/65536-8                  100000            143565 ns/op
  BenchmarkDoubleLinkedList_Delete/131072-8                 100000            165863 ns/op
  BenchmarkDoubleLinkedList_Delete/262144-8                 100000            161544 ns/op
  BenchmarkDoubleLinkedList_Delete/524288-8                 100000            155565 ns/op
  BenchmarkDoubleLinkedList_Delete/1048576-8                100000            140267 ns/op
  PASS
  ok      go-doublelinkedlist/doublelinkedlist    120.970s
</pre>

# LICENSE
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