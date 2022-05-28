package dao

import (
	"SimpleTikTok/commom"
	"SimpleTikTok/model"
	"errors"
	"sync"
)

type VideoDao struct {
}

var videoDao *VideoDao
var publishOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	publishOnce.Do(
		func() {
			videoDao = new(VideoDao)
		})

	return videoDao
}

// func (*PostDao) QueryPostById(id int64) (*Post, error) {
// 	var post Post
// 	err := db.Where("id = ?", id).Find(&post).Error
// 	if err == gorm.ErrRecordNotFound {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		util.Logger.Error("find post by id err:" + err.Error())
// 		return nil, err
// 	}
// 	return &post, nil
// }

func (*VideoDao) QueryVideoByUserId(userId int64) ([]*model.Videos, error) {
	var videos []*model.Videos
	db := commom.GetDB()
	err := db.Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		return nil, errors.New("无上传视屏")
	}
	return videos, nil
}

func (*VideoDao) CreateVideo(video *model.Videos) error {
	db := commom.GetDB()
	if err := db.Create(video).Error; err != nil {
		return errors.New("视频添加失败")
	}
	return nil
}