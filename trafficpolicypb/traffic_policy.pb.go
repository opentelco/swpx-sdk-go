// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v5.26.1
// source: traffic_policy.proto

package trafficpolicypb

import (
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

// display current-configuration interface GigabitEthernet0/0/5 | include traffic-policy|shaping
// traffic-policy KBPS_0000 inbound
// traffic-policy KBPS_100000 outbound
// qos queue 0 shaping cir 110000 pir 110000 cbs 2500000 pbs 2500000
type ConfiguredTrafficPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inbound  string                       `protobuf:"bytes,1,opt,name=inbound,proto3" json:"inbound,omitempty" bson:"inbound"`
	Outbound string                       `protobuf:"bytes,2,opt,name=outbound,proto3" json:"outbound,omitempty" bson:"outbound"`
	Qos      *ConfiguredTrafficPolicy_QOS `protobuf:"bytes,3,opt,name=qos,proto3" json:"qos,omitempty" bson:"qos"`
	// display traffic policy statistics interface GigabitEthernet0/0/5 inbound verbose classifier-base
	InboundStatistics *ConfiguredTrafficPolicy_Statistics `protobuf:"bytes,4,opt,name=inbound_statistics,json=inboundStatistics,proto3" json:"inbound_statistics,omitempty" bson:"inbound_statistics"`
}

func (x *ConfiguredTrafficPolicy) Reset() {
	*x = ConfiguredTrafficPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfiguredTrafficPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfiguredTrafficPolicy) ProtoMessage() {}

func (x *ConfiguredTrafficPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfiguredTrafficPolicy.ProtoReflect.Descriptor instead.
func (*ConfiguredTrafficPolicy) Descriptor() ([]byte, []int) {
	return file_traffic_policy_proto_rawDescGZIP(), []int{0}
}

func (x *ConfiguredTrafficPolicy) GetInbound() string {
	if x != nil {
		return x.Inbound
	}
	return ""
}

func (x *ConfiguredTrafficPolicy) GetOutbound() string {
	if x != nil {
		return x.Outbound
	}
	return ""
}

func (x *ConfiguredTrafficPolicy) GetQos() *ConfiguredTrafficPolicy_QOS {
	if x != nil {
		return x.Qos
	}
	return nil
}

func (x *ConfiguredTrafficPolicy) GetInboundStatistics() *ConfiguredTrafficPolicy_Statistics {
	if x != nil {
		return x.InboundStatistics
	}
	return nil
}

type QOS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueStatistics []*QOS_QueueStatistics `protobuf:"bytes,1,rep,name=queue_statistics,json=queueStatistics,proto3" json:"queue_statistics,omitempty" bson:"queue_statistics"`
}

func (x *QOS) Reset() {
	*x = QOS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QOS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QOS) ProtoMessage() {}

func (x *QOS) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QOS.ProtoReflect.Descriptor instead.
func (*QOS) Descriptor() ([]byte, []int) {
	return file_traffic_policy_proto_rawDescGZIP(), []int{1}
}

func (x *QOS) GetQueueStatistics() []*QOS_QueueStatistics {
	if x != nil {
		return x.QueueStatistics
	}
	return nil
}

type ConfiguredTrafficPolicy_QOS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue   int64                                `protobuf:"varint,1,opt,name=queue,proto3" json:"queue,omitempty" bson:"queue"`
	Shaping *ConfiguredTrafficPolicy_QOS_Shaping `protobuf:"bytes,2,opt,name=shaping,proto3" json:"shaping,omitempty" bson:"shaping"`
}

func (x *ConfiguredTrafficPolicy_QOS) Reset() {
	*x = ConfiguredTrafficPolicy_QOS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfiguredTrafficPolicy_QOS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfiguredTrafficPolicy_QOS) ProtoMessage() {}

func (x *ConfiguredTrafficPolicy_QOS) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfiguredTrafficPolicy_QOS.ProtoReflect.Descriptor instead.
func (*ConfiguredTrafficPolicy_QOS) Descriptor() ([]byte, []int) {
	return file_traffic_policy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConfiguredTrafficPolicy_QOS) GetQueue() int64 {
	if x != nil {
		return x.Queue
	}
	return 0
}

