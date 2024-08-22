local-dev-env-docker-compose-file := deploy/docker-compose/docker-compose.yml
circleci-docker-compose-file := deploy/docker-compose/docker-compose-circleci.yml
local-dev-env-docker-compose-project := sample_app
proto_dir := ./api/proto

# .protoからGolang向けのgRPCコードを生成
protoc:
	protoc \
		--go_out=. \
		--go-grpc_out=. \
		$(proto_dir)/*.proto

# ローカル環境のdocker-composeコマンドエイリアス
# 全コンテナをイメージのビルドから始めて再起動
docker-compose-full-reload: docker-compose-down docker-compose-build docker-compose-up
docker-compose-full-reload-d: docker-compose-down docker-compose-build docker-compose-up-d
# 全コンテナイメージのビルド
docker-compose-build:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} build --no-cache
# 全コンテナを起動
docker-compose-up:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} up
	# ヘルスチェックは起動時のみに行うため削除
	make docker-compose-down-spanner-emulator-healthcheck
	make docker-compose-down-api-healthcheck
# 全コンテナをバックグラウンドで起動
docker-compose-up-d:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} up -d
	# ヘルスチェックは起動時のみに行うため削除
	make docker-compose-down-spanner-emulator-healthcheck
	make docker-compose-down-api-healthcheck
# 全コンテナを削除
docker-compose-down:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} down
# spanner-emulatorコンテナのみを削除
docker-compose-down-spanner-emulator:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} rm -fsv spanner-emulator
# spanner-emulator-healthcheckコンテナのみを削除
docker-compose-down-spanner-emulator-healthcheck:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} rm -fsv spanner-emulator-healthcheck
# api-healthcheckコンテナのみを削除
docker-compose-down-api-healthcheck:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} rm -fsv api-healthcheck
# 全コンテナの状態を表示
docker-compose-ps:
	watch -n 0.05 'docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} ps -a --format "table {{.Service}}\t{{.Status}}"'

# CircleCI用のdocker-composeコマンドエイリアス
# 全コンテナをバックグラウンドで起動
docker-compose-circleci-up-d:
	docker-compose -f ${circleci-docker-compose-file} -p ${local-dev-env-docker-compose-project} up -d
	# ヘルスチェックは起動時のみに行うため削除
	docker-compose -f ${circleci-docker-compose-file} -p ${local-dev-env-docker-compose-project} rm -fsv spanner-emulator-healthcheck
	docker-compose -f ${circleci-docker-compose-file} -p ${local-dev-env-docker-compose-project} rm -fsv api-healthcheck
# 全コンテナを削除
docker-compose-circleci-down:
	docker-compose -f ${circleci-docker-compose-file} -p ${local-dev-env-docker-compose-project} down
# 全コンテナの状態を表示
docker-compose-circleci-ps:
	watch -n 0.05 'docker-compose -f ${circleci-docker-compose-file} -p ${local-dev-env-docker-compose-project} ps -a --format "table {{.Service}}\t{{.Status}}"'


# dockerコマンドエイリアス
# apiコンテナへログインしてbash起動
docker-login-api:
	docker exec -it `docker ps -qf name=api` /bin/bash
# apiコンテナ上でgo testを実行
docker-go-test:
	docker exec -it `docker ps -qf name=api` go test ./...
# apiコンテナ上でgo testを実行(直列)
docker-go-test-serial:
	docker exec -it `docker ps -qf name=api` go test -p 1 ./...

.PHONY: docker-compose-full-reload docker-compose-build docker-compose-up docker-compose-down docker-compose-ps docker-login-api docker-go-test api
