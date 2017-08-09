package lib

import "fmt"

type ErrWrapper struct {
    err error
}

func (ew *ErrWrapper) do(f func() error) bool {
    ew.err = f();
    if ew.err != nil {
        fmt.Println(ew.err.Error())
        return false
    }
    return true
}
