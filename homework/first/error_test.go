package first

import (
	"testing"
)

func TestQueryRow(t *testing.T) {
	row, err := GetQueryRow(12)
	t.Logf("%+v\n", row)
	t.Logf("%+v\n", err)
}
