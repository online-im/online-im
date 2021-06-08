kubectl create ns glory
kubectl create -f ./cmd/redis.yaml
cd ./cmd/instance
sh build.sh
cd ../gateway
sh build.sh