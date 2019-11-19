output "bucket_name" {
  value = google_storage_bucket.terratest.name
}

output "file_path" {
  value = "${google_storage_bucket.terratest.name}/${google_storage_bucket_object.test_file.output_name}"
}