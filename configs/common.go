package configs

const (
	// TIMEOUT string for timeout handler response
	TIMEOUT = "timeout"

	// ClustersKey and the tree like: /clusters/{clusterID}/{serverID}
	ClustersKey = "/clusters/"

	// PluginsKey and the tree like: /plugins/{pluginNamespace}/{key1}/{key2}/...
	PluginsKey = "/plugins/"
	// RbacKey "rbac/" means the root key, it contains "users/" "roles/" "permissions/"
	RbacKey = "rbac/"
	// RbacUsersKey users/
	RbacUsersKey = "users/"
	// RbacRolesKey roles/
	RbacRolesKey = "roles/"
	// RbacPermissionKey permission/
	RbacPermissionKey = "permission/"
	// CacheKey "cache/" means the root key
	CacheKey = "cache/"
)

/*
 * /plugins/roles/{roleID} is the key to value of Role{ID: roleID, Permissions: PermissionList}
 */
