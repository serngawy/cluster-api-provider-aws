#!/bin/bash

version="1.14.5"
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-03-cilium-ciliumconfigs-crd.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00000-cilium-namespace.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00001-cilium-olm-serviceaccount.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00002-cilium-olm-deployment.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00003-cilium-olm-service.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00004-cilium-olm-leader-election-role.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00005-cilium-olm-role.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00006-leader-election-rolebinding.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00007-cilium-olm-rolebinding.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00008-cilium-cilium-olm-clusterrole.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00009-cilium-cilium-clusterrole.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00010-cilium-cilium-olm-clusterrolebinding.yaml
oc --kubeconfig ./kubeconfigRosa apply -f https://raw.githubusercontent.com/isovalent/olm-for-cilium/main/manifests/cilium.v${version}/cluster-network-06-cilium-00011-cilium-cilium-clusterrolebinding.yaml
oc --kubeconfig ./kubeconfigRosa apply -f ./ciliumConfig.yaml

