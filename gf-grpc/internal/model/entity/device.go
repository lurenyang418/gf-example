// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure for table device.
type Device struct {
	Id        uint        `json:"id"        orm:"id"         ` // 主键ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	DeletedAt uint64      `json:"deletedAt" orm:"deleted_at" ` // 软删除
	Name      string      `json:"name"      orm:"name"       ` // 名称
}
