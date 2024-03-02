package mongodb

import "net/url"

type Config struct {
	Host        string `env:"MONGO_HOST" json:"host"`
	Port        string `env:"MONGO_PORT" json:"port"`
	User        string `env:"MONGO_USER" json:"user"`
	Password    string `env:"MONGO_PASSWORD" json:"-"`
	Name        string `env:"MONGO_DB_NAME" json:"name"`
	MaxPoolSize string `env:"MONGO_MAX_POOL_SIZE,default=20" json:"max_pool_size"`
	Write       string `env:"MONGO_WRITE,default=majority" json:"write"`
}

func (c *Config) ConnectionString() string {
	if c == nil {
		return ""
	}

	host := c.Host
	if v := c.Port; v != "" {
		host = host + ":" + v
	}

	u := &url.URL{
		Scheme: "mongodb",
		Host:   host,
		Path:   c.Name,
	}

	if c.User != "" || c.Password != "" {
		u.User = url.UserPassword(c.User, c.Password)
	}

	q := u.Query()
	if v := c.MaxPoolSize; v != "" {
		q.Add("maxPoolSize", v)
	}
	if v := c.Write; v != "" {
		q.Add("w", v)
	}
	u.RawQuery = q.Encode()

	return u.String()
}
