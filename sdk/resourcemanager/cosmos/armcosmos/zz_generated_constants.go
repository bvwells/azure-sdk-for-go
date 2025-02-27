//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

const (
	module  = "armcosmos"
	version = "v0.1.0"
)

// APIType - Enum to indicate the API type of the restorable database account.
type APIType string

const (
	APITypeCassandra APIType = "Cassandra"
	APITypeGremlin   APIType = "Gremlin"
	APITypeGremlinV2 APIType = "GremlinV2"
	APITypeMongoDB   APIType = "MongoDB"
	APITypeSQL       APIType = "Sql"
	APITypeTable     APIType = "Table"
)

// PossibleAPITypeValues returns the possible values for the APIType const type.
func PossibleAPITypeValues() []APIType {
	return []APIType{
		APITypeCassandra,
		APITypeGremlin,
		APITypeGremlinV2,
		APITypeMongoDB,
		APITypeSQL,
		APITypeTable,
	}
}

// ToPtr returns a *APIType pointing to the current value.
func (c APIType) ToPtr() *APIType {
	return &c
}

// AnalyticalStorageSchemaType - Describes the types of schema for analytical storage.
type AnalyticalStorageSchemaType string

const (
	AnalyticalStorageSchemaTypeFullFidelity AnalyticalStorageSchemaType = "FullFidelity"
	AnalyticalStorageSchemaTypeWellDefined  AnalyticalStorageSchemaType = "WellDefined"
)

// PossibleAnalyticalStorageSchemaTypeValues returns the possible values for the AnalyticalStorageSchemaType const type.
func PossibleAnalyticalStorageSchemaTypeValues() []AnalyticalStorageSchemaType {
	return []AnalyticalStorageSchemaType{
		AnalyticalStorageSchemaTypeFullFidelity,
		AnalyticalStorageSchemaTypeWellDefined,
	}
}

// ToPtr returns a *AnalyticalStorageSchemaType pointing to the current value.
func (c AnalyticalStorageSchemaType) ToPtr() *AnalyticalStorageSchemaType {
	return &c
}

// AuthenticationMethod - Which authentication method Cassandra should use to authenticate clients. 'None' turns off authentication, so should not be used
// except in emergencies. 'Cassandra' is the default password based
// authentication. The default is 'Cassandra'.
type AuthenticationMethod string

const (
	AuthenticationMethodCassandra AuthenticationMethod = "Cassandra"
	AuthenticationMethodNone      AuthenticationMethod = "None"
)

// PossibleAuthenticationMethodValues returns the possible values for the AuthenticationMethod const type.
func PossibleAuthenticationMethodValues() []AuthenticationMethod {
	return []AuthenticationMethod{
		AuthenticationMethodCassandra,
		AuthenticationMethodNone,
	}
}

// ToPtr returns a *AuthenticationMethod pointing to the current value.
func (c AuthenticationMethod) ToPtr() *AuthenticationMethod {
	return &c
}

// BackupPolicyMigrationStatus - Describes the status of migration between backup policy types.
type BackupPolicyMigrationStatus string

const (
	BackupPolicyMigrationStatusCompleted  BackupPolicyMigrationStatus = "Completed"
	BackupPolicyMigrationStatusFailed     BackupPolicyMigrationStatus = "Failed"
	BackupPolicyMigrationStatusInProgress BackupPolicyMigrationStatus = "InProgress"
	BackupPolicyMigrationStatusInvalid    BackupPolicyMigrationStatus = "Invalid"
)

// PossibleBackupPolicyMigrationStatusValues returns the possible values for the BackupPolicyMigrationStatus const type.
func PossibleBackupPolicyMigrationStatusValues() []BackupPolicyMigrationStatus {
	return []BackupPolicyMigrationStatus{
		BackupPolicyMigrationStatusCompleted,
		BackupPolicyMigrationStatusFailed,
		BackupPolicyMigrationStatusInProgress,
		BackupPolicyMigrationStatusInvalid,
	}
}

// ToPtr returns a *BackupPolicyMigrationStatus pointing to the current value.
func (c BackupPolicyMigrationStatus) ToPtr() *BackupPolicyMigrationStatus {
	return &c
}

// BackupPolicyType - Describes the mode of backups.
type BackupPolicyType string

