package parser

import (

"../soap"
"errors"
"encoding/xml"
"../wsdlgo"
"context"
"fmt"
"reflect"
)

var Namespace = "http://mgt.user.carbon.wso2.org"

type UserAdminPortType struct {
	Cli soap.Client
}

/*func NewUserAdminPortType(cli soap.Client) UserAdminType {
	return &UserAdminPortType{cli}
}*/

type UserAdminType interface {

	GetAllPermittedRoleNames(α *wsdlgo.GetAllPermittedRoleNames) (β *wsdlgo.GetAllPermittedRoleNamesResponse, err error)
	GetAllRolesNames(α *wsdlgo.GetAllRolesNames) (β *wsdlgo.GetAllRolesNamesResponse, err error)
	GetAllSharedRoleNames(α *wsdlgo.GetAllSharedRoleNames) (β *wsdlgo.GetAllSharedRoleNamesResponse, err error)
	GetAllUIPermissions(α *wsdlgo.GetAllUIPermissions) (β *wsdlgo.GetAllUIPermissionsResponse, err error)
	GetRolePermissions(α *wsdlgo.GetRolePermissions) (β *wsdlgo.GetRolePermissionsResponse, err error)
	GetRolesOfCurrentUser(α *wsdlgo.GetRolesOfCurrentUser) (β *wsdlgo.GetRolesOfCurrentUserResponse, err error)
	GetRolesOfUser(α *wsdlgo.GetRolesOfUser) (β *wsdlgo.GetRolesOfUserResponse, err error)
	GetUserRealmInfo(α *wsdlgo.GetUserRealmInfo) (β *wsdlgo.GetUserRealmInfoResponse, err error)
	GetUsersOfRole(α *wsdlgo.GetUsersOfRole) (β *wsdlgo.GetUsersOfRoleResponse, err error)
	HasMultipleUserStores(α *wsdlgo.HasMultipleUserStores) (β *wsdlgo.HasMultipleUserStoresResponse, err error)
	IsSharedRolesEnabled(α *wsdlgo.IsSharedRolesEnabled) (β *wsdlgo.IsSharedRolesEnabledResponse, err error)
	ListAllUsers(α *wsdlgo.ListAllUsers) (β *wsdlgo.ListAllUsersResponse, err error)
	ListAllUsersWithPermission(α *wsdlgo.ListAllUsersWithPermission) (β *wsdlgo.ListAllUsersWithPermissionResponse, err error)
	ListUserByClaim(α *wsdlgo.ListUserByClaim) (β *wsdlgo.ListUserByClaimResponse, err error)
	ListUserByClaimWithPermission(α *wsdlgo.ListUserByClaimWithPermission) (β *wsdlgo.ListUserByClaimWithPermissionResponse, err error)
	ListUsers(α *wsdlgo.ListUsers) (β *wsdlgo.ListUsersResponse, err error)
}

// AddInternalRole
func (p *UserAdminPortType) AddInternalRole(α *wsdlgo.AddInternalRole) (err error){
	
	if err = p.Cli.OneWayTrip(α, "addInternalRole"); err != nil {
		return err
	}
	return nil
}

// AddRemoveRolesOfUser
func (p *UserAdminPortType) AddRemoveRolesOfUser(α *wsdlgo.AddRemoveRolesOfUser) (err error) {
	
	if err = p.Cli.OneWayTrip(α, "addRemoveRolesOfUser"); err != nil {
		return err
	}
	return nil
}

// AddRemoveUsersOfRole
func (p *UserAdminPortType) AddRemoveUsersOfRole(α *wsdlgo.AddRemoveUsersOfRole) (err error) {
	
	if err = p.Cli.OneWayTrip(α, "addRemoveUsersOfRole"); err != nil {
		return err
	}
	return nil
}

// AddRole
func (p *UserAdminPortType) AddRole(α *wsdlgo.AddRole) (err error) {
	
	if err = p.Cli.OneWayTrip(α, "addRole"); err != nil {
		return err
	}
	return nil
}

