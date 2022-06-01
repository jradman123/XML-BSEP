
--Admin sistema ,lozinka je Admin123!
INSERT INTO public.agent_user(
	date_of_birth, email, first_name, gender, is_confirmed, last_name, password, phone_number, recovery_email, role, username)
	VALUES ('1999-06-12 00:00:00', 'admin123@gmail.com', 'Anđela', 0, true, 'Bojčić', '$2a$10$q/8SyLsTjkr8yNvB8/L3jeNvQ7PsCqDS57cjGHld.1OdeOYMbAIFO', '066545545', 'psw.company2@gmail.com',0, 'admin123');