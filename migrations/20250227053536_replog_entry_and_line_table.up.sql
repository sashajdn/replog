CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "replog_channel" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL UNIQUE,
  messaging_provider messinging_provider NOT NULL DEFAULT 'discord',
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
)

CREATE TABLE IF NOT EXISTS "replog_entry" (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id VARCHAR(255) REFERENCES "user"(id) NOT NULL,
  channel_id VARCHAR(255) REFERENCES "channel"(id) NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS  "replog_entry_link" (
  parent_entry_id UUID NOT NULL REFERENCES "replog_entry"(id) ON DELETE CASCADE,
  child_entry_id UUID NOT NULL REFERENCES "replog_entry"(id) ON DELETE CASCADE,
  PRIMARY KEY (parent_entry_id, child_entry_id),
);