const (
	BackupPolicyTypeContinuous BackupPolicyType = "Continuous"
	BackupPolicyTypePeriodic   BackupPolicyType = "Periodic"
)

// PossibleBackupPolicyTypeValues returns the possible values for the BackupPolicyType const type.
func PossibleBackupPolicyTypeValues() []BackupPolicyType {
	return []BackupPolicyType{
		BackupPolicyTypeContinuous,
		BackupPolicyTypePeriodic,
	}
}

// ToPtr returns a *BackupPolicyType pointing to the current value.
func (c BackupPolicyType) ToPtr() *BackupPolicyType {
	return &c
}

// BackupStorageRedundancy - Enum to indicate type of backup storage redundancy.
type BackupStorageRedundancy string

const (
	BackupStorageRedundancyGeo   BackupStorageRedundancy = "Geo"
	BackupStorageRedundancyLocal BackupStorageRedundancy = "Local"
	BackupStorageRedundancyZone  BackupStorageRedundancy = "Zone"
)

// PossibleBackupStorageRedundancyValues returns the possible values for the BackupStorageRedundancy const type.
func PossibleBackupStorageRedundancyValues() []BackupStorageRedundancy {
	return []BackupStorageRedundancy{
		BackupStorageRedundancyGeo,
		BackupStorageRedundancyLocal,
		BackupStorageRedundancyZone,
	}
}

// ToPtr returns a *BackupStorageRedundancy pointing to the current value.
func (c BackupStorageRedundancy) ToPtr() *BackupStorageRedundancy {
	return &c
}

// CompositePathSortOrder - Sort order for composite paths.
type CompositePathSortOrder string

const (
	CompositePathSortOrderAscending  CompositePathSortOrder = "ascending"
	CompositePathSortOrderDescending CompositePathSortOrder = "descending"
)

// PossibleCompositePathSortOrderValues returns the possible values for the CompositePathSortOrder const type.
func PossibleCompositePathSortOrderValues() []CompositePathSortOrder {
	return []CompositePathSortOrder{
		CompositePathSortOrderAscending,
		CompositePathSortOrderDescending,
	}
}

// ToPtr returns a *CompositePathSortOrder pointing to the current value.
func (c CompositePathSortOrder) ToPtr() *CompositePathSortOrder {
	return &c
}

// ConflictResolutionMode - Indicates the conflict resolution mode.
type ConflictResolutionMode string

const (
	ConflictResolutionModeCustom         ConflictResolutionMode = "Custom"
	ConflictResolutionModeLastWriterWins ConflictResolutionMode = "LastWriterWins"
)

// PossibleConflictResolutionModeValues returns the possible values for the ConflictResolutionMode const type.
func PossibleConflictResolutionModeValues() []ConflictResolutionMode {
	return []ConflictResolutionMode{
		ConflictResolutionModeCustom,
		ConflictResolutionModeLastWriterWins,
	}
}

// ToPtr returns a *ConflictResolutionMode pointing to the current value.
func (c ConflictResolutionMode) ToPtr() *ConflictResolutionMode {
	return &c
}

// ConnectorOffer - The cassandra connector offer type for the Cosmos DB C* database account.
type ConnectorOffer string

const (
	ConnectorOfferSmall ConnectorOffer = "Small"
)

// PossibleConnectorOfferValues returns the possible values for the ConnectorOffer const type.
func PossibleConnectorOfferValues() []ConnectorOffer {
	return []ConnectorOffer{
		ConnectorOfferSmall,
	}
}

// ToPtr returns a *ConnectorOffer pointing to the current value.
func (c ConnectorOffer) ToPtr() *ConnectorOffer {
	return &c
}

// CreateMode - Enum to indicate the mode of account creation.
type CreateMode string

const (
	CreateModeDefault CreateMode = "Default"
	CreateModeRestore CreateMode = "Restore"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeDefault,
		CreateModeRestore,
	}
}

