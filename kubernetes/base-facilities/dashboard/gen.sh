# 生成证书
openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout dashboard.chenyingqiao.com.key -out dashboard.chenyingqiao.com.pem -subj "/CN=dashboard.chenyingqiao.com"
# 生成secret 注意 命名空间需要和dashboard 所在空间一致
kubectl create secret tls dashboard-secret --namespace=kube-system --cert dashboard.chenyingqiao.com.pem --key dashboard.chenyingqiao.com.key
# 192.168.10.5 是外网可以访问ingress 边缘node的地址，添加这个记录仅仅是为了curl 测试
sudo echo "192.168.10.5 dashboard.chenyingqiao.com" >> /etc/hosts
echo " please use curl -I -k https//dashboard.chenyingqiao.com  to test access "
echo " use kubectl -n kube-system describe secret \$(kubectl -n kube-system get secret | grep admin-user | awk '{print $1}') to get login token"