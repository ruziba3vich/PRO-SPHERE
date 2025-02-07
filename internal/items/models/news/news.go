package newsmodels

type (
	News struct {
		NewsId        string `json:"news_id"`
		Name          string `json:"name"`
		Slug          string `json:"slug"`
		Icon          string `json:"icon_url"`
		CategoryOrder int    `json:"category_order"`
	}

	Feed struct {
		FeedId       string `json:"feed_id"`
		FeedName     string `json:"feed_name"`
		FeedImage    string `json:"feed_image"`
		FeedLink     string `json:"feed_link"`
		FeedCategory string `json:"feed_category"`
		FeedOrder    int    `json:"feed_order"`
	}

	CreateNewsRequest struct {
		Name          string `json:"name"`
		Slug          string `json:"slug"`
		CategoryOrder int    `json:"category_order"`
	}

	CreateFeedRequest struct {
		FeedName     string `json:"feed_name"`
		FeedCategory string `json:"feed_category"`
		FeedOrder    int    `json:"feed_order"`
	}

	GetNewsById struct {
		NewsId string `json:"news_id"`
	}

	DeleteNewsById struct {
		NewsId string `json:"news_id"`
	}

	GetFeedById struct {
		FeedId string `json:"feed_id"`
	}

	DeleteFeedById struct {
		FeedId string `json:"feed_id"`
	}
)

/*
	CreateCategory()
	UpdateCategory()
	GetCategoryById()
	GetCategoryByName()
	GetCategoryByOrderNumber()
*/
