# kube/pkg
> kubernetes addition util package

1. image_api
   - docker_api
   - harbor_api

2. k8s_api
   - test

3. log_api
   - stdlog
   - zaplog

4. storage_api
   - minio_api
   - s3_api
   - db
     - kv: key-value
     - rdb: relational

5. util
   1. safe_cast
   2. safe_go
   3. safe_json
   4. safe_reflect
   5. safe_validator
   6. .


| Service \\ Connect | Method                             | Path \| Loc         | Go Global Client | Reason           |
|--------------------|------------------------------------|---------------------|------------------|------------------|
| Kubernetes         | HTTP Host Config File              | REST API            | ✅                | InCluster        |
| Harbor             | HTTP Credential: username&password | Console & REST API  | ✅ \| Both        | Maybe Migrate    |
| Docker             | HTTP No Credential                 | Console & REST API  | ✅ \| Both        | Maybe Migrate    |
| Minio              | HTTP Credential: username&password | Console & REST API  | ✅ \| Both        | Maybe Migrate    |
| StdLog             | App Config File                    |                     | ✅                | Only Local files |
| ZapLog             | App Config File                    |                     | ✅                | Only Local files |