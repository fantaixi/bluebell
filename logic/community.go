package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community,error) {
	//查询数据库并且返回
	return mysql.GetCommunityList()
}

//根据id查询社区分类查询
func GetCommunityDetail(id int64) (*models.CommunityDetail,error){
	return mysql.GetCommunityDetailByID(id)
}
