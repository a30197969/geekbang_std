package first

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// QueryRow
// 毛剑老师第二周作业：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 答：不应该Wrap这个error，因为sql.ErrNoRows是数据库执行QueryRow查询的时候结果为空返回的，是Go定义的一个特殊的常量，这需要作为特殊情况处理，不能将该空结果作为错误，在判断中应该检查错误是否为这个特殊常量，不是才可以Wrap这个错误
// 代码如下，查询 order 表中是否有 id = 100 的数据，如果该行不存在，则需要特殊处理sql.ErrNoRows
func QueryRow(id int) error {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/market")
	if err != nil {
		return errors.Wrap(err, "连接 market 库报错")
	}
	defer db.Close()
	var price float32
	err = db.QueryRow("select price from `order` where id = ?", id).Scan(price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		} else {
			return errors.Wrap(err, "查询 price 表报错")
		}
	}
	return nil
}
