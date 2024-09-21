1. Desktop or Server 
   1. unix
   
      ```bash
      echo 192.168.31.19 lb.kubesphere.local >> /etc/hosts
      echo 192.168.31.242 docker.local >> /etc/hosts
      echo 192.168.31.242 harbor.local >> /etc/hosts
      echo 192.168.31.110 minio.local >> /etc/hosts
      ```
   2. windows

       C:\\Windows\\System32\\drivers\\etc\\hosts
      ```shell
      192.168.31.19 lb.kubesphere.local
      192.168.31.242 docker.local
      192.168.31.242 harbor.local
      192.168.31.110 minio.local
      ```
