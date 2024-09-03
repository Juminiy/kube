package minio_api

/*type AdminStatement struct {
	s3_api.SASRStatement
}

type AccessKeyWithOneBucketRWStatement struct {
	s3_api.MASRStatement
}

type BucketStatement struct {
	s3_api.SAMRStatement
}

// WARNING: not api expose, only test for internal use
func AdminPolicy(policyName ...string) string {
	p := s3_api.Policy{
		Version: Version,
		StatementList: s3_api.StatementList{
			&s3_api.AdminStatement{
				SASRStatement: s3_api.SASRStatement{
					Sid:       s3_api.GetSid(policyName...),
					Effect:    Allow,
					Action:    S3All,
					Resource:  s3_api.ResourceAll,
					Condition: nil,
				},
			},
		},
	}
	return p.String()
}

func AccessKeyWithBucketRWPolicy(bucketName string, policyName ...string) string {
	p := s3_api.Policy{
		Version: Version,
		StatementList: s3_api.StatementList{
			&s3_api.AccessKeyWithOneBucketRWStatement{
				MASRStatement: s3_api.MASRStatement{
					Sid:    s3_api.GetSid(policyName...),
					Effect: Allow,
					Action: []string{
						s3_api.S3DeleteObject,
						S3GetObject,
						s3_api.S3ListBucket,
						S3PutObject,
					},
					Resource: s3_api.GetBucketResource(bucketName),
				},
			},
		},
	}
	return p.String()
}

func BucketRWPolicy(bucketName string, policyName ...string) string {
	p := s3_api.Policy{
		Version: Version,
		StatementList: s3_api.StatementList{
			&s3_api.BucketStatement{
				s3_api.SAMRStatement{
					Sid:    s3_api.GetSid(policyName...),
					Effect: Allow,
					Action: S3All,
					Resource: []string{
						s3_api.GetBucketResource(bucketName),
						s3_api.GetBucketAnyResource(bucketName),
					},
				},
			},
		},
	}
	return p.String()
}
*/
