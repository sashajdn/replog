CREATE TABLE messaging_provider_message_last_seen (
  message_id VARCHAR(255),
  channel_id VARCHAR(255) REFERENCES "channel"(id) NOT NULL,
  user_id VARCHAR(255) REFERENCES "user"(id) NOT NULL,
  messaging_provider messinging_provider NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (channel_id, user_id, messaging_provider)
)
