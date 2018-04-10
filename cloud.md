## Install IBM Private Cloud

1. Get the configuration files
  `docker run -e LICENSE=accept --rm -v "$(pwd)":/data ibmcom/icp-inception cp -r cluster /data`
