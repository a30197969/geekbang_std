package first

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // 数据库驱动
	"github.com/pkg/errors"
)

// AlbumData 存储作品表的结构体
type AlbumData struct {
	id          int    // 主键ID
	title       string // 作品标题
	description string // 作品描述
	userid      int    // 用户ID
	dateline    string // 发布时间
}

// GetQueryRow
// 毛剑老师第二周作业：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 答：不应该Wrap这个error，因为sql.ErrNoRows是数据库执行QueryRow查询的时候结果为空返回的，是Go定义的一个特殊的常量，这需要作为特殊情况处理，不能将该空结果作为错误，在判断中应该检查错误是否为这个特殊常量，不是才可以Wrap这个错误
// 代码如下，查询 `album` 表中是否有 id = 12 的数据，如果该行不存在，则需要特殊处理sql.ErrNoRows
// 返回结果：{id:12 title:王者的风范 description:摄于曲靖 userid:15 dateline:2021-06-04 10:37:19},<nil>
func GetQueryRow(id int) (AlbumData, error) {
	var (
		db        *sql.DB
		err       error
		albumData AlbumData
	)

	db, err = sql.Open("mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/fn_bird")
	if err != nil {
		return albumData, errors.Wrap(err, "连接 fn_bird 库报错")
	}
	defer db.Close()

	err = db.QueryRow("select id,title,description,userid,dateline from `album` where id = ? limit 1", id).Scan(&albumData.id, &albumData.title, &albumData.description, &albumData.userid, &albumData.dateline)
	if err != nil {
		if err != sql.ErrNoRows {
			return albumData, errors.Wrap(err, "查询 fn_bird.album 表报错")
		}
	}
	return albumData, nil
}
