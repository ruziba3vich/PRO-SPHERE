/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-09-21 19:53:48
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-11-22 23:27:43
 * @FilePath: /sfere_backend/internal/items/storage/db.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/projects/pro-sphere-backend/admin/internal/items/config"

	_ "github.com/lib/pq"
)

func ConnectDB(config *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.Database.User,
		config.Database.DBName,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
	)
	log.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
