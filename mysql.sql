CREATE TABLE account (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  owner VARCHAR(255) NOT NULL,
  balance BIGINT NOT NULL,
  currency VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE entry (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  account_id BIGINT NOT NULL,
  amount BIGINT NOT NULL COMMENT 'Can be negative or positive',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transfer (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  from_account_id BIGINT NOT NULL,
  to_account_id BIGINT NOT NULL,
  amount BIGINT NOT NULL COMMENT 'Must be positive',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX account_index_0 ON account (owner);
CREATE INDEX entry_index_1 ON entry (account_id);
CREATE INDEX transfer_index_2 ON transfer (from_account_id);
CREATE INDEX transfer_index_3 ON transfer (to_account_id);
CREATE INDEX transfer_index_4 ON transfer (from_account_id, to_account_id);

ALTER TABLE entry 
ADD CONSTRAINT fk_entry_account 
FOREIGN KEY (account_id) REFERENCES account(id);

ALTER TABLE transfer 
ADD CONSTRAINT fk_transfer_from 
FOREIGN KEY (from_account_id) REFERENCES account(id);

ALTER TABLE transfer 
ADD CONSTRAINT fk_transfer_to 
FOREIGN KEY (to_account_id) REFERENCES account(id);