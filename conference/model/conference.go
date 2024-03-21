package model

import "time"

type Conference struct {
	ConferenceName         string    `json:"conference_name" gorm:"primary_key length:50"`
	ConferenceAbbr         string    `json:"conference_abbr" gorm:"length:50"`
	ConferenceStartAt      time.Time `json:"conference_start_at"`
	ConferenceEndAt        time.Time `json:"conference_end_at"`
	Venue                  string    `json:"venue" gorm:"length:100"`
	SubmissionDeadline     time.Time `json:"submission_deadline"`
	ReviewResultAnnounceAt time.Time `json:"reviewResult_announce_at"`
	ConferenceStatus       string    `json:"conference_status" gorm:"length:10"`
	CreatedAt              time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt              time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type ConferenceApplication struct {
	ID                int       `json:"id" gorm:"primary_key autoIncrement"`
	ConferenceName    string    `json:"conference_name" gorm:"length:50"`
	ConferenceAbbr    string    `json:"conference_abbr" gorm:"length:50"`
	ConferenceStartAt time.Time `json:"conference_start_at"`
	ConferenceEndAt   time.Time `json:"conference_end_at"`
	Venue             string    `json:"venue" gorm:"length:100"`
	ApplicationStatus string    `json:"application_status" gorm:"length:10"`
	ApplicantUsername string    `json:"applicant_username" gorm:"length:50"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

//func GetConferenceByName(name string) (conference Conference, err error) {
//	err = db.Where("conference_name = ?", name).First(&conference).Error
//	return
//}
