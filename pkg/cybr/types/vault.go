package types

import (
	"encoding/json"
	"log"
)

/*
=========================================
* Structs
=========================================
*/

// Permission Struct

type Permission struct {
	ManageSafe								bool `json:"manageSafe"`
	ManageSafeMembers						bool `json:"manageSafeMembers"`
	ViewSafeMembers							bool `json:"viewSafeMembers"`
	ViewAuditLog							bool `json:"viewAuditLog"`
	UseAccounts								bool `json:"useAccounts"`
	RetrieveAccounts						bool `json:"retrieveAccounts"`
	ListAccounts							bool `json:"listAccounts"`
	AddAccounts								bool `json:"addAccounts"`
	UpdateAccountContent					bool `json:"updateAccountContent"`
	UpdateAccountProperties					bool `json:"updateAccountProperties"`
	RenameAccounts							bool `json:"renameAccounts"`
	DeleteAccounts							bool `json:"deleteAccounts"`
	UnlockAccounts							bool `json:"unlockAccounts"`
	InitiateCPMAccountManagementOperations 	bool `json:"initiateCPMAccountManagementOperations"`
	SpecifyNextAccountContent				bool `json:"specifyNextAccountContent"`
	BackupSafe								bool `json:"backupSafe"`
	AccessWithoutConfirmation				bool `json:"accessWithoutConfirmation"`
	CreateFolders							bool `json:"createFolders"`
	DeleteFolders							bool `json:"deleteFolders"`
	MoveAccountsAndFolders					bool `json:"moveAccountsAndFolders"`
	RequestsAuthorizationLevel1				bool `json:"requestsAuthorizationLevel1"`
	RequestsAuthorizationLevel2				bool `json:"requestsAuthorizationLevel2"`
}

// ACL Struct

type Member struct {
	Member			string `json:"memberName"`
	MemberType 		string `json:"memberType"`
	Perm 			Permission `json:"permissions"`
}

// Shared Services Structs

type CloudConfig struct {
	Tenant 		 string // Shared Services Tenant
	ClientID 	 string // Email address of onboarding services account
	ClientSecret string // Related password to onboarding services account
	Domain		 string // Base domain for privilege cloud
}

type Token struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int    `json:"expires_in"`
}

type Credential struct {

	Name		string `json:"name"` // Custom Account Name of the credential
	Address 	string `json:"address"` // Address of where the credential is used
	UserName 	string `json:"userName"` // Username value
	Platform	string `json:"platformId"` // Required: Management platform
	SafeName    string `json:"safeName"` // Required: Target Safe
	SecretType	string `json:"secretType"` // Type of secret (use password)
	Secret		string `json:"secret"` // Password Value
	SecretMgmt  SecretManagement `json:"secretManagement"`
	Props		AccountProps `json:"platformAccountProperties"`
}

type AccountProps struct {
	Port   string `json:"port"`
	DBName string `json:"database"`
}

type SecretManagement struct {
	AutomaticManagement bool `json:"automaticManagementEnabled"`
	ManualManagementReason string `json:"manualManagementReason"`
}

type RemoteAccess struct {
	RemoteMachines 	  string `json:"remoteMachines"`
	RestrictToMachine string `json:"accessRestrictedToRemoteMachines"` 
}

type Safe struct {

}

/*
=========================================
* Functions
=========================================
*/

// Get Full Administrator Permissions
// intakes a user type string and user string to bundle permissions
func FullAdmin(UserType string, User string) ([]byte, error) {

	Perm := Permission {
	ManageSafe:								true,
	ManageSafeMembers:						true,
	ViewSafeMembers:						true,
	ViewAuditLog:							true,
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AddAccounts:							true,
	UpdateAccountContent:					true,
	UpdateAccountProperties:				true,
	RenameAccounts:							true,
	DeleteAccounts:							true,
	UnlockAccounts:							true,
	InitiateCPMAccountManagementOperations:	true,
	SpecifyNextAccountContent:				true,
	BackupSafe:								true,
	AccessWithoutConfirmation:				true,
	CreateFolders:							true,
	DeleteFolders:							true,
	MoveAccountsAndFolders:					true,
	RequestsAuthorizationLevel1:			true,
	RequestsAuthorizationLevel2:			true,
	}

	userBlock := Member {
		Member: User,
		MemberType: UserType,
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, err

}

// Get Read Only Permissions
// intakes a user type string and user string to bundle permissions
func ReadOnly(UserType string, User string) ([]byte, error) {

	Perm := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	}

	userBlock := Member {
		Member: User,
		MemberType: UserType,
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, err

}

// Get Approver Permissions
// intakes a user type string and user string to bundle permissions
func Approver(UserType string, User string) ([]byte, error) {

	Perm := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	ViewSafeMembers:						true,
	ManageSafeMembers:						true,
	}

	userBlock := Member {
		Member: User,
		MemberType: UserType,
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, err

}

// Get Safe Manager Permissions
// intakes a user type string and user string to bundle permissions
func Manager(UserType string, User string) ([]byte, error)  {

	Perm := Permission {
	ManageSafeMembers:						true,
	ViewSafeMembers:						true,
	ViewAuditLog:							true,
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AddAccounts:							true,
	UpdateAccountContent:					true,
	UpdateAccountProperties:				true,
	RenameAccounts:							true,
	DeleteAccounts:							true,
	UnlockAccounts:							true,
	InitiateCPMAccountManagementOperations:	true,
	SpecifyNextAccountContent:				true,
	AccessWithoutConfirmation:				true,
	}

	userBlock := Member {
		Member: User,
		MemberType: UserType,
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, err

}

// Get Conjur Component User Permissions
func ConjurSync() ([]byte, error) {

	Perm := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AccessWithoutConfirmation:				true,
	}

	userBlock := Member {
		Member: "ConjurSync",
		MemberType: "User",
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, err

}

// Get Secrets Hub Component User Permissions
func SecretsHub() ([]byte, error) {

	Perm := Permission {
	ViewSafeMembers:						true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AccessWithoutConfirmation:				true,
	}

	userBlock := Member {
		Member: "SecretsHub",
		MemberType: "User",
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, err

}

// Build credential object for onboarding via API
func Cred(cred *Credential) ([]byte, error) {

	d := &cred;

	jsonData, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData, err
}