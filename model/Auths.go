package model

import "time"

// 权限表
type Auths struct { 
	Id  int  `json:"id"`  
	AuthName  string  `json:"auth_name"`  // 权限名称
	CreatedAt  time.Time  `json:"created_at"`  ad
	UpdatedAt  time.Time  `json:"updated_at"`  
	State  int  `json:"state"`  // 状态（有效-1/失效-0）
}