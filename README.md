# Command Line Tool For WSO2 [IS]

[![N|Solid](http://static.wixstatic.com/media/ae45cc_4fad18bcbc154f2da70a90cd93d74a4e.png_256)](https://nodesource.com/products/nsolid)
### Installation

CLI requires [GO](https://golang.org/) to run.

***Linux Installation***


```sh
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
$ cli config --base-url [base_url] --user-name [user_name] --password [password] --soap11 [soap11_endpoint] --soap12 [soap12_endpoint]
```

### Usage

CLI is currently extended with the following capabilities.

| Command | Arguements |
| ------ | ------ |
| addInternalRole | **[role-name] [user-list] [permissions]** |
| addRemoveRolesOfUser | **[user-name] [new-roles] [deleted-roles]** |
| addRemoveUsersOfRole | **[role-name] [new-users] [deleted-users]** |
| addRole | **[role-name] [user-list] [permissions] [shared]** |

### Samples

***Add Internal Method***

```sh
$ cli addInternalRole --role-name sample --user-list isuranga 
```

