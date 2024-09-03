# Kube pkg

1. image_api
   1. harbor_api
   2. docker_api

2. k8s_api
   1. index_api
   2. pod_api
   3. volume_api
   4. service_api
   5. instance_api

3. log_api
   1. stdlog
   2. zerolog
   3. zaplog

4. storage_api
   1. minio_api
   2. s3_api
   3. mq_api

5. util_api
   1. util/xxx_api 
   2. more refer to k8s.io/apimachinery/pkg/util 
   3. and more refer to xxx.xxx/xxx/pkg/../../util/../...



| Service \ Connect | Method                        | Path \| Loc                       | Go Global Singleton Type | Why                              |
| ----------------- | ----------------------------- | --------------------------------- | ------------------------ | -------------------------------- |
| K8s               | Host Config File              | root@k8s_master_ip:~/.kube/config | ✅                        | Self High Available              |
| Harbor            | Credential: username&pasword  | Harbor console                    | ✅ \| Both                | Maybe migrate                    |
| Docker            |                               |                                   | ✅ \| Both                | Maybe migrate                    |
| Minio             | Credential: username&password | Minio console                     | ✅ \| Both                | Maybe migrate                    |
| NFS               |                               |                                   |                          |                                  |
| zaplog            | App config file               |                                   | ✅                        | not a resource, only local files |
| zerolog           | App config file               |                                   | ✅                        | not a resource, only local files |
|                   |                               |                                   |                          |                                  |
|                   |                               |                                   |                          |                                  |