// AddUser
func (p *UserAdminPortType) AddUser(ctx context.Context, parameters *wsdlgo.AddUser) (err error) {
	return errors.New("not implemented")
}

// BulkImportUsers
func (p *UserAdminPortType) BulkImportUsers(ctx context.Context, parameters *wsdlgo.BulkImportUsers) (err error) {
	return errors.New("not implemented")
}

// ChangePassword
func (p *UserAdminPortType) ChangePassword(ctx context.Context, parameters *wsdlgo.ChangePassword) (err error) {
	return errors.New("not implemented")
}

// ChangePasswordByUser
func (p *UserAdminPortType) ChangePasswordByUser(ctx context.Context, parameters *wsdlgo.ChangePasswordByUser) (err error) {
	return errors.New("not implemented")
}

// DeleteRole
func (p *UserAdminPortType) DeleteRole(α *wsdlgo.DeleteRole) (err error) {
	
	if err = p.Cli.OneWayTrip(α, "deleteRole"); err != nil {
		return err
	}
	return nil
}

// DeleteUser
func (p *UserAdminPortType) DeleteUser(α *wsdlgo.DeleteUser) (err error) {
	
	if err = p.Cli.OneWayTrip(α, "deleteUser"); err != nil {
		return err
	}
	return nil
}

// GetAllPermittedRoleNames
/*func (p *UserAdminPortType) GetAllPermittedRoleNames(α *wsdlgo.GetAllPermittedRoleNames) (β *wsdlgo.GetAllPermittedRoleNamesResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetAllPermittedRoleNamesResponse `xml:"getAllPermittedRoleNamesResponse"`
		}
		}{}
		if err = p.Cli.RoundTrip(α, &γ); err != nil {
			return nil, err
		}
		return &γ.Body.M, nil
	}

// GetAllRolesNames
	func (p *UserAdminPortType) GetAllRolesNames(α *wsdlgo.GetAllRolesNames) (β *wsdlgo.GetAllRolesNamesResponse, err error) {
		γ := struct {
			XMLName xml.Name `xml:"Envelope"`
			Body    struct {
				M wsdlgo.GetAllRolesNamesResponse `xml:"getAllRolesNamesResponse"`
			}
			}{}
			if err = p.Cli.RoundTrip(α, &γ); err != nil {
				return nil, err
			}
			return &γ.Body.M, nil
		}

// GetAllSharedRoleNames
		func (p *UserAdminPortType) GetAllSharedRoleNames(α *wsdlgo.GetAllSharedRoleNames) (β *wsdlgo.GetAllSharedRoleNamesResponse, err error) {
			γ := struct {
				XMLName xml.Name `xml:"Envelope"`
				Body    struct {
					M wsdlgo.GetAllSharedRoleNamesResponse `xml:"getAllSharedRoleNamesResponse"`
				}
				}{}
				if err = p.Cli.RoundTrip(α, &γ); err != nil {
					return nil, err
				}
				return &γ.Body.M, nil
			}*/


func (p *UserAdminPortType) GetAllUIPermissions(α *wsdlgo.GetAllUIPermissions) (β *wsdlgo.GetAllUIPermissionsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetAllUIPermissionsResponse `xml:"getAllUIPermissionsResponse"`
		}
		}{}
		if err = p.Cli.RoundTrip(α, &γ, "getAllUIPermissions"); err != nil {
			return nil, err
		}
		return &γ.Body.M, nil
}


func (p *UserAdminPortType) GetRolePermissions(α *wsdlgo.GetRolePermissions) (β *wsdlgo.GetRolePermissionsResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.GetRolePermissionsResponse `xml:"getRolePermissionsResponse"`
		}
	}{}
	if err = p.Cli.RoundTrip(α, &γ, "getRolePermissions"); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}

