alter table users
add column if not exists role_id INTEGER
references user_role(id);

alter table users
add constraint fk_user_role
foreign key (role_id)
references user_role(id)
on delete set NULL;