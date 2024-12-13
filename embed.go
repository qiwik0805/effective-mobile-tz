package effective_mobile_tz

import "embed"

//go:embed migrations/*.sql
var Migrations embed.FS
