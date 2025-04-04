BEGIN;

ALTER TABLE vacancies DROP CONSTRAINT fk_vacancies_created_by;

ALTER TABLE sla_rules 
DROP CONSTRAINT fk_sla_rules_vacancy,
DROP CONSTRAINT fk_sla_rules_stage,
DROP CONSTRAINT fk_sla_rules_created_by;

ALTER TABLE resumes 
DROP CONSTRAINT fk_resumes_vacancy,
DROP CONSTRAINT fk_resumes_stage;

ALTER TABLE resume_histories 
DROP CONSTRAINT fk_histories_resume,
DROP CONSTRAINT fk_histories_stage,
DROP CONSTRAINT fk_histories_user;

COMMIT;