// ToPtr returns a *CreateMode pointing to the current value.
func (c CreateMode) ToPtr() *CreateMode {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// DataType - The datatype for which the indexing behavior is applied to.
type DataType string

const (
	DataTypeLineString   DataType = "LineString"
	DataTypeMultiPolygon DataType = "MultiPolygon"
	DataTypeNumber       DataType = "Number"
	DataTypePoint        DataType = "Point"
	DataTypePolygon      DataType = "Polygon"
	DataTypeString       DataType = "String"
)

// PossibleDataTypeValues returns the possible values for the DataType const type.
func PossibleDataTypeValues() []DataType {
	return []DataType{
		DataTypeLineString,
		DataTypeMultiPolygon,
		DataTypeNumber,
		DataTypePoint,
		DataTypePolygon,
		DataTypeString,
	}
}

// ToPtr returns a *DataType pointing to the current value.
func (c DataType) ToPtr() *DataType {
	return &c
}

// DatabaseAccountKind - Indicates the type of database account. This can only be set at database account creation.
type DatabaseAccountKind string

const (
	DatabaseAccountKindGlobalDocumentDB DatabaseAccountKind = "GlobalDocumentDB"
	DatabaseAccountKindMongoDB          DatabaseAccountKind = "MongoDB"
	DatabaseAccountKindParse            DatabaseAccountKind = "Parse"
)

// PossibleDatabaseAccountKindValues returns the possible values for the DatabaseAccountKind const type.
func PossibleDatabaseAccountKindValues() []DatabaseAccountKind {
	return []DatabaseAccountKind{
		DatabaseAccountKindGlobalDocumentDB,
		DatabaseAccountKindMongoDB,
		DatabaseAccountKindParse,
	}
}

// ToPtr returns a *DatabaseAccountKind pointing to the current value.
func (c DatabaseAccountKind) ToPtr() *DatabaseAccountKind {
	return &c
}

// DefaultConsistencyLevel - The default consistency level and configuration settings of the Cosmos DB account.
type DefaultConsistencyLevel string

const (
	DefaultConsistencyLevelEventual         DefaultConsistencyLevel = "Eventual"
	DefaultConsistencyLevelSession          DefaultConsistencyLevel = "Session"
	DefaultConsistencyLevelBoundedStaleness DefaultConsistencyLevel = "BoundedStaleness"
	DefaultConsistencyLevelStrong           DefaultConsistencyLevel = "Strong"
	DefaultConsistencyLevelConsistentPrefix DefaultConsistencyLevel = "ConsistentPrefix"
)

// PossibleDefaultConsistencyLevelValues returns the possible values for the DefaultConsistencyLevel const type.
func PossibleDefaultConsistencyLevelValues() []DefaultConsistencyLevel {
	return []DefaultConsistencyLevel{
		DefaultConsistencyLevelEventual,
		DefaultConsistencyLevelSession,
		DefaultConsistencyLevelBoundedStaleness,
		DefaultConsistencyLevelStrong,
		DefaultConsistencyLevelConsistentPrefix,
	}
}

// ToPtr returns a *DefaultConsistencyLevel pointing to the current value.
func (c DefaultConsistencyLevel) ToPtr() *DefaultConsistencyLevel {
	return &c
}

// EnableFullTextQuery - Describe the level of detail with which queries are to be logged.
type EnableFullTextQuery string

const (
	EnableFullTextQueryNone  EnableFullTextQuery = "None"
	EnableFullTextQueryTrue  EnableFullTextQuery = "True"
	EnableFullTextQueryFalse EnableFullTextQuery = "False"
)

// PossibleEnableFullTextQueryValues returns the possible values for the EnableFullTextQuery const type.
func PossibleEnableFullTextQueryValues() []EnableFullTextQuery {
	return []EnableFullTextQuery{
		EnableFullTextQueryNone,
		EnableFullTextQueryTrue,
		EnableFullTextQueryFalse,
	}
}

// ToPtr returns a *EnableFullTextQuery pointing to the current value.
func (c EnableFullTextQuery) ToPtr() *EnableFullTextQuery {
	return &c
}

// IndexKind - Indicates the type of index.
type IndexKind string

const (
	IndexKindHash    IndexKind = "Hash"
	IndexKindRange   IndexKind = "Range"
	IndexKindSpatial IndexKind = "Spatial"
)

// PossibleIndexKindValues returns the possible values for the IndexKind const type.
func PossibleIndexKindValues() []IndexKind {
	return []IndexKind{
		IndexKindHash,
		IndexKindRange,
		IndexKindSpatial,
	}
}

// ToPtr returns a *IndexKind pointing to the current value.
func (c IndexKind) ToPtr() *IndexKind {
	return &c
}

// IndexingMode - Indicates the indexing mode.
type IndexingMode string

const (
	IndexingModeConsistent IndexingMode = "consistent"
	IndexingModeLazy       IndexingMode = "lazy"
	IndexingModeNone       IndexingMode = "none"
)

// PossibleIndexingModeValues returns the possible values for the IndexingMode const type.
func PossibleIndexingModeValues() []IndexingMode {
	return []IndexingMode{
		IndexingModeConsistent,
		IndexingModeLazy,
		IndexingModeNone,
	}
}

// ToPtr returns a *IndexingMode pointing to the current value.
func (c IndexingMode) ToPtr() *IndexingMode {
	return &c
}

// KeyKind - The access key to regenerate.
type KeyKind string

const (
	KeyKindPrimary           KeyKind = "primary"
	KeyKindPrimaryReadonly   KeyKind = "primaryReadonly"
	KeyKindSecondary         KeyKind = "secondary"
	KeyKindSecondaryReadonly KeyKind = "secondaryReadonly"
)

// PossibleKeyKindValues returns the possible values for the KeyKind const type.
func PossibleKeyKindValues() []KeyKind {
	return []KeyKind{
		KeyKindPrimary,
		KeyKindPrimaryReadonly,
		KeyKindSecondary,
		KeyKindSecondaryReadonly,
	}
}

// ToPtr returns a *KeyKind pointing to the current value.
func (c KeyKind) ToPtr() *KeyKind {
	return &c
}

// ManagedCassandraProvisioningState - The status of the resource at the time the operation was called.
type ManagedCassandraProvisioningState string

const (
	ManagedCassandraProvisioningStateCanceled  ManagedCassandraProvisioningState = "Canceled"
	ManagedCassandraProvisioningStateCreating  ManagedCassandraProvisioningState = "Creating"
	ManagedCassandraProvisioningStateDeleting  ManagedCassandraProvisioningState = "Deleting"
	ManagedCassandraProvisioningStateFailed    ManagedCassandraProvisioningState = "Failed"
	ManagedCassandraProvisioningStateSucceeded ManagedCassandraProvisioningState = "Succeeded"
	ManagedCassandraProvisioningStateUpdating  ManagedCassandraProvisioningState = "Updating"
)

// PossibleManagedCassandraProvisioningStateValues returns the possible values for the ManagedCassandraProvisioningState const type.
func PossibleManagedCassandraProvisioningStateValues() []ManagedCassandraProvisioningState {
	return []ManagedCassandraProvisioningState{
		ManagedCassandraProvisioningStateCanceled,
		ManagedCassandraProvisioningStateCreating,
		ManagedCassandraProvisioningStateDeleting,
		ManagedCassandraProvisioningStateFailed,
		ManagedCassandraProvisioningStateSucceeded,
		ManagedCassandraProvisioningStateUpdating,
	}
}

// ToPtr returns a *ManagedCassandraProvisioningState pointing to the current value.
func (c ManagedCassandraProvisioningState) ToPtr() *ManagedCassandraProvisioningState {
	return &c
}

// NetworkACLBypass - Indicates what services are allowed to bypass firewall checks.
type NetworkACLBypass string

const (
	NetworkACLBypassNone          NetworkACLBypass = "None"
	NetworkACLBypassAzureServices NetworkACLBypass = "AzureServices"
)

// PossibleNetworkACLBypassValues returns the possible values for the NetworkACLBypass const type.
func PossibleNetworkACLBypassValues() []NetworkACLBypass {
	return []NetworkACLBypass{
		NetworkACLBypassNone,
		NetworkACLBypassAzureServices,
	}
}

// ToPtr returns a *NetworkACLBypass pointing to the current value.
func (c NetworkACLBypass) ToPtr() *NetworkACLBypass {
	return &c
}

// NodeState - The state of the node in relation to the cluster.
type NodeState string

const (
	NodeStateJoining NodeState = "Joining"
	NodeStateLeaving NodeState = "Leaving"
	NodeStateMoving  NodeState = "Moving"
	NodeStateNormal  NodeState = "Normal"
	NodeStateStopped NodeState = "Stopped"
)

// PossibleNodeStateValues returns the possible values for the NodeState const type.
func PossibleNodeStateValues() []NodeState {
	return []NodeState{
		NodeStateJoining,
		NodeStateLeaving,
		NodeStateMoving,
		NodeStateNormal,
		NodeStateStopped,
	}
}

// ToPtr returns a *NodeState pointing to the current value.
func (c NodeState) ToPtr() *NodeState {
	return &c
}

// NodeStatus - Indicates whether the node is functioning or not.
type NodeStatus string

const (
	NodeStatusDown NodeStatus = "Down"
	NodeStatusUp   NodeStatus = "Up"
)

// PossibleNodeStatusValues returns the possible values for the NodeStatus const type.
func PossibleNodeStatusValues() []NodeStatus {
	return []NodeStatus{
		NodeStatusDown,
		NodeStatusUp,
	}
}

// ToPtr returns a *NodeStatus pointing to the current value.
func (c NodeStatus) ToPtr() *NodeStatus {
	return &c
}

type NotebookWorkspaceName string

const (
	NotebookWorkspaceNameDefault NotebookWorkspaceName = "default"
)

// PossibleNotebookWorkspaceNameValues returns the possible values for the NotebookWorkspaceName const type.
func PossibleNotebookWorkspaceNameValues() []NotebookWorkspaceName {
	return []NotebookWorkspaceName{
		NotebookWorkspaceNameDefault,
	}
}

// ToPtr returns a *NotebookWorkspaceName pointing to the current value.
func (c NotebookWorkspaceName) ToPtr() *NotebookWorkspaceName {
	return &c
}

// OperationType - Enum to indicate the operation type of the event.
type OperationType string

const (
	OperationTypeCreate          OperationType = "Create"
	OperationTypeDelete          OperationType = "Delete"
	OperationTypeReplace         OperationType = "Replace"
	OperationTypeSystemOperation OperationType = "SystemOperation"
)

// PossibleOperationTypeValues returns the possible values for the OperationType const type.
func PossibleOperationTypeValues() []OperationType {
	return []OperationType{
		OperationTypeCreate,
		OperationTypeDelete,
		OperationTypeReplace,
		OperationTypeSystemOperation,
	}
}

// ToPtr returns a *OperationType pointing to the current value.
func (c OperationType) ToPtr() *OperationType {
	return &c
}

// PartitionKind - Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum) are supported for
// container create
type PartitionKind string

const (
	PartitionKindHash      PartitionKind = "Hash"
	PartitionKindMultiHash PartitionKind = "MultiHash"
	PartitionKindRange     PartitionKind = "Range"
)

// PossiblePartitionKindValues returns the possible values for the PartitionKind const type.
func PossiblePartitionKindValues() []PartitionKind {
	return []PartitionKind{
		PartitionKindHash,
		PartitionKindMultiHash,
		PartitionKindRange,
	}
}

// ToPtr returns a *PartitionKind pointing to the current value.
func (c PartitionKind) ToPtr() *PartitionKind {
	return &c
}

// PrimaryAggregationType - The primary aggregation type of the metric.
type PrimaryAggregationType string

const (
	PrimaryAggregationTypeAverage PrimaryAggregationType = "Average"
	PrimaryAggregationTypeLast    PrimaryAggregationType = "Last"
	PrimaryAggregationTypeMaximum PrimaryAggregationType = "Maximum"
	PrimaryAggregationTypeMinimum PrimaryAggregationType = "Minimum"
	PrimaryAggregationTypeNone    PrimaryAggregationType = "None"
	PrimaryAggregationTypeTotal   PrimaryAggregationType = "Total"
)

// PossiblePrimaryAggregationTypeValues returns the possible values for the PrimaryAggregationType const type.
func PossiblePrimaryAggregationTypeValues() []PrimaryAggregationType {
	return []PrimaryAggregationType{
		PrimaryAggregationTypeAverage,
		PrimaryAggregationTypeLast,
		PrimaryAggregationTypeMaximum,
		PrimaryAggregationTypeMinimum,
		PrimaryAggregationTypeNone,
		PrimaryAggregationTypeTotal,
	}
}

// ToPtr returns a *PrimaryAggregationType pointing to the current value.
func (c PrimaryAggregationType) ToPtr() *PrimaryAggregationType {
	return &c
}

// PublicNetworkAccess - Whether requests from Public Network are allowed
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ToPtr returns a *PublicNetworkAccess pointing to the current value.
func (c PublicNetworkAccess) ToPtr() *PublicNetworkAccess {
	return &c
}

// ResourceIdentityType - The type of identity used for the resource. The type 'SystemAssigned,UserAssigned' includes both an implicitly created identity
// and a set of user assigned identities. The type 'None' will remove any
// identities from the service.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned,UserAssigned"
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// ToPtr returns a *ResourceIdentityType pointing to the current value.
func (c ResourceIdentityType) ToPtr() *ResourceIdentityType {
	return &c
}

// RestoreMode - Describes the mode of the restore.
type RestoreMode string

const (
	RestoreModePointInTime RestoreMode = "PointInTime"
)

// PossibleRestoreModeValues returns the possible values for the RestoreMode const type.
func PossibleRestoreModeValues() []RestoreMode {
	return []RestoreMode{
		RestoreModePointInTime,
	}
}

// ToPtr returns a *RestoreMode pointing to the current value.
func (c RestoreMode) ToPtr() *RestoreMode {
	return &c
}

// RoleDefinitionType - Indicates whether the Role Definition was built-in or user created.
type RoleDefinitionType string

const (
	RoleDefinitionTypeBuiltInRole RoleDefinitionType = "BuiltInRole"
	RoleDefinitionTypeCustomRole  RoleDefinitionType = "CustomRole"
)

// PossibleRoleDefinitionTypeValues returns the possible values for the RoleDefinitionType const type.
func PossibleRoleDefinitionTypeValues() []RoleDefinitionType {
	return []RoleDefinitionType{
		RoleDefinitionTypeBuiltInRole,
		RoleDefinitionTypeCustomRole,
	}
}

// ToPtr returns a *RoleDefinitionType pointing to the current value.
func (c RoleDefinitionType) ToPtr() *RoleDefinitionType {
	return &c
}

// ServerVersion - Describes the ServerVersion of an a MongoDB account.
type ServerVersion string

const (
	ServerVersionFour0  ServerVersion = "4.0"
	ServerVersionThree2 ServerVersion = "3.2"
	ServerVersionThree6 ServerVersion = "3.6"
)

// PossibleServerVersionValues returns the possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{
		ServerVersionFour0,
		ServerVersionThree2,
		ServerVersionThree6,
	}
}

