package mysql

import (
	"bluebell/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

//创建帖子
func CreatePost(p *models.Post) (err error){
	sqlStr := `insert into post(
	post_id,title,content,author_id,community_id) 
    values(?,?,?,?,?)`

	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorId, p.CommunityID)
	return
}

//根据id查询单个帖子
func GetPostByID(pid int64) (post *models.Post,err error) {
	post = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return
}

//获取帖子详情
func GetPostList(page,size int64) (posts []*models.Post,err error){
	sqlStr := `select post_id,title,content,author_id,community_id,create_time 
		from post 
		ORDER BY create_time DESC 
		limit ?,?`
	posts = make([]*models.Post,0,2)
	err = db.Select(&posts,sqlStr,(page-1)*size,size)
	return
}

//根据给定的id列表查询数据
func GetPostListByIDs(ids []string) (postList []*models.Post,err error){
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id,?)
    `
	query,args,err := sqlx.In(sqlStr,ids,strings.Join(ids,","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList,query,args...)
	return
}
