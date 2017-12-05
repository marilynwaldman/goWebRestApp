This is a test GO RESTFUL API used to test Jenkins on a Google Cloud platform.

Push to Google Cloud without CORS first. Check with Curl.

Tutorial from :

https://medium.com/google-cloud/running-a-simple-kubernetes-frontend-backend-service-on-google-cloud-platform-85eb0346f600
Manual Instruction for testing; follow with Jenkins workflow

gcloud config set project jenkins-187820 gcloud config set compute/zone us-west1-b

gcloud container clusters create kubetest --num-nodes=3

git clone https://github.com/marilynwaldman/webappDogPark.git

cd web*

docker build -t backend api 
docker tag backend gcr.io/jenkins-187820/backend 
gcloud docker -- push gcr.io/jenkins-187820/backend

cd ..

kubectl create -f be-rc.yaml kubectl create -f be-srv.yaml

kubectl get service

curl http://....

kubectl logs -f POD-NAME

kubectl delete -f be-srv.yaml kubectl delete -f be-rc.yaml

Jenkins stuff

docker build -t be . 
docker run --rm -p 3000:3000 dog

Work around until yml files work:

https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app?authuser=1

kubectl run be-rc --image=gcr.io/testkube-187517/be --port 3000

kubectl expose deployment be-rc --type=LoadBalancer --port 80 --target-port 3000 http://35.203.157.252/dogparks Utilities

###!/bin/bash

Delete all containers

docker rm $(docker ps -a -q)

Delete all images

docker rmi $(docker images -q)


