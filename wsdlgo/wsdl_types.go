package wsdlgo

import (

	"encoding/xml"
)

// ClaimValue was auto-generated from WSDL.
type ClaimValue struct {
	ClaimURI string `xml:"claimURI,omitempty" json:"claimURI,omitempty" yaml:"claimURI,omitempty"`
	Value    string `xml:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty"`
}

// FlaggedName was auto-generated from WSDL.
type FlaggedName struct {
	Dn              string `xml:"dn,omitempty" json:"dn,omitempty" yaml:"dn,omitempty"`
	DomainName      string `xml:"domainName,omitempty" json:"domainName,omitempty" yaml:"domainName,omitempty"`
	Editable        bool   `xml:"editable,omitempty" json:"editable,omitempty" yaml:"editable,omitempty"`
	ItemDisplayName string `xml:"itemDisplayName,omitempty" json:"itemDisplayName,omitempty" yaml:"itemDisplayName,omitempty"`
	ItemName        string `xml:"itemName,omitempty" json:"itemName,omitempty" yaml:"itemName,omitempty"`
	ReadOnly        bool   `xml:"readOnly,omitempty" json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	RoleType        string `xml:"roleType,omitempty" json:"roleType,omitempty" yaml:"roleType,omitempty"`
	Selected        bool   `xml:"selected,omitempty" json:"selected,omitempty" yaml:"selected,omitempty"`
	Shared          bool   `xml:"shared,omitempty" json:"shared,omitempty" yaml:"shared,omitempty"`
}

// UIPermissionNode was auto-generated from WSDL.
type UIPermissionNode struct {
	DisplayName  string              `xml:"displayName,omitempty" json:"displayName,omitempty" yaml:"displayName,omitempty"`
	NodeList     []*UIPermissionNode `xml:"nodeList,omitempty" json:"nodeList,omitempty" yaml:"nodeList,omitempty"`
	ResourcePath string              `xml:"resourcePath,omitempty" json:"resourcePath,omitempty" yaml:"resourcePath,omitempty"`
	Selected     bool                `xml:"selected,omitempty" json:"selected,omitempty" yaml:"selected,omitempty"`
}

// UserAdminException was auto-generated from WSDL.
type UserAdminException struct {
	Message string `xml:"message,omitempty" json:"message,omitempty" yaml:"message,omitempty"`
}

// UserAdminUserAdminException was auto-generated from WSDL.
type UserAdminUserAdminException struct {
	UserAdminException *UserAdminException `xml:"UserAdminException,omitempty" json:"UserAdminException,omitempty" yaml:"UserAdminException,omitempty"`
}

// UserRealmInfo was auto-generated from WSDL.
type UserRealmInfo struct {
	AdminRole            string           `xml:"adminRole,omitempty" json:"adminRole,omitempty" yaml:"adminRole,omitempty"`
	AdminUser            string           `xml:"adminUser,omitempty" json:"adminUser,omitempty" yaml:"adminUser,omitempty"`
	BulkImportSupported  bool             `xml:"bulkImportSupported,omitempty" json:"bulkImportSupported,omitempty" yaml:"bulkImportSupported,omitempty"`
	DefaultUserClaims    []string         `xml:"defaultUserClaims,omitempty" json:"defaultUserClaims,omitempty" yaml:"defaultUserClaims,omitempty"`
	DomainNames          []string         `xml:"domainNames,omitempty" json:"domainNames,omitempty" yaml:"domainNames,omitempty"`
	EnableUIPageCache    bool             `xml:"enableUIPageCache,omitempty" json:"enableUIPageCache,omitempty" yaml:"enableUIPageCache,omitempty"`
	EveryOneRole         string           `xml:"everyOneRole,omitempty" json:"everyOneRole,omitempty" yaml:"everyOneRole,omitempty"`
	MaxItemsPerUIPage    int              `xml:"maxItemsPerUIPage,omitempty" json:"maxItemsPerUIPage,omitempty" yaml:"maxItemsPerUIPage,omitempty"`
	MaxUIPagesInCache    int              `xml:"maxUIPagesInCache,omitempty" json:"maxUIPagesInCache,omitempty" yaml:"maxUIPagesInCache,omitempty"`
	MultipleUserStore    bool             `xml:"multipleUserStore,omitempty" json:"multipleUserStore,omitempty" yaml:"multipleUserStore,omitempty"`
	PrimaryUserStoreInfo *UserStoreInfo   `xml:"primaryUserStoreInfo,omitempty" json:"primaryUserStoreInfo,omitempty" yaml:"primaryUserStoreInfo,omitempty"`
	RequiredUserClaims   []string         `xml:"requiredUserClaims,omitempty" json:"requiredUserClaims,omitempty" yaml:"requiredUserClaims,omitempty"`
	UserClaims           []string         `xml:"userClaims,omitempty" json:"userClaims,omitempty" yaml:"userClaims,omitempty"`
	UserStoresInfo       []*UserStoreInfo `xml:"userStoresInfo,omitempty" json:"userStoresInfo,omitempty" yaml:"userStoresInfo,omitempty"`
}

// UserStoreInfo was auto-generated from WSDL.
type UserStoreInfo struct {
	BulkImportSupported            bool   `xml:"bulkImportSupported,omitempty" json:"bulkImportSupported,omitempty" yaml:"bulkImportSupported,omitempty"`
	CaseSensitiveUsername          bool   `xml:"caseSensitiveUsername,omitempty" json:"caseSensitiveUsername,omitempty" yaml:"caseSensitiveUsername,omitempty"`
	DomainName                     string `xml:"domainName,omitempty" json:"domainName,omitempty" yaml:"domainName,omitempty"`
	ExternalIdP                    string `xml:"externalIdP,omitempty" json:"externalIdP,omitempty" yaml:"externalIdP,omitempty"`
	MaxRoleLimit                   int    `xml:"maxRoleLimit,omitempty" json:"maxRoleLimit,omitempty" yaml:"maxRoleLimit,omitempty"`
	MaxUserLimit                   int    `xml:"maxUserLimit,omitempty" json:"maxUserLimit,omitempty" yaml:"maxUserLimit,omitempty"`
	PasswordRegEx                  string `xml:"passwordRegEx,omitempty" json:"passwordRegEx,omitempty" yaml:"passwordRegEx,omitempty"`
	PasswordRegExViolationErrorMsg string `xml:"passwordRegExViolationErrorMsg,omitempty" json:"passwordRegExViolationErrorMsg,omitempty" yaml:"passwordRegExViolationErrorMsg,omitempty"`
	PasswordsExternallyManaged     bool   `xml:"passwordsExternallyManaged,omitempty" json:"passwordsExternallyManaged,omitempty" yaml:"passwordsExternallyManaged,omitempty"`
	ReadGroupsEnabled              bool   `xml:"readGroupsEnabled,omitempty" json:"readGroupsEnabled,omitempty" yaml:"readGroupsEnabled,omitempty"`
	ReadOnly                       bool   `xml:"readOnly,omitempty" json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	RoleNameRegEx                  string `xml:"roleNameRegEx,omitempty" json:"roleNameRegEx,omitempty" yaml:"roleNameRegEx,omitempty"`
	UserNameRegEx                  string `xml:"userNameRegEx,omitempty" json:"userNameRegEx,omitempty" yaml:"userNameRegEx,omitempty"`
	UsernameRegExViolationErrorMsg string `xml:"usernameRegExViolationErrorMsg,omitempty" json:"usernameRegExViolationErrorMsg,omitempty" yaml:"usernameRegExViolationErrorMsg,omitempty"`
	WriteGroupsEnabled             bool   `xml:"writeGroupsEnabled,omitempty" json:"writeGroupsEnabled,omitempty" yaml:"writeGroupsEnabled,omitempty"`
}