// ToPtr returns a *ServerVersion pointing to the current value.
func (c ServerVersion) ToPtr() *ServerVersion {
	return &c
}

// ServiceSize - Instance type for the service.
type ServiceSize string

const (
	ServiceSizeCosmosD16S ServiceSize = "Cosmos.D16s"
	ServiceSizeCosmosD4S  ServiceSize = "Cosmos.D4s"
	ServiceSizeCosmosD8S  ServiceSize = "Cosmos.D8s"
)

// PossibleServiceSizeValues returns the possible values for the ServiceSize const type.
func PossibleServiceSizeValues() []ServiceSize {
	return []ServiceSize{
		ServiceSizeCosmosD16S,
		ServiceSizeCosmosD4S,
		ServiceSizeCosmosD8S,
	}
}

// ToPtr returns a *ServiceSize pointing to the current value.
func (c ServiceSize) ToPtr() *ServiceSize {
	return &c
}

// ServiceStatus - Describes the status of a service.
type ServiceStatus string

const (
	ServiceStatusCreating ServiceStatus = "Creating"
	ServiceStatusDeleting ServiceStatus = "Deleting"
	ServiceStatusError    ServiceStatus = "Error"
	ServiceStatusRunning  ServiceStatus = "Running"
	ServiceStatusStopped  ServiceStatus = "Stopped"
	ServiceStatusUpdating ServiceStatus = "Updating"
)

