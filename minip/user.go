package minip

const (
	UserGenderUnkown int = 0
	UserGenderMale   int = 1
	UserGenderFemale int = 2
)

type UserInfo struct {
	NicknName string `json:"nickName"`

	//用户头像图片的 URL。URL 最后一个数值代表正方形头像大小（有 0、46、64、96、132 数值可选，0 代表 640x640 的正方形头像，46 表示 46x46 的正方形头像，剩余数值以此类推。默认132），用户没有头像时该项为空。若用户更换头像，原有头像 URL 将失效。
	AvatarURL string `json:"avatarUrl"`

	// 0	未知
	// 1	男性
	// 2	女性
	Gender int `json:"gender"`

	// 用户所在国家
	Country string `json:"contry"`

	// 用户所在省份
	Province string `json:"province"`

	// 用户所在城市
	City string `json:"city"`

	// 显示 country，province，city 所用的语言
	// en	英文
	// zh_CN	简体中文
	// zh_TW	繁体中文
	Language string `json:"language"`
}