// AddInternalRole was auto-generated from WSDL.
type AddInternalRole struct {
	XMLName     xml.Name `xml:"http://mgt.user.carbon.wso2.org addInternalRole" json:"-" yaml:"-"`
	RoleName    string   `xml:"roleName,omitempty" json:"roleName,omitempty" yaml:"roleName,omitempty"`
	UserList    []string `xml:"userList,omitempty" json:"userList,omitempty" yaml:"userList,omitempty"`
	Permissions []string `xml:"permissions,omitempty" json:"permissions,omitempty" yaml:"permissions,omitempty"`
}

// AddRemoveRolesOfUser was auto-generated from WSDL.
type AddRemoveRolesOfUser struct {
	XMLName      xml.Name `xml:"http://mgt.user.carbon.wso2.org addRemoveRolesOfUser" json:"-" yaml:"-"`
	UserName     string   `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	NewRoles     []string `xml:"newRoles,omitempty" json:"newRoles,omitempty" yaml:"newRoles,omitempty"`
	DeletedRoles []string `xml:"deletedRoles,omitempty" json:"deletedRoles,omitempty" yaml:"deletedRoles,omitempty"`
}

// AddRemoveUsersOfRole was auto-generated from WSDL.
type AddRemoveUsersOfRole struct {
	XMLName      xml.Name `xml:"http://mgt.user.carbon.wso2.org addRemoveUsersOfRole" json:"-" yaml:"-"`
	RoleName     string   `xml:"roleName,omitempty" json:"roleName,omitempty" yaml:"roleName,omitempty"`
	NewUsers     []string `xml:"newUsers,omitempty" json:"newUsers,omitempty" yaml:"newUsers,omitempty"`
	DeletedUsers []string `xml:"deletedUsers,omitempty" json:"deletedUsers,omitempty" yaml:"deletedUsers,omitempty"`
}

// AddRole was auto-generated from WSDL.
type AddRole struct {
	XMLName      xml.Name `xml:"http://mgt.user.carbon.wso2.org addRole" json:"-" yaml:"-"`
	RoleName     string   `xml:"roleName,omitempty" json:"roleName,omitempty" yaml:"roleName,omitempty"`
	UserList     []string `xml:"userList,omitempty" json:"userList,omitempty" yaml:"userList,omitempty"`
	Permissions  []string `xml:"permissions,omitempty" json:"permissions,omitempty" yaml:"permissions,omitempty"`
	IsSharedRole bool     `xml:"isSharedRole,omitempty" json:"isSharedRole,omitempty" yaml:"isSharedRole,omitempty"`
}

// AddUser was auto-generated from WSDL.
type AddUser struct {
	XMLName     xml.Name      `xml:"http://mgt.user.carbon.wso2.org addUser" json:"-" yaml:"-"`
	UserName    string        `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	Password    string        `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Roles       []string      `xml:"roles,omitempty" json:"roles,omitempty" yaml:"roles,omitempty"`
	Claims      []*ClaimValue `xml:"claims,omitempty" json:"claims,omitempty" yaml:"claims,omitempty"`
	ProfileName string        `xml:"profileName,omitempty" json:"profileName,omitempty" yaml:"profileName,omitempty"`
}

// BulkImportUsers was auto-generated from WSDL.
type BulkImportUsers struct {
	XMLName         xml.Name `xml:"http://mgt.user.carbon.wso2.org bulkImportUsers" json:"-" yaml:"-"`
	UserStoreDomain string   `xml:"userStoreDomain,omitempty" json:"userStoreDomain,omitempty" yaml:"userStoreDomain,omitempty"`
	FileName        string   `xml:"fileName,omitempty" json:"fileName,omitempty" yaml:"fileName,omitempty"`
	Handler         []byte   `xml:"handler,omitempty" json:"handler,omitempty" yaml:"handler,omitempty"`
	DefaultPassword string   `xml:"defaultPassword,omitempty" json:"defaultPassword,omitempty" yaml:"defaultPassword,omitempty"`
}

// ChangePassword was auto-generated from WSDL.
type ChangePassword struct {
	XMLName     xml.Name `xml:"http://mgt.user.carbon.wso2.org changePassword" json:"-" yaml:"-"`
	UserName    string   `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	NewPassword string   `xml:"newPassword,omitempty" json:"newPassword,omitempty" yaml:"newPassword,omitempty"`
}

