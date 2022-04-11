package helper

import (
	"database/sql"
	"fmt"
	"github.com/tejashwikalptaru/tutorial/database"
	"github.com/tejashwikalptaru/tutorial/models"
)

func CreateUser(name, email, password string) (string, error) {
	// language=SQL
	SQLEsist := `SELECT id ,name ,email, password From users Where email = $1 and password = $2`
	var user models.User
	err := database.Tutorial.Get(&user, SQLEsist, email, password)
	if user.Email == email && user.Password == password {
		return "Data Already Exist", nil
	}

	SQL := `INSERT INTO users(name, email , password) VALUES ($1, $2 , $3) RETURNING id;`
	var userID string
	err = database.Tutorial.Get(&userID, SQL, name, email, password)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func GetUserBySession(sessionId string) (*models.User, error) {

	SQL := `select users.id , name , email , created_at , archived_at , password from
			users join sessions on users.id = sessions.uid where sessions.id = $1 and sessions.log_out is NULL`

	var user models.User
	err := database.Tutorial.Get(&user, SQL, sessionId)

	if err != nil {
		fmt.Println("GetUserBySession Error")
		return nil, err
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}

func GetUser(userID string) (*models.User, error) {
	//language=SQL
	SQL := `SELECT id, name, email, created_at, archived_at FROM users WHERE archived_at IS NULL AND id = $1`
	var user models.User
	err := database.Tutorial.Get(&user, SQL, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &user, nil
}

func LogOutDetails() {

}

func LoginDetails(email, pass string) (bool, string) {
	SQL := `SELECT id ,name ,email, password From users Where email = $1 and password = $2`
	var user models.User
	err := database.Tutorial.Get(&user, SQL, email, pass)
	if err != nil {
		//	User is not present in the user Table
		return false, " "
	}
	return true, user.ID
}

func SessionLogin(userid string) (error, string) {
	SQL := `INSERT INTO sessions(uid) values($1) returning id`
	var token string
	err := database.Tutorial.Get(&token, SQL, userid)

	if err != nil {
		return err, ""
	}
	return err, token
}

func CreateTask(task models.Task, user models.User) (string, error) {

	SQL := `INSERT INTO task(taskname , taskID , completed_upto) values($1 , $2 , $3) returning uniqueid`
	var TaskId string
	err := database.Tutorial.Get(&TaskId, SQL, task.TaskName, user.ID, "2008-11-09 15:45:21")

	if err != nil {
		return " ", err
	}
	return TaskId, err
}

func GetTask(user *models.User) []string {
	SQL := `SELECT taskName from task where taskid = $1`
	var taskName = make([]string, 0)
	err := database.Tutorial.Select(&taskName, SQL, user.ID)
	if err != nil {
		fmt.Println("Error to get data from table")
		return nil
	}
	return taskName

}

func SetLogout(user *models.User) error {
	// language=sql
	SQL := `Update sessions set log_out = now() where uid = $1`
	_, err := database.Tutorial.Exec(SQL, user.ID)
	if err != nil {

		return err
	}
	return nil
}
