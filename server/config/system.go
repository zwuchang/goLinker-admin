package config

type System struct {
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`    // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`    // 多点登录拦截
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                   // 使用redis
	UseMongo      bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`                   // 使用mongo
	UseStrictAuth bool   `mapstructure:"use-strict-auth" json:"use-strict-auth" yaml:"use-strict-auth"` // 使用树形角色分配模式
	InviteBaseURL string `mapstructure:"invite-base-url" json:"invite-base-url" yaml:"invite-base-url"` // 邀请链接基础URL
	HTTPProxy     string `mapstructure:"http-proxy" json:"http-proxy" yaml:"http-proxy"`                // http代理
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                     // 环境 dev|prod
	IpDbPath      string `mapstructure:"ip-db-path" json:"ip-db-path" yaml:"ip-db-path"`                // IP地址库路径
	// 是否开启 1 号超级用户
	IsSuperUser bool `mapstructure:"is-super-user" json:"is-super-user" yaml:"is-super-user"` // 是否开启 1 号超级用户
}
