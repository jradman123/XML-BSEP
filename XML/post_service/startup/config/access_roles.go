package config

func AccessibleRoles() map[string][]string {
	const postService = "/post_service.PostService/"

	return map[string][]string{
		postService + "CreatePost":    {"Regular"},
		postService + "CreateComment": {"Regular"},
		postService + "LikePost":      {"Regular"},
		postService + "DislikePost":   {"Regular"},
	}
}
