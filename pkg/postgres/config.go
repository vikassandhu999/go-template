package postgres

import (
	"net/url"
)

type Config struct {
	Host     string `env:"PG_DB_HOST" json:"host"`
	Port     string `env:"PG_DB_PORT" json:"port"`
	Name     string `env:"PG_DB_NAME" json:"name"`
	User     string `env:"PG_DB_USER" json:"user"`
	Password string `env:"PG_DB_PASSWORD" json:"-"`
	SSLMode  string `env:"PG_DB_SSLMODE" json:"ssl_mode"`
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
		Scheme: "postgres",
		Host:   host,
		Path:   c.Name,
	}

	if c.User != "" || c.Password != "" {
		u.User = url.UserPassword(c.User, c.Password)
	}

	q := u.Query()
	if v := c.SSLMode; v != "" {
		q.Add("sslmode", v)
	}
	u.RawQuery = q.Encode()

	return u.String()
}
