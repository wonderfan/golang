## Install IBM Private Cloud

1. Get the configuration files
  ```
  docker run -e LICENSE=accept --rm -v "$(pwd)":/data ibmcom/icp-inception cp -r cluster /data
  ```
2. Modify the host and config file

3. Generate the ssh key
  ```
  ssh-keygen -b 4096 -t rsa -f ~/.ssh/master.id_rsa -N ""
  cat ~/.ssh/master.id_rsa.pub >> ~/.ssh/authorized_keys
  cp ~/.ssh/master.id_rsa cluster/ssh_key;
  ```
  
4. Run the install command
  ```
  docker run -e LICENSE=accept  --net=host --rm -t -v "$(pwd)":/installer/cluster ibmcom/icp-inception install
  ```
