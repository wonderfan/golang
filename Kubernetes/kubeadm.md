# vim  /etc/yum.repos.d/kubernetes.repo
[kubernetes] 
name=Kubernetes 
baseurl=http://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64
enabled=1 
gpgcheck=1 
repo_gpgcheck=1 
gpgkey=http://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg
http://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg

# yum makecache

# yum install -y  kubelet kubeadm kubectl kubernetes-cni 

# systemctl enable docker && systemctl start docker 
# systemctl enable kubelet && systemctl start kubelet 

kubectl proxy --port=8080 --address=0.0.0.0 --accept-hosts=^*$ &

scp root@9.111.215.74:/etc/kubernetes/admin.conf .
kubectl --kubeconfig ./admin.conf proxy --port=5000 --address=0.0.0.0 --accept-hosts=^*$ &

kubeadm join  9.111.111.55 --token d1a11f.6b46833bb5c4a498

[[ ubuntu ]]

apt-get update
apt-get install -y docker.io

apt-get update && apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -

cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF

apt-get update
apt-get install -y kubelet kubeadm kubectl

kubeadm init --pod-network-cidr=10.244.0.0/16  --kubernetes-version=1.6

mkdir -p $HOME/.kube
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
chown $(id -u):$(id -g) $HOME/.kube/config
export KUBECONFIG=/etc/kubernetes/admin.conf

kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/v0.9.1/Documentation/kube-flannel.yml

kubectl taint nodes --all node-role.kubernetes.io/master-

apiVersion: v1
kind: ConfigMap
metadata:
  name: kube-dns
  namespace: kube-system
data:
  upstreamNameservers: |
    ["9.111.221.77", "8.8.8.8", "8.8.4.4"]

kubeadm reset

[[helm]]

curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get > get_helm.sh
chmod 700 get_helm.sh
./get_helm.sh

kubectl create clusterrolebinding helm-admin --clusterrole=cluster-admin --serviceaccount=kube-system:default

helm install --name my-release stable/spinnaker  

helm install --name my-release --version 0.3.5 --timeout 300 -f values.yaml stable/spinnaker

helm fetch --untar --version 0.9.0  stable/jenkins

kubectl create clusterrolebinding system:anonymous --clusterrole=cluster-admin --user=system:anonymous

curl -Lo values.yaml https://raw.githubusercontent.com/kubernetes/charts/master/stable/spinnaker/values.yaml

[[kubectl]]
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl;
chmod +x ./kubectl;
sudo mv ./kubectl /usr/local/bin/kubectl; 
kubectl create clusterrolebinding apps-deployment --clusterrole=cluster-admin --serviceaccount=kube-system:default 

