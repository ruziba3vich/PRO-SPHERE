/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-26 02:37:19
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-29 14:24:37
 * @FilePath: /admin/internal/items/models/feeds.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

type Feed struct {
	Id            int           `json:"id"`
	Priority      int           `json:"priority"`
	MaxItems      int           `json:"max_items"`
	BaseUrl       string        `json:"base_url"`
	LogoUrl       string        `json:"logo_url"`
	LogoUrlId     string        `json:"logo_url_id"`
	Translations  []Translation `json:"translations"`
	LastRefreshed string        `json:"last_refreshed"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}

type Translation struct {
	Id          int    `json:"id"`
	FeedId      int    `json:"feed_int"`
	Lang        string `json:"lang"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FeedContent struct {
	Id         int    `json:"id"`
	FeedId     int    `json:"feed_id"`
	Lang       string `json:"lang"`
	Link       string `json:"link"`
	CategoryId int    `json:"category_id"`
}

type FeedLogoInfo struct {
	LogoUrl   string `json:"logo_url"`
	LogoUrlId string `json:"logo_url_id"`
}
