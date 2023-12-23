package postgres

import (
	"net/url"
)

type Config struct {
	Host     string `env:"PG_DB_HOST"`
	Port     string `env:"PG_DB_PORT"`
	Name     string `env:"PG_DB_NAME"`
	User     string `env:"PG_DB_USER"`
	Password string `env:"PG_DB_PASSWORD"`
	SSLMode  string `env:"PG_DB_SSLMODE"`
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
