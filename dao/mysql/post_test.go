package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "123456",
		DbName:       "bluebell",
		Port:         3306,
		MaxOpenConns: 10,
		MaxIdleConns: 10,
		MinIdleConns: 10,
	}
	err := Init(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          123,
		AuthorId:    1,
		CommunityID: 1,
		Status:      1,
		Title:       "asdas",
		Content:     "asdasd",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("post failed,err:%v\n",err)
	}
	t.Logf("success")
}