// PossibleServiceStatusValues returns the possible values for the ServiceStatus const type.
func PossibleServiceStatusValues() []ServiceStatus {
	return []ServiceStatus{
		ServiceStatusCreating,
		ServiceStatusDeleting,
		ServiceStatusError,
		ServiceStatusRunning,
		ServiceStatusStopped,
		ServiceStatusUpdating,
	}
}

// ToPtr returns a *ServiceStatus pointing to the current value.
func (c ServiceStatus) ToPtr() *ServiceStatus {
	return &c
}

// ServiceType - ServiceType for the service.
type ServiceType string

const (
	ServiceTypeDataTransfer        ServiceType = "DataTransfer"
	ServiceTypeGraphAPICompute     ServiceType = "GraphAPICompute"
	ServiceTypeSQLDedicatedGateway ServiceType = "SqlDedicatedGateway"
)

// PossibleServiceTypeValues returns the possible values for the ServiceType const type.
func PossibleServiceTypeValues() []ServiceType {
	return []ServiceType{
		ServiceTypeDataTransfer,
		ServiceTypeGraphAPICompute,
		ServiceTypeSQLDedicatedGateway,
	}
}

// ToPtr returns a *ServiceType pointing to the current value.
func (c ServiceType) ToPtr() *ServiceType {
	return &c
}

