/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-08 23:45:29
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-08 23:45:31
 * @FilePath: /auth/pkg/redis/connection.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/projects/pro-sphere-backend/auth/internal/items/config"
)

func ConnectDB(cnf *config.RedisConfig) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     cnf.Host + ":" + cnf.Port,
			Password: "",
			DB:       0,
		},
	)
}
