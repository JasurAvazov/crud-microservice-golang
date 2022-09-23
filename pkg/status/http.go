package status

var (
	// NoError - Success code
	NoError = 0

	// ErrorCodeValidation - for http 422
	ErrorCodeValidation = -10

	// ErrorCodeValidationDateFormat - for http 422 on date error
	ErrorCodeValidationDateFormat = -12

	// ErrorCodeDB - for http 500
	ErrorCodeDB = -30
	// ErrorCodeEntityNotFound - for http 500 could not retrieve from DB
	ErrorCodeEntityNotFound = -31

	// ErrorCodeRemoteDBO - DBO returned error
	ErrorCodeRemoteDBO = -40
	// ErrorCodeRemoteDBOEntityNotFound - DBO returned error not found
	ErrorCodeRemoteDBOEntityNotFound = -41

	// ErrorCodeRemoteESB - ESB returned error
	ErrorCodeRemoteESB = -45
	// ErrorCodeRemoteESB - ESB returned empty data
	ErrorCodeRemoteESBEntityNotFound = -46

	// ErrorCodeRemoteCRM - CRM returned error
	ErrorCodeRemoteCRM = -50
	// ErrorCodeRemoteCRMEntityNotFound - CRM returned error
	ErrorCodeRemoteCRMEntityNotFound = -51

	// ErrorCodeRemoteOther - Remote resource returned error
	ErrorCodeRemoteOther = -55
	// ErrorCodeRemoteOtherEntityNotFound - Remote resource returned error
	ErrorCodeRemoteOtherEntityNotFound = -56

	// ErrorCodeRMQ - RMQ returned error
	ErrorCodeRMQ = -60
)

var (
	// Success ...
	Success = "Success"
	// Failure ...
	Failure = "Failure"
)
