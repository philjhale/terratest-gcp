# terratest-gcp
Playing with Terratest on GCP.

# Prerequisites

* A Mac
* Go `brew install go`
* Dep `brew install dep`
* Terraform `brew install terraform`
* Google Cloud SDK `brew cask install google-cloud-sdk`
* A GCP project
* GCP credentials


# Configuration steps

## Terraform

Configure environment variables.
```
export GOOGLE_PROJECT=my-project-id
export TF_VAR_project_id=my-project-id
```

Create a GCP project and switch to it.
```
gcloud iam service-accounts create terraform --display-name=Terraform

gcloud projects add-iam-policy-binding $GOOGLE_PROJECT \
  --member serviceAccount:terraform@$GOOGLE_PROJECT.iam.gserviceaccount.com \
  --role roles/owner

gcloud iam service-accounts keys create ./terraform-key.json \
  --iam-account terraform@$GOOGLE_PROJECT.iam.gserviceaccount.com

# Manually link billing account
gsutil mb -b on -p $GOOGLE_PROJECT gs://${GOOGLE_PROJECT}_terraform
```

Initialise Terraform.
```
export GOOGLE_APPLICATION_CREDENTIALS=$(pwd)/terraform-key.json
terraform init -backend-config=bucket=${GOOGLE_PROJECT}_terraform
```

## Terratest

```
cd tests
dep init
go test -v
```