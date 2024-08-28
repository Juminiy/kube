# kubectl exec
kubectl exec -it -n default $(kubectl get pods -n default | grep "ubuntu2204" | tail -n 1 | awk -F'   ' '{print $1}') -- /bin/bash

# ssh exec : remote
ssh -p 30022 root@192.168.31.19