package configs

const (
	// TIMEOUT string for timeout handler response
	TIMEOUT = "timeout"

	ClustersKey = "/clusters/" // ClustersKey and the tree like: /clusters/{clusterID}/{serverID}
	APIsKey     = "/apis/"     // APIsKey configs of API config /apis/{apiCfgID}
	RoutingsKey = "/routings/" // RoutingsKey configs of routing /routings/{routingID}

	CacheKey          = "/plugins/cache/" // CacheKey "cache/" means the root key
	RbacKey           = "/plugins/rbac/"  // RbacKey "rbac/" means the root key, it contains "users/" "roles/" "permissions/"
	RbacUsersKey      = "users/"          // RbacUsersKey users/
	RbacRolesKey      = "roles/"          // RbacRolesKey roles/
	RbacPermissionKey = "permission/"     // RbacPermissionKey permission/
)

/*
 * /plugins/roles/{roleID} is the key to value of Role{ID: roleID, Permissions: PermissionList}
 */
