/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-20 20:44:01
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-20 20:44:03
 * @FilePath: /admin/pkg/redisconn/redis.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package redisconn

import (
	"github.com/go-redis/redis/v8"
	"github.com/projects/pro-sphere-backend/admin/internal/items/config"
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
