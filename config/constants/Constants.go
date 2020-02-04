package constants

const (
	// InsuficientLimitError represents that the transaction amount should not exceed the available limit
	InsuficientLimitError = "insuficient-limit"
	// AccountAlreadyInitialized represents that an account has already initialized
	AccountAlreadyInitialized = "account-already-initialized"
	// AccountIsNotInitialized no transaction should be accepted without a properly initialized account
	AccountIsNotInitialized = "account-not-initialized"
	// AccountCardIsNotActive represents no transaction should be accepted when the card is not active
	AccountCardIsNotActive = "card-not-active"
	// TransactionHighFrequency represents that there should not be more than 3 transactions on a 2-minute interval
	TransactionHighFrequency = "high-frequency-small-interval"
	// DoubleTransaction represents that there should not be more than 1 similar transactions
	DoubleTransaction = "double-transaction"
)
