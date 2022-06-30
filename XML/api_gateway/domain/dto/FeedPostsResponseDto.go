package dto

import "gateway/module/domain/model"

type FeedPostsResponseDto struct {
	Feed []model.Post
}
