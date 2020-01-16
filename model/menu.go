package model

import "time"

// 后台菜单表
type Menu struct {
	Id           string    `json:"id"`             // 菜单ID
	MenuName     string    `json:"menu_name"`      // 菜单名称
	MenuUrl      string    `json:"menu_url"`       // 菜单url
	ParentMenuId int       `json:"parent_menu_id"` // 父菜单ID
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	State        int       `json:"state"` // 状态
}
