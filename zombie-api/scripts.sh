#!/bin/bash

echo "Creating mysql container..."

docker pull mysql/mysql-server:5.7

docker run -d -p 3306:3306 --name mysql-zombie \
-e MYSQL_ROOT_PASSWORD=password \
-e MYSQL_ROOT_HOST=% \
-e MYSQL_DATABASE=zombie_data mysql/mysql-server:5.7

echo "Sleep 20 seconds"
sleep 20

echo "Importing sql scripts..."
mysql -hlocalhost -P3306 --protocol=tcp -u root -ppassword \
-Dzombie_data < ./sql-scripts.sql