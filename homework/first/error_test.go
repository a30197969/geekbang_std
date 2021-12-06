package first

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := errors.New("这是一个错误")
	err2 := errors.New("这是一个错误")
	fmt.Printf("%#v\n", err)
	fmt.Printf("%p\n", err)
	fmt.Printf("%p\n", err2)
}
