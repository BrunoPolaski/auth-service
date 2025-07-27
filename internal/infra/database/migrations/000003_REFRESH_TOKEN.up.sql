CREATE TABLE IF NOT EXISTS tokens (
    uuid VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    tenant_id VARCHAR(36) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(uuid),
    FOREIGN KEY (tenant_id) REFERENCES tenants(uuid)
);