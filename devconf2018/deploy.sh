#!/bin/bash

SUPER_SECRET_TOKEN=abcdef.1234567890123456
ADVERTISED_MASTER_IP=10.96.0.15 # can be a k8s service IP

# Install docker and kubeadm
dnf install -y docker kubernetes-kubeadm kubernetes-client

# Start the kubelet and docker
systemctl enable docker && systemctl start docker
systemctl enable kubelet && systemctl start kubelet

# Let kubeadm do the final bootrapping and the registration
kubeadm join --token $SUPER_SECRET_TOKEN $ADVERTISED_MASTER_IP:6443 \
    --ignore-preflight-errors=all --discovery-token-unsafe-skip-ca-verification=true
