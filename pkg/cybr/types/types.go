package types

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

// Shared Services Structs

type Authn struct {
	Tenant 		 string
	ClientID 	 string
	ClientSecret string
	GrantType	 string `default:"client_credentials"`
}

type Token struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int    `json:"expires_in"`
}

/*
=========================================
* Functions
=========================================
*/

func FullAdmin() *Permission {

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

	return &Perm

}

func ReadOnly() *Permission {

	Perm := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	}

	return &Perm

}

func Approver() *Permission {

	Perm := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	}

	return &Perm

}

func Manager() *Permission {

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

	return &Perm

}

func ConjurSync() *Permission {

	Perm := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AccessWithoutConfirmation:				true,
	}

	return &Perm

}

func SecretsHub() *Permission {

	Perm := Permission {
	ViewSafeMembers:						true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AccessWithoutConfirmation:				true,
	}

	return &Perm

}

func SetAuthn(t string, cid string, csec string) *Authn {

	AuthnObj := Authn {
		Tenant: t,
		ClientID: cid,
		ClientSecret: csec,
	}

	return &AuthnObj

}