/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:02:16
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-27 03:33:00
 * @FilePath: /admin/internal/items/repo/categories.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%A
 */
package repo

import (
	"context"

	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
)

type FeedCategoriesRepository interface {
	CreateFeedCategory(ctx context.Context, category *models.FeedCategory) (*models.FeedCategory, error)
	DeleteFeedCategory(ctx context.Context, id int64) error
	GetAllFeedCategories(ctx context.Context, page int, limit int, lang string) ([]*models.FeedCategory, error)
	GetFeedCategoryByID(ctx context.Context, id int64, lang string) (*models.FeedCategory, error)
	UpdateFeedCategory(ctx context.Context, category *models.FeedCategory) (*models.FeedCategory, error)
	GetFeedCategoryIconInfo(ctx context.Context, id int) (*models.FeedIconInfo, error)
}
