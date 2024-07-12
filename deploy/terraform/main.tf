# Google APIの有効化
variable "gcp_service_list" {
  description = "GCPで有効にするサービスのリスト"
  type        = list(string)
  default     = [
    "run.googleapis.com",
  ]
}

resource "google_project_service" "gcp_services" {
  for_each = toset(var.gcp_service_list)
  service  = each.key
}

# Cloud Run関連
# 参考: https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service
resource "google_cloud_run_v2_service" "api" {
  # サービス名
  name     = "api"
  description = "外部公開するAPIのCloud Run services設定"
  location = var.default_region
  # 外部からのトラフィックの受け入れ許可設定
  ingress = "INGRESS_TRAFFIC_ALL"
  # このサービスにリビジョンを作成する時に使われるテンプレート設定
  template {
    # リビジョンインスタンスを実行する環境の世代
    execution_environment = "EXECUTION_ENVIRONMENT_GEN2"
    # リビジョンインスタンスのオートスケーリング設定
    scaling {
      min_instance_count = var.cloud_run_api.min_instance_count
      max_instance_count = var.cloud_run_api.max_instance_count
    }
    containers {
      name = "go-application"
      # イメージのURL(ダミー)
      image = "us-docker.pkg.dev/cloudrun/container/hello"
      resources {
        # 上限設定(CPU/メモリお互いに関係あり。以下参考)
        # - https://cloud.google.com/run/docs/configuring/cpu?hl=ja
        # - https://cloud.google.com/run/docs/configuring/memory-limits?hl=ja
        limits = {
          cpu    = var.cloud_run_api.limit_cpu
          memory = var.cloud_run_api.limit_memory
        }
        # リクエストがあるときだけCPUを割り当てるか(=コールドスタートを許容するか)
        cpu_idle = var.cloud_run_api.cpu_idle
        # CPUブーストするか(コールドスタート時のレイテンシを低減する)
        startup_cpu_boost = var.cloud_run_api.startup_cpu_boost
      }
      ports {
        # プロトコル
        name = "h2c"
        # コンテナのポート番号(外部から内部への転送先。コンテナの環境変数PORTとしても設定される)
        container_port = var.cloud_run_api.container_port
      }
    }
  }
  depends_on = [google_project_service.gcp_services["run.googleapis.com"]]
}
# Cloud Run services(API) の公開設定
# 参考:
# - https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service_iam
# - https://zenn.dev/t_shunsuke/articles/50a4ff8dd37c77
resource "google_cloud_run_v2_service_iam_binding" "api_all_users" {
  name = google_cloud_run_v2_service.api.name
  role = "roles/run.invoker"
  members = ["allUsers"]
}

resource "google_artifact_registry_repository" "run-image" {
  project       = var.default_project_id
  location      = var.default_region
  repository_id = "api"
  description = "Cloud Run services (API)のイメージを格納するArtifact Registryのリポジトリ"
  format        = "DOCKER"
}
