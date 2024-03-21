package schema

type ApplyConferenceRequest struct {
	ConferenceName    string `json:"conference_name" validate:"required"`
	ConferenceAbbr    string `json:"conference_abbr" validate:"required"`
	ConferenceStartAt string `json:"conference_start_at" validate:"required,datetime=2006-01-02T15:04:05"`
	ConferenceEndAt   string `json:"conference_end_at" validate:"required,datetime=2006-01-02T15:04:05"`
	Venue             string `json:"venue" validate:"required"`
}

type AuditConferenceRequest struct {
	ApplicationId   int  `json:"application_id" validate:"required"`
	ApplicationPass bool `json:"application_pass" validate:"required"`
}

type StartSubmissionRequest struct {
	ConferenceName         int    `json:"conference_name" validate:"required"`
	SubmissionDeadline     string `json:"submission_deadline" validate:"required,datetime=2006-01-02T15:04:05"`
	ReviewResultAnnounceAt string `json:"review_result_announce_at" validate:"required,datetime=2006-01-02T15:04:05"`
}

type ConferenceName struct {
	ConferenceName string `json:"conference_name" validate:"required"`
}
