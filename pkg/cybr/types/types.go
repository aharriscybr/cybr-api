package cyberarkapi_types

import (
	"encoding/json"
	"log"
)

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