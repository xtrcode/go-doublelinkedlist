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

type IDoubleLinkedListNode interface {
	SetValue(value interface{})
	GetValue() interface{}
	GetPrevious() IDoubleLinkedListNode
	GetNext() IDoubleLinkedListNode
	SetPrevious(prev IDoubleLinkedListNode)
	SetNext(next IDoubleLinkedListNode)
	Delete()
	Append(node IDoubleLinkedListNode)
	Compare(value interface{}) bool
}

type DoubleLinkedList struct {
	root *DoubleLinkedListNode
	len  int
}

type DoubleLinkedListNode struct {
	value interface{}
	prev  IDoubleLinkedListNode
	next  IDoubleLinkedListNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	var root = NewDoubleLinkedListNode(nil)
	root.SetNext(root)
	root.SetPrevious(root)

	return &DoubleLinkedList{root, 0}
}

func (dll *DoubleLinkedList) Append(value interface{}) {
	dll.root.GetPrevious().Append(NewDoubleLinkedListNode(value))
	dll.len++
}

func (dll *DoubleLinkedList) Prepend(value interface{}) {
	dll.root.Append(NewDoubleLinkedListNode(value))
	dll.len++
}

func (dll *DoubleLinkedList) Delete(value interface{}) {
	n := dll.Find(value)
	if n != nil {
		n.Delete()
		dll.len--
	}
}

func (dll *DoubleLinkedList) Len() int {
	return dll.len
}

func (dll *DoubleLinkedList) Find(value interface{}) *DoubleLinkedListNode {
	var n *DoubleLinkedListNode

	dll.root.SetValue(value)

	for n = dll.root.GetNext().(*DoubleLinkedListNode);
		!n.Compare(value); n = n.GetNext().(*DoubleLinkedListNode) {
	}

	dll.root.SetValue(nil)
	if n == dll.root {
		return nil
	}

	return n
}

func NewDoubleLinkedListNode(value interface{}) *DoubleLinkedListNode {
	return &DoubleLinkedListNode{value, nil, nil}
}

func (dlln *DoubleLinkedListNode) Compare(value interface{}) bool {
	return dlln.value == value
}

func (dlln *DoubleLinkedListNode) SetValue(value interface{}) {
	dlln.value = value
}

func (dlln *DoubleLinkedListNode) Delete() {
	dlln.GetPrevious().SetNext(dlln.GetNext())
	dlln.GetNext().SetPrevious(dlln.GetPrevious())
}

func (dlln *DoubleLinkedListNode) Append(newnode IDoubleLinkedListNode) {
	newnode.SetNext(dlln.GetNext())
	newnode.SetPrevious(dlln)
	dlln.GetNext().SetPrevious(newnode)
	dlln.SetNext(newnode)
}

func (dlln *DoubleLinkedListNode) GetValue() interface{} {
	return dlln.value
}

func (dlln *DoubleLinkedListNode) GetPrevious() IDoubleLinkedListNode {
	return dlln.prev
}

func (dlln *DoubleLinkedListNode) GetNext() IDoubleLinkedListNode {
	return dlln.next
}

func (dlln *DoubleLinkedListNode) SetPrevious(prev IDoubleLinkedListNode) {
	dlln.prev = prev
}

func (dlln *DoubleLinkedListNode) SetNext(next IDoubleLinkedListNode) {
	dlln.next = next
}
