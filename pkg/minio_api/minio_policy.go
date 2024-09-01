package minio_api

import "github.com/Juminiy/kube/pkg/s3_api"

const (
	Version = s3_api.Version
)

const (
	Allow = s3_api.Allow
	Deny  = s3_api.Deny
)

// referred from:
// https://min.io/docs/minio/linux/administration/identity-access-management/policy-based-access-control.html#id3
// Minio Built-In Policies
// assigning to users or groups

// consoleAdmin
const (
	S3All    = s3_api.S3All
	adminAll = AdminAll
)

// readonly
const (
	S3GetBucketLocation = s3_api.S3GetBucketLocation
	S3GetObject         = s3_api.S3GetObject
)

// readwrite
const (
	S3RW = S3All
	s3RW
)

// diagnostics
const (
	adminServerTrace      = AdminServerTrace
	adminProfiling        = AdminProfiling
	adminConsoleLog       = AdminConsoleLog
	adminServerInfo       = AdminServerInfo
	adminTopLocksInfo     = AdminTopLocksInfo
	adminOBDInfo          = AdminOBDInfo
	adminBandwidthMonitor = AdminBandwidthMonitor
	adminPrometheus       = AdminPrometheus
)

// writeonly
const (
	S3PutObject = s3_api.S3PutObject
	s3PutObject
)

// referred from:
// https://min.io/docs/minio/linux/administration/identity-access-management/policy-based-access-control.html#id7
// minio admin Policy Action Keys
// minio admin API exclusive, not compatible with S3 API
const (
	AdminAll                     = "admin:*"
	AdminHeal                    = "admin:Heal"
	AdminStorageInfo             = "admin:StorageInfo"
	AdminDataUsageInfo           = "admin:DataUsageInfo"
	AdminTopLocksInfo            = "admin:TopLocksInfo"
	AdminProfiling               = "admin:Profiling"
	AdminServerTrace             = "admin:ServerTrace"
	AdminConsoleLog              = "admin:ConsoleLog"
	AdminKMSCreateKey            = "admin:KMSCreateKey"
	AdminKMSKeyStatus            = "admin:KMSKeyStatus"
	AdminServerInfo              = "admin:ServerInfo"
	AdminOBDInfo                 = "admin:OBDInfo"
	AdminServerUpdate            = "admin:ServerUpdate"
	AdminServiceRestart          = "admin:ServiceRestart"
	AdminServiceStop             = "admin:ServiceStop"
	AdminConfigUpdate            = "admin:ConfigUpdate"
	AdminCreateUser              = "admin:CreateUser"
	AdminDeleteUser              = "admin:DeleteUser"
	AdminListUsers               = "admin:ListUsers"
	AdminEnableUser              = "admin:EnableUser"
	AdminDisableUser             = "admin:DisableUser"
	AdminGetUser                 = "admin:GetUser"
	AdminAddUserToGroup          = "admin:AddUserToGroup"
	AdminRemoveUserFromGroup     = "admin:RemoveUserFromGroup"
	AdminGetGroup                = "admin:GetGroup"
	AdminListGroups              = "admin:ListGroups"
	AdminEnableGroup             = "admin:EnableGroup"
	AdminDisableGroup            = "admin:DisableGroup"
	AdminCreatePolicy            = "admin:CreatePolicy"
	AdminDeletePolicy            = "admin:DeletePolicy"
	AdminGetPolicy               = "admin:GetPolicy"
	AdminAttachUserOrGroupPolicy = "admin:AttachUserOrGroupPolicy"
	AdminListUserPolicies        = "admin:ListUserPolicies"
	AdminCreateServiceAccount    = "admin:CreateServiceAccount"
	AdminUpdateServiceAccount    = "admin:UpdateServiceAccount"
	AdminRemoveServiceAccount    = "admin:RemoveServiceAccount"
	AdminListServiceAccounts     = "admin:ListServiceAccounts"
	AdminSetBucketQuota          = "admin:SetBucketQuota"
	AdminGetBucketQuota          = "admin:GetBucketQuota"
	AdminSetBucketTarget         = "admin:SetBucketTarget"
	AdminGetBucketTarget         = "admin:GetBucketTarget"
	AdminSetTier                 = "admin:SetTier"
	AdminListTier                = "admin:ListTier"
	AdminBandwidthMonitor        = "admin:BandwidthMonitor"
	AdminPrometheus              = "admin:Prometheus"
	AdminListBatchJobs           = "admin:ListBatchJobs"
	AdminDescribeBatchJobs       = "admin:DescribeBatchJobs"
	AdminStartBatchJob           = "admin:StartBatchJob"
	AdminCancelBatchJob          = "admin:CancelBatchJob"
	AdminRebalance               = "admin:Rebalance"
)

// referred from:
// https://min.io/docs/minio/linux/administration/identity-access-management/policy-based-access-control.html#id8
// minio admin Policy Condition Keys
// compatible with S3 API
const (
	AWSReferer         = s3_api.AWSReferer
	AWSSourceIp        = s3_api.AWSSourceIp
	AWSUserAgent       = s3_api.AWSUserAgent
	AWSSecureTransport = s3_api.AWSSecureTransport
	AWSCurrentTime     = s3_api.AWSCurrentTime
	AWSEpochTime       = s3_api.AWSEpochTime
)

const (
	KMSAll = "kms:*"
	kMSAll
)

const (
	ResourceARNS3All = s3_api.ResourceARNS3All
)
