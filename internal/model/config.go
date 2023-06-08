package model

import "time"

type (
	Config struct {
		App            App            `json:"app"`
		PostgresClient PostgresClient `json:"postgres_client"`
	}

	App struct {
		Env             string        `json:"env"`
		Port            int           `json:"port"`
		Name            string        `json:"name"`
		LogOption       string        `json:"log_option"`
		LogLevel        string        `json:"log_level"`
		HTTPTimeout     time.Duration `json:"http_timeout"`
		GracefulTimeout time.Duration `json:"graceful_timeout"`
		RPCAddress      string        `json:"rpc_address"`
		RPCInsecure     bool          `json:"rpc_insecure"`
	}

	PostgresClient struct {
		Db          string `json:"db"`
		Host        string `json:"host"`
		Username    string `json:"username"`
		Password    string `json:"password"`
		Port        string `json:"port"`
		SslMode     string `json:"ssl_mode"`
		MaxIdleConn int    `json:"max_idle_conn"`
		MaxOpenConn int    `json:"max_open_conn"`
		DebugMode   bool   `json:"debug_mode"`
	}
)
