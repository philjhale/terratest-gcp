resource "google_storage_bucket" "terratest" {
  name               = "${var.project_id}_terratest"
  location           = "EU"
  bucket_policy_only = false
}

resource "google_storage_bucket_object" "test_file" {
  name   = "test.txt"
  source = "${path.module}/test.txt"
  bucket = google_storage_bucket.terratest.name
}

resource "google_storage_object_access_control" "public_rule" {
  object = google_storage_bucket_object.test_file.output_name
  bucket = google_storage_bucket.terratest.name
  role   = "READER"
  entity = "allUsers"
}