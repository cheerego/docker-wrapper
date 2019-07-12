# Intruduce
是否因为 Pull Docker Image 而烦恼。
万恶的 `k8s.gcr.io` `gcr.io` `quay.io`

# Requirement
docker in your system and docker command in your environment

# Usage
```bash
# k8s.gcr.io
docker-wrapper pull k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1

# gcr.io
docker-wrapper pull gcr.io/jenkinsxio/prow/crier:v20190510-3f2c8d788

# quay.io
docker-wrapper pull quay.io/coreos/etcd:v3.3

# normal
docker-wrapper pull nginx
```

# Install

## Source build
```
git clone xxx  
cd xxx
make
```