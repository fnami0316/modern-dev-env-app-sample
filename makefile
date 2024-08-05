local-dev-env-docker-compose-file := deploy/docker-compose/docker-compose.yml
local-dev-env-docker-compose-project := sample_app

# ローカル環境のdocker-composeコマンドエイリアス
# 全コンテナをイメージのビルドから始めて再起動
docker-compose-full-reload: docker-compose-down docker-compose-build docker-compose-up
# 全コンテナイメージのビルド
docker-compose-build:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} build --no-cache
# 全コンテナを起動
docker-compose-up:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} up -d
# 全コンテナを削除
docker-compose-down:
	docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} down
# 全コンテナの状態を表示
docker-compose-ps:
	watch -n 0.05 'docker-compose -f ${local-dev-env-docker-compose-file} -p ${local-dev-env-docker-compose-project} ps -a --format "table {{.Service}}\t{{.Status}}"'

# ローカル環境のdockerコマンドエイリアス
# 指定したサービス名を持つコンテナへログインしてbash起動
docker-login-api:
	docker exec -it `docker ps -qf name=${*}` /bin/bash
# apiのコンテナ上でgo testを実行
docker-go-test:
	docker exec -it `docker ps -qf name=api` go test ./...

.PHONY: docker-compose-full-reload docker-compose-build docker-compose-up docker-compose-down docker-compose-ps docker-login-api docker-go-test