// ChangePasswordByUser was auto-generated from WSDL.
type ChangePasswordByUser struct {
	XMLName     xml.Name `xml:"http://mgt.user.carbon.wso2.org changePasswordByUser" json:"-" yaml:"-"`
	UserName    string   `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	OldPassword string   `xml:"oldPassword,omitempty" json:"oldPassword,omitempty" yaml:"oldPassword,omitempty"`
	NewPassword string   `xml:"newPassword,omitempty" json:"newPassword,omitempty" yaml:"newPassword,omitempty"`
}

// DeleteRole was auto-generated from WSDL.
type DeleteRole struct {
	XMLName  xml.Name `xml:"http://mgt.user.carbon.wso2.org deleteRole" json:"-" yaml:"-"`
	RoleName string   `xml:"roleName,omitempty" json:"roleName,omitempty" yaml:"roleName,omitempty"`
}

// DeleteUser was auto-generated from WSDL.
type DeleteUser struct {
	XMLName  xml.Name `xml:"http://mgt.user.carbon.wso2.org deleteUser" json:"-" yaml:"-"`
	UserName string   `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
}

// GetAllPermittedRoleNames was auto-generated from WSDL.
type GetAllPermittedRoleNames struct {
	XMLName    xml.Name `xml:"http://mgt.user.carbon.wso2.org getAllPermittedRoleNames" json:"-" yaml:"-"`
	Filter     string   `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Permission string   `xml:"permission,omitempty" json:"permission,omitempty" yaml:"permission,omitempty"`
	Limit      int      `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// GetAllPermittedRoleNamesResponse was auto-generated from WSDL.
type GetAllPermittedRoleNamesResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAllRolesNames was auto-generated from WSDL.
type GetAllRolesNames struct {
	XMLName xml.Name `xml:"http://mgt.user.carbon.wso2.org getAllRolesNames" json:"-" yaml:"-"`
	Filter  string   `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Limit   int      `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// GetAllRolesNamesResponse was auto-generated from WSDL.
type GetAllRolesNamesResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAllSharedRoleNames was auto-generated from WSDL.
type GetAllSharedRoleNames struct {
	XMLName xml.Name `xml:"http://mgt.user.carbon.wso2.org getAllSharedRoleNames" json:"-" yaml:"-"`
	Filter  string   `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Limit   int      `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// GetAllSharedRoleNamesResponse was auto-generated from WSDL.
type GetAllSharedRoleNamesResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetAllUIPermissions was auto-generated from WSDL.
type GetAllUIPermissions struct{}

// GetAllUIPermissionsResponse was auto-generated from WSDL.
type GetAllUIPermissionsResponse struct {
	Return *UIPermissionNode `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetRolePermissions was auto-generated from WSDL.
type GetRolePermissions struct {
	XMLName  xml.Name `xml:"http://mgt.user.carbon.wso2.org getRolePermissions" json:"-" yaml:"-"`
	RoleName string   `xml:"roleName,omitempty" json:"roleName,omitempty" yaml:"roleName,omitempty"`
}

// GetRolePermissionsResponse was auto-generated from WSDL.
type GetRolePermissionsResponse struct {
	Return *UIPermissionNode `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetRolesOfCurrentUser was auto-generated from WSDL.
type GetRolesOfCurrentUser struct{}

// GetRolesOfCurrentUserResponse was auto-generated from WSDL.
type GetRolesOfCurrentUserResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetRolesOfUser was auto-generated from WSDL.
type GetRolesOfUser struct {
	XMLName  xml.Name `xml:"http://mgt.user.carbon.wso2.org getRolesOfUser" json:"-" yaml:"-"`
	UserName string   `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	Filter   string   `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Limit    int      `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// GetRolesOfUserResponse was auto-generated from WSDL.
type GetRolesOfUserResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetUserRealmInfo was auto-generated from WSDL.
type GetUserRealmInfo struct{}

// GetUserRealmInfoResponse was auto-generated from WSDL.
type GetUserRealmInfoResponse struct {
	Return *UserRealmInfo `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// GetUsersOfRole was auto-generated from WSDL.
type GetUsersOfRole struct {
	XMLName  xml.Name `xml:"http://mgt.user.carbon.wso2.org getUsersOfRole" json:"-" yaml:"-"`
	RoleName string   `xml:"roleName,omitempty" json:"roleName,omitempty" yaml:"roleName,omitempty"`
	Filter   string   `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Limit    int      `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// GetUsersOfRoleResponse was auto-generated from WSDL.
type GetUsersOfRoleResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// HasMultipleUserStores was auto-generated from WSDL.
type HasMultipleUserStores struct{}

// HasMultipleUserStoresResponse was auto-generated from WSDL.
type HasMultipleUserStoresResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// IsSharedRolesEnabled was auto-generated from WSDL.
type IsSharedRolesEnabled struct{}

// IsSharedRolesEnabledResponse was auto-generated from WSDL.
type IsSharedRolesEnabledResponse struct {
	Return bool `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ListAllUsers was auto-generated from WSDL.
type ListAllUsers struct {
	XMLName xml.Name `xml:"xsd:listAllUsers" json:"-" yaml:"-"`
	Filter  string   `xml:"xsd:filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Limit   int      `xml:"xsd:limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// ListAllUsersResponse was auto-generated from WSDL.
type ListAllUsersResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ListAllUsersWithPermission was auto-generated from WSDL.
type ListAllUsersWithPermission struct {
	XMLName    xml.Name `xml:"http://mgt.user.carbon.wso2.org listAllUsersWithPermission" json:"-" yaml:"-"`
	Filter     string   `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Permission string   `xml:"permission,omitempty" json:"permission,omitempty" yaml:"permission,omitempty"`
	Limit      int      `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// ListAllUsersWithPermissionResponse was auto-generated from WSDL.
type ListAllUsersWithPermissionResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ListUserByClaim was auto-generated from WSDL.
type ListUserByClaim struct {
	XMLName    xml.Name    `xml:"http://mgt.user.carbon.wso2.org listUserByClaim" json:"-" yaml:"-"`
	ClaimValue *ClaimValue `xml:"claimValue,omitempty" json:"claimValue,omitempty" yaml:"claimValue,omitempty"`
	Filter     string      `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	MaxLimit   int         `xml:"maxLimit,omitempty" json:"maxLimit,omitempty" yaml:"maxLimit,omitempty"`
}

// ListUserByClaimResponse was auto-generated from WSDL.
type ListUserByClaimResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ListUserByClaimWithPermission was auto-generated from WSDL.
type ListUserByClaimWithPermission struct {
	XMLName    xml.Name    `xml:"http://mgt.user.carbon.wso2.org listUserByClaimWithPermission" json:"-" yaml:"-"`
	ClaimValue *ClaimValue `xml:"claimValue,omitempty" json:"claimValue,omitempty" yaml:"claimValue,omitempty"`
	Filter     string      `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Permission string      `xml:"permission,omitempty" json:"permission,omitempty" yaml:"permission,omitempty"`
	MaxLimit   int         `xml:"maxLimit,omitempty" json:"maxLimit,omitempty" yaml:"maxLimit,omitempty"`
}

// ListUserByClaimWithPermissionResponse was auto-generated from
// WSDL.
type ListUserByClaimWithPermissionResponse struct {
	Return []*FlaggedName `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// ListUsers was auto-generated from WSDL.
type ListUsers struct {
	XMLName xml.Name `xml:"http://mgt.user.carbon.wso2.org listUsers" json:"-" yaml:"-"`
	Filter  string   `xml:"filter,omitempty" json:"filter,omitempty" yaml:"filter,omitempty"`
	Limit   int      `xml:"limit,omitempty" json:"limit,omitempty" yaml:"limit,omitempty"`
}

// ListUsersResponse was auto-generated from WSDL.
type ListUsersResponse struct {
	Return []string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// SetRoleUIPermission was auto-generated from WSDL.
type SetRoleUIPermission struct {
	XMLName      xml.Name `xml:"http://mgt.user.carbon.wso2.org setRoleUIPermission" json:"-" yaml:"-"`
	RoleName     string   `xml:"roleName,omitempty" json:"roleName,omitempty" yaml:"roleName,omitempty"`
	RawResources []string `xml:"rawResources,omitempty" json:"rawResources,omitempty" yaml:"rawResources,omitempty"`
}

// UpdateRoleName was auto-generated from WSDL.
type UpdateRoleName struct {
	XMLName     xml.Name `xml:"http://mgt.user.carbon.wso2.org updateRoleName" json:"-" yaml:"-"`
	RoleName    string   `xml:"roleName,omitempty" json:"roleName,omitempty" yaml:"roleName,omitempty"`
	NewRoleName string   `xml:"newRoleName,omitempty" json:"newRoleName,omitempty" yaml:"newRoleName,omitempty"`
}

// UpdateRolesOfUser was auto-generated from WSDL.
type UpdateRolesOfUser struct {
	XMLName     xml.Name `xml:"http://mgt.user.carbon.wso2.org updateRolesOfUser" json:"-" yaml:"-"`
	UserName    string   `xml:"userName,omitempty" json:"userName,omitempty" yaml:"userName,omitempty"`
	NewRoleList []string `xml:"newRoleList,omitempty" json:"newRoleList,omitempty" yaml:"newRoleList,omitempty"`
}

// UpdateUsersOfRole was auto-generated from WSDL.
type UpdateUsersOfRole struct {
	XMLName  xml.Name       `xml:"http://mgt.user.carbon.wso2.org updateUsersOfRole" json:"-" yaml:"-"`
	RoleName string         `xml:"roleName,omitempty" json:"roleName,omitempty" yaml:"roleName,omitempty"`
	UserList []*FlaggedName `xml:"userList,omitempty" json:"userList,omitempty" yaml:"userList,omitempty"`
}