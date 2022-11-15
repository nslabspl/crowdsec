package database

import "errors"

var (

	// General errors
	UserExists        			= errors.New("user already exist")
	UserNotExists     			= errors.New("user doesn't exist")
	HashError         			= errors.New("unable to hash")
	InsertFail        			= errors.New("unable to insert row")
	QueryFail         			= errors.New("unable to query")
	UpdateFail        			= errors.New("unable to update")
	DeleteFail        			= errors.New("unable to delete")
	ItemNotFound      			= errors.New("object not found")
	ParseTimeFail     			= errors.New("unable to parse time")
	ParseDurationFail 			= errors.New("unable to parse duration")
	MarshalFail       			= errors.New("unable to marshal")
	UnmarshalFail     			= errors.New("unable to unmarshal")
	BulkError         			= errors.New("unable to insert bulk")
	ParseType         			= errors.New("unable to parse type")

	// Filtering errors
	InvalidFilter     			= errors.New("invalid filter")
	InvalidDecision	  			= errors.New("Decision not valid due to errors (either codelike or semantic)")
	Ip4NotInRange      			= errors.New("Given IP is out of range for IPv4")
	RangeNotValidOrNotSpecified = errors.New("Provided range is not valid CIDR range")
	FilterNotApplied			= errors.New("Filter has not been applied")
)
