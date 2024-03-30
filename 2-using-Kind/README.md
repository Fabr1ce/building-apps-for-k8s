# building-apps-for-k8s-l4-using-Kind

This repository is used to build a _Kubernetes_ cluster locally for testing using _Kind_ and run a pod on that cluster using a docker image that was created in the [1-using-Docker section](https://github.com/Fabr1ce/building-apps-for-k8s/tree/cool-dev/1-using-Docker). The assumption here is that docker and kind are aleady installed locally, if not the following links are used to [install docker](https://docs.docker.com/engine/install/) and to [install kind](https://kubernetes.io/docs/tasks/tools/). For more on _Kind_, see the docs here https://kind.sigs.k8s.io/.

The pod.yaml file is the only file needed.

1 - First _kind_ is used to spin up a k8's cluster, _minikube_ or _docker desktop_ can also be used here:

	kind create cluster
  
The created cluster is verified using the following cmds:
	
	kind get clusters 

which return the name of the cluster (kind)

	kubectl get cluster-info 

returns the k8's master and kubeDNS info
  
	kubectl config get-contexts 

returns inactive and active contexts (kind-kind)
  
	kubectl get nodes 

returns the kube-control-plane
  
	kubectl get all -A 
    
returns all the nodes spun up by _kind_ to launch the k8's cluster

2 - Second, k8's is used with the yaml file to launch a pod:
    
	kubectl apply -f pod.yaml 
		
which returns the name of the pod created
    
	kubectl get pods 
			
which lists the new pod and its status

3 - Clean up:

Delete pods
   
	 kubectl get pods

then   
	 
	 kubeclt delete pod building-apps-pod

or 
	 
	 kubectl delete -f pod.yaml

Delete cluster
   
	 kind get clusters
	 
then

	 kind delete cluster
