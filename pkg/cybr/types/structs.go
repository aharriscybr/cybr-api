package cyberarkapi_types

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
	Member			*string `json:"memberName"`
	MemberType 		*string `json:"memberType"`
	Perm 			Permission `json:"permissions"`
}

// Shared Services Structs

type CloudConfig struct {
	Tenant 		 *string `tfsdk:"tenant"` // Shared Services Tenant
	ClientID 	 *string `tfsdk:"clientid"` // Email address of onboarding services account
	ClientSecret *string `tfsdk:"clientsecret"` // Related password to onboarding services account
	Domain		 *string `tfsdk:"domain"` // Base domain for privilege cloud
}

type Token struct {
	Access_token *string `json:"access_token"`
	Token_type   *string `json:"token_type"`
	Expires_in   *int    `json:"expires_in"`
}

// Vault API Structs

type Credential struct {

	Name		*string `json:"name"` // Custom Account Name of the credential
	Address 	*string `json:"address"` // Address of where the credential is used
	UserName 	*string `json:"userName"` // Username value
	Platform	*string `json:"platformId"` // Required: Management platform
	SafeName    *string `json:"safeName"` // Required: Target Safe
	SecretType	*string `json:"secretType"` // Type of secret (use password)
	Secret		*string `json:"secret"` // Password Value
	SecretMgmt  SecretManagement `json:"secretManagement"`
	Props		*AccountProps `json:"platformAccountProperties"`
}

type AccountProps struct {

	/*
	* Generic - Types that map to additional platforms
	*/

	Port   *string `json:"port,omitempty"`

	/*
	* DB Handlers
	*/
	
	DBName *string `json:"database,omitempty"`
	DSN *string `json:"dsn,omitempty"`

	/*
	* AWS
	*/

	AWSKID *string `json:"AWSAccessKeyID,omitempty"`
	AWSAccount *string `json:"AWSAccountID,omitempty"`
	Alias *string `json:"AWSAccountAliasName,omitempty"`
	Region *string `json:"Region,omitempty"`

	/*
	* Microsoft
	*/

	MAppID *string `json:"ApplicationID,omitempty"`
	MAppObjectID *string `json:"ApplicationObjectID,omitempty"`
	MKID *string `json:"KeyID,omitempty"`
	MADID *string `json:"ActiveDirectoryID,omitempty"`
	MDur *string `json:"Duration,omitempty"`
	MPop *string `json:"PopulateIfNotExist,omitempty"`
	MKeyDesc *string `json:"KeyDescription,omitempty"`

	// template t *string `json:"t,omitempty"`
}

type SecretManagement struct {
	AutomaticManagement bool `json:"automaticManagementEnabled"`
	ManualManagementReason *string `json:"manualManagementReason"`
}

type RemoteAccess struct {
	RemoteMachines 	  *string `json:"remoteMachines"`
	RestrictToMachine *string `json:"accessRestrictedToRemoteMachines"` 
}

type CredentialResponse struct {
	CredID *string `json:"id,omitempty"`
	CreationTime *string `json:"lastModifiedTime,omitempty"`
	ModifiedTime *string `json:"createdTime,omitempty"`
}

type Safe struct {

}