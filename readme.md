# Intruduce
是否因为 Pull Docker Image 而烦恼。
万恶的 `k8s.gcr.io` `gcr.io` `quay.io`

# Requirement
docker in your system and docker command in your envirment

# Usage
```bash
docker-wrapper pull quay.io/coreos/etcd:v3.3
docker-wrapper pull k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1
docker-wrapper pull gcr.io/jenkinsxio/prow/crier:v20190510-3f2c8d788
```

# Install

## Source build
```
git clone xxx  
cd xxx
make
```