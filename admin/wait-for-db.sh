#!/bin/bash
###
 # @Author: javohir-a abdusamatovjavohir@gmail.com
 # @Date: 2024-11-02 14:16:16
 # @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 # @LastEditTime: 2024-11-02 14:16:20
 # @FilePath: /sphere_posts/wait-for-db.sh
 # @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
### 

# Wait for the database to be ready
until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER"; do
  echo "Waiting for database..."
  sleep 2
done

echo "Database is ready!"
