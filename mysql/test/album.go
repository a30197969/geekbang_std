package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"strconv"
)

type Album struct {
	id            int
	albumType     int `json:"type"`
	title         string
	description   string
	userid        int
	picCount      int
	commentCount  int
	zanCount      int
	pvCount       int
	ip            string
	source        int
	lastCommentId int
	isJinghua     int
	isDel         int
	dateline      string
}

var (
	Db        *sql.DB
	DbName    = "fn_bird"
	TableName = "album"
)

func (album *Album) Conn() error {
	var err error
	Db, err = sql.Open("mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/"+DbName+"?charset=utf8mb4&parseTime=true&loc=Local") //  返回 *sql.DB 数据库操作的指针对象，并不是一个连接
	return err
}

// Insert 插入行
func (album *Album) Insert() (int64, error) {
	err := album.Conn()
	if err != nil {
		return 0, errors.Wrap(err, "连接 "+DbName+" 库失败")
	}
	defer Db.Close()

	rs, err := Db.Exec("insert into "+TableName+" (`type`,`title`,`description`,`userid`,`pic_count`,`ip`,`source`,`dateline`) values (?,?,?,?,?,?,?,?)",
		album.albumType, album.title, album.description, album.userid, album.picCount, album.ip, album.source, album.dateline)
	if err != nil {
		return 0, errors.Wrap(err, "插入 "+DbName+"."+TableName+" 表失败")
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, errors.Wrap(err, "获取"+DbName+"."+TableName+"表最新插入的ID失败")
	}
	return id, nil
}

// Delete 删除数据
func (album *Album) Delete() (int64, error) {
	err := album.Conn()
	if err != nil {
		return 0, errors.Wrap(err, "连接 "+DbName+" 库失败")
	}
	defer Db.Close()
	rs, err := Db.Exec("delete from "+TableName+" where id=?", album.id)
	if err != nil {
		return 0, errors.Wrap(err, "删除"+DbName+"."+TableName+"表ID="+strconv.Itoa(album.id)+"的行失败")
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "获取"+DbName+"."+TableName+"表ID="+strconv.Itoa(album.id)+"删除的受影响行失败")
	}
	return rows, nil
}

func MultiSelect(userid int) (res []Album, err error) {
	Db, err = sql.Open("mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/"+DbName) //  返回 *sql.DB 数据库操作的指针对象，并不是一个连接
	if err != nil {
		return res, errors.Wrap(err, "连接 "+DbName+" 库失败")
	}
	defer Db.Close()
	rows, err := Db.Query("select id,title,description,dateline from "+TableName+" where userid=?", userid)
	if err != nil {
		return res, errors.Wrap(err, "查询"+DbName+"."+TableName+"表userid="+strconv.Itoa(userid)+"失败")
	}
	defer rows.Close()
	for rows.Next() {
		album := Album{}
		err := rows.Scan(&album.id, &album.title, &album.description, &album.dateline)
		if err != nil {
			if err = rows.Err(); err != nil {
				return res, errors.Wrap(err, "扫描"+DbName+"."+TableName+"表userid="+strconv.Itoa(userid)+"时报错")
			}
		}
		res = append(res, album)
	}
	return
}
func SelectAll() (res []map[string]string, err error) {
	Db, err = sql.Open("mysql", "fengniao:fengniao123@tcp(172.16.151.61:3306)/"+DbName) //  返回 *sql.DB 数据库操作的指针对象，并不是一个连接
	if err != nil {
		return res, errors.Wrap(err, "连接 "+DbName+" 库失败")
	}
	defer Db.Close()
	rows, err := Db.Query("select * from " + TableName + " order by id desc limit 10")
	if err != nil {
		return res, errors.Wrap(err, "查询"+DbName+"."+TableName+"表失败")
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return res, errors.Wrap(err, "获取"+DbName+"."+TableName+"表的字段名失败")
	}
	vals := make([][]byte, len(cols))
	scans := make([]interface{}, len(cols))
	for i := range vals {
		scans[i] = &vals[i]
	}
	fmt.Println(cols)
	for rows.Next() {
		err := rows.Scan(scans...)
		if err != nil {
			if err = rows.Err(); err != nil {
				return res, errors.Wrap(err, "扫描"+DbName+"."+TableName+"表时报错")
			}
		}
		row := make(map[string]string)
		for k, v := range vals {
			key := cols[k]
			row[key] = string(v)
		}
		res = append(res, row)
	}
	for i, re := range res {
		fmt.Println(i, re)
	}
	return
}
