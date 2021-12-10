package test

import (
	"testing"
	"time"
)

func TestSelectAlbum(t *testing.T) {
	res, err := MultiSelect(58)
	SelectAll()
	t.Logf("%v, %v\n", res, err)
}

func TestInsertAlbum(t *testing.T) {
	now := time.Now().Format("2006-01-02 15:04:05")
	//var r *http.Request
	//ip := exnet.ClientPublicIP(r)
	//if ip == "" {
	//	ip = exnet.ClientIP(r)
	//}
	album := Album{
		albumType:   1,
		title:       "Go 测试标题",
		description: "Go 测试描述",
		userid:      7970743,
		picCount:    3,
		ip:          "192.168.80.254",
		source:      1,
		dateline:    now,
	}
	id, err := album.Insert()
	t.Logf("%v, %v\n", id, err)
}
func TestDeleteAlbum(t *testing.T) {
	album := Album{
		id: 163,
	}
	rows, err := album.Delete()
	t.Logf("%v, %v\n", rows, err)
}
