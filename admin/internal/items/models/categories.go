/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-10-04 16:03:55
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-22 05:21:45
 * @FilePath: /sphere_posts/internal/items/models/categories.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

// FeedCategory represents the main feed category.
type FeedCategory struct {
	ID           int64                     `json:"id"`
	IconURL      string                    `json:"icon_url"`
	IconID       string                    `json:"icon_id"`
	Translations []FeedCategoryTranslation `json:"translations"`
}

type FeedIconInfo struct {
	IconUrl string `json:"icon_url"`
	IconID  string `json:"icon_id"`
}

// FeedCategoryTranslation represents a translation for a feed category.
type FeedCategoryTranslation struct {
	Lang string `json:"lang"` // Language code (e.g., "uz", "ru", "en")
	Name string `json:"name"` // Translated name of the category
}
