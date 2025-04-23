package enums

type ResumeSortField string

const (
	CreatedAt ResumeSortField = "created_at"
	SLADeadline ResumeSortField = "sla_deadline"
)

type ResumeSortDirection string

const (
	Asc ResumeSortDirection = "asc"
	Desc ResumeSortDirection = "desc"
)

type ResumeStage string

const (
	Opened ResumeStage = "opened"
	Studied ResumeStage = "studied"
	Interview ResumeStage = "interview"
	PassedInterview ResumeStage = "passed_interview"
	TechnicalInterviewScheduled ResumeStage = "technical_interview_scheduled"
	TechnicalInterviewPassed ResumeStage = "technical_interview_passed"
	OfferSent ResumeStage = "offer_sent"
)