CREATE TABLE api_keys (
    uuid VARCHAR(36) PRIMARY KEY,
    description VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    tenant_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (tenant_id) REFERENCES tenants (uuid)
);