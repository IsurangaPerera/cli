package main

import (
	
	"./soap"
	"errors"
	"encoding/xml"
	"./wsdlgo"
	"context"
)

var Namespace = "http://mgt.user.carbon.wso2.org"

// userAdminPortType implements the UserAdminPortType interface.
type userAdminPortType struct {
	cli *soap.Client
}

// NewUserAdminPortType creates an initializes a UserAdminPortType.
func NewUserAdminPortType(cli *soap.Client) UserAdminPortType {
	return &userAdminPortType{cli}
}

// UserAdminPortType was auto-generated from WSDL
type UserAdminPortType interface {
	// GetAllPermittedRoleNames was auto-generated from WSDL.
	//GetAllPermittedRoleNames(α *GetAllPermittedRoleNames) (β *GetAllPermittedRoleNamesResponse, err error)

	// GetAllRolesNames was auto-generated from WSDL.
	GetAllRolesNames(α *wsdlgo.GetAllRolesNames) (β *wsdlgo.GetAllRolesNamesResponse, err error)

	// GetAllSharedRoleNames was auto-generated from WSDL.
	GetAllSharedRoleNames(α *wsdlgo.GetAllSharedRoleNames) (β *wsdlgo.GetAllSharedRoleNamesResponse, err error)

	// GetAllUIPermissions was auto-generated from WSDL.
	GetAllUIPermissions(α *wsdlgo.GetAllUIPermissions) (β *wsdlgo.GetAllUIPermissionsResponse, err error)

	// GetRolePermissions was auto-generated from WSDL.
	GetRolePermissions(α *wsdlgo.GetRolePermissions) (β *wsdlgo.GetRolePermissionsResponse, err error)

	// GetRolesOfCurrentUser was auto-generated from WSDL.
	GetRolesOfCurrentUser(α *wsdlgo.GetRolesOfCurrentUser) (β *wsdlgo.GetRolesOfCurrentUserResponse, err error)

	// GetRolesOfUser was auto-generated from WSDL.
	GetRolesOfUser(α *wsdlgo.GetRolesOfUser) (β *wsdlgo.GetRolesOfUserResponse, err error)

	// GetUserRealmInfo was auto-generated from WSDL.
	GetUserRealmInfo(α *wsdlgo.GetUserRealmInfo) (β *wsdlgo.GetUserRealmInfoResponse, err error)

	// GetUsersOfRole was auto-generated from WSDL.
	GetUsersOfRole(α *wsdlgo.GetUsersOfRole) (β *wsdlgo.GetUsersOfRoleResponse, err error)

	// HasMultipleUserStores was auto-generated from WSDL.
	HasMultipleUserStores(α *wsdlgo.HasMultipleUserStores) (β *wsdlgo.HasMultipleUserStoresResponse, err error)

	// IsSharedRolesEnabled was auto-generated from WSDL.
	IsSharedRolesEnabled(α *wsdlgo.IsSharedRolesEnabled) (β *wsdlgo.IsSharedRolesEnabledResponse, err error)

	// ListAllUsers was auto-generated from WSDL.
	ListAllUsers(α *wsdlgo.ListAllUsers) (β *wsdlgo.ListAllUsersResponse, err error)

	// ListAllUsersWithPermission was auto-generated from WSDL.
	ListAllUsersWithPermission(α *wsdlgo.ListAllUsersWithPermission) (β *wsdlgo.ListAllUsersWithPermissionResponse, err error)

	// ListUserByClaim was auto-generated from WSDL.
	ListUserByClaim(α *wsdlgo.ListUserByClaim) (β *wsdlgo.ListUserByClaimResponse, err error)

	// ListUserByClaimWithPermission was auto-generated from WSDL.
	ListUserByClaimWithPermission(α *wsdlgo.ListUserByClaimWithPermission) (β *wsdlgo.ListUserByClaimWithPermissionResponse, err error)

	// ListUsers was auto-generated from WSDL.
	ListUsers(α *wsdlgo.ListUsers) (β *wsdlgo.ListUsersResponse, err error)
}

