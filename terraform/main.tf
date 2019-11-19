module "terratest_test" {
  source     = "./modules/bucket-upload"
  project_id = var.project_id
}