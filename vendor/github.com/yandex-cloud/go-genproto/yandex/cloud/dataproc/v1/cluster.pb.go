// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: yandex/cloud/dataproc/v1/cluster.proto

package dataproc

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Cluster_Status int32

const (
	// Cluster state is unknown.
	Cluster_STATUS_UNKNOWN Cluster_Status = 0
	// Cluster is being created.
	Cluster_CREATING Cluster_Status = 1
	// Cluster is running normally.
	Cluster_RUNNING Cluster_Status = 2
	// Cluster encountered a problem and cannot operate.
	Cluster_ERROR Cluster_Status = 3
	// Cluster is stopping.
	Cluster_STOPPING Cluster_Status = 4
	// Cluster stopped.
	Cluster_STOPPED Cluster_Status = 5
	// Cluster is starting.
	Cluster_STARTING Cluster_Status = 6
)

// Enum value maps for Cluster_Status.
var (
	Cluster_Status_name = map[int32]string{
		0: "STATUS_UNKNOWN",
		1: "CREATING",
		2: "RUNNING",
		3: "ERROR",
		4: "STOPPING",
		5: "STOPPED",
		6: "STARTING",
	}
	Cluster_Status_value = map[string]int32{
		"STATUS_UNKNOWN": 0,
		"CREATING":       1,
		"RUNNING":        2,
		"ERROR":          3,
		"STOPPING":       4,
		"STOPPED":        5,
		"STARTING":       6,
	}
)

func (x Cluster_Status) Enum() *Cluster_Status {
	p := new(Cluster_Status)
	*p = x
	return p
}

func (x Cluster_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_dataproc_v1_cluster_proto_enumTypes[0].Descriptor()
}

func (Cluster_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_dataproc_v1_cluster_proto_enumTypes[0]
}

func (x Cluster_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_Status.Descriptor instead.
func (Cluster_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_cluster_proto_rawDescGZIP(), []int{0, 0}
}

type HadoopConfig_Service int32

const (
	HadoopConfig_SERVICE_UNSPECIFIED HadoopConfig_Service = 0
	HadoopConfig_HDFS                HadoopConfig_Service = 1
	HadoopConfig_YARN                HadoopConfig_Service = 2
	HadoopConfig_MAPREDUCE           HadoopConfig_Service = 3
	HadoopConfig_HIVE                HadoopConfig_Service = 4
	HadoopConfig_TEZ                 HadoopConfig_Service = 5
	HadoopConfig_ZOOKEEPER           HadoopConfig_Service = 6
	HadoopConfig_HBASE               HadoopConfig_Service = 7
	HadoopConfig_SQOOP               HadoopConfig_Service = 8
	HadoopConfig_FLUME               HadoopConfig_Service = 9
	HadoopConfig_SPARK               HadoopConfig_Service = 10
	HadoopConfig_ZEPPELIN            HadoopConfig_Service = 11
	HadoopConfig_OOZIE               HadoopConfig_Service = 12
	HadoopConfig_LIVY                HadoopConfig_Service = 13
)

// Enum value maps for HadoopConfig_Service.
var (
	HadoopConfig_Service_name = map[int32]string{
		0:  "SERVICE_UNSPECIFIED",
		1:  "HDFS",
		2:  "YARN",
		3:  "MAPREDUCE",
		4:  "HIVE",
		5:  "TEZ",
		6:  "ZOOKEEPER",
		7:  "HBASE",
		8:  "SQOOP",
		9:  "FLUME",
		10: "SPARK",
		11: "ZEPPELIN",
		12: "OOZIE",
		13: "LIVY",
	}
	HadoopConfig_Service_value = map[string]int32{
		"SERVICE_UNSPECIFIED": 0,
		"HDFS":                1,
		"YARN":                2,
		"MAPREDUCE":           3,
		"HIVE":                4,
		"TEZ":                 5,
		"ZOOKEEPER":           6,
		"HBASE":               7,
		"SQOOP":               8,
		"FLUME":               9,
		"SPARK":               10,
		"ZEPPELIN":            11,
		"OOZIE":               12,
		"LIVY":                13,
	}
)

func (x HadoopConfig_Service) Enum() *HadoopConfig_Service {
	p := new(HadoopConfig_Service)
	*p = x
	return p
}

func (x HadoopConfig_Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HadoopConfig_Service) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_dataproc_v1_cluster_proto_enumTypes[1].Descriptor()
}

func (HadoopConfig_Service) Type() protoreflect.EnumType {
	return &file_yandex_cloud_dataproc_v1_cluster_proto_enumTypes[1]
}

func (x HadoopConfig_Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HadoopConfig_Service.Descriptor instead.
func (HadoopConfig_Service) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_cluster_proto_rawDescGZIP(), []int{2, 0}
}

