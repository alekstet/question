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
	UserNickname string `json:"user_nickname"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
}

type UserQuestion struct {
	QuestionId   string `json:"question_id"`
	UserNickname string `json:"user_nickname"`
	Answer       string `json:"answer"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
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

func (u *UserQuestion) Valid() bool {
	if u.QuestionId != "" && u.UserNickname != "" && u.Answer != "" {
		return true
	}
	return false
}

func (u *UserQuestion) Init(questionId, userNickname, answer, createdAt, updatedAt string) {
	u.QuestionId = questionId
	u.UserNickname = userNickname
	u.Answer = answer
	u.CreatedAt = createdAt
	u.UpdatedAt = updatedAt
}

func (q *Question) Valid() bool {
	if q.Date != "" && q.Question != "" {
		return true
	}

	return false
}

func (q *Question) Init(date, question string) {
	q.Date = date
	q.Question = question
}