func (x *ConfiguredTrafficPolicy_QOS) GetShaping() *ConfiguredTrafficPolicy_QOS_Shaping {
	if x != nil {
		return x.Shaping
	}
	return nil
}

type ConfiguredTrafficPolicy_Statistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrafficPolicy string `protobuf:"bytes,1,opt,name=traffic_policy,json=trafficPolicy,proto3" json:"traffic_policy,omitempty" bson:"traffic_policy"`
	RuleNumber    int64  `protobuf:"varint,2,opt,name=rule_number,json=ruleNumber,proto3" json:"rule_number,omitempty" bson:"rule_number"`
	Status        string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty" bson:"status"`
	Interval      int64  `protobuf:"varint,4,opt,name=interval,proto3" json:"interval,omitempty" bson:"interval"`
	// AS A TABLE
	//
	//	PERMIT       drop        0 pass        0 for KBPS_0000
	//	DISCARD      drop        0 pass        0 for DISCARD
	//	VOIP         drop        0 pass        0 for VOIP
	//	OTHER        drop        0 pass        0 for KBPS_0000
	Classifiers map[string]*ConfiguredTrafficPolicy_Statistics_Classifier `protobuf:"bytes,5,rep,name=classifiers,proto3" json:"classifiers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"classifiers" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfiguredTrafficPolicy_Statistics) Reset() {
	*x = ConfiguredTrafficPolicy_Statistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfiguredTrafficPolicy_Statistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfiguredTrafficPolicy_Statistics) ProtoMessage() {}

func (x *ConfiguredTrafficPolicy_Statistics) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfiguredTrafficPolicy_Statistics.ProtoReflect.Descriptor instead.
func (*ConfiguredTrafficPolicy_Statistics) Descriptor() ([]byte, []int) {
	return file_traffic_policy_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ConfiguredTrafficPolicy_Statistics) GetTrafficPolicy() string {
	if x != nil {
		return x.TrafficPolicy
	}
	return ""
}

func (x *ConfiguredTrafficPolicy_Statistics) GetRuleNumber() int64 {
	if x != nil {
		return x.RuleNumber
	}
	return 0
}

func (x *ConfiguredTrafficPolicy_Statistics) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ConfiguredTrafficPolicy_Statistics) GetInterval() int64 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *ConfiguredTrafficPolicy_Statistics) GetClassifiers() map[string]*ConfiguredTrafficPolicy_Statistics_Classifier {
	if x != nil {
		return x.Classifiers
	}
	return nil
}

type ConfiguredTrafficPolicy_QOS_Shaping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cir float64 `protobuf:"fixed64,1,opt,name=cir,proto3" json:"cir,omitempty" bson:"cir"`
	Pir float64 `protobuf:"fixed64,2,opt,name=pir,proto3" json:"pir,omitempty" bson:"pir"`
	Cbs float64 `protobuf:"fixed64,3,opt,name=cbs,proto3" json:"cbs,omitempty" bson:"cbs"`
	Pbs float64 `protobuf:"fixed64,4,opt,name=pbs,proto3" json:"pbs,omitempty" bson:"pbs"`
}

func (x *ConfiguredTrafficPolicy_QOS_Shaping) Reset() {
	*x = ConfiguredTrafficPolicy_QOS_Shaping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfiguredTrafficPolicy_QOS_Shaping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfiguredTrafficPolicy_QOS_Shaping) ProtoMessage() {}

func (x *ConfiguredTrafficPolicy_QOS_Shaping) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfiguredTrafficPolicy_QOS_Shaping.ProtoReflect.Descriptor instead.
func (*ConfiguredTrafficPolicy_QOS_Shaping) Descriptor() ([]byte, []int) {
	return file_traffic_policy_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ConfiguredTrafficPolicy_QOS_Shaping) GetCir() float64 {
	if x != nil {
		return x.Cir
	}
	return 0
}

func (x *ConfiguredTrafficPolicy_QOS_Shaping) GetPir() float64 {
	if x != nil {
		return x.Pir
	}
	return 0
}

func (x *ConfiguredTrafficPolicy_QOS_Shaping) GetCbs() float64 {
	if x != nil {
		return x.Cbs
	}
	return 0
}

func (x *ConfiguredTrafficPolicy_QOS_Shaping) GetPbs() float64 {
	if x != nil {
		return x.Pbs
	}
	return 0
}

type ConfiguredTrafficPolicy_Statistics_Classifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Classifier string                                                           `protobuf:"bytes,1,opt,name=classifier,proto3" json:"classifier,omitempty" bson:"classifier"`
	Behavior   string                                                           `protobuf:"bytes,2,opt,name=behavior,proto3" json:"behavior,omitempty" bson:"behavior"`
	Board      string                                                           `protobuf:"bytes,3,opt,name=board,proto3" json:"board,omitempty" bson:"board"`
	Metrics    map[string]*ConfiguredTrafficPolicy_Statistics_Classifier_Metric `protobuf:"bytes,4,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"metrics" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier) Reset() {
	*x = ConfiguredTrafficPolicy_Statistics_Classifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_policy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfiguredTrafficPolicy_Statistics_Classifier) ProtoMessage() {}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_policy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfiguredTrafficPolicy_Statistics_Classifier.ProtoReflect.Descriptor instead.
