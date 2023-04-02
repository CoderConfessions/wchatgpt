package mysqlwrapper

import (
	"database/sql"
	"fmt"
	"sort"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sashabaranov/go-openai"
	"k8s.io/klog/v2"
)

var pool *sql.DB

type DBConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Schema   string `json:"schema"`
	Port     int    `json:"port"`
}

func InitPool(conf DBConfig) error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Schema)
	pool, err = sql.Open("mysql", dsn)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		klog.Fatal("unable to use data source name", err)
	}
	return err
}

func ReleasePool() {
	pool.Close()
}

type TableUserRecord struct {
	userid          string
	total_use_count int
	version         int
}

func GetUserByUserId(userid string) (*TableUserRecord, error) {
	rows, err := pool.Query("SELECT * FROM user WHERE userid = $1", userid)
	if err != nil {
		return nil, err
	}
	var rec TableUserRecord
	rows.Next()
	err = rows.Scan(&rec.userid, &rec.total_use_count, &rec.version)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%v\n", rec)
	return &rec, nil
}

func UpdateChatID(userUID, chatID string) error {
	_, err := pool.Exec("INSERT chat_id (user_id, chat_id, creation_time) VALUES (?, ?, ?) ", userUID, chatID, time.Now())
	return err
}

func GetUserUIDByChatID(chatID string) (userUID string, err error) {
	rows := pool.QueryRow("SELECT user_id FROM chat_id WHERE chat_id = ?", chatID)
	err = rows.Scan(&userUID)
	return
}

func GetHistoryMessageByChatID(chatID string) ([]openai.ChatCompletionMessage, error) {
	rows, err := pool.Query("SELECT role, content FROM chat_data WHERE chat_id = ? ORDER BY idx DESC LIMIT 10 ", chatID)
	if err != nil {
		return nil, err
	}
	messages := make([]openai.ChatCompletionMessage, 0)
	for rows.Next() {
		message := openai.ChatCompletionMessage{}
		rows.Scan(&message.Role, &message.Content)
		messages = append(messages, message)
	}
	sort.SliceStable(messages, func(i, j int) bool {
		return i > j
	})
	return messages, nil
}

func UpdateHistoryMessageByChatID(chatID string, newMessages []openai.ChatCompletionMessage) error {
	for _, m := range newMessages {
		_, err := pool.Exec("INSERT chat_data(chat_id, role, content, creation_time) VALUES(? , ?, ?, ?)", chatID, m.Role, m.Content, time.Now())
		if err != nil {
			return err
		}
	}
	return nil
}
