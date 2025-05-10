UPDATE resumes SET sladeadline = NULL;

ALTER TABLE resumes
ALTER COLUMN sladeadline TYPE interval USING NULL; 