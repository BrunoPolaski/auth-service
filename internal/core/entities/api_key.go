package entities

type ApiKey struct {
	uuid        int64
	description string
	tenantId    int64
}

func NewApiKey(
	uuid int64,
	description string,
	tenantId int64,
) *ApiKey {
	return &ApiKey{
		uuid:        uuid,
		description: description,
		tenantId:    tenantId,
	}
}

func (u *ApiKey) Id() int64 {
	return u.uuid
}

func (u *ApiKey) Description() string {
	return u.description
}

func (u *ApiKey) TenantId() int64 {
	return u.tenantId
}
