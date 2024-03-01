package types


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

func FullAdmin() Permission {

	// ACL USER
	permissionBlock := Permission {
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

	return permissionBlock

}

func ReadOnly() Permission {

	// ACL USER
	permissionBlock := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	}

	return permissionBlock

}

func Approver() Permission {

	// ACL USER
	permissionBlock := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	}

	return permissionBlock

}

func Manager() Permission {

	// ACL USER
	permissionBlock := Permission {
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

	return permissionBlock

}

func ConjurSync() Permission {

	// ACL USER
	permissionBlock := Permission {
	UseAccounts:							true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AccessWithoutConfirmation:				true,
	}

	return permissionBlock

}

func SecretsHub() Permission {

	// ACL USER
	permissionBlock := Permission {
	ViewSafeMembers:						true,
	RetrieveAccounts:						true,
	ListAccounts:							true,
	AccessWithoutConfirmation:				true,
	}

	return permissionBlock

}