func (*ConfiguredTrafficPolicy_Statistics_Classifier) Descriptor() ([]byte, []int) {
	return file_traffic_policy_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier) GetClassifier() string {
	if x != nil {
		return x.Classifier
	}
	return ""
}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier) GetBehavior() string {
	if x != nil {
		return x.Behavior
	}
	return ""
}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier) GetBoard() string {
	if x != nil {
		return x.Board
	}
	return ""
}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier) GetMetrics() map[string]*ConfiguredTrafficPolicy_Statistics_Classifier_Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type ConfiguredTrafficPolicy_Statistics_Classifier_Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values map[string]float64 `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3" bson:"values" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier_Metric) Reset() {
	*x = ConfiguredTrafficPolicy_Statistics_Classifier_Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_policy_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier_Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfiguredTrafficPolicy_Statistics_Classifier_Metric) ProtoMessage() {}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier_Metric) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_policy_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfiguredTrafficPolicy_Statistics_Classifier_Metric.ProtoReflect.Descriptor instead.
func (*ConfiguredTrafficPolicy_Statistics_Classifier_Metric) Descriptor() ([]byte, []int) {
	return file_traffic_policy_proto_rawDescGZIP(), []int{0, 1, 1, 1}
}

func (x *ConfiguredTrafficPolicy_Statistics_Classifier_Metric) GetValues() map[string]float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

type QOS_QueueStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId        int64   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty" bson:"queue_id"`
	Cir            float64 `protobuf:"fixed64,2,opt,name=cir,proto3" json:"cir,omitempty" bson:"cir"` // kbps
	Pir            float64 `protobuf:"fixed64,3,opt,name=pir,proto3" json:"pir,omitempty" bson:"pir"` // kbps
	PassedPackets  int64   `protobuf:"varint,4,opt,name=passed_packets,json=passedPackets,proto3" json:"passed_packets,omitempty" bson:"passed_packets"`
	PassedRatePps  float64 `protobuf:"fixed64,5,opt,name=passed_rate_pps,json=passedRatePps,proto3" json:"passed_rate_pps,omitempty" bson:"passed_rate_pps"` //pps
	PassedBytes    int64   `protobuf:"varint,6,opt,name=passed_bytes,json=passedBytes,proto3" json:"passed_bytes,omitempty" bson:"passed_bytes"`
	PassedRateBps  float64 `protobuf:"fixed64,7,opt,name=passed_rate_bps,json=passedRateBps,proto3" json:"passed_rate_bps,omitempty" bson:"passed_rate_bps"`
	DroppedPackets int64   `protobuf:"varint,8,opt,name=dropped_packets,json=droppedPackets,proto3" json:"dropped_packets,omitempty" bson:"dropped_packets"`
	DroppedRatePps float64 `protobuf:"fixed64,9,opt,name=dropped_rate_pps,json=droppedRatePps,proto3" json:"dropped_rate_pps,omitempty" bson:"dropped_rate_pps"` //pps
	DroppedBytes   int64   `protobuf:"varint,10,opt,name=dropped_bytes,json=droppedBytes,proto3" json:"dropped_bytes,omitempty" bson:"dropped_bytes"`
	DroppedRateBps float64 `protobuf:"fixed64,11,opt,name=dropped_rate_bps,json=droppedRateBps,proto3" json:"dropped_rate_bps,omitempty" bson:"dropped_rate_bps"`
}

