# Introduction

> 在K8S的使用过程中，总是有一些镜像拉取不下来。

> 灵感来源于 https://ieevee.com/tech/2019/03/02/azure-gcr-proxy.html 

迫于拉取镜像很难受。

频繁重复的`docker pull` `docker tag` 很麻烦。

编写shell每次都要改。

所以编写了这个工具。

万恶的 `k8s.gcr.io` `gcr.io` `quay.io` .

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

## Download
[release](https://github.com/cheerego/docker-wrapper/releases)

or

go get github.com/cheerego/docker-wrapper

## Source build

```
git clone https://github.com/cheerego/docker-wrapper.git  
cd docker-wrapper
go get 
make
```

## 我的另一个小工具，方便你切换 `kubeconfig` 
https://github.com/cheerego/kube-switch