// GetRolesOfCurrentUser
					func (p *UserAdminPortType) GetRolesOfCurrentUser(α *wsdlgo.GetRolesOfCurrentUser) (β *wsdlgo.GetRolesOfCurrentUserResponse, err error) {
						γ := struct {
							XMLName xml.Name `xml:"Envelope"`
							Body    struct {
								M wsdlgo.GetRolesOfCurrentUserResponse `xml:"getRolesOfCurrentUserResponse"`
							}
							}{}
							if err = p.Cli.RoundTrip(α, &γ, "getRolesOfCurrentUser"); err != nil {
								return nil, err
							}
							return &γ.Body.M, nil
						}

// GetRolesOfUser
						func (p *UserAdminPortType) GetRolesOfUser(α *wsdlgo.GetRolesOfUser) (β *wsdlgo.GetRolesOfUserResponse, err error) {
							γ := struct {
								XMLName xml.Name `xml:"Envelope"`
								Body    struct {
									M wsdlgo.GetRolesOfUserResponse `xml:"getRolesOfUserResponse"`
								}
								}{}
								if err = p.Cli.RoundTrip(α, &γ, "getRolesOfUser"); err != nil {
									return nil, err
								}
								return &γ.Body.M, nil
							}

// GetUserRealmInfo
							func (p *UserAdminPortType) GetUserRealmInfo(α *wsdlgo.GetUserRealmInfo) (β *wsdlgo.GetUserRealmInfoResponse, err error) {
								γ := struct {
									XMLName xml.Name `xml:"Envelope"`
									Body    struct {
										M wsdlgo.GetUserRealmInfoResponse `xml:"getUserRealmInfoResponse"`
									}
									}{}
									if err = p.Cli.RoundTrip(α, &γ, "getUserRealmInfo"); err != nil {
										return nil, err
									}
									return &γ.Body.M, nil
								}

// GetUsersOfRole
								func (p *UserAdminPortType) GetUsersOfRole(α *wsdlgo.GetUsersOfRole) (β *wsdlgo.GetUsersOfRoleResponse, err error) {
									γ := struct {
										XMLName xml.Name `xml:"Envelope"`
										Body    struct {
											M wsdlgo.GetUsersOfRoleResponse `xml:"getUsersOfRoleResponse"`
										}
										}{}
										if err = p.Cli.RoundTrip(α, &γ, "getUsersOfRole"); err != nil {
											return nil, err
										}
										return &γ.Body.M, nil
									}

// HasMultipleUserStores
									func (p *UserAdminPortType) HasMultipleUserStores(α *wsdlgo.HasMultipleUserStores) (β *wsdlgo.HasMultipleUserStoresResponse, err error) {
										γ := struct {
											XMLName xml.Name `xml:"Envelope"`
											Body    struct {
												M wsdlgo.HasMultipleUserStoresResponse `xml:"hasMultipleUserStoresResponse"`
											}
											}{}
											if err = p.Cli.RoundTrip(α, &γ, "hasMultipleUserStores"); err != nil {
												return nil, err
											}
											return &γ.Body.M, nil
										}

// IsSharedRolesEnabled
										func (p *UserAdminPortType) IsSharedRolesEnabled(α *wsdlgo.IsSharedRolesEnabled) (β *wsdlgo.IsSharedRolesEnabledResponse, err error) {
											γ := struct {
												XMLName xml.Name `xml:"Envelope"`
												Body    struct {
													M wsdlgo.IsSharedRolesEnabledResponse `xml:"isSharedRolesEnabledResponse"`
												}
												}{}
												if err = p.Cli.RoundTrip(α, &γ, "isSharedRolesEnabled"); err != nil {
													return nil, err
												}
												return &γ.Body.M, nil
											}


