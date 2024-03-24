# Improving Development Workflow with Skaffold
Skaffold builds k8s manifest, detects changes in manifest and automatically updates the deployments. It provides a tight feedback loop.

## 1 - Install Skaffold
The command below is for Linux x86_64, for other systems, see the [docs](https://skaffold.dev/docs/install/).

```
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
sudo install skaffold /usr/local/bin/
```

## 2 - Setup

```sh
skaffold init
```
View the created `skaffold.yaml` file.

 - Skaffold needs a cluster to run the app on, so start a cluster with minikube or kind (steps in this project inside the [../2-using-Kind](https://github.com/Fabr1ce/building-apps-for-k8s/tree/main/2-using-Kind) directory).

 - Before running skaffold, verify which cluster `kubectl` is configured with.

```sh
kubectl config current-context
```

## Start Skaffold
Start skaffold in development mode.

```sh
skaffold dev
```

Verify Deployment created Pods.

```sh
kubectl get pods
```

Update `src/main.go` to reflect a new version number.

```go
const version = "0.2"`
```

View logs to verify the new version has been deployed.

How do we verify the change by accessing the API?

Stop `skaffold` and rerun with the `--port-forward` flag.

```sh
skaffold dev --port-forward
```

Notice the port that gets reported in the logs.

```sh
# NOTE: Your port may differ.
curl 127.0.0.1:4503
```

What happens when we update the Kubernetes manifests?
Update the `k8s/deployment.yaml` manifest to specify 5 replicas.

```yaml
spec:
  replicas: 5
```

Note that in the skaffold logs, we are not rebuilding our image.
Because our Kubernetes manifests live outside of our configured Docker context, skaffold knows that it does not need to rebuild our images.

