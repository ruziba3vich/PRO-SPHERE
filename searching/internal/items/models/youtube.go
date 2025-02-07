/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-20 20:57:02
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-04 16:23:42
 * @FilePath: /searching/internal/items/models/youtube.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

type VideoSearchRequest struct {
	Query         string `json:"query"`
	NextPageToken string `json:"nextPageToken"`
	MaxResults    string `json:"maxResults"`
}

type YouTubeSearchResponse struct {
	Kind          string       `json:"kind"`
	Etag          string       `json:"etag"`
	NextPageToken string       `json:"nextPageToken"`
	RegionCode    string       `json:"regionCode"`
	PageInfo      PageInfo     `json:"pageInfo"`
	Items         []SearchItem `json:"items"`
}

type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type SearchItem struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      ItemID  `json:"id"`
	Snippet Snippet `json:"snippet"`
}

type ItemID struct {
	Kind    string `json:"kind"`
	VideoID string `json:"videoId,omitempty"`
}

type Snippet struct {
	PublishedAt          string     `json:"publishedAt"`
	ChannelID            string     `json:"channelId"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Thumbnails           Thumbnails `json:"thumbnails"`
	ChannelTitle         string     `json:"channelTitle"`
	LiveBroadcastContent string     `json:"liveBroadcastContent"`
	PublishTime          string     `json:"publishTime"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
	Medium  Thumbnail `json:"medium"`
	High    Thumbnail `json:"high"`
}

type (
	SearchVideoRes struct {
		Title        string      `json:"title"`
		ChannelTitle string      `json:"youtubeChannelTitle"`
		ChanneLink   string      `json:"channelLink"`
		VideoLink    string      `json:"videoLink"`
		PublishTime  string      `json:"publishTime"`
		Images       VideoImages `json:"images"`
	}

	VideoImages struct {
		Default VideoImage `json:"default"`
		Medium  VideoImage `json:"medium"`
		High    VideoImage `json:"high"`
	}

	VideoImage struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	}
)

type SearchVideoResponse struct {
	NextPageToken string `json:"nextPageToken"`
	Results       []SearchVideoRes
}