func (p *UserAdminPortType) ListAllUsers(α *wsdlgo.ListAllUsers) (β *wsdlgo.ListAllUsersResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.ListAllUsersResponse `xml:"listAllUsersResponse"`
		}
	}{}
	if err = p.Cli.RoundTrip(α, &γ, "listAllUsers"); err != nil {
		return nil, err
	}

	l := γ.Body.M.Return

	for i := 0;i < len(l); i++ {
		fmt.Println("========================")
		fmt.Printf("User #%d\n", i+1)
		fmt.Println("========================")
		s := reflect.ValueOf(l[i]).Elem()
		typeOfT := s.Type()
		for j := 0;j < s.NumField(); j++ {
			f := s.Field(j)
			fmt.Printf("%s = %v\n", typeOfT.Field(j).Name, f.Interface())	
		}
		fmt.Println()
	}
	return &γ.Body.M, nil
}


func (p *UserAdminPortType) ListAllUsersWithPermission(α *wsdlgo.ListAllUsersWithPermission) (β *wsdlgo.ListAllUsersWithPermissionResponse, err error) {
	γ := struct {
		XMLName xml.Name `xml:"Envelope"`
		Body    struct {
			M wsdlgo.ListAllUsersWithPermissionResponse `xml:"listAllUsersWithPermissionResponse"`
		}
	}{}
	if err = p.Cli.RoundTrip(α, &γ, "listAllUsersWithPermission"); err != nil {
		return nil, err
	}
	return &γ.Body.M, nil
}


													func (p *UserAdminPortType) ListUserByClaim(α *wsdlgo.ListUserByClaim) (β *wsdlgo.ListUserByClaimResponse, err error) {
														γ := struct {
															XMLName xml.Name `xml:"Envelope"`
															Body    struct {
																M wsdlgo.ListUserByClaimResponse `xml:"listUserByClaimResponse"`
															}
															}{}
															if err = p.Cli.RoundTrip(α, &γ, "listUserByClaim"); err != nil {
																return nil, err
															}
															return &γ.Body.M, nil
														}

// ListUserByClaimWithPermission
														func (p *UserAdminPortType) ListUserByClaimWithPermission(α *wsdlgo.ListUserByClaimWithPermission) (β *wsdlgo.ListUserByClaimWithPermissionResponse, err error) {
															γ := struct {
																XMLName xml.Name `xml:"Envelope"`
																Body    struct {
																	M wsdlgo.ListUserByClaimWithPermissionResponse `xml:"listUserByClaimWithPermissionResponse"`
																}
																}{}
																if err = p.Cli.RoundTrip(α, &γ, "listUserByClaimWithPermission"); err != nil {
																	return nil, err
																}
																return &γ.Body.M, nil
															}

// ListUsers
															func (p *UserAdminPortType) ListUsers(α *wsdlgo.ListUsers) (β *wsdlgo.ListUsersResponse, err error) {
																γ := struct {
																	XMLName xml.Name `xml:"Envelope"`
																	Body    struct {
																		M wsdlgo.ListUsersResponse `xml:"listUsersResponse"`
																	}
																	}{}
																	if err = p.Cli.RoundTrip(α, &γ, "listUsers"); err != nil {
																		return nil, err
																	}
																	return &γ.Body.M, nil
																}

// SetRoleUIPermission
																func SetRoleUIPermission(ctx context.Context, parameters *wsdlgo.SetRoleUIPermission) (err error) {
																	return errors.New("not implemented")
																}

// UpdateRoleName
																func UpdateRoleName(ctx context.Context, parameters *wsdlgo.UpdateRoleName) (err error) {
																	return errors.New("not implemented")
																}

// UpdateRolesOfUser
																func UpdateRolesOfUser(ctx context.Context, parameters *wsdlgo.UpdateRolesOfUser) (err error) {
																	return errors.New("not implemented")
																}

// UpdateUsersOfRole
																func UpdateUsersOfRole(ctx context.Context, parameters *wsdlgo.UpdateUsersOfRole) (err error) {
																	return errors.New("not implemented")
																}