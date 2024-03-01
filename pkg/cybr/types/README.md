## Types

*types* library is to specify json permission block mappings for onboarding user types into the CyberArk Vault API. This will return a properly marshalled json body ready for permissions to be applied on the object inside CyberArk Vault.

### Extending this library:

Declare new function returning the **struct** and define the permissions in the **Perm** section

    func NewType(UserType string, User string) ([]byte, error) {

	Perm := Permission {
        UseAccounts:    true,
        RetrieveAccounts:   true,
        ListAccounts:   true,
        AccessWithoutConfirmation:  true,
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

        return thisBlock, nil

    }