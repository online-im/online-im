kubectl create ns glory
kubectl create -f ./cmd/redis.yaml
cd ./cmd/instance
sh build.sh
echo "setup in k8s procedure finished"