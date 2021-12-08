package first

import (
	"testing"
)

func TestQueryRow(t *testing.T) {
	row, err := GetQueryRow(12)
	t.Logf("%+v,%+v\n", row, err)
}
