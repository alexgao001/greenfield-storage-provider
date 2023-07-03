package sqldb

import "fmt"

// MigrateSubscribeProgressTable table schema.
type MigrateSubscribeProgressTable struct {
	EventName                 string `gorm:"primary_key"`
	LastSubscribedBlockHeight uint64
}

// TableName is used to set MigrateSubscribeEventProgressTable Schema's table name in database.
func (MigrateSubscribeProgressTable) TableName() string {
	return MigrateSubscribeProgressTableName
}

// MigrateGVGTable table schema.
// sp exit, bucket migrate
type MigrateGVGTable struct {
	MigrateKey             string `gorm:"primary_key"`
	GlobalVirtualGroupID   uint32 `gorm:"index:gvg_index"`
	VirtualGroupFamilyID   uint32 `gorm:"index:vgf_index"`
	MigrateRedundancyIndex int32
	BucketID               uint64 `gorm:"index:bucket_index"`
	IsSecondaryGVG         bool   `gorm:"index:secondary_index"`
	SrcSPID                uint32
	DestSPID               uint32
	LastMigrateObjectID    uint64
	MigrateStatus          int
	CheckStatus            int
}

// MigrateGVGPrimaryKey defines MigrateGVGTable primary key.
func MigrateGVGPrimaryKey(m *MigrateGVGTable) string {
	return fmt.Sprintf("gvg_id[%d]-vgf_id[%d]-reduncdancy_idx[%d]-bucket_id[%d]-is_secondary[%d]",
		m.GlobalVirtualGroupID, m.VirtualGroupFamilyID, m.MigrateRedundancyIndex, m.BucketID, m.BucketID)
}

// TableName is used to set MigrateGVGTable Schema's table name in database.
func (MigrateGVGTable) TableName() string {
	return MigrateGVGTableName
}
