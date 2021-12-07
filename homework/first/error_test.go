package first

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"testing"
)

func TestQueryRow(t *testing.T) {
	title, err := GetAlbumTitle(13)
	fmt.Printf("%+v,%+v\n", title, err)
}
func GetAlbumTitle(userid int) (string, error) {
	db, err := sql.Open("mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/fn_bird")
	if err != nil {
		return "", errors.Wrap(err, "连接 iam 库报错")
	}
	defer db.Close()

	var title string
	err = db.QueryRow("select title from `album` where userid = ?", userid).Scan(&title)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			return "", errors.Wrap(err, "查询 fn_bird.album 表报错")
		}
	}
	return title, nil
}
