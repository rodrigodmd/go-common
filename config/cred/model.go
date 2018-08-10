package cred

type CredType string

const (
	BITBUCKET       CredType = "bitbucket"
	DIGITAL_CONSOLE          = "digital-console"
	LEGACY_CONSOLE           = "legacy-console"
	SSH                      = "ssh"
	CUSTOM                   = "custom"
)

type CredModel struct {
	Username string
	Password string
}

type ConfigCredsModel struct {
	Version  string
	CredList map[CredType]CredModel
}
