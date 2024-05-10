CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE gender as enum (
    'Male',
    'Female'
);

CREATE TABLE users(
    id  UUID DEFAULT gen_random_uuid() PRIMARY KEY ,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    gender gender NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_time timestamp with time zone  DEFAULT NOW() NOT NULL,
    updated_time timestamp with time zone 
);

CREATE TABLE sessions (
    id uuid PRIMARY KEY NOT NULL,
    username varchar NOT NULL,
    refresh_token varchar NOT NULL,
    expires_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    deleted_at timestamptz
);

--  di monogodb aja ini
-- CREATE TABLE likes(
--     id  UUID DEFAULT gen_random_uuid() PRIMARY KEY ,
--     user_id UUID,
--     tweets_id UUID
-- );

-- CREATE TABLE tweets (
--     id  UUID DEFAULT gen_random_uuid() PRIMARY KEY ,
--     user_id UUID NOT NULL,
--     hashtag VARCHAR(255),
--     content  TEXT,
--     created_time timestamp with time zone  DEFAULT NOW() NOT NULL,
--     updated_time timestamp with time zone 
-- );

-- CREATE TABLE videos (
--     id  UUID DEFAULT gen_random_uuid() PRIMARY KEY ,
--     name VARCHAR(255) NOT NULL,
--     url VARCHAR(255) ,
--     length int NOT NULL,
--     size double NOT NULL,
--     tweets_id UUID NOT NULL,
--     created_time timestamp with time zone  DEFAULT NOW() NOT NULL,
--     updated_time timestamp with time zone 
-- );



-- CREATE TABLE images (
--     id  UUID DEFAULT gen_random_uuid() PRIMARY KEY ,
--     name VARCHAR(255) NOT NULL,
--     url VARCHAR(255),
--     size double NOT NULL,
--     tweets_id UUID NOT NULL,
--     created_time timestamp with time zone  DEFAULT NOW() NOT NULL,
--     updated_time timestamp with time zone 
-- );

-- ALTER TABLE tweets ADD CONSTRAINT fk_tweets_users
--     FOREIGN KEY(user_id) 
--     REFERENCES users(id);

-- ALTER TABLE likes ADD CONSTRAINT fk_likes_users
--     FOREIGN KEY (user_id)
--     REFERENCES users(id);


-- ALTER TABLE likes ADD CONSTRAINT fk_likes_tweets
--     FOREIGN KEY (tweets_id)
--     REFERENCES tweets(id);

-- ALTER TABLE videos ADD CONSTRAINT fk_video_tweets
--     FOREIGN KEY (tweets_id)
--     REFERENCES tweets(id);

-- ALTER TABLE images ADD CONSTRAINT fk_images_tweets
--     FOREIGN KEY (tweets_id)
--     REFERENCES tweets(id);





-- ALTER TABLE container_lifecycles ADD  CONSTRAINT fk_lifecycles_containers
--     FOREIGN KEY (container_id)
--     REFERENCES containers (id);