// A Data Proc cluster. For details about the concept, see [documentation](/docs/data-proc/concepts/).
type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the cluster. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the cluster belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the cluster. The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the cluster.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Cluster labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Monitoring systems relevant to the cluster.
	Monitoring []*Monitoring `protobuf:"bytes,7,rep,name=monitoring,proto3" json:"monitoring,omitempty"`
	// Configuration of the cluster.
	Config *ClusterConfig `protobuf:"bytes,8,opt,name=config,proto3" json:"config,omitempty"`
	// Aggregated cluster health.
	Health Health `protobuf:"varint,9,opt,name=health,proto3,enum=yandex.cloud.dataproc.v1.Health" json:"health,omitempty"`
	// Cluster status.
	Status Cluster_Status `protobuf:"varint,10,opt,name=status,proto3,enum=yandex.cloud.dataproc.v1.Cluster_Status" json:"status,omitempty"`
	// ID of the availability zone where the cluster resides.
	ZoneId string `protobuf:"bytes,11,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// ID of service account for the Data Proc manager agent.
	ServiceAccountId string `protobuf:"bytes,12,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Object Storage bucket to be used for Data Proc jobs that are run in the cluster.
	Bucket string `protobuf:"bytes,13,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Whether UI Proxy feature is enabled.
	UiProxy bool `protobuf:"varint,14,opt,name=ui_proxy,json=uiProxy,proto3" json:"ui_proxy,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cluster) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Cluster) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Cluster) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Cluster) GetMonitoring() []*Monitoring {
	if x != nil {
		return x.Monitoring
	}
	return nil
}

func (x *Cluster) GetConfig() *ClusterConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Cluster) GetHealth() Health {
	if x != nil {
		return x.Health
	}
	return Health_HEALTH_UNKNOWN
}

func (x *Cluster) GetStatus() Cluster_Status {
	if x != nil {
		return x.Status
	}
	return Cluster_STATUS_UNKNOWN
}

func (x *Cluster) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *Cluster) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *Cluster) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Cluster) GetUiProxy() bool {
	if x != nil {
		return x.UiProxy
	}
	return false
}

// Metadata of a monitoring system for a Data Proc cluster.
type Monitoring struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the monitoring system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the monitoring system.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Link to the monitoring system.
	Link string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *Monitoring) Reset() {
	*x = Monitoring{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Monitoring) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitoring) ProtoMessage() {}

func (x *Monitoring) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitoring.ProtoReflect.Descriptor instead.
func (*Monitoring) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *Monitoring) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Monitoring) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Monitoring) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

// Hadoop configuration that describes services installed in a cluster,
// their properties and settings.
type HadoopConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Set of services used in the cluster (if empty, the default set is used).
	Services []HadoopConfig_Service `protobuf:"varint,1,rep,packed,name=services,proto3,enum=yandex.cloud.dataproc.v1.HadoopConfig_Service" json:"services,omitempty"`
	// Properties set for all hosts in `*-site.xml` configurations. The key should indicate
	// the service and the property.
	//
	// For example, use the key 'hdfs:dfs.replication' to set the `dfs.replication` property
	// in the file `/etc/hadoop/conf/hdfs-site.xml`.
	Properties map[string]string `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// List of public SSH keys to access to cluster hosts.
	SshPublicKeys []string `protobuf:"bytes,3,rep,name=ssh_public_keys,json=sshPublicKeys,proto3" json:"ssh_public_keys,omitempty"`
}

func (x *HadoopConfig) Reset() {
	*x = HadoopConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HadoopConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HadoopConfig) ProtoMessage() {}

func (x *HadoopConfig) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HadoopConfig.ProtoReflect.Descriptor instead.
func (*HadoopConfig) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *HadoopConfig) GetServices() []HadoopConfig_Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *HadoopConfig) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *HadoopConfig) GetSshPublicKeys() []string {
	if x != nil {
		return x.SshPublicKeys
	}
	return nil
}

type ClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Image version for cluster provisioning.
	// All available versions are listed in the [documentation](/docs/managed-hadoop/concepts/image-versions).
	VersionId string `protobuf:"bytes,1,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Data Proc specific configuration options.
	Hadoop *HadoopConfig `protobuf:"bytes,2,opt,name=hadoop,proto3" json:"hadoop,omitempty"`
}

func (x *ClusterConfig) Reset() {
	*x = ClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfig) ProtoMessage() {}

func (x *ClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterConfig.ProtoReflect.Descriptor instead.
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterConfig) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *ClusterConfig) GetHadoop() *HadoopConfig {
	if x != nil {
		return x.Hadoop
	}
	return nil
}

var File_yandex_cloud_dataproc_v1_cluster_proto protoreflect.FileDescriptor

