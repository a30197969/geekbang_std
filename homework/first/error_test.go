package first

import (
	"fmt"
	"testing"
)

func TestQueryRow(t *testing.T) {
	row, err := GetQueryRow(12)
	fmt.Printf("%+v,%+v\n", row, err)
}
