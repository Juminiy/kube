package minio_api

import (
	s3apiv2 "github.com/Juminiy/kube/pkg/s3_api/v2"
)

var consoleAdmin = s3apiv2.IBAPolicy{
	Version: Version,
	Statement: []s3apiv2.Statement{
		s3apiv2.Statement{
			Effect: Allow,
			Action: []string{AdminAll},
		},
		s3apiv2.Statement{
			Effect: Allow,
			Action: []string{KMSAll},
		},
		s3apiv2.Statement{
			Effect:   Allow,
			Action:   []string{S3All},
			Resource: []string{ResourceARNS3All},
		},
	},
}

var diagnostics = s3apiv2.IBAPolicy{
	Version: Version,
	Statement: []s3apiv2.Statement{
		s3apiv2.Statement{
			Effect: Allow,
			Action: []string{
				AdminOBDInfo,
				AdminProfiling,
				AdminPrometheus,
				AdminServerInfo,
				AdminServerTrace,
				AdminTopLocksInfo,
				AdminBandwidthMonitor,
				AdminConsoleLog,
			},
			Resource: []string{ResourceARNS3All},
		},
	},
}

var readOnly = s3apiv2.IBAPolicy{
	Version: Version,
	Statement: []s3apiv2.Statement{
		s3apiv2.Statement{
			Effect: Allow,
			Action: []string{
				S3GetBucketLocation,
				S3GetObject,
			},
			Resource: []string{ResourceARNS3All},
		},
	},
}

var readWrite = s3apiv2.IBAPolicy{
	Version: Version,
	Statement: []s3apiv2.Statement{
		s3apiv2.Statement{
			Effect:   Allow,
			Action:   []string{S3All},
			Resource: []string{ResourceARNS3All},
		},
	},
}

var writeOnly = s3apiv2.IBAPolicy{
	Version: Version,
	Statement: []s3apiv2.Statement{
		s3apiv2.Statement{
			Effect:   Allow,
			Action:   []string{S3PutObject},
			Resource: []string{ResourceARNS3All},
		},
	},
}
