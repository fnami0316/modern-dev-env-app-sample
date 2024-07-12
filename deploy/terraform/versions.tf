terraform {
  # terraform自体のバージョン制約
  required_version = "1.9.1"

  # 使用するproviderとバージョン制約
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "5.36.0"
    }
  }
}