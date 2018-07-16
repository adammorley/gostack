/*
    a stack implementation for stuff (using interface{}) in golang
*/

package stack

import "errors"
import "reflect"

type Stack struct {
    top *element
}

type element struct {
    value interface{}
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
func (s *Stack) Peek() (interface{}, error) {
    var e *element = s.top
    if e == nil {
        return 0, errors.New("empty stack")
    }
    return e.value, nil
}

// pop an element from the top of the stack
func (s *Stack) Pop() (interface{}, error) {
    var e *element = s.top
    if e == nil {
        return 0, errors.New("empty stack")
    }
    s.top = e.under
    return e.value, nil
}

// push an element on the top of the stack
// enforces that the types match as elements are pushed
func (s *Stack) Push(i interface{}) {
    var e *element = new(element)
    if s.top != nil && reflect.TypeOf(s.top.value) != reflect.TypeOf(i) {
        panic("type mismatch")
    }
    e.value = i
    e.under = s.top
    s.top = e
}
