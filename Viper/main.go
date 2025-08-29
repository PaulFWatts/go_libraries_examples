package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Configuration structure that matches our config files
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	Logging  LoggingConfig  `mapstructure:"logging"`
	Features FeatureFlags   `mapstructure:"features"`
	Security SecurityConfig `mapstructure:"security"`
}

type ServerConfig struct {
	Host           string        `mapstructure:"host"`
	Port           int           `mapstructure:"port"`
	ReadTimeout    time.Duration `mapstructure:"read_timeout"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout"`
	MaxConnections int           `mapstructure:"max_connections"`
	TLS            TLSConfig     `mapstructure:"tls"`
}

type TLSConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	CertFile string `mapstructure:"cert_file"`
	KeyFile  string `mapstructure:"key_file"`
}

type DatabaseConfig struct {
	Driver          string        `mapstructure:"driver"`
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Database        string        `mapstructure:"database"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	MaxConnections  int           `mapstructure:"max_connections"`
	MaxIdleTime     time.Duration `mapstructure:"max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
	PoolSize int    `mapstructure:"pool_size"`
}

type LoggingConfig struct {
	Level      string `mapstructure:"level"`
	Format     string `mapstructure:"format"`
	Output     string `mapstructure:"output"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type FeatureFlags struct {
	EnableMetrics   bool `mapstructure:"enable_metrics"`
	EnableTracing   bool `mapstructure:"enable_tracing"`
	EnableProfiling bool `mapstructure:"enable_profiling"`
	EnableCaching   bool `mapstructure:"enable_caching"`
	BetaFeatures    bool `mapstructure:"beta_features"`
}

type SecurityConfig struct {
	JWTSecret       string        `mapstructure:"jwt_secret"`
	JWTExpiration   time.Duration `mapstructure:"jwt_expiration"`
	RateLimitRPS    int           `mapstructure:"rate_limit_rps"`
	RateLimitBurst  int           `mapstructure:"rate_limit_burst"`
	CORSOrigins     []string      `mapstructure:"cors_origins"`
	CSRFSecret      string        `mapstructure:"csrf_secret"`
	EnableHTTPSOnly bool          `mapstructure:"enable_https_only"`
}

var (
	cfgFile    string
	config     Config
	configType string
	envPrefix  string
)

// Root command
var rootCmd = &cobra.Command{
	Use:   "viper-demo",
	Short: "Viper configuration management demonstration",
	Long: `A comprehensive demonstration of Viper configuration management library showing:
- Loading configuration from multiple file formats (JSON, YAML, TOML)
- Environment variable override support
- Command-line flag integration
- Configuration hot-reloading
- Configuration validation and defaults
- Complex nested configuration structures`,
	Run: func(cmd *cobra.Command, args []string) {
		runDemo()
	},
}

// Commands
var showConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Show current configuration",
	Long:  "Display the current configuration with all values resolved from files, environment variables, and defaults",
	Run: func(cmd *cobra.Command, args []string) {
		showConfiguration()
	},
}

var validateConfigCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate configuration",
	Long:  "Validate the current configuration against business rules and constraints",
	Run: func(cmd *cobra.Command, args []string) {
		validateConfiguration()
	},
}

var watchConfigCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch configuration changes",
	Long:  "Watch for configuration file changes and display updates in real-time",
	Run: func(cmd *cobra.Command, args []string) {
		watchConfiguration()
	},
}

var createSampleCmd = &cobra.Command{
	Use:   "create-samples",
	Short: "Create sample configuration files",
	Long:  "Create sample configuration files in different formats (JSON, YAML, TOML)",
	Run: func(cmd *cobra.Command, args []string) {
		createSampleConfigs()
	},
}

var envDemoCmd = &cobra.Command{
	Use:   "env-demo",
	Short: "Environment variable demonstration",
	Long:  "Show how environment variables can override configuration file values",
	Run: func(cmd *cobra.Command, args []string) {
		environmentDemo()
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default searches for config.{json,yaml,yml,toml} in current directory)")
	rootCmd.PersistentFlags().StringVar(&configType, "type", "yaml", "config file type (json, yaml, toml)")
	rootCmd.PersistentFlags().StringVar(&envPrefix, "env-prefix", "VIPERAPP", "environment variable prefix")

	// Server flags
	rootCmd.PersistentFlags().String("server.host", "localhost", "server host")
	rootCmd.PersistentFlags().Int("server.port", 8080, "server port")
	rootCmd.PersistentFlags().Bool("server.tls.enabled", false, "enable TLS")

	// Database flags
	rootCmd.PersistentFlags().String("database.driver", "postgres", "database driver")
	rootCmd.PersistentFlags().String("database.host", "localhost", "database host")
	rootCmd.PersistentFlags().Int("database.port", 5432, "database port")

	// Feature flags
	rootCmd.PersistentFlags().Bool("features.enable-metrics", false, "enable metrics collection")
	rootCmd.PersistentFlags().Bool("features.beta-features", false, "enable beta features")

	// Bind flags to viper
	viper.BindPFlags(rootCmd.PersistentFlags())

	// Add commands
	rootCmd.AddCommand(showConfigCmd)
	rootCmd.AddCommand(validateConfigCmd)
	rootCmd.AddCommand(watchConfigCmd)
	rootCmd.AddCommand(createSampleCmd)
	rootCmd.AddCommand(envDemoCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("$HOME/.viperapp")
		viper.SetConfigName("config")
		viper.SetConfigType(configType)
	}

	// Environment variables
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Set defaults
	setDefaults()

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("⚠️  No config file found, using defaults and environment variables")
		} else {
			fmt.Printf("❌ Error reading config file: %v\n", err)
		}
	} else {
		fmt.Printf("✅ Using config file: %s\n", viper.ConfigFileUsed())
	}

	// Unmarshal into struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}
}

func setDefaults() {
	// Server defaults
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.read_timeout", "30s")
	viper.SetDefault("server.write_timeout", "30s")
	viper.SetDefault("server.max_connections", 1000)
	viper.SetDefault("server.tls.enabled", false)
	viper.SetDefault("server.tls.cert_file", "")
	viper.SetDefault("server.tls.key_file", "")

	// Database defaults
	viper.SetDefault("database.driver", "postgres")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.username", "user")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.database", "myapp")
	viper.SetDefault("database.ssl_mode", "disable")
	viper.SetDefault("database.max_connections", 25)
	viper.SetDefault("database.max_idle_time", "15m")
	viper.SetDefault("database.conn_max_lifetime", "1h")

	// Redis defaults
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.database", 0)
	viper.SetDefault("redis.pool_size", 10)

	// Logging defaults
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.format", "json")
	viper.SetDefault("logging.output", "stdout")
	viper.SetDefault("logging.max_size", 100)
	viper.SetDefault("logging.max_backups", 3)
	viper.SetDefault("logging.max_age", 7)
	viper.SetDefault("logging.compress", true)

	// Feature flags defaults
	viper.SetDefault("features.enable_metrics", true)
	viper.SetDefault("features.enable_tracing", false)
	viper.SetDefault("features.enable_profiling", false)
	viper.SetDefault("features.enable_caching", true)
	viper.SetDefault("features.beta_features", false)

	// Security defaults
	viper.SetDefault("security.jwt_secret", "your-secret-key")
	viper.SetDefault("security.jwt_expiration", "24h")
	viper.SetDefault("security.rate_limit_rps", 100)
	viper.SetDefault("security.rate_limit_burst", 200)
	viper.SetDefault("security.cors_origins", []string{"http://localhost:3000"})
	viper.SetDefault("security.csrf_secret", "csrf-secret-key")
	viper.SetDefault("security.enable_https_only", false)
}

func runDemo() {
	fmt.Println("🚀 Viper Configuration Management Demo")
	fmt.Println("=====================================")
	fmt.Println()

	// Show configuration sources
	fmt.Println("📋 Configuration Sources:")
	fmt.Printf("   Config File: %s\n", getConfigFileInfo())
	fmt.Printf("   Environment Prefix: %s\n", envPrefix)
	fmt.Printf("   Command-line Flags: Available\n")
	fmt.Println()

	// Show basic configuration
	showBasicConfig()

	// Show environment override example
	fmt.Println("🌍 Environment Variable Override Examples:")
	fmt.Printf("   Set %s_SERVER_PORT=9090 to change server port\n", envPrefix)
	fmt.Printf("   Set %s_DATABASE_HOST=remote-db to change database host\n", envPrefix)
	fmt.Printf("   Set %s_FEATURES_BETA_FEATURES=true to enable beta features\n", envPrefix)
	fmt.Println()

	// Show available commands
	fmt.Println("🛠️  Available Commands:")
	fmt.Println("   viper-demo show              - Show full configuration")
	fmt.Println("   viper-demo validate          - Validate configuration")
	fmt.Println("   viper-demo watch             - Watch for config changes")
	fmt.Println("   viper-demo create-samples    - Create sample config files")
	fmt.Println("   viper-demo env-demo          - Environment variable demo")
	fmt.Println()

	// Show configuration precedence
	fmt.Println("🔄 Configuration Precedence (highest to lowest):")
	fmt.Println("   1. Command-line flags")
	fmt.Println("   2. Environment variables")
	fmt.Println("   3. Configuration file")
	fmt.Println("   4. Default values")
	fmt.Println()

	// Show some dynamic access examples
	showDynamicAccess()
}

func showBasicConfig() {
	fmt.Println("⚙️  Current Configuration Summary:")
	fmt.Printf("   Server: %s:%d (TLS: %v)\n", config.Server.Host, config.Server.Port, config.Server.TLS.Enabled)
	fmt.Printf("   Database: %s://%s:%d/%s\n", config.Database.Driver, config.Database.Host, config.Database.Port, config.Database.Database)
	fmt.Printf("   Redis: %s:%d (DB: %d)\n", config.Redis.Host, config.Redis.Port, config.Redis.Database)
	fmt.Printf("   Logging: %s level, %s format\n", config.Logging.Level, config.Logging.Format)
	fmt.Printf("   Features: Metrics=%v, Tracing=%v, Beta=%v\n", config.Features.EnableMetrics, config.Features.EnableTracing, config.Features.BetaFeatures)
	fmt.Println()
}

func showConfiguration() {
	fmt.Println("📊 Complete Configuration")
	fmt.Println("=========================")
	fmt.Println()

	fmt.Printf("Config File: %s\n", getConfigFileInfo())
	fmt.Printf("Environment Prefix: %s\n", envPrefix)
	fmt.Println()

	// Server Configuration
	fmt.Println("🌐 Server Configuration:")
	fmt.Printf("  Host: %s\n", config.Server.Host)
	fmt.Printf("  Port: %d\n", config.Server.Port)
	fmt.Printf("  Read Timeout: %v\n", config.Server.ReadTimeout)
	fmt.Printf("  Write Timeout: %v\n", config.Server.WriteTimeout)
	fmt.Printf("  Max Connections: %d\n", config.Server.MaxConnections)
	fmt.Printf("  TLS Enabled: %v\n", config.Server.TLS.Enabled)
	if config.Server.TLS.Enabled {
		fmt.Printf("  TLS Cert File: %s\n", config.Server.TLS.CertFile)
		fmt.Printf("  TLS Key File: %s\n", config.Server.TLS.KeyFile)
	}
	fmt.Println()

	// Database Configuration
	fmt.Println("🗄️  Database Configuration:")
	fmt.Printf("  Driver: %s\n", config.Database.Driver)
	fmt.Printf("  Host: %s\n", config.Database.Host)
	fmt.Printf("  Port: %d\n", config.Database.Port)
	fmt.Printf("  Username: %s\n", config.Database.Username)
	fmt.Printf("  Password: %s\n", maskPassword(config.Database.Password))
	fmt.Printf("  Database: %s\n", config.Database.Database)
	fmt.Printf("  SSL Mode: %s\n", config.Database.SSLMode)
	fmt.Printf("  Max Connections: %d\n", config.Database.MaxConnections)
	fmt.Printf("  Max Idle Time: %v\n", config.Database.MaxIdleTime)
	fmt.Printf("  Connection Max Lifetime: %v\n", config.Database.ConnMaxLifetime)
	fmt.Println()

	// Redis Configuration
	fmt.Println("🔴 Redis Configuration:")
	fmt.Printf("  Host: %s\n", config.Redis.Host)
	fmt.Printf("  Port: %d\n", config.Redis.Port)
	fmt.Printf("  Password: %s\n", maskPassword(config.Redis.Password))
	fmt.Printf("  Database: %d\n", config.Redis.Database)
	fmt.Printf("  Pool Size: %d\n", config.Redis.PoolSize)
	fmt.Println()

	// Logging Configuration
	fmt.Println("📝 Logging Configuration:")
	fmt.Printf("  Level: %s\n", config.Logging.Level)
	fmt.Printf("  Format: %s\n", config.Logging.Format)
	fmt.Printf("  Output: %s\n", config.Logging.Output)
	fmt.Printf("  Max Size: %d MB\n", config.Logging.MaxSize)
	fmt.Printf("  Max Backups: %d\n", config.Logging.MaxBackups)
	fmt.Printf("  Max Age: %d days\n", config.Logging.MaxAge)
	fmt.Printf("  Compress: %v\n", config.Logging.Compress)
	fmt.Println()

	// Feature Flags
	fmt.Println("🚩 Feature Flags:")
	fmt.Printf("  Enable Metrics: %v\n", config.Features.EnableMetrics)
	fmt.Printf("  Enable Tracing: %v\n", config.Features.EnableTracing)
	fmt.Printf("  Enable Profiling: %v\n", config.Features.EnableProfiling)
	fmt.Printf("  Enable Caching: %v\n", config.Features.EnableCaching)
	fmt.Printf("  Beta Features: %v\n", config.Features.BetaFeatures)
	fmt.Println()

	// Security Configuration
	fmt.Println("🔐 Security Configuration:")
	fmt.Printf("  JWT Secret: %s\n", maskPassword(config.Security.JWTSecret))
	fmt.Printf("  JWT Expiration: %v\n", config.Security.JWTExpiration)
	fmt.Printf("  Rate Limit RPS: %d\n", config.Security.RateLimitRPS)
	fmt.Printf("  Rate Limit Burst: %d\n", config.Security.RateLimitBurst)
	fmt.Printf("  CORS Origins: %v\n", config.Security.CORSOrigins)
	fmt.Printf("  CSRF Secret: %s\n", maskPassword(config.Security.CSRFSecret))
	fmt.Printf("  HTTPS Only: %v\n", config.Security.EnableHTTPSOnly)
	fmt.Println()

	// Show all viper keys
	fmt.Println("🔑 All Configuration Keys:")
	keys := viper.AllKeys()
	for _, key := range keys {
		value := viper.Get(key)
		if strings.Contains(strings.ToLower(key), "password") || strings.Contains(strings.ToLower(key), "secret") {
			value = maskPassword(fmt.Sprintf("%v", value))
		}
		fmt.Printf("  %s = %v\n", key, value)
	}
}

func validateConfiguration() {
	fmt.Println("✅ Configuration Validation")
	fmt.Println("===========================")
	fmt.Println()

	valid := true
	issues := []string{}

	// Validate server configuration
	if config.Server.Port < 1 || config.Server.Port > 65535 {
		issues = append(issues, "Server port must be between 1 and 65535")
		valid = false
	}

	if config.Server.MaxConnections < 1 {
		issues = append(issues, "Server max_connections must be positive")
		valid = false
	}

	if config.Server.TLS.Enabled {
		if config.Server.TLS.CertFile == "" || config.Server.TLS.KeyFile == "" {
			issues = append(issues, "TLS cert_file and key_file are required when TLS is enabled")
			valid = false
		}
	}

	// Validate database configuration
	if config.Database.Port < 1 || config.Database.Port > 65535 {
		issues = append(issues, "Database port must be between 1 and 65535")
		valid = false
	}

	if config.Database.MaxConnections < 1 {
		issues = append(issues, "Database max_connections must be positive")
		valid = false
	}

	// Validate Redis configuration
	if config.Redis.Port < 1 || config.Redis.Port > 65535 {
		issues = append(issues, "Redis port must be between 1 and 65535")
		valid = false
	}

	if config.Redis.Database < 0 || config.Redis.Database > 15 {
		issues = append(issues, "Redis database must be between 0 and 15")
		valid = false
	}

	// Validate logging configuration
	validLogLevels := []string{"debug", "info", "warn", "error", "fatal"}
	validLevel := false
	for _, level := range validLogLevels {
		if strings.ToLower(config.Logging.Level) == level {
			validLevel = true
			break
		}
	}
	if !validLevel {
		issues = append(issues, "Logging level must be one of: debug, info, warn, error, fatal")
		valid = false
	}

	// Validate security configuration
	if config.Security.RateLimitRPS < 1 {
		issues = append(issues, "Security rate_limit_rps must be positive")
		valid = false
	}

	if len(config.Security.JWTSecret) < 32 {
		issues = append(issues, "Security jwt_secret should be at least 32 characters for security")
		valid = false
	}

	// Display results
	if valid {
		fmt.Println("✅ Configuration is valid!")
		fmt.Println()
		fmt.Println("📋 Configuration Summary:")
		fmt.Printf("  Server will run on: %s:%d\n", config.Server.Host, config.Server.Port)
		fmt.Printf("  Database connection: %s://%s:%d\n", config.Database.Driver, config.Database.Host, config.Database.Port)
		fmt.Printf("  Logging level: %s\n", config.Logging.Level)
		fmt.Printf("  Security features: JWT=%v, HTTPS-Only=%v\n", len(config.Security.JWTSecret) > 0, config.Security.EnableHTTPSOnly)
	} else {
		fmt.Println("❌ Configuration validation failed!")
		fmt.Println()
		fmt.Println("Issues found:")
		for i, issue := range issues {
			fmt.Printf("  %d. %s\n", i+1, issue)
		}
	}
}

func watchConfiguration() {
	fmt.Println("👀 Watching Configuration Changes")
	fmt.Println("=================================")
	fmt.Println()

	if viper.ConfigFileUsed() == "" {
		fmt.Println("❌ No configuration file is being used. Cannot watch for changes.")
		fmt.Println("💡 Create a config file or specify one with --config flag")
		return
	}

	fmt.Printf("📁 Watching file: %s\n", viper.ConfigFileUsed())
	fmt.Println("🔄 Make changes to the config file to see live updates...")
	fmt.Println("Press Ctrl+C to stop watching")
	fmt.Println()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("🔔 Config file changed: %s\n", e.Name)

		// Reload configuration
		oldConfig := config
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Printf("❌ Error reloading config: %v\n", err)
			return
		}

		// Show what changed
		fmt.Println("🔄 Configuration updated:")
		compareConfigs(oldConfig, config)
		fmt.Println("---")
	})

	// Keep the program running
	select {}
}

func createSampleConfigs() {
	fmt.Println("📄 Creating Sample Configuration Files")
	fmt.Println("=====================================")
	fmt.Println()

	// Create YAML config
	yamlConfig := `# Viper Demo Configuration - YAML Format
server:
  host: "localhost"
  port: 8080
  read_timeout: "30s"
  write_timeout: "30s"
  max_connections: 1000
  tls:
    enabled: false
    cert_file: "/path/to/cert.pem"
    key_file: "/path/to/key.pem"

database:
  driver: "postgres"
  host: "localhost"
  port: 5432
  username: "myapp_user"
  password: "secure_password"
  database: "myapp_db"
  ssl_mode: "require"
  max_connections: 25
  max_idle_time: "15m"
  conn_max_lifetime: "1h"

redis:
  host: "localhost"
  port: 6379
  password: ""
  database: 0
  pool_size: 10

logging:
  level: "info"
  format: "json"
  output: "stdout"
  max_size: 100
  max_backups: 3
  max_age: 7
  compress: true

features:
  enable_metrics: true
  enable_tracing: false
  enable_profiling: false
  enable_caching: true
  beta_features: false

security:
  jwt_secret: "your-super-secret-jwt-key-here-make-it-long"
  jwt_expiration: "24h"
  rate_limit_rps: 100
  rate_limit_burst: 200
  cors_origins:
    - "http://localhost:3000"
    - "https://myapp.com"
  csrf_secret: "csrf-secret-key-here"
  enable_https_only: false
`

	// Create JSON config
	jsonConfig := `{
  "server": {
    "host": "localhost",
    "port": 8080,
    "read_timeout": "30s",
    "write_timeout": "30s",
    "max_connections": 1000,
    "tls": {
      "enabled": false,
      "cert_file": "/path/to/cert.pem",
      "key_file": "/path/to/key.pem"
    }
  },
  "database": {
    "driver": "postgres",
    "host": "localhost",
    "port": 5432,
    "username": "myapp_user",
    "password": "secure_password",
    "database": "myapp_db",
    "ssl_mode": "require",
    "max_connections": 25,
    "max_idle_time": "15m",
    "conn_max_lifetime": "1h"
  },
  "redis": {
    "host": "localhost",
    "port": 6379,
    "password": "",
    "database": 0,
    "pool_size": 10
  },
  "logging": {
    "level": "info",
    "format": "json",
    "output": "stdout",
    "max_size": 100,
    "max_backups": 3,
    "max_age": 7,
    "compress": true
  },
  "features": {
    "enable_metrics": true,
    "enable_tracing": false,
    "enable_profiling": false,
    "enable_caching": true,
    "beta_features": false
  },
  "security": {
    "jwt_secret": "your-super-secret-jwt-key-here-make-it-long",
    "jwt_expiration": "24h",
    "rate_limit_rps": 100,
    "rate_limit_burst": 200,
    "cors_origins": [
      "http://localhost:3000",
      "https://myapp.com"
    ],
    "csrf_secret": "csrf-secret-key-here",
    "enable_https_only": false
  }
}
`

	// Create TOML config
	tomlConfig := `# Viper Demo Configuration - TOML Format

[server]
host = "localhost"
port = 8080
read_timeout = "30s"
write_timeout = "30s"
max_connections = 1000

[server.tls]
enabled = false
cert_file = "/path/to/cert.pem"
key_file = "/path/to/key.pem"

[database]
driver = "postgres"
host = "localhost"
port = 5432
username = "myapp_user"
password = "secure_password"
database = "myapp_db"
ssl_mode = "require"
max_connections = 25
max_idle_time = "15m"
conn_max_lifetime = "1h"

[redis]
host = "localhost"
port = 6379
password = ""
database = 0
pool_size = 10

[logging]
level = "info"
format = "json"
output = "stdout"
max_size = 100
max_backups = 3
max_age = 7
compress = true

[features]
enable_metrics = true
enable_tracing = false
enable_profiling = false
enable_caching = true
beta_features = false

[security]
jwt_secret = "your-super-secret-jwt-key-here-make-it-long"
jwt_expiration = "24h"
rate_limit_rps = 100
rate_limit_burst = 200
cors_origins = [
  "http://localhost:3000",
  "https://myapp.com"
]
csrf_secret = "csrf-secret-key-here"
enable_https_only = false
`

	// Write files
	configs := map[string]string{
		"config.yaml": yamlConfig,
		"config.json": jsonConfig,
		"config.toml": tomlConfig,
	}

	for filename, content := range configs {
		if err := os.WriteFile(filename, []byte(content), 0644); err != nil {
			fmt.Printf("❌ Error creating %s: %v\n", filename, err)
		} else {
			fmt.Printf("✅ Created %s\n", filename)
		}
	}

	fmt.Println()
	fmt.Println("📝 Sample configuration files created!")
	fmt.Println("Try running the demo with different config files:")
	fmt.Println("  viper-demo --config config.yaml")
	fmt.Println("  viper-demo --config config.json")
	fmt.Println("  viper-demo --config config.toml")
}

func environmentDemo() {
	fmt.Println("🌍 Environment Variable Demonstration")
	fmt.Println("====================================")
	fmt.Println()

	envPrefix := viper.GetString("env-prefix")
	fmt.Printf("Environment variable prefix: %s_\n", envPrefix)
	fmt.Println()

	// Show current values and their environment variable names
	fmt.Println("🔧 Configuration Values and Environment Overrides:")

	configMappings := []struct {
		key     string
		envVar  string
		value   interface{}
		example string
	}{
		{"server.host", "SERVER_HOST", viper.Get("server.host"), "localhost"},
		{"server.port", "SERVER_PORT", viper.Get("server.port"), "9090"},
		{"server.tls.enabled", "SERVER_TLS_ENABLED", viper.Get("server.tls.enabled"), "true"},
		{"database.host", "DATABASE_HOST", viper.Get("database.host"), "db-server.com"},
		{"database.port", "DATABASE_PORT", viper.Get("database.port"), "5433"},
		{"database.password", "DATABASE_PASSWORD", maskPassword(fmt.Sprintf("%v", viper.Get("database.password"))), "new_secure_password"},
		{"redis.host", "REDIS_HOST", viper.Get("redis.host"), "redis-server.com"},
		{"redis.database", "REDIS_DATABASE", viper.Get("redis.database"), "1"},
		{"logging.level", "LOGGING_LEVEL", viper.Get("logging.level"), "debug"},
		{"features.beta_features", "FEATURES_BETA_FEATURES", viper.Get("features.beta_features"), "true"},
		{"security.jwt_expiration", "SECURITY_JWT_EXPIRATION", viper.Get("security.jwt_expiration"), "48h"},
	}

	for _, mapping := range configMappings {
		fmt.Printf("  %s\n", mapping.key)
		fmt.Printf("    Current value: %v\n", mapping.value)
		fmt.Printf("    Environment variable: %s_%s\n", envPrefix, mapping.envVar)
		fmt.Printf("    Example: export %s_%s=%s\n", envPrefix, mapping.envVar, mapping.example)
		fmt.Println()
	}

	// Show environment variables that are currently set
	fmt.Println("🌱 Currently Set Environment Variables:")
	foundAny := false
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, envPrefix+"_") {
			fmt.Printf("  %s\n", env)
			foundAny = true
		}
	}
	if !foundAny {
		fmt.Printf("  No %s_* environment variables are currently set\n", envPrefix)
	}
	fmt.Println()

	// Provide instructions
	fmt.Println("💡 Try setting environment variables and running commands:")
	fmt.Printf("  export %s_SERVER_PORT=9090\n", envPrefix)
	fmt.Printf("  export %s_DATABASE_HOST=remote-db\n", envPrefix)
	fmt.Printf("  export %s_LOGGING_LEVEL=debug\n", envPrefix)
	fmt.Println("  viper-demo show")
}

func showDynamicAccess() {
	fmt.Println("🔍 Dynamic Configuration Access Examples:")
	fmt.Println("   Get String: viper.GetString(\"server.host\") = ", viper.GetString("server.host"))
	fmt.Println("   Get Int: viper.GetInt(\"server.port\") = ", viper.GetInt("server.port"))
	fmt.Println("   Get Bool: viper.GetBool(\"server.tls.enabled\") = ", viper.GetBool("server.tls.enabled"))
	fmt.Println("   Get Duration: viper.GetDuration(\"server.read_timeout\") = ", viper.GetDuration("server.read_timeout"))
	fmt.Println("   Get StringSlice: viper.GetStringSlice(\"security.cors_origins\") = ", viper.GetStringSlice("security.cors_origins"))
	fmt.Println()
}

// Utility functions
func getConfigFileInfo() string {
	if configFile := viper.ConfigFileUsed(); configFile != "" {
		abs, _ := filepath.Abs(configFile)
		return fmt.Sprintf("%s (%s)", configFile, abs)
	}
	return "Not using config file"
}

func maskPassword(password string) string {
	if password == "" {
		return "(empty)"
	}
	if len(password) <= 4 {
		return strings.Repeat("*", len(password))
	}
	return password[:2] + strings.Repeat("*", len(password)-4) + password[len(password)-2:]
}

func compareConfigs(old, new Config) {
	// Simple comparison - in a real app you might use reflection or a proper diff library
	if old.Server.Port != new.Server.Port {
		fmt.Printf("  Server Port: %d → %d\n", old.Server.Port, new.Server.Port)
	}
	if old.Server.Host != new.Server.Host {
		fmt.Printf("  Server Host: %s → %s\n", old.Server.Host, new.Server.Host)
	}
	if old.Database.Host != new.Database.Host {
		fmt.Printf("  Database Host: %s → %s\n", old.Database.Host, new.Database.Host)
	}
	if old.Logging.Level != new.Logging.Level {
		fmt.Printf("  Logging Level: %s → %s\n", old.Logging.Level, new.Logging.Level)
	}
	if old.Features.BetaFeatures != new.Features.BetaFeatures {
		fmt.Printf("  Beta Features: %v → %v\n", old.Features.BetaFeatures, new.Features.BetaFeatures)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
