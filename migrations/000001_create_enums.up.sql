CREATE TYPE "vacancy_status" AS ENUM (
  'opened',
  'closed'
);

CREATE TYPE "user_role" AS ENUM (
  'hr',
  'team_lead_hr'
);

CREATE TYPE "resume_stages" AS ENUM (
  'opened',
  'studied',
  'interview',
  'passed_interview',
  'technical_interview_scheduled',
  'technical_interview_passed',
  'offer_sent'
);
