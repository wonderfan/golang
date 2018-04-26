## Install IBM Private Cloud

1. Download docker and images
  ```
  URL: http://pokgsa.ibm.com/projects/i/icp-2.1.0.1
  Location: 9.123.114.44:/root/icp-installer
  ```
2. Install docker
  ```
  ./icp_docker_17.09_rhel_x86_64.bin
  ```
3. Load the images
  ```
  tar xf ibm-cloud-private-x86_64-2.1.0.1.tar.gz -O | docker load
  ```
4. Get the configuration files
  ```
  docker run -e LICENSE=accept --rm -v "$(pwd)":/data ibmcom/icp-inception:2.1.0.1-ee cp -r cluster /data
  ```
5. Modify the host and config file
  ```
  change cluster_access_ip to floating IP
  disabled_management_services: ["metering", "monitoring"]
  ```
6. Generate the ssh key
  ```
  ssh-keygen -b 4096 -t rsa -f ~/.ssh/master.id_rsa -N ""
  cat ~/.ssh/master.id_rsa.pub >> ~/.ssh/authorized_keys
  cd cluster
  cp ~/.ssh/master.id_rsa ssh_key
  ssh-copy-id -i ~/.ssh/master.id_rsa root@other-node-ip
  ssh-keyscan 161.202.102.236 | sudo tee -a /root/.ssh/known_hosts
  ```
7. Disable firewall
  ```
  systemctl stop firewalld
  systemctl disable firewalld
  ```
8. Run the install command
  ```
  docker run -e LICENSE=accept --net=host --rm -v /usr/local/bin:/data ibmcom/kubernetes:v1.8.3-ee cp /kubectl /data
  docker run -e LICENSE=accept --net=host --rm -t -v "$(pwd)":/installer/cluster ibmcom/icp-inception:2.1.0.1-ee install
  ```
