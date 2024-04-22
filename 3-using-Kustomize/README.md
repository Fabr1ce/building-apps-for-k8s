# 3-using-Kustomize

This repo is used to practice k8s deployments and service manifests using Kustomize as outlined in this [KubeAcademy course](https://kube.academy/courses/building-applications-for-kubernetes/lessons/deploying-your-application). Concepts covered here are also explained in the [k8s docs](https://kubernetes.io/docs/concepts/workloads/):

1. pods (groud of one or more containers with single IP)

2. replicaSet (way to deploy multiple identical pods at once)

3. deployment (manages replicasets plus rolling updates)

4. service (provides a stable network address for apps including local IP that allows pods to be destroyed while the remaining pods have this IP where traffic is directed, DNS within the cluster, it load balances traffic accross pods)

5. [kustomize](https://kustomize.io/) manages k8s apps without the need of templates.

## Prerequisites
1. Install Kubernetes
2. Install Kustomize

## Steps to create a k8s deployment and service:
1. Create a cluster: refer to the section [2-using-Kind](https://github.com/Fabr1ce/building-apps-for-k8s/tree/main/2-using-Kind) for steps on how to create a kubernetes cluster.

2. Write YAML files to pass to k8s or use the ones in this repo and use the following cmds:

- **Without using Kustomize:**
   1. Create the deployment and verify:
  
     ```
     kubectl apply -f deployment.yaml
     ```
     
     ```
     kubectl get pods
     ```
  
   2. Create the service and verify:
  
     ```
     kubectl apply -f service.yaml
     ```
     
     ```
     kubectl get services
     ```
  
  - service IP can be accessed through cURL like this `exec <outside-pod-name> curl <cluster-ip>`
  - pods can be accessed through cURL like this `exec <outside-pod-name> curl <pod-ip>:<port-in-manifest>`. Notice the difference between Port (port service listens on and forwards to NodePort) and NodePort (port pods listens on).
  - clean up:
  
     ```
     kubectl delete -f service.yaml
     ```
     
     ```
     kubectl delete -f deployment.yaml
     ```

- **Using Kustomize:**
  Kustomize takes all the yaml files and creates all the components/objects/environments using the following cmds:
  
     ```
     kustomize build base
     ```
     
     ```
     kustomize build overlays/production
     ```
     
     ```
     kustomize build overlays/production | kubectl apply -f -
     ```

In summary, Kustomize took the base resource manifests and patched them with overlays.
