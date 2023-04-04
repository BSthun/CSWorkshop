package hub

import (
	"time"

	"github.com/getsentry/sentry-go"

	"github.com/go-co-op/gocron"
	"gorm.io/gorm"

	"backend/modules/config"
	"backend/modules/db"
	"backend/modules/db/model"
	"backend/utils/text"
)

var Hub *HubModel

type HubModel struct {
	Branches map[uint64]*Branch `json:"branches"`
}

type Branch struct {
	DBDsn             string      `json:"dbDsn"`
	DB                *gorm.DB    `json:"db"`
	Profile           *model.User `json:"profile"`
	AccessToken       string      `json:"accessToken"`
	AccessTokenExpire time.Time   `json:"accessTokenExpire"`
}

func Init() {
	// Defer recover
	defer func() {
		if err := recover(); err != nil {
			_ = sentry.Recover()
			if config.C.Environment == 1 {
				panic(err)
			}
		}
	}()

	// * Create hub
	Hub = &HubModel{
		Branches: make(map[uint64]*Branch),
	}

	// * Load branches
	var users []*model.User
	if result := db.DB.Preload("Client").Find(&users); result.Error != nil {
		panic(result.Error)
	}

	// * Apply branches
	for _, user := range users {
		dsn, db := Create(user)
		if db == nil {
			continue
		}
		Hub.Branches[*user.Id] = &Branch{
			DBDsn:             dsn,
			DB:                db,
			Profile:           user,
			AccessToken:       "",
			AccessTokenExpire: time.Time{},
		}
	}

	// * Schedule update

	s := gocron.NewScheduler(text.BangkokTime)
	_, _ = s.Every(1 * time.Minute).Do(Scrape)
	s.StartAsync()
}
