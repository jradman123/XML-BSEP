package config

func AccessibleRoles() map[string][]string {
	const postService = "/post_service.PostService/"

	return map[string][]string{
		postService + "create":        {"Regular"},
		postService + "createComment": {"Regular"},
		postService + "likePost":      {"Regular"},
		postService + "dislikePost":   {"Regular"},
	}
}
