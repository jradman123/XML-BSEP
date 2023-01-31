
--Admin sistema ,lozinka je Admin12345!
INSERT INTO public.agent_user(
	date_of_birth, email, first_name, gender, is_confirmed, last_name, password, phone_number, recovery_email, role, username)
	VALUES ('1999-06-12 00:00:00', 'admin123@gmail.com', 'Anđela', 1, true, 'Bojčić', '$2a$10$SqMC5jOgeB8f0qmQgJ685O7E6NvdXbKDV9qPAmMygg1R0que5xhjS', '066545545', 'psw.company2@gmail.com',0, 'admin123');