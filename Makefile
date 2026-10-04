

gen-oapi:
	cd backend && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.8.0 -config ./api/oapi-codegen.yaml -o ./internal/rest/gen/gen.go ./api/openapi.yaml
	cd ./clients/shared && npm run gen:api
migrate-create:
	goose -s -dir ./backend/repository/migrations create temp sql
