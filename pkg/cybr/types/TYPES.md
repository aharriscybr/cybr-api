## Types

*types* library is to specify json permission block mappings for onboarding user types into the CyberArk Vault API. 

### Extending this library:

Declare new function returning the **struct**:

    func NewType() Permission {

        // ACL USER
        permissionBlock := Permission {
        UseAccounts:							true,
        RetrieveAccounts:						true,
        ListAccounts:							true,
        AccessWithoutConfirmation:				true,
        }

        return permissionBlock

    }
