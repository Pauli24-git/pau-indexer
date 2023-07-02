package models

type DBHandler interface {
	SearchQuery(cred Credentials, query Search) (ZSResponse, error)
}

type ZincAuthHandler interface {
	ValidateAuthDbUser() (Credentials, error)
	UserExists(credentials Credentials) error
}
