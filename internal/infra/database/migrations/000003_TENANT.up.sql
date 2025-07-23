CREATE TABLE tenants (
    uuid VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE tokens
ADD CONSTRAINT fk_tokens_tenants
FOREIGN KEY (tenant_id) REFERENCES tenants(uuid);

ALTER TABLE users
ADD CONSTRAINT fk_users_tenants
FOREIGN KEY (tenant_id) REFERENCES tenants(uuid);