// SpatialType - Indicates the spatial type of index.
type SpatialType string

const (
	SpatialTypeLineString   SpatialType = "LineString"
	SpatialTypeMultiPolygon SpatialType = "MultiPolygon"
	SpatialTypePoint        SpatialType = "Point"
	SpatialTypePolygon      SpatialType = "Polygon"
)

// PossibleSpatialTypeValues returns the possible values for the SpatialType const type.
func PossibleSpatialTypeValues() []SpatialType {
	return []SpatialType{
		SpatialTypeLineString,
		SpatialTypeMultiPolygon,
		SpatialTypePoint,
		SpatialTypePolygon,
	}
}

// ToPtr returns a *SpatialType pointing to the current value.
func (c SpatialType) ToPtr() *SpatialType {
	return &c
}

// TriggerOperation - The operation the trigger is associated with
type TriggerOperation string

const (
	TriggerOperationAll     TriggerOperation = "All"
	TriggerOperationCreate  TriggerOperation = "Create"
	TriggerOperationDelete  TriggerOperation = "Delete"
	TriggerOperationReplace TriggerOperation = "Replace"
	TriggerOperationUpdate  TriggerOperation = "Update"
)

// PossibleTriggerOperationValues returns the possible values for the TriggerOperation const type.
func PossibleTriggerOperationValues() []TriggerOperation {
	return []TriggerOperation{
		TriggerOperationAll,
		TriggerOperationCreate,
		TriggerOperationDelete,
		TriggerOperationReplace,
		TriggerOperationUpdate,
	}
}

