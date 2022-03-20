package models

type Question struct {
	Date     string `json:"date"`
	Question string `json:"question"`
}

type UsersAuth struct {
	Nickname string `json:"nickname"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UsersData struct {
	UserNickname string `json:"userNickname"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
}

type User_question struct {
	QuestionId   string `json:"questionId"`
	UserNickname string `json:"userNickname"`
	Answer       string `json:"answer"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type TodaysInfo struct {
	Question string         `json:"question"`
	Answers  []TodaysAnswer `json:"answers"`
}

type TodaysAnswer struct {
	Nickname string `json:"nickname"`
	Answer   string `json:"answer"`
}

type UserInfo struct {
	Nickname string     `json:"nickname"`
	Name     string     `json:"name"`
	Sex      string     `json:"sex"`
	Answers  []UserAnsw `json:"answers"`
}

type UserAnsw struct {
	Date     string `json:"date"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
