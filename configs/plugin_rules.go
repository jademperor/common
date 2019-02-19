package configs

// // PermissionModel struct for RBAC plugins
// type PermissionModel struct {
// 	Idx      string
// 	Action   string
// 	Resource string
// }

// // Equal judge a permission is equal to another or not
// func (p *PermissionModel) Equal(perm *PermissionModel) bool {
// 	if perm.Idx != "" && p.Idx == perm.Idx {
// 		return true
// 	}

// 	if perm.Action != "" {
// 		if perm.Action == p.Action && perm.Resource == p.Resource {
// 			return true
// 		}
// 	}

// 	return false
// }

// // RoleModel ...
// type RoleModel struct {
// 	Idx         string
// 	Permissions map[string]*PermissionModel
// 	Name        string
// }

// // Permit ...
// func (r *RoleModel) Permit(target *PermissionModel) bool {
// 	for _, perm := range r.Permissions {
// 		if target.Equal(perm) {
// 			return true
// 		}
// 	}
// 	return false
// }

// // UserModel ...
// type UserModel struct {
// 	Idx    string
// 	UserID string
// 	Roles  map[string]*RoleModel
// }

// // PermitURLModel ...
// type PermitURLModel struct {
// 	Idx        string
// 	Permission *PermissionModel
// 	URI        string
// }