// ToPtr returns a *TriggerOperation pointing to the current value.
func (c TriggerOperation) ToPtr() *TriggerOperation {
	return &c
}

// TriggerType - Type of the Trigger
type TriggerType string

const (
	TriggerTypePost TriggerType = "Post"
	TriggerTypePre  TriggerType = "Pre"
)

// PossibleTriggerTypeValues returns the possible values for the TriggerType const type.
func PossibleTriggerTypeValues() []TriggerType {
	return []TriggerType{
		TriggerTypePost,
		TriggerTypePre,
	}
}

// ToPtr returns a *TriggerType pointing to the current value.
func (c TriggerType) ToPtr() *TriggerType {
	return &c
}

// UnitType - The unit of the metric.
type UnitType string

const (
	UnitTypeBytes          UnitType = "Bytes"
	UnitTypeBytesPerSecond UnitType = "BytesPerSecond"
	UnitTypeCount          UnitType = "Count"
	UnitTypeCountPerSecond UnitType = "CountPerSecond"
	UnitTypeMilliseconds   UnitType = "Milliseconds"
	UnitTypePercent        UnitType = "Percent"
	UnitTypeSeconds        UnitType = "Seconds"
)

// PossibleUnitTypeValues returns the possible values for the UnitType const type.
func PossibleUnitTypeValues() []UnitType {
	return []UnitType{
		UnitTypeBytes,
		UnitTypeBytesPerSecond,
		UnitTypeCount,
		UnitTypeCountPerSecond,
		UnitTypeMilliseconds,
		UnitTypePercent,
		UnitTypeSeconds,
	}
}

// ToPtr returns a *UnitType pointing to the current value.
func (c UnitType) ToPtr() *UnitType {
	return &c
}
