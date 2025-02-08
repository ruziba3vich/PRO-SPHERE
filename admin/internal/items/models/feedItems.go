/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:02:16
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-20 11:04:48
 * @FilePath: /admin/internal/items/models/feedItems.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

type (
	FeedItem struct {
		FeedID      int    `json:"feed_id"`
		Title       string `json:"title"`
		Link        string `json:"link"`
		ImageURL    string `json:"image_url"`
		Description string `json:"description"`
		PublishedAt string `json:"published_at"`
	}
	FeedItemResponse struct {
		ID          int    `json:"id"`
		FeedID      int    `json:"feed_id"`
		Link        string `json:"link"`
		Title       string `json:"title"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		PublishedAt string `json:"published_at"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
	UpdateItem struct {
		Id          int
		Title       string `json:"title"`
		Description string `json:"description"`
		PublishedAt string `json:"published_at"`
		ImageURL    string `json:"image_url"`
	}
)
