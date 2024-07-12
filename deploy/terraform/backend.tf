terraform {
  backend "gcs" {
    bucket = "tfstate-94696b5b-d077-4c53-a0be-7dbac69be4b7"
  }
}

# tfstateを保存するためのGCSバケット
resource "google_storage_bucket" "tfstate" {
  name     = "tfstate-94696b5b-d077-4c53-a0be-7dbac69be4b7"
  location = var.default_region
  storage_class = "STANDARD"
  versioning {
    enabled = true
  }
}