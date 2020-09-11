echo "工作目录"$HOME
cd $HOME
apt update

apt install git

echo "安装docker"

curl -fsSL https://get.docker.com | bash -s docker

echo "安装minikube"

curl -Lo minikube https://storage.googleapis.com/minikube/releases/v1.12.2/minikube-linux-amd64 \
  && chmod +x minikube
sudo mkdir -p /usr/local/bin/
sudo install minikube /usr/local/bin/
apt-get install conntrack

echo "安装kubectl"

curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x kubectl && mv kubectl /usr/bin

echo "安装启动minikube"

minikube start --driver=none --extra-config=kubeadm.ignore-preflight-errors=NumCPU --force --cpus 1


echo "安装istio"

curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.6.5 sh -
echo 'export PATH=$HOME/istio-1.6.5/bin:$PATH' >> $HOME/.bashrc
source $HOME/.bashrc
cd istio-1.6.5/
istioctl install --set profile=demo
cd ..

echo "开启ingress-controller"

minikube addons enable ingress

echo "克隆项目"

git clone https://github.com/chenyingqiao/movie-micro.git

echo "构建项目"

apt install make

cd movie-micro
make load-base
make load