## Install IBM Private Cloud

1. Download docker and images
  ```
  URL: http://pokgsa.ibm.com/projects/i/icp-2.1.0.1
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

6. Generate the ssh key
  ```
  ssh-keygen -b 4096 -t rsa -f ~/.ssh/master.id_rsa -N ""
  cat ~/.ssh/master.id_rsa.pub >> ~/.ssh/authorized_keys
  cp ~/.ssh/master.id_rsa cluster/ssh_key;
  ```
  
7. Run the install command
  ```
  docker run -e LICENSE=accept  --net=host --rm -t -v "$(pwd)":/installer/cluster ibmcom/icp-inception:2.1.0.1-ee install
  ```
