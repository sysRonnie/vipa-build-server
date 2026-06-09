1. gcloud auth configure-docker us-east1-docker.pkg.dev
2. gcloud artifacts repositories create vipa-build 

gcloud artifacts repositories create vipa-build \
  --repository-format=docker \
  --location=us-east1

docker tag vipa-build-server \
us-east1-docker.pkg.dev/vipa-build-prod/vipa-build/vipa-build-server:latest

docker push \
us-east1-docker.pkg.dev/vipa-build-prod/vipa-build/vipa-build-server:latest