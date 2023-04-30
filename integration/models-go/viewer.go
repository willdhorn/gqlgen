package models

import "github.com/willdhorn/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
