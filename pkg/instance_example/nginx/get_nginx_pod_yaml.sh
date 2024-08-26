#!/bin/bash

action_arg=$1
pod_name="nginx-pod-example"

function kube_edit() {
  kubectl edit deployment $pod_name -n default
}

function list_pod_yaml() {
  pod_list=$(kubectl get pod -n default | grep $pod_name | awk -F'   ' '{print $1}')

  for pod_name0 in $pod_list
    do
      {
        kubectl get pods -n default "$pod_name0" -o yaml > "$pod_name0".yaml
      }&
  done
  wait
}

if [ "$action_arg" = "list_yaml" ]; then
  list_pod_yaml
elif [ "$action_arg" = "edit" ]; then
  kube_edit
else
  echo "error fmt arg: edit | list_yaml"
fi
