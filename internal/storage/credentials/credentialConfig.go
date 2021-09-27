package credentials

import "fmt"

type CredentialConfig struct {
	server string
	port   int
	dbName string

	user, password,
	connectionString string
}

func (c *CredentialConfig) GetConnectionString() string {
	return c.connectionString
}

type CredentialBuilder struct {
	credential *CredentialConfig
}

type ServerInfoBuilder struct {
	CredentialBuilder
}

type UserCredentialBuilder struct {
	CredentialBuilder
}

func NewCredentialBuilder() *CredentialBuilder {
	return &CredentialBuilder{credential: &CredentialConfig{}}
}

func (b *CredentialBuilder) ServerInfo() *ServerInfoBuilder {
	return &ServerInfoBuilder{*b}
}

func (b *CredentialBuilder) UserAuthInfo() *UserCredentialBuilder {
	return &UserCredentialBuilder{*b}
}

func (a *ServerInfoBuilder) WithServerName(serverName string) *ServerInfoBuilder {
	a.credential.server = serverName
	return a
}

func (a *ServerInfoBuilder) WithPortNumber(port int) *ServerInfoBuilder {
	a.credential.port = port
	return a
}

func (j *ServerInfoBuilder) WithDbName(dbName string) *ServerInfoBuilder {
	j.credential.dbName = dbName
	return j
}

func (j *UserCredentialBuilder) WithUserName(uName string) *UserCredentialBuilder {
	j.credential.user = uName
	return j
}

func (j *UserCredentialBuilder) WithPassword(pwd string) *UserCredentialBuilder {
	j.credential.password = pwd
	return j
}

func (j *UserCredentialBuilder) GetConnectionString() *UserCredentialBuilder {
	var sqlFormat = "server=%s;user id=%s;password=%s;port=%d;database=%s;"
	j.credential.connectionString = fmt.Sprintf(sqlFormat, j.credential.server, j.credential.user,
		j.credential.password, j.credential.port,
		j.credential.dbName)
	return j
}

//Build builds a person from PersonBuilder
func (b *CredentialBuilder) Build() *CredentialConfig {
	return b.credential
}
