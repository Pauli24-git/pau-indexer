package models

type AuthHandlerMock struct {
}

func (a *AuthHandlerMock) ValidateAuthDbUser() (Credentials, error) {
	c := Credentials{}
	return c, nil
}

func (a *AuthHandlerMock) UserExists(credentials Credentials) error {
	return nil
}

type DbHandlerMock struct {
}

func (a *DbHandlerMock) SearchQuery(credentials Credentials, query Search) (ZSResponse, error) {
	resp := ZSResponse{}

	src := Source{Message_ID: "ABC123", From: "pau@enron.com", Subject: "Urgente"}
	hit := Hits{Source: src}

	resp.Hits.Hits = append(resp.Hits.Hits, hit)
	return resp, nil
}
