BEGIN;

ALTER TABLE vacancies 
ADD CONSTRAINT fk_vacancies_created_by 
FOREIGN KEY (created_by) REFERENCES users(id) 
ON DELETE CASCADE;

ALTER TABLE sla_rules 
ADD CONSTRAINT fk_sla_rules_vacancy 
FOREIGN KEY (vacancy_id) REFERENCES vacancies(id) 
ON DELETE CASCADE;

ALTER TABLE sla_rules 
ADD CONSTRAINT fk_sla_rules_stage 
FOREIGN KEY (stage_id) REFERENCES stages(id) 
ON DELETE CASCADE;

ALTER TABLE sla_rules 
ADD CONSTRAINT fk_sla_rules_created_by 
FOREIGN KEY (created_by) REFERENCES users(id) 
ON DELETE CASCADE;

ALTER TABLE resumes 
ADD CONSTRAINT fk_resumes_vacancy 
FOREIGN KEY (vacancy_id) REFERENCES vacancies(id) 
ON DELETE CASCADE;

ALTER TABLE resumes 
ADD CONSTRAINT fk_resumes_stage 
FOREIGN KEY (current_stage_id) REFERENCES stages(id) 
ON DELETE CASCADE;

ALTER TABLE resume_histories 
ADD CONSTRAINT fk_histories_resume 
FOREIGN KEY (resume_id) REFERENCES resumes(id) 
ON DELETE CASCADE;

ALTER TABLE resume_histories 
ADD CONSTRAINT fk_histories_stage 
FOREIGN KEY (stage_id) REFERENCES stages(id) 
ON DELETE CASCADE;

ALTER TABLE resume_histories 
ADD CONSTRAINT fk_histories_user 
FOREIGN KEY (user_id) REFERENCES users(id) 
ON DELETE CASCADE;

COMMIT;