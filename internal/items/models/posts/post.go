package posts

type (
	Post struct {
		PostId            string `json:"post_id"`
		PostTitle         string `json:"post_title"`
		PostCategory      string `json:"post_category"`
		PostShortContent  string `json:"post_short_content"`
		PostContent       string `json:"post_content"`
		PostFeaturedImage string `json:"post_featured_image"`
		PostStource       string `json:"post_source"`
		ImportedData      string `json:"imported_data"`
		PostViews         int    `json:"views"`
		Deleted           bool   `json:"-"`
	}

	CreatePostRequest struct {
		PostTitle        string `json:"post_title"`
		PostCategory     string `json:"post_category"`
		PostShortContent string `json:"post_short_content"`
		PostContent      string `json:"post_content"`
		PostStource      string `json:"post_source"`
		ImportedData     string `json:"imported_data"`
	}

	GetPostByIdRequest struct {
		PostId string `json:"post_id"`
	}

	UpdatePostRequest struct {
		PostTitle        string `json:"post_title"`
		PostCategory     string `json:"post_category"`
		PostShortContent string `json:"post_short_content"`
		PostContent      string `json:"post_content"`
		PostStource      string `json:"post_source"`
		ImportedData     string `json:"imported_data"`
	}

	DeletePostRequest struct {
		PostId string `json:"post_id"`
	}
)

/*
	CreatePost()
	GetPostById()
	UpdatePost()
	DeletePost()
*/
