-- up
CREATE TABLE users (
  uuid       UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email      VARCHAR(255) NOT NULL,
  password   VARCHAR(100) NOT NULL,
  created_at TIMESTAMP    NOT NULL,
  updated_at TIMESTAMP    NOT NULL
);

CREATE TABLE vcards (
  uuid      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_uuid UUID        NOT NULL,
  name      VARCHAR(50) NOT NULL,
  path      UUID        NOT NULL UNIQUE
);

ALTER TABLE vcards
  ADD CONSTRAINT fk_user_uuid FOREIGN KEY (user_uuid) REFERENCES users (uuid) ON DELETE RESTRICT;

-- down
ALTER TABLE vcards
  DROP CONSTRAINT fk_user_uuid;

DROP TABLE vcards;
DROP TABLE users;