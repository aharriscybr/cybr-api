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
		Member: &User,
		MemberType: &UserType,
		Perm: Perm,
	}

	if &User == nil || &UserType == nil {

		log.Fatal("User or User Type is null")
		return nil, nil
		
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, nil

}

// Get Read Only Permissions
// intakes a user type string and user string to bundle permissions
func ReadOnly(UserType string, User string) ([]byte, error) {

	Perm := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	}

	if &User == nil || &UserType == nil {

		log.Fatal("User or User Type is null")
		return nil, nil
		
	}

	userBlock := Member {
		Member: &User,
		MemberType: &UserType,
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, nil

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

	if &User == nil || &UserType == nil {

		log.Fatal("User or User Type is null")
		return nil, nil
		
	}

	userBlock := Member {
		Member: &User,
		MemberType: &UserType,
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, nil

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

	if &User == nil || &UserType == nil {

		log.Fatal("User or User Type is null")
		return nil, nil
		
	}

	userBlock := Member {
		Member: &User,
		MemberType: &UserType,
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, nil

}

// Get Conjur Component User Permissions
func ConjurSync() ([]byte, error) {

	Perm := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AccessWithoutConfirmation:				true,
	}

	US := "ConjurSync"
	UT := "User"

	userBlock := Member {
		Member: &US,
		MemberType: &UT,
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, nil

}

// Get Secrets Hub Component User Permissions
func SecretsHub() ([]byte, error) {

	Perm := Permission {
	ViewSafeMembers:						true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AccessWithoutConfirmation:				true,
	}

	US := "SecretsHub"
	UT := "User"

	userBlock := Member {
		Member: &US,
		MemberType: &UT,
		Perm: Perm,
	}

	thisBlock, err := json.Marshal(userBlock)
	if err != nil {
		log.Fatal(err)
	}

	return thisBlock, nil

}

