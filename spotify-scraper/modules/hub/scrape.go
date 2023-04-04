package hub

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"

	"backend/functions"
)

func Scrape() {
	for _, branch := range Hub.Branches {
		// * Check access token
		if time.Now().After(branch.AccessTokenExpire) {
			// * Refresh access token
			cred, err := functions.SpotifyGetRefreshToken(nil, branch.Profile.Client, "refresh_token", *branch.Profile.RefreshToken)
			if err != nil {
				sentry.CaptureException(err)
				continue
			}
			branch.AccessToken = *cred.AccessToken
			branch.AccessTokenExpire = time.Now().Add(time.Duration(*cred.ExpiresIn) * time.Second)
		}

		// * Get playback state
		state, err := functions.SpotifyGetPlaybackState(nil, branch.AccessToken)
		if err != nil {
			logrus.WithField("e", err).WithField("u", *branch.Profile.Id).Warn("Unable to get Spotify playback state")
			sentry.CaptureException(err)
			continue
		}

		// * Check if playback is active
		if state == nil || state.Item == nil || !state.IsPlaying || *state.Item.Type != "track" {
			continue
		}

		// * Record playback
		if err := RecordLv1(branch.DB, state); err != nil {
			sentry.CaptureException(err)
			continue
		}
		if err := RecordLv2(branch.DB, state); err != nil {
			sentry.CaptureException(err)
			continue
		}
		if err := RecordLv3(branch.DB, state); err != nil {
			sentry.CaptureException(err)
			continue
		}
	}
}
