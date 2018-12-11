package config

// AppType
type AppType string

const (
	PlatformMP          AppType = "mp"
	PlatformMiniProgram AppType = "minip"
	PlatformMiniGame    AppType = "minig"
)

type AppConfig struct {
	AppType   AppType `json:"app_type"`
	AppID     string  `json:"app_id"`
	AppSecret string  `json:"app_secret"`
}
