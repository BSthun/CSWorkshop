package ihub

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"

	"backend/types/extern"
	"backend/types/model"
)

func GradeSchedule() {
	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Every(1 * time.Second).Do(GradeScheduleAction)
	s.StartAsync()
}

func GradeScheduleAction() {
	// * Query sql slow log to struct
	rows, err := b.SqlDB.Query("SELECT * FROM mysql.slow_log")
	if err != nil {
		logrus.Warn(err)
		return
	}

	// * Construct debounce
	type bounce struct {
		Session    *extern.Session
		Submission *model.Submission
	}
	debounce := make(map[uint64]*bounce)
	var wg sync.WaitGroup

	// * Iterate rows
	for rows.Next() {
		// * Scan row to struct
		var (
			startTime    time.Time
			userHost     string
			queryTime    string
			lockTime     string
			rowsSent     int
			rowsExamined int
			db           string
			lastInsertId int
			insertId     int
			serverId     int
			sqlText      string
			threadId     int
			rowsAffected int
		)
		if err := rows.Scan(&startTime, &userHost, &queryTime, &lockTime, &rowsSent, &rowsExamined, &db, &lastInsertId, &insertId, &serverId, &sqlText, &threadId, &rowsAffected); err != nil {
			logrus.Warn(err)
			return
		}

		if strings.Contains(userHost, "[root]") {
			continue
		}

		if !strings.HasPrefix(db, "lab_") {
			continue
		}

		if strings.Contains(strings.ToLower(sqlText), "information_schema") {
			continue
		}

		if strings.Contains(strings.ToLower(sqlText), "@@tx_isolation") {
			continue
		}

		if strings.Contains(strings.ToLower(sqlText), "select database()") {
			continue
		}

		if !strings.Contains(strings.ToLower(sqlText), "select") {
			continue
		}

		session, ok := h.SessionDbNameMap[db]
		if !ok {
			continue
		}

		if _, ok := h.Mocks[*session.Id]; ok {
			continue
		}

		// * Dump all data
		fmt.Println(startTime, userHost, queryTime, lockTime, rowsSent, rowsExamined, db, lastInsertId, insertId, serverId, sqlText, threadId, rowsAffected)

		// * Grade
		wg.Add(1)
		go func() {
			bn := &bounce{
				session,
				GradeSubmit(session, startTime, sqlText),
			}

			if bn.Session.CurrentTask == nil {
				wg.Done()
				return
			}
			debounce[*bn.Session.Id] = bn
			wg.Done()
		}()
	}

	// * Clear sql general log
	if _, err := b.SqlDB.Exec("TRUNCATE mysql.slow_log"); err != nil {
		spew.Dump("TRUNCATE SLOW LOG", err)
		return
	}

	wg.Wait()

	// * Grade debounce
	for _, bn := range debounce {
		GradePasser(bn.Session, bn.Submission)
	}
}