func (x *QOS_QueueStatistics) Reset() {
	*x = QOS_QueueStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_traffic_policy_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QOS_QueueStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QOS_QueueStatistics) ProtoMessage() {}

func (x *QOS_QueueStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_traffic_policy_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QOS_QueueStatistics.ProtoReflect.Descriptor instead.
func (*QOS_QueueStatistics) Descriptor() ([]byte, []int) {
	return file_traffic_policy_proto_rawDescGZIP(), []int{1, 0}
}

func (x *QOS_QueueStatistics) GetQueueId() int64 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *QOS_QueueStatistics) GetCir() float64 {
	if x != nil {
		return x.Cir
	}
	return 0
}

func (x *QOS_QueueStatistics) GetPir() float64 {
	if x != nil {
		return x.Pir
	}
	return 0
}

func (x *QOS_QueueStatistics) GetPassedPackets() int64 {
	if x != nil {
		return x.PassedPackets
	}
	return 0
}

func (x *QOS_QueueStatistics) GetPassedRatePps() float64 {
	if x != nil {
		return x.PassedRatePps
	}
	return 0
}

func (x *QOS_QueueStatistics) GetPassedBytes() int64 {
	if x != nil {
		return x.PassedBytes
	}
	return 0
}

func (x *QOS_QueueStatistics) GetPassedRateBps() float64 {
	if x != nil {
		return x.PassedRateBps
	}
	return 0
}

func (x *QOS_QueueStatistics) GetDroppedPackets() int64 {
	if x != nil {
		return x.DroppedPackets
	}
	return 0
}

func (x *QOS_QueueStatistics) GetDroppedRatePps() float64 {
	if x != nil {
		return x.DroppedRatePps
	}
	return 0
}

func (x *QOS_QueueStatistics) GetDroppedBytes() int64 {
	if x != nil {
		return x.DroppedBytes
	}
	return 0
}

func (x *QOS_QueueStatistics) GetDroppedRateBps() float64 {
	if x != nil {
		return x.DroppedRateBps
	}
	return 0
}

var File_traffic_policy_proto protoreflect.FileDescriptor

var file_traffic_policy_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x9c, 0x0a, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x3d, 0x0a, 0x03, 0x71, 0x6f, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x51,
	0x4f, 0x53, 0x52, 0x03, 0x71, 0x6f, 0x73, 0x12, 0x61, 0x0a, 0x12, 0x69, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x11, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0xbd, 0x01, 0x0a, 0x03, 0x51,
	0x4f, 0x53, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x4d, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x70,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x51, 0x4f, 0x53, 0x2e, 0x53, 0x68, 0x61, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x07,
	0x73, 0x68, 0x61, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x51, 0x0a, 0x07, 0x53, 0x68, 0x61, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x63, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x70, 0x69, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x62, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x63, 0x62, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x62, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x70, 0x62, 0x73, 0x1a, 0xe8, 0x06, 0x0a, 0x0a, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x75, 0x6c, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x75, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x65, 0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x43, 0x2e, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x1a, 0x7d, 0x0a, 0x10,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x53, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3d, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xf7, 0x03, 0x0a, 0x0a,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x64, 0x0a, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e,
	0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x1a, 0x80, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xad, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x12, 0x68, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x50, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe4, 0x03, 0x0a, 0x03, 0x51, 0x4f, 0x53, 0x12, 0x4e, 0x0a,
	0x10, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x51, 0x4f, 0x53, 0x2e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x0f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x8c, 0x03,
	0x0a, 0x0f, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x73, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x63, 0x69, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x70, 0x69, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x70, 0x69, 0x72,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x65,
	0x64, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x70, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x50, 0x70, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x62, 0x70, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x70, 0x61, 0x73,
	0x73, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x42, 0x70, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x72,
	0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x70, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x64,
	0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x50, 0x70, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x5f, 0x62, 0x70, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x64, 0x72,
	0x6f, 0x70, 0x70, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x42, 0x70, 0x73, 0x42, 0x32, 0x5a, 0x30,
	0x67, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x63, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x67, 0x6f, 0x2d, 0x73, 0x77, 0x70, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_traffic_policy_proto_rawDescOnce sync.Once
	file_traffic_policy_proto_rawDescData = file_traffic_policy_proto_rawDesc
)