var file_yandex_cloud_dataproc_v1_cluster_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x06, 0x0a, 0x07, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31,
	0x04, 0x31, 0x2d, 0x36, 0x33, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x30, 0x2d, 0x32, 0x35, 0x36, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0x82, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x36,
	0x34, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x44, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x0a, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x3f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x7a,
	0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x69, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75,
	0x69, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x6b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x05,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x22, 0x56,
	0x0a, 0x0a, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xd2, 0x03, 0x0a, 0x0c, 0x48, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4a, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x73,
	0x73, 0x68, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xb6, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17,
	0x0a, 0x13, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x44, 0x46, 0x53, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x59, 0x41, 0x52, 0x4e, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4d,
	0x41, 0x50, 0x52, 0x45, 0x44, 0x55, 0x43, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x49,
	0x56, 0x45, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x45, 0x5a, 0x10, 0x05, 0x12, 0x0d, 0x0a,
	0x09, 0x5a, 0x4f, 0x4f, 0x4b, 0x45, 0x45, 0x50, 0x45, 0x52, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05,
	0x48, 0x42, 0x41, 0x53, 0x45, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x51, 0x4f, 0x4f, 0x50,
	0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x55, 0x4d, 0x45, 0x10, 0x09, 0x12, 0x09, 0x0a,
	0x05, 0x53, 0x50, 0x41, 0x52, 0x4b, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x5a, 0x45, 0x50, 0x50,
	0x45, 0x4c, 0x49, 0x4e, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x4f, 0x5a, 0x49, 0x45, 0x10,
	0x0c, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x56, 0x59, 0x10, 0x0d, 0x22, 0x6e, 0x0a, 0x0d, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x06, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x42, 0x65, 0x0a, 0x1c, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x31, 0x5a, 0x45, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_dataproc_v1_cluster_proto_rawDescOnce sync.Once
	file_yandex_cloud_dataproc_v1_cluster_proto_rawDescData = file_yandex_cloud_dataproc_v1_cluster_proto_rawDesc
)

func file_yandex_cloud_dataproc_v1_cluster_proto_rawDescGZIP() []byte {
	file_yandex_cloud_dataproc_v1_cluster_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_dataproc_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_dataproc_v1_cluster_proto_rawDescData)
	})
	return file_yandex_cloud_dataproc_v1_cluster_proto_rawDescData
}

var file_yandex_cloud_dataproc_v1_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_yandex_cloud_dataproc_v1_cluster_proto_goTypes = []interface{}{
	(Cluster_Status)(0),         // 0: yandex.cloud.dataproc.v1.Cluster.Status
	(HadoopConfig_Service)(0),   // 1: yandex.cloud.dataproc.v1.HadoopConfig.Service
	(*Cluster)(nil),             // 2: yandex.cloud.dataproc.v1.Cluster
	(*Monitoring)(nil),          // 3: yandex.cloud.dataproc.v1.Monitoring
	(*HadoopConfig)(nil),        // 4: yandex.cloud.dataproc.v1.HadoopConfig
	(*ClusterConfig)(nil),       // 5: yandex.cloud.dataproc.v1.ClusterConfig
	nil,                         // 6: yandex.cloud.dataproc.v1.Cluster.LabelsEntry
	nil,                         // 7: yandex.cloud.dataproc.v1.HadoopConfig.PropertiesEntry
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(Health)(0),                 // 9: yandex.cloud.dataproc.v1.Health
}
var file_yandex_cloud_dataproc_v1_cluster_proto_depIdxs = []int32{
	8, // 0: yandex.cloud.dataproc.v1.Cluster.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: yandex.cloud.dataproc.v1.Cluster.labels:type_name -> yandex.cloud.dataproc.v1.Cluster.LabelsEntry
	3, // 2: yandex.cloud.dataproc.v1.Cluster.monitoring:type_name -> yandex.cloud.dataproc.v1.Monitoring
	5, // 3: yandex.cloud.dataproc.v1.Cluster.config:type_name -> yandex.cloud.dataproc.v1.ClusterConfig
	9, // 4: yandex.cloud.dataproc.v1.Cluster.health:type_name -> yandex.cloud.dataproc.v1.Health
	0, // 5: yandex.cloud.dataproc.v1.Cluster.status:type_name -> yandex.cloud.dataproc.v1.Cluster.Status
	1, // 6: yandex.cloud.dataproc.v1.HadoopConfig.services:type_name -> yandex.cloud.dataproc.v1.HadoopConfig.Service
	7, // 7: yandex.cloud.dataproc.v1.HadoopConfig.properties:type_name -> yandex.cloud.dataproc.v1.HadoopConfig.PropertiesEntry
	4, // 8: yandex.cloud.dataproc.v1.ClusterConfig.hadoop:type_name -> yandex.cloud.dataproc.v1.HadoopConfig
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_dataproc_v1_cluster_proto_init() }
func file_yandex_cloud_dataproc_v1_cluster_proto_init() {
	if File_yandex_cloud_dataproc_v1_cluster_proto != nil {
		return
	}
	file_yandex_cloud_dataproc_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Monitoring); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HadoopConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_dataproc_v1_cluster_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_dataproc_v1_cluster_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_dataproc_v1_cluster_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_dataproc_v1_cluster_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_dataproc_v1_cluster_proto_msgTypes,
	}.Build()
	File_yandex_cloud_dataproc_v1_cluster_proto = out.File
	file_yandex_cloud_dataproc_v1_cluster_proto_rawDesc = nil
	file_yandex_cloud_dataproc_v1_cluster_proto_goTypes = nil
	file_yandex_cloud_dataproc_v1_cluster_proto_depIdxs = nil
}
