package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ScheduleResponse struct {
	ErrorCode    int           `json:"error_code"`
	ErrorMessage string        `json:"error_message"`
	Result       []DaySchedule `json:"result"`
}

type DaySchedule struct {
	Date    time.Time `json:"section_id"`
	Lessons []Lesson  `json:"lessons"`
}

type Lesson struct {
	Id            int           `json:"id"`
	Date          time.Time     `json:"date"`
	DateEnd       time.Time     `json:"date_end"`
	SectionId     int           `json:"section_id"`
	SectionName   string        `json:"section_name"`
	SectionLevel  int           `json:"section_level"`
	LessonGroupId int           `json:"lesson_group_id"`
	LessonLevel   int           `json:"lesson_level"`
	TypeId        int           `json:"type_id"`
	BuildingId    int           `json:"building_id"`
	RoomId        int           `json:"room_id"`
	RoomName      string        `json:"room_name"`
	LinkUrl       string        `json:"link_url"`
	Limit         int           `json:"limit"`
	Available     int           `json:"available"`
	Comment       string        `json:"comment"`
	TimeSlotId    int           `json:"time_slot_id"`
	TimeSlotStart string        `json:"time_slot_start"`
	TimeSlotEnd   string        `json:"time_slot_end"`
	Intersection  bool          `json:"intersection"`
	CanSignIn     CanSignIn     `json:"can_sign_in"`
	OtherLessons  []OtherLesson `json:"other_lessons"`
	Signed        bool          `json:"signed"`
	TeacherIsu    int           `json:"teacher_isu"`
	TeacherFio    string        `json:"teacher_fio"`
}

type CanSignIn struct {
	CanSignIn          bool     `json:"can_sign_in"`
	UnavailableReasons []string `json:"unavailable_reasons"`
}

type OtherLesson struct {
	Id           int       `json:"id"`
	DateStart    time.Time `json:"date_start"`
	DateEnd      time.Time `json:"date_end"`
	Signed       bool      `json:"signed"`
	Intersection bool      `json:"intersection"`
	CanSignIn    CanSignIn `json:"can_sign_in"`
}

func (cfg *AuthConfig) getSchedule() ([]DaySchedule, error) {
	client := &http.Client{}

	buildingId := "273"
	dateStart := "2024-09-02"
	dateEnd := "2024-09-29"
	req, err := http.NewRequest("GET", fmt.Sprintf("https://my.itmo.ru/api/sport/sign/schedule?building_id=%s&date_start=%s&date_end=%s", buildingId, dateStart, dateEnd), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+cfg.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		err = cfg.Refresh()
		fmt.Println("Trying to refresh.")
		if err != nil {
			return nil, err
		}
	}

	var schedule ScheduleResponse
	err = json.NewDecoder(resp.Body).Decode(&schedule)
	if err != nil {
		return nil, err
	}

	return schedule.Result, err
}
