package schema

//Credentials stores all the credentials along with channelID
type Credentials struct {
	ChannelID string             `json:"channel_id,omitempty" bson:"channel_id,omitempty"`
	Twitter   TwitterCredentials `json:"twitter,omitempty" bson:"twitter,omitempty"`
}

// TwitterCredentials stores the twitter credentials
type TwitterCredentials struct {
	AccessToken    string            `json:"access_token,omitempty" bson:"access_token,omitempty"`
	AccessSecret   string            `json:"access_secret,omitempty" bson:"access_secret,omitempty"`
	AdditionalData map[string]string `json:"additional_data,omitempty" bson:"additional_data,omitempty"`
	RequestToken   RequestToken      `json:"request_token,omitempty" bson:"request_token,omitempty"`
}

// RequestToken stores the OAuth token and OAuth verifier.
type RequestToken struct {
	OAuthToken    string `json:"oauth_token" bson:"oauth_token"`
	OAuthVerifier string `json:"oauth_verifier" bson:"oauth_verifier"`
}
