alter table users
add column if not exists role_id INTEGER
references user_role(id);