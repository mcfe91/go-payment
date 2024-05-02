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