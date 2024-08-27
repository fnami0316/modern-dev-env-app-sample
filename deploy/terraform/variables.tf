variable "default_project_id" {
  type = string
  description = "各リソースのデフォルトのプロジェクトID"
  default = "imposing-sentry-429112-d8"
}
variable "default_region" {
  type = string
  description = "各リージョンリソースのデフォルトのリージョン"
  default = "us-central1"
}
variable "default_zone" {
  type = string
  description = "各ゾーンリソースのデフォルトのゾーン"
  default = "us-central1-a"
}

variable "cloud_run_api" {
  type = object({
      min_instance_count = number
      max_instance_count = number
      container_port = number
      limit_cpu = string
      limit_memory = string
      cpu_idle = bool
      startup_cpu_boost = bool
    })
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

variable "spanner_instance_dev" {
  description = "Cloud Spannerのインスタンスの設定値"
  type = object({
    name = string
    config = string
    display_name = string
    num_nodes = number
  })
  default = {
    # インスタンス名(ID)
    name = "sentry-429112-dev"
    # インスタンス構成
    config = "regional-us-central1"
    # 表示名
    display_name = "sentry-429112-dev"
    # ノード数
    num_nodes = 1
  }
}

variable "spanner_database_dev" {
  description = "Cloud Spannerのデータベースの設定値"
  type = object({
    instance = string
    name = string
  })
  default = {
    # インスタンス名(ID)
    instance = "sentry-429112-dev"
    # データベース名(ID)
    name = "sentry-429112-dev-1"
  }
}
