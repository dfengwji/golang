package model

import "eas.cloud/proxy"

const (
	RoleAdminSuper  uint8 = 0
	RoleAdminNormal uint8 = 1
	RoleTeacher     uint8 = 2
)

type AdminInfo struct {
	BaseInfo
	ID       uint32 `json:"id"`
	Role     uint8  `json:"type"`
	Password string `json:"psw"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (a *AdminInfo) CopyAdmin(admin *proxy.Admin) {
	a.UID = admin.UID.Hex()
	a.Name = admin.Name
	a.Password = admin.Password
	a.Role = admin.Role
	a.Phone = admin.Phone
	a.ID = admin.ID
	a.Email = admin.Email
}