// AddInternalRole was auto-generated from WSDL.
func AddInternalRole(ctx context.Context, parameters *wsdlgo.AddInternalRole) (err error) {
	return errors.New("not implemented")
}

// AddRemoveRolesOfUser was auto-generated from WSDL.
func AddRemoveRolesOfUser(ctx context.Context, parameters *wsdlgo.AddRemoveRolesOfUser) (err error) {
	return errors.New("not implemented")
}

// AddRemoveUsersOfRole was auto-generated from WSDL.
func AddRemoveUsersOfRole(ctx context.Context, parameters *wsdlgo.AddRemoveUsersOfRole) (err error) {
	return errors.New("not implemented")
}

// AddRole was auto-generated from WSDL.
func AddRole(ctx context.Context, parameters *wsdlgo.AddRole) (err error) {
	return errors.New("not implemented")
}

// AddUser was auto-generated from WSDL.
func AddUser(ctx context.Context, parameters *wsdlgo.AddUser) (err error) {
	return errors.New("not implemented")
}

// BulkImportUsers was auto-generated from WSDL.
func BulkImportUsers(ctx context.Context, parameters *wsdlgo.BulkImportUsers) (err error) {
	return errors.New("not implemented")
}

// ChangePassword was auto-generated from WSDL.
func ChangePassword(ctx context.Context, parameters *wsdlgo.ChangePassword) (err error) {
	return errors.New("not implemented")
}

// ChangePasswordByUser was auto-generated from WSDL.
func ChangePasswordByUser(ctx context.Context, parameters *wsdlgo.ChangePasswordByUser) (err error) {
	return errors.New("not implemented")
}

// DeleteRole was auto-generated from WSDL.
func DeleteRole(ctx context.Context, parameters *wsdlgo.DeleteRole) (err error) {
	return errors.New("not implemented")
}

// DeleteUser was auto-generated from WSDL.
func DeleteUser(ctx context.Context, parameters *wsdlgo.DeleteUser) (err error) {
	return errors.New("not implemented")
}

