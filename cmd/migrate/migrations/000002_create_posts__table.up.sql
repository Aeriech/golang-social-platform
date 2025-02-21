CREATE TABLE IF NOT EXISTS posts(
   id bigserial PRIMARY KEY,
   content text NOT NULL,
   title text NOT NULL,
   user_id bigint NOT NULL,
   created_at timestamp with time zone NOT NULL DEFAULT NOW(),
   updated_at timestamp with time zone NOT NULL DEFAULT NOW()
);