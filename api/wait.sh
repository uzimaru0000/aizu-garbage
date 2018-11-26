#!/bin/sh

echo "Waiting for mysql"

until mysql --host="$MYSQL_HOST" --user="$MYSQL_USER" --password="$MYSQL_PASSWORD" &> /dev/null
do
    echo "waiting..."
    sleep 1
done

echo "MySQL is up"
./api
