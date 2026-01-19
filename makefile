# airを使ってホットリロードを実行
shop-air:
	docker compose exec shop-go air -c .air.shop.toml
#--------------------- migration start
migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: name parameter is required. Usage: make migration name=<migration_name>"; \
		exit 1; \
	fi
	docker compose exec shop-go migrate create -ext sql -dir migrations $(name)

# マイグレーションを適用
migrate-up:
	docker compose exec shop-go migrate -path migrations -database "mysql://user:user_password@tcp(shop-db:3306)/my_database" up

# マイグレーションをロールバック
migrate-down:
	docker compose exec shop-go migrate -path migrations -database "mysql://user:user_password@tcp(shop-db:3306)/my_database" down -all

migrate-down-one:
	docker compose exec shop-go migrate -path migrations -database "mysql://user:user_password@tcp(shop-db:3306)/my_database" down 1

migrate-refresh:
	@make migrate-down
	@make migrate-up

migrate-fresh:
	docker compose exec shop-db mysql --user=root --password=root_password -e 'DROP DATABASE IF EXISTS `my_database`;'
	docker compose exec shop-db mysql --user=root --password=root_password -e 'CREATE DATABASE `my_database`;'
	@make migrate-up
#--------------------- wire start
wire:
	docker compose exec shop-go wire gen ./internal/di
	docker compose exec shop-go wire gen -output_file_prefix=test_ ./internal/di/wire-test.go ./internal/di/shop-wire-test.go