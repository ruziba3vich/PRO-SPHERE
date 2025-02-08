/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:02:16
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-01 00:54:40
 * @FilePath: /pro-sphere-backend/admin/internal/items/repo/feedItems.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package repo

import (
	"context"

	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
)

type (
	FeedItemsRepository interface {
		CreateFeedItem(ctx context.Context, item *models.FeedItem) (*models.FeedItemResponse, error)
		GetFeedItemByID(ctx context.Context, id int) (*models.FeedItemResponse, error)
		UpdateFeedItem(ctx context.Context, updatedItem *models.UpdateItem) (*models.FeedItemResponse, error)
		DeleteFeedItem(ctx context.Context, id int) error
		GetAllFeedItems(ctx context.Context, limit, page int) ([]*models.FeedItemResponse, error)
		GetAllFeedItemsByFeedId(ctx context.Context, feedId, limit, page int) ([]*models.FeedItemResponse, error)
		GetAllFeedItemsByFeedCategoryId(ctx context.Context, catergoryId, limit, page int) ([]*models.FeedItemResponse, error)
	}
)
