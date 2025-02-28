CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE messinging_provider AS ENUM ('discord');

CREATE TABLE "user" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  username VARCHAR(255) NOT NULL,
  messaging_provider_id VARCHAR(255) UNIQUE,
  messaging_provider messinging_provider NOT NULL DEFAULT 'discord',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);
