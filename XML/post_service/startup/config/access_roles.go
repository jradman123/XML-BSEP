package config

func AccessibleRoles() map[string][]string {
	const postService = "/post_service.PostService/"

	return map[string][]string{
		postService + "CreatePost":    {"User"},
		postService + "CreateComment": {"User"},
		postService + "LikePost":      {"User"},
		postService + "DislikePost":   {"User"},
	}
}
