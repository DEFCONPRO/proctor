package middleware

import (
	"fmt"
	"github.com/hashicorp/go-version"
	"net/http"
	"proctor/proctord/config"
	"proctor/proctord/logger"
	utility "proctor/shared/constant"
)

func ValidateClientVersion(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		requestHeaderClientVersion := r.Header.Get(utility.ClientVersionHeaderKey)

		if requestHeaderClientVersion != "" {
			clientVersion, err := version.NewVersion(requestHeaderClientVersion)
			if err != nil {
				logger.Error("Error while creating requestHeaderClientVersion", err.Error())
			}

			minClientVersion, err := version.NewVersion(config.MinClientVersion())
			if err != nil {
				logger.Error("Error while creating minClientVersion", err.Error())
			}

			if clientVersion.LessThan(minClientVersion) {
				w.WriteHeader(400)
				w.Write([]byte(fmt.Sprintf(utility.ClientOutdatedErrorMessage, clientVersion)))
				return
			}
			next.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	}
}
