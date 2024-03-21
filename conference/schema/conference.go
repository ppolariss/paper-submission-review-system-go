package schema

import "time"

var STATUS = []string{"准备中", "投稿中", "审稿中", "终评中", "审稿结束"}

type Conference struct {
	ConferenceName         string    `json:"conference_name"`
	ConferenceAbbr         string    `json:"conference_abbr"`
	ConferenceStartAt      time.Time `json:"conference_start_at"`
	ConferenceEndAt        time.Time `json:"conference_end_at"`
	Venue                  string    `json:"venue"`
	SubmissionDeadline     time.Time `json:"submission_deadline"`
	ReviewResultAnnounceAt string    `json:"review_result_announce"`
	ConferenceStatus       string    `json:"conference_status"`
}

type ConferenceApplication struct {
	ApplicationID     int       `json:"application_id"`
	ConferenceName    string    `json:"conference_name"`
	ConferenceAbbr    string    `json:"conference_abbr"`
	ConferenceStartAt time.Time `json:"conference_start_at"`
	ConferenceEndAt   time.Time `json:"conference_end_at"`
	Venue             string    `json:"venue"`
	ApplicationStatus string    `json:"application_status"`
	ApplicantUsername string    `json:"applicant_username"`
}

//private Integer applicationID;     // 申请ID，主键
//
//private String conferenceName; // 会议全称
//
//private String conferenceAbbr; // 会议简称
//
//private LocalDateTime conferenceStartAt;    // 会议举办时间-左端点
//
//private LocalDateTime conferenceEndAt;      // 会议举办时间-右端点
//
//private String venue;          // 会议举办地点
//
//private String applicationStatus;          // 申请状态（审核中，已通过，未通过）
//
//private String applicantUsername;          // 申请人用户名
