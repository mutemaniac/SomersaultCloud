package SomersaultCloud

import "SomersaultCloud/utils/environment"

var kongAdminURL string

func init() {
	kongAdminURL = environment.GetEnv("KONG_ADMIN_URL", "http://localhost:8001")
}
