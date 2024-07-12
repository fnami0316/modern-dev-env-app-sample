variable "default_project_id" {
  description = "各リソースのデフォルトのプロジェクトID"
  default = "imposing-sentry-429112-d8"
}
variable "default_region" {
  description = "各リージョンリソースのデフォルトのリージョン"
  default = "us-central1"
}
variable "default_zone" {
  description = "各ゾーンリソースのデフォルトのゾーン"
  default = "us-central1-a"
}

variable "cloud_run_api" {
  description = "Cloud Run services (サービス名:API)の設定値"
  default = {
    # リビジョンインスタンス数最小値
    min_instance_count = 0
    # リビジョンインスタンス数最大値
    max_instance_count = 1
    # コンテナのポート番号(外部から内部への転送先)
    container_port = 8080
    # CPU上限
    limit_cpu = "1"
    # メモリ上限
    limit_memory = "512Mi"
    # リクエストがあるときだけCPU割り当てるか
    cpu_idle = true
    # CPUブーストするか(コールドスタート時のレイテンシを低減する)
    startup_cpu_boost = false
  }
}