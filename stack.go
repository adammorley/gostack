/*
    a stack implementation for integers in golang
*/

package stack

import "errors"

type Stack struct {
    top *element
}

type element struct {
    value int
    under *element
}

// create a new stack
func New() *Stack {
    var s *Stack = new(Stack)
    s.top = nil
    return s
}

// return the number of elements on the stack
func (s *Stack) Count() (c uint) {
    var e *element = s.top
    for e != nil {
        c++
        e = e.under
    }
    return
}

// peek at the element on the top of the stack
func (s *Stack) Peek() (int, error) {
    var e *element = s.top
    if e == nil {
        return 0, errors.New("empty stack")
    }
    return e.value, nil
}

// pop an element from the top of the stack
func (s *Stack) Pop() (int, error) {
    var e *element = s.top
    if e == nil {
        return 0, errors.New("empty stack")
    }
    s.top = e.under
    return e.value, nil
}

// push an element on the top of the stack
func (s *Stack) Push(v int) {
    var e *element = new(element)
    e.value = v
    e.under = s.top
    s.top = e
}
