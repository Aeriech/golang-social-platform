CREATE TABLE IF NOT EXISTS users(
   id bigserial PRIMARY KEY,
   username VARCHAR(255) UNIQUE NOT NULL,
   email VARCHAR(255) UNIQUE NOT NULL,
   password bytea NOT NULL,
   created_at timestamp with time zone NOT NULL DEFAULT NOW(),
   updated_at timestamp with time zone NOT NULL DEFAULT NOW()
);