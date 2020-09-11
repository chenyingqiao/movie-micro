cd ../$HOME
apt update
curl -fsSL https://get.docker.com | bash -s docker
curl -Lo minikube https://storage.googleapis.com/minikube/releases/v1.12.2/minikube-linux-amd64 \
  && chmod +x minikube
sudo mkdir -p /usr/local/bin/
sudo install minikube /usr/local/bin/
apt-get install conntrack
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x kubectl && mv kubectl /usr/bin
curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.6.5 sh -
echo 'export PATH=$HOME/istio-1.6.5/bin:$PATH' >> $HOME/.bashrc
source $HOME/.bashrc
istioctl install --set profile=demo

minikube addons enable ingress

git clone https://github.com/chenyingqiao/movie-micro.git

apt install make

cd movie-micro
make load-base
make load