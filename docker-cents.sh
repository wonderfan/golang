# https://docs.docker.com/engine/installation/linux/docker-ce/centos/#os-requirements

yum install docker-ce 
systemctl start docker
systemctl is-enabled firewalld
systemctl disable firewalld
systemctl is-active firewalld
systemctl stop firewalld
sysctl vm.max_map_count
sysctl -w vm.max_map_count=262144
