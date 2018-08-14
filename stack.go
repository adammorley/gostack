/*
    a stack implementation for stuff (using interface{}) in golang
*/

package stack

import "fmt"
import "reflect"

// a stack is an in-order pile of things; in this particular case, whatever was last added is on the top.
type Stack struct {
    top *element
    count uint
}

// an error can be returned
type StackError struct {
    msg string
}
func (e *StackError) Error() string {
    return fmt.Sprintf("stack %v", e.msg)
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
func (s *Stack) Count() uint {
    return s.count
}

// peek at the element on the top of the stack
func (s *Stack) Peek() (interface{}, error) {
    var e *element = s.top
    if e == nil {
        return 0, &StackError{msg: "empty"}
    }
    return e.value, nil
}

// pop an element from the top of the stack
func (s *Stack) Pop() (interface{}, error) {
    var e *element = s.top
    if e == nil {
        return 0, &StackError{msg: "empty"}
    }
    s.top = e.under
    s.count -= 1
    return e.value, nil
}

// push an element on the top of the stack
// enforces that the types match as elements are pushed
func (s *Stack) Push(i interface{}) error {
    if s.top != nil && reflect.TypeOf(s.top.value) != reflect.TypeOf(i) {
        return &StackError{msg: "type mismatch"}
    }
    var e *element = new(element)
    e.value = i
    e.under = s.top
    s.top = e
    s.count += 1
    return nil
}
