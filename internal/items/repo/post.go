package repo

type (
	PostRepository interface {
		CreatePost()
		GetPostById()
		UpdatePost()
		DeletePost()
	}
)
