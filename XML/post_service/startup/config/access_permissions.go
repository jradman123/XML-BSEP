package config

func AccessiblePermissions() map[string]string {
	const postService = "/post_service_proto.PostService/"

	return map[string]string{
		postService + "CreatePost":    "createPostPermission",
		postService + "CreateComment": "createCommentPermission",
		postService + "LikePost":      "likePostPermission",
		postService + "DislikePost":   "dislikePostPermission",
	}
}
