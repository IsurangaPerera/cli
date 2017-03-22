# Command Line Tool For WSO2 [IS]

[![N|Solid](http://static.wixstatic.com/media/ae45cc_4fad18bcbc154f2da70a90cd93d74a4e.png_256)](https://nodesource.com/products/nsolid)
### Installation

CLI requires [GO](https://golang.org/) to run.

***Linux Installation***


```sh
$ go get github.com/IsurangaPerera/mow.cli
$ git clone https://github.com/IsurangaPerera/cli.git
$ cd cli
$ go build .
```
### Optional Steps
Create a Symlink to the cli directory inorder to run from anywhere

### Configuration
Following attributes should be set inorder to run properly
* **Base URL** (eg. localhost:9443)
* **User Credentials** (username & password)
* **Soap11 Endpoint** (eg. services/UserAdmin.UserAdminHttpsSoap11Endpoint/)
* **Soap12 Endpoint** (eg. services/UserAdmin.UserAdminHttpsSoap12Endpoint/)
```sh
$ cli config -url [base_url] -u [user_name] -p [password] -s1 [soap11_endpoint] -s2 [soap12_endpoint]
```

### Usage

CLI is currently extended with the following capabilities.

| Command | Arguements | Description | Type |
| ------ | ------ | ------| ------ |
| addInternalRole | **-r** | **role name** | string |
| | **-u**| **user list** | string [ ] |
| | **-p** | **permissions** | string [ ] |
| addRemoveRolesOfUser | **-u** | **user name** | string |
| | **-r** | **new roles** | string [ ] |
| | **-d** | **deleted roles** | string [ ] |
| addRemoveUsersOfRole | **-r** | **role name** | string |
| | **-u** | **new users** | string [ ] |
| | **-d** | **deleted users** | string [ ] |
| addRole | **-r** | **role name** | string |
| | **-u** | **user list** | string [ ] |
| | **-p** | **permissions** | string [ ] |
| | **-s** | **shared** | boolean |
| listAllUsers | **-f** | **filter** | string |
| | **-l** | limit | int |


### Samples

***Add Internal Role***

```sh
$ cli addInternalRole -r "sample" -u [isuranga harsha] 
```
***List All Users***

```sh
$ cli listAllUsers
```
 OR

```sh
$ cli listAllUsers -f "*" -l 10
```

