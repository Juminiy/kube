package internal

import (
	s3apiv2 "github.com/Juminiy/kube/pkg/storage_api/s3_api/v2"
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

func ConsoleAdminPolicy() string {
	consoleAdminPolicy, _ := consoleAdmin.String()
	return consoleAdminPolicy
}

func DiagnosticsPolicy() string {
	diagnosticsPolicy, _ := diagnostics.String()
	return diagnosticsPolicy
}

func ReadOnlyPolicy() string {
	readOnlyPolicy, _ := readOnly.String()
	return readOnlyPolicy
}

func ReadWritePolicy() string {
	readWritePolicy, _ := readWrite.String()
	return readWritePolicy
}

func WriteOnlyPolicy() string {
	writeOnlyPolicy, _ := writeOnly.String()
	return writeOnlyPolicy
}
