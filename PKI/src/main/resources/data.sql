insert into authorities(id, name) values (1, 'ADMIN');
insert into authorities(id, name) values (2, 'USER_ROOT');
insert into authorities(id, name) values (3, 'USER_INTERMEDIATE');
insert into authorities(id, name) values (4, 'USER_END_ENTITY');

-- role: admin, user_root, user_ica, user_ee
--password AdminUser123!
INSERT INTO public.users(
    common_name, country, email, is_activated, locality, organization, organization_unit, password, recovery_email, role)
    VALUES ('Admin Main', 'SRB', 'admin123@gmail.com', true, 'NS', 'UNS', 'FTN', '$2a$10$RIVr/hKkPlf/fq/INofwse926OZPkElV9HXnPwCmnUZHiHgXtniim', 'psw.company2@gmail.com', 0);
--password UserAdmin123!
INSERT INTO public.users(
    common_name, country, email, is_activated, locality, organization, organization_unit, password, recovery_email, role)
    VALUES ('Marko Markovic', 'SRB', 'markomarkovic123@gmail.com', true, 'NS', 'UNS', 'FTN', '$2a$10$.sq6OHYCOBy3nDql82vsX.E0srnWWHZwSD31SJmAm02Xku5fZ8k9y', 'raandmjenale@gmail.com', 1);

INSERT INTO public.user_authorities(user_id, authority_id) VALUES (1, 1);
INSERT INTO public.user_authorities(user_id, authority_id) VALUES (2, 4);
