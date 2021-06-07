export GOOS="linux"
export GOARCH="amd64"
go build -o online-im-instance .
docker build --no-cache -t online-im-instance-image  .
rm ./online-im-instance
kubectl create namespace glory
kubectl delete -f ./instance.yaml
kubectl create -f ./instance.yaml
