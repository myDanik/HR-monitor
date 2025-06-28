ALTER TABLE resumes
ADD COLUMN sladeadline timestamp DEFAULT (now());
