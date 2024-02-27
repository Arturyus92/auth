#!/bin/bash
source .env

export LOCAL_MIGRATION_DIR=${MIGRATION_DIR}
export LOCAL_MIGRATION_DSN="host=localhost port=$PG_PORT dbname=$PG_DATABASE_NAME user=$PG_USER password=$PG_PASSWORD sslmode=disable"


$(pwd)/bin/goose.exe -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v