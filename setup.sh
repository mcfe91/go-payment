# minikube
# minikube start --memory=4096  
# minikube mount /Users/maxmcferren/go/src/go-micro:/go-micro
# minikube tunnel

# Add kafka quick start
# kubectl create namespace kafka
# kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
# kubectl get pod -n kafka --watch
# kubectl apply -f https://strimzi.io/examples/latest/kafka/kafka-persistent-single.yaml -n kafka
# kubectl wait kafka/my-cluster --for=condition=Ready --timeout=300s -n kafka

# Clean up kafka
# kubectl -n kafka delete $(kubectl get strimzi -o name -n kafka)
# kubectl -n kafka delete -f 'https://strimzi.io/install/latest?namespace=kafka'

# Ingress
# minikube addons enable ingress
# kubectl get pods -n ingress-nginx
# (ingress yaml is in api_gateway/)

# access mysql database
# kubectl port-forward -n kafka mysql-money-movement-7597c855cc-s2psx 3306:3306
# mysql -u root -h 127.0.0.1 -p (make sure local mysql service is stopped)

# fix issue with GKE arch and docker build
# docker buildx build --tag mmcferren/money-movement:latest ./ --platform linux/amd64