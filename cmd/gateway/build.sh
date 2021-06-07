export GOOS="linux"
export GOARCH="amd64"
go build -o online-im-gateway .
docker build --no-cache -t online-im-gateway-image  .
rm ./online-im-gateway
kubectl create namespace glory
kubectl delete -f ./gateway.yaml
kubectl create -f ./gateway.yaml
