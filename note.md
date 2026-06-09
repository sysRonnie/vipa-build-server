docker run --rm \
  -p 8080:8080 \
  --env-file .env \
  -e DATABASE_DSN="postgres://sandboxadmin@host.docker.internal:5432/vipa_build_db_local?sslmode=disable" \
  vipa-build-server