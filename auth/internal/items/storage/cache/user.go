/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-09 02:13:58
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-18 13:51:51
 * @FilePath: /auth/internal/items/storage/cache/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/projects/pro-sphere-backend/auth/internal/items/models"
	"go.uber.org/zap"
)

type (
	UserCache struct {
		rClient *redis.Client
		logger  *zap.Logger
	}
)

func NewUserCache(rClient *redis.Client, logger *zap.Logger) *UserCache {
	return &UserCache{
		rClient: rClient,
	}
}

func (c *UserCache) SaveUserWithCode(ctx context.Context, tokens *models.Tokens, code string) error {
	cachekey := fmt.Sprintf("user_code:%s", code)
	cachVal, err := json.Marshal(tokens)
	if err != nil {
		c.logger.Error("caching user with code failed while unmarshaling", zap.Error(err))
		return err
	}

	if err = c.rClient.Set(ctx, cachekey, cachVal, time.Minute*2).Err(); err != nil {
		return err
	}
	return nil
}

func (c *UserCache) GetUserTokensByCode(ctx context.Context, code string) (*models.Tokens, error) {
	cachekey := fmt.Sprintf("user_code:%s", code)
	cacheVal, err := c.rClient.Get(ctx, cachekey).Result()
	if err != nil {
		c.logger.Error("Code must be expired or invalid", zap.Error(err))
		return nil, err
	}
	tokens := &models.Tokens{}
	if err := json.Unmarshal([]byte(cacheVal), &tokens); err == nil {
		return tokens, nil
	}

	return tokens, err
}