func file_traffic_policy_proto_rawDescGZIP() []byte {
	file_traffic_policy_proto_rawDescOnce.Do(func() {
		file_traffic_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_traffic_policy_proto_rawDescData)
	})
	return file_traffic_policy_proto_rawDescData
}

var file_traffic_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_traffic_policy_proto_goTypes = []interface{}{
	(*ConfiguredTrafficPolicy)(nil),             // 0: traffic_policy.ConfiguredTrafficPolicy
	(*QOS)(nil),                                 // 1: traffic_policy.QOS
	(*ConfiguredTrafficPolicy_QOS)(nil),         // 2: traffic_policy.ConfiguredTrafficPolicy.QOS
	(*ConfiguredTrafficPolicy_Statistics)(nil),  // 3: traffic_policy.ConfiguredTrafficPolicy.Statistics
	(*ConfiguredTrafficPolicy_QOS_Shaping)(nil), // 4: traffic_policy.ConfiguredTrafficPolicy.QOS.Shaping
	nil, // 5: traffic_policy.ConfiguredTrafficPolicy.Statistics.ClassifiersEntry
	(*ConfiguredTrafficPolicy_Statistics_Classifier)(nil), // 6: traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier
	nil, // 7: traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier.MetricsEntry
	(*ConfiguredTrafficPolicy_Statistics_Classifier_Metric)(nil), // 8: traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier.Metric
	nil,                         // 9: traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier.Metric.ValuesEntry
	(*QOS_QueueStatistics)(nil), // 10: traffic_policy.QOS.QueueStatistics
}
var file_traffic_policy_proto_depIdxs = []int32{
	2,  // 0: traffic_policy.ConfiguredTrafficPolicy.qos:type_name -> traffic_policy.ConfiguredTrafficPolicy.QOS
	3,  // 1: traffic_policy.ConfiguredTrafficPolicy.inbound_statistics:type_name -> traffic_policy.ConfiguredTrafficPolicy.Statistics
	10, // 2: traffic_policy.QOS.queue_statistics:type_name -> traffic_policy.QOS.QueueStatistics
	4,  // 3: traffic_policy.ConfiguredTrafficPolicy.QOS.shaping:type_name -> traffic_policy.ConfiguredTrafficPolicy.QOS.Shaping
	5,  // 4: traffic_policy.ConfiguredTrafficPolicy.Statistics.classifiers:type_name -> traffic_policy.ConfiguredTrafficPolicy.Statistics.ClassifiersEntry
	6,  // 5: traffic_policy.ConfiguredTrafficPolicy.Statistics.ClassifiersEntry.value:type_name -> traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier
	7,  // 6: traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier.metrics:type_name -> traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier.MetricsEntry
	8,  // 7: traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier.MetricsEntry.value:type_name -> traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier.Metric
	9,  // 8: traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier.Metric.values:type_name -> traffic_policy.ConfiguredTrafficPolicy.Statistics.Classifier.Metric.ValuesEntry
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_traffic_policy_proto_init() }
func file_traffic_policy_proto_init() {
	if File_traffic_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_traffic_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfiguredTrafficPolicy); i {
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
		file_traffic_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QOS); i {
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
		file_traffic_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfiguredTrafficPolicy_QOS); i {
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
		file_traffic_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfiguredTrafficPolicy_Statistics); i {
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
		file_traffic_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfiguredTrafficPolicy_QOS_Shaping); i {
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
		file_traffic_policy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfiguredTrafficPolicy_Statistics_Classifier); i {
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
		file_traffic_policy_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfiguredTrafficPolicy_Statistics_Classifier_Metric); i {
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
		file_traffic_policy_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QOS_QueueStatistics); i {
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
			RawDescriptor: file_traffic_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_traffic_policy_proto_goTypes,
		DependencyIndexes: file_traffic_policy_proto_depIdxs,
		MessageInfos:      file_traffic_policy_proto_msgTypes,
	}.Build()
	File_traffic_policy_proto = out.File
	file_traffic_policy_proto_rawDesc = nil
	file_traffic_policy_proto_goTypes = nil
	file_traffic_policy_proto_depIdxs = nil
}
