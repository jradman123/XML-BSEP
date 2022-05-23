insert into authorities(id, name) values (1, 'ADMIN');
insert into authorities(id, name) values (2, 'USER_ROOT');
insert into authorities(id, name) values (3, 'USER_INTERMEDIATE');
insert into authorities(id, name) values (4, 'USER_END_ENTITY');

-- role: admin, user_root, user_ica, user_ee

INSERT INTO public.users(email, is_activated, password, recovery_email, role)
VALUES ( 'admin123@gmail.com', true, '$2a$10$C8y2GXceEjgc48nr1r5rt.KP0m1OBEfubQbagvEYDr/2DQE0Bddka', 'psw.company2@gmail.com', 0);

INSERT INTO public.user_authorities(user_id, authority_id) VALUES (1, 1);