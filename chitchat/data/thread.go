package data

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func GetThreads() ([]Thread, error) {
	rows, err := Db.Query("SELECT * FROM threads ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	threads := make([]Thread, 0)
	for rows.Next() {
		th := Thread{}
		// 因为 Query 时写的是 *，所以 Scan 时列名要和 Thread 结构对应
		if err := rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			return nil, err
		}
		threads = append(threads, th)
	}
	return threads, nil
}
