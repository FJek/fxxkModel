package model

// 用户角色表
type SysUserRole struct { 
	Id  int  `json:"id"`  
	UserId  int  `json:"user_id"`  // 用户id
	RoleId  int  `json:"role_id"`  // 角色id
	CreatedAt  time.Time  `json:"created_at"`  
	UpdatedAt  time.Time  `json:"updated_at"`  
	State  int  `json:"state"`  // 状态（失效-0/有效-1）
}