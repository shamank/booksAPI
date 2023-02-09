package models

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleID   int    `json:"role_id" db:"role_id"`
	RoleName string `json:"role_name" db:"role_name"`
}
