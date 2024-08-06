package posts

type (
	Post struct {
		PostId            string `json:"post_id"`
		PublisherId       string `json:"publisher_id"`
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
		PublisherId      string `json:"publisher_id"`
		PostTitle        string `json:"post_title"`
		PostCategory     string `json:"post_category"`
		PostShortContent string `json:"post_short_content"`
		PostContent      string `json:"post_content"`
		PostStource      string `json:"post_source"`
		ImportedData     string `json:"imported_data"`
	}

	// gets

	GetPostByPublisherIdRequest struct {
		PublisherId string `json:"publisher_id"`
	}

	GetPostByIdRequest struct {
		PostId string `json:"post_id"`
	}

	GetAllPostsRequest struct {
		Page  int
		Limit int
	}

	// udates

	UpdatePostRequest struct {
		PostId           string `json:"post_id"`
		PostTitle        string `json:"post_title"`
		PostCategory     string `json:"post_category"`
		PostShortContent string `json:"post_short_content"`
		PostContent      string `json:"post_content"`
		PostStource      string `json:"post_source"`
		ImportedData     string `json:"imported_data"`
	}

	// deletions

	DeletePostRequest struct {
		PostId string `json:"post_id"`
	}

	// responses

	GetPostsResponse struct {
		Posts []*Post `json:"response"`
	}
)

/*
	CreatePost()
	GetPostById()
	UpdatePost()
	DeletePost()
*/