// GetAllPermittedRoleNames was auto-generated from WSDL.
func (p *userAdminPortType) GetAllPermittedRoleNames(α *wsdlgo.GetAllPermittedRoleNames) (β *wsdlgo.GetAllPermittedRoleNamesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetAllPermittedRoleNamesResponse `xml:"getAllPermittedRoleNamesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAllRolesNames was auto-generated from WSDL.
func (p *userAdminPortType) GetAllRolesNames(α *wsdlgo.GetAllRolesNames) (β *wsdlgo.GetAllRolesNamesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetAllRolesNamesResponse `xml:"getAllRolesNamesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAllSharedRoleNames was auto-generated from WSDL.
func (p *userAdminPortType) GetAllSharedRoleNames(α *wsdlgo.GetAllSharedRoleNames) (β *wsdlgo.GetAllSharedRoleNamesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetAllSharedRoleNamesResponse `xml:"getAllSharedRoleNamesResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetAllUIPermissions was auto-generated from WSDL.
func (p *userAdminPortType) GetAllUIPermissions(α *wsdlgo.GetAllUIPermissions) (β *wsdlgo.GetAllUIPermissionsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetAllUIPermissionsResponse `xml:"getAllUIPermissionsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRolePermissions was auto-generated from WSDL.
func (p *userAdminPortType) GetRolePermissions(α *wsdlgo.GetRolePermissions) (β *wsdlgo.GetRolePermissionsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetRolePermissionsResponse `xml:"getRolePermissionsResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRolesOfCurrentUser was auto-generated from WSDL.
func (p *userAdminPortType) GetRolesOfCurrentUser(α *wsdlgo.GetRolesOfCurrentUser) (β *wsdlgo.GetRolesOfCurrentUserResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetRolesOfCurrentUserResponse `xml:"getRolesOfCurrentUserResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRolesOfUser was auto-generated from WSDL.
func (p *userAdminPortType) GetRolesOfUser(α *wsdlgo.GetRolesOfUser) (β *wsdlgo.GetRolesOfUserResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetRolesOfUserResponse `xml:"getRolesOfUserResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetUserRealmInfo was auto-generated from WSDL.
func (p *userAdminPortType) GetUserRealmInfo(α *wsdlgo.GetUserRealmInfo) (β *wsdlgo.GetUserRealmInfoResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetUserRealmInfoResponse `xml:"getUserRealmInfoResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetUsersOfRole was auto-generated from WSDL.
func (p *userAdminPortType) GetUsersOfRole(α *wsdlgo.GetUsersOfRole) (β *wsdlgo.GetUsersOfRoleResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetUsersOfRoleResponse `xml:"getUsersOfRoleResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// HasMultipleUserStores was auto-generated from WSDL.
func (p *userAdminPortType) HasMultipleUserStores(α *wsdlgo.HasMultipleUserStores) (β *wsdlgo.HasMultipleUserStoresResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.HasMultipleUserStoresResponse `xml:"hasMultipleUserStoresResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// IsSharedRolesEnabled was auto-generated from WSDL.
func (p *userAdminPortType) IsSharedRolesEnabled(α *wsdlgo.IsSharedRolesEnabled) (β *wsdlgo.IsSharedRolesEnabledResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.IsSharedRolesEnabledResponse `xml:"isSharedRolesEnabledResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ListAllUsers was auto-generated from WSDL.
func (p *userAdminPortType) ListAllUsers(α *wsdlgo.ListAllUsers) (β *wsdlgo.ListAllUsersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.ListAllUsersResponse `xml:"listAllUsersResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ListAllUsersWithPermission was auto-generated from WSDL.
func (p *userAdminPortType) ListAllUsersWithPermission(α *wsdlgo.ListAllUsersWithPermission) (β *wsdlgo.ListAllUsersWithPermissionResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.ListAllUsersWithPermissionResponse `xml:"listAllUsersWithPermissionResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ListUserByClaim was auto-generated from WSDL.
func (p *userAdminPortType) ListUserByClaim(α *wsdlgo.ListUserByClaim) (β *wsdlgo.ListUserByClaimResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.ListUserByClaimResponse `xml:"listUserByClaimResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ListUserByClaimWithPermission was auto-generated from WSDL.
func (p *userAdminPortType) ListUserByClaimWithPermission(α *wsdlgo.ListUserByClaimWithPermission) (β *wsdlgo.ListUserByClaimWithPermissionResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.ListUserByClaimWithPermissionResponse `xml:"listUserByClaimWithPermissionResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// ListUsers was auto-generated from WSDL.
func (p *userAdminPortType) ListUsers(α *wsdlgo.ListUsers) (β *wsdlgo.ListUsersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.ListUsersResponse `xml:"listUsersResponse"`
		}
	}{}
	if err = p.cli.RoundTrip(α, &γ); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// SetRoleUIPermission was auto-generated from WSDL.
func SetRoleUIPermission(ctx context.Context, parameters *wsdlgo.SetRoleUIPermission) (err error) {
	return errors.New("not implemented")
}

// UpdateRoleName was auto-generated from WSDL.
func UpdateRoleName(ctx context.Context, parameters *wsdlgo.UpdateRoleName) (err error) {
	return errors.New("not implemented")
}

// UpdateRolesOfUser was auto-generated from WSDL.
func UpdateRolesOfUser(ctx context.Context, parameters *wsdlgo.UpdateRolesOfUser) (err error) {
	return errors.New("not implemented")
}

// UpdateUsersOfRole was auto-generated from WSDL.
func UpdateUsersOfRole(ctx context.Context, parameters *wsdlgo.UpdateUsersOfRole) (err error) {
	return errors.New("not implemented")
}