package kong

import "service-cloud/utils/environment"

var kongAdminURL string

func init() {
	kongAdminURL = environment.GetEnv("KONG_ADMIN_URL", "http://hnadokku.cloudapp.net:8001")
}
