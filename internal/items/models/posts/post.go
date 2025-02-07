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
		PostSource        string `json:"post_source"`
		ImportedData      string `json:"imported_data"`
		PostViews         int    `json:"views"`
		Deleted           bool   `json:"deleted"`
	}

	CreatePostRequest struct {
		PublisherId      string `json:"publisher_id"`
		PostTitle        string `json:"post_title"`
		PostCategory     string `json:"post_category"`
		PostShortContent string `json:"post_short_content"`
		PostContent      string `json:"post_content"`
		PostSource       string `json:"post_source"`
		ImportedData     string `json:"imported_data"`
		PostViews        int    `json:"views"`
	}

	// gets

	GetPostByPublisherIdRequest struct {
		PublisherId string `json:"publisher_id"`
	}

	GetPostByIdRequest struct {
		PostId string `json:"post_id"`
	}

	GetAllPostsRequest struct {
		Page  string
		Limit string
	}

	// udates

	UpdatePostRequest struct {
		PostId           string `json:"post_id"`
		PostTitle        string `json:"post_title"`
		PostCategory     string `json:"post_category"`
		PostShortContent string `json:"post_short_content"`
		PostContent      string `json:"post_content"`
		PostSource       string `json:"post_source"`
		PostViews        int    `json:"views"`
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

	//

	AddPostView struct {
		PostId string `json:"post_id"`
	}
)

/*
	CreatePost()
	GetPostById()
	UpdatePost()
	DeletePost()
*/
