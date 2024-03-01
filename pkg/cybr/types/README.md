## Types

*types* library is to specify json permission block mappings for onboarding user types into the CyberArk Vault API. 

### Extending this library:

Declare new function returning the **struct**:

    func NewType() *Permission {

        perm := Permission {
        UseAccounts:    true,
        RetrieveAccounts:   true,
        ListAccounts:   true,
        AccessWithoutConfirmation:  true,
        }

        // Return pointer of the new permissions
        return &perm

    }