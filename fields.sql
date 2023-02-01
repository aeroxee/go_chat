CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    first_name CHARACTER VARYING(50) NOT NULL,
    last_name CHARACTER VARYING(50) NOT NULL,
    username CHARACTER VARYING(100) NOT NULL,
    email CHARACTER VARYING(100) NOT NULL,
    password CHARACTER VARYING(128) NOT NULL,
    last_login TIMESTAMP WITH TIME ZONE NULL,
    date_joined TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE UNIQUE INDEX idx_users_username ON users(username);
CREATE UNIQUE INDEX idx_users_email ON users(email);

CREATE TABLE groups (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    title CHARACTER VARYING(30) NOT NULL,
    logo CHARACTER VARYING(255) NULL,
    description CHARACTER VARYING(255) NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT fk_groups_user FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE UNIQUE INDEX idx_groups_title ON groups(title);

CREATE TABLE group_members (
    user_id BIGINT NOT NULL,
    group_id BIGINT NOT NULL,
    CONSTRAINT fk_group_members_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_group_members_group FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE chats (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT NOT NULL,
    group_id BIGINT NOT NULL,
    text TEXT NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT fk_chats_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_chats_group FOREIGN KEY (group_id) REFERENCES groups(id)
);
--
-- DROP TABLE chats;
-- DROP TABLE group_members;
-- DROP TABLE groups;
-- DROP TABLE users;