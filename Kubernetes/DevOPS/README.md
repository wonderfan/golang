# Exploration

### Existing Solutions

1. [kontinuous](https://github.com/AcalephStorage/kontinuous)
   
    - define the pipeline kind resource and specification;
    - use the job kind resource to do the work
    - various agents for the detail work
    - REST API with the clients for key/value from background.
    - The content of deployment:
        * use the kubernetes namespace,deployment and service resource on the basis of image build
    - The pipeline build:
        * how to trigger the build?
    - Use the kubernetes job to do the command task;
    - Leverage the key/value storage and contruct necessary data schema and models
    - reshape the model on the base of existing models
    
2. Ansible Integration

    * the main point is to use ansible module to manage kubernetes's resources
  
3. [AppController](https://github.com/Mirantis/k8s-AppController) 

4. to-do-list
