# Viper Configuration Management Demo

A comprehensive demonstration of [Viper](https://github.com/spf13/viper) - the complete configuration solution for Go applications.

## Overview

Viper is a powerful configuration management library that supports:
- **Multiple Config Formats**: JSON, TOML, YAML, HCL, envfile, Java properties
- **Environment Variables**: Automatic prefix support with nested key mapping
- **Command-line Flags**: Integration with pflag/Cobra
- **Remote Config**: etcd, Consul support
- **Live Watching**: Automatic config reload on file changes
- **Configuration Precedence**: Clear override hierarchy

## Features Demonstrated

### 🔧 Core Configuration Management
- [x] Reading from multiple config file formats (YAML, JSON, TOML)
- [x] Environment variable integration with prefixes
- [x] Command-line flag support via Cobra
- [x] Configuration validation and error handling
- [x] Default values and fallback handling
- [x] Live configuration watching (file changes)

### 🌍 Environment Integration
- [x] Automatic environment variable mapping
- [x] Custom prefix support (`VIPERAPP_`)
- [x] Nested key mapping (e.g., `VIPERAPP_SERVER_PORT`)
- [x] Override demonstration with precedence

### 🏗️ Advanced Features
- [x] Complex nested configuration structures
- [x] Type-safe configuration access
- [x] Multiple configuration sources
- [x] Configuration marshaling/unmarshaling
- [x] Real-time configuration updates

## Project Structure

```
Viper/
├── main.go                 # Comprehensive CLI application
├── go.mod                 # Module dependencies
├── go.sum                 # Dependency checksums
├── README.md              # This documentation
├── config.yaml            # Primary config file
├── config.json            # JSON format example
└── config.toml            # TOML format example
```

## Configuration Structure

The demo uses a realistic application configuration with:

```yaml
server:
  host: "localhost"
  port: 8080
  tls:
    enabled: false
    cert_file: ""
    key_file: ""
  timeouts:
    read_timeout: "30s"
    write_timeout: "30s"
    idle_timeout: "60s"

database:
  host: "localhost"
  port: 5432
  name: "myapp_db"
  user: "postgres"
  password: "secret123"  # Masked in display
  ssl_mode: "require"
  max_connections: 25

redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0
  pool_size: 10

logging:
  level: "info"
  format: "json"
  output: "stdout"

features:
  metrics_enabled: true
  tracing_enabled: false
  beta_features: false

security:
  jwt_secret: "super-secret-key"  # Masked in display
  cors_origins:
    - "http://localhost:3000"
    - "https://myapp.com"
  rate_limit: 100
```

## Installation & Setup

1. **Clone and navigate:**
   ```bash
   cd Viper
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Create sample configuration files:**
   ```bash
   go run main.go create-samples
   ```

## Usage Examples

### Basic Configuration Display

```bash
# Show complete configuration
go run main.go --config config.yaml show

# Show configuration from different formats
go run main.go --config config.json show
go run main.go --config config.toml show
```

### Environment Variable Overrides

```bash
# Override server port via environment variable
$env:VIPERAPP_SERVER_PORT="9090"
go run main.go --config config.yaml

# Override database host
$env:VIPERAPP_DATABASE_HOST="production-db.company.com"
go run main.go --config config.yaml show
```

### Configuration Validation

```bash
# Validate current configuration
go run main.go --config config.yaml validate

# Validate with overrides
$env:VIPERAPP_SERVER_PORT="9090"
go run main.go --config config.yaml validate
```

### Live Configuration Watching

```bash
# Watch for config file changes (runs in background)
go run main.go --config config.yaml watch

# In another terminal, modify config.yaml to see live updates
```

### Environment Variable Demo

```bash
# See all available environment variable mappings
go run main.go env-demo
```

## Key Code Patterns

### 1. Viper Initialization
```go
viper.SetConfigName("config")
viper.SetConfigType("yaml")
viper.AddConfigPath(".")
viper.SetEnvPrefix("VIPERAPP")
viper.AutomaticEnv()
viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
```

### 2. Configuration Structure Binding
```go
type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    // ... other fields
}

var config Config
viper.Unmarshal(&config)
```

### 3. Type-Safe Access
```go
// Direct access with type conversion
host := viper.GetString("server.host")
port := viper.GetInt("server.port")
timeout := viper.GetDuration("server.timeouts.read_timeout")
origins := viper.GetStringSlice("security.cors_origins")
```

### 4. Environment Variable Mapping
```go
// VIPERAPP_SERVER_PORT maps to server.port
// VIPERAPP_DATABASE_HOST maps to database.host
viper.SetEnvPrefix("VIPERAPP")
viper.AutomaticEnv()
viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
```

### 5. Live Configuration Watching
```go
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
    fmt.Printf("Config file changed: %s\n", e.Name)
    // Handle configuration reload
})
```

## Configuration Precedence

Viper follows this precedence order (highest to lowest):

1. **Explicit calls** (viper.Set())
2. **Command-line flags**
3. **Environment variables**
4. **Configuration file**
5. **Key/Value store**
6. **Default values**

## Environment Variable Examples

All configuration keys can be overridden via environment variables with the `VIPERAPP_` prefix:

```bash
# Server configuration
VIPERAPP_SERVER_HOST=0.0.0.0
VIPERAPP_SERVER_PORT=9090
VIPERAPP_SERVER_TLS_ENABLED=true

# Database configuration
VIPERAPP_DATABASE_HOST=prod-db.company.com
VIPERAPP_DATABASE_PORT=5432
VIPERAPP_DATABASE_NAME=production_db

# Feature flags
VIPERAPP_FEATURES_METRICS_ENABLED=true
VIPERAPP_FEATURES_BETA_FEATURES=true
```

## Advanced Features

### Configuration Validation
The demo includes comprehensive validation:
- Required field checking
- Value range validation
- Connection string validation
- Security configuration validation

### Password Masking
Sensitive values are automatically masked in output:
- Database passwords
- JWT secrets
- API keys

### Multiple Format Support
Create configuration files in any supported format:
```bash
# Creates config.yaml, config.json, and config.toml
go run main.go create-samples
```

## Integration Patterns

### With Cobra CLI
```go
var cfgFile string
rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
```

### With Dependency Injection
```go
func NewDatabaseConnection(config *Config) *sql.DB {
    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        config.Database.Host,
        config.Database.Port,
        config.Database.User,
        config.Database.Password,
        config.Database.Name,
        config.Database.SSLMode)
    // ... connection logic
}
```

## Testing the Demo

1. **Basic functionality:**
   ```bash
   go run main.go --help
   ```

2. **Configuration display:**
   ```bash
   go run main.go --config config.yaml show
   ```

3. **Environment overrides:**
   ```bash
   $env:VIPERAPP_SERVER_PORT="9090"
   go run main.go --config config.yaml
   ```

4. **Live watching** (requires two terminals):
   ```bash
   # Terminal 1: Start watching
   go run main.go --config config.yaml watch
   
   # Terminal 2: Modify config.yaml
   # See live updates in Terminal 1
   ```

## Key Benefits Demonstrated

### 🚀 **Developer Productivity**
- **Zero Configuration**: Works out of the box with sensible defaults
- **Multiple Formats**: Support for all common config formats
- **Environment Integration**: Seamless 12-factor app compliance

### 🔧 **Production Ready**
- **Live Reloading**: No application restarts needed for config changes
- **Override Hierarchy**: Clear precedence for different config sources
- **Validation**: Built-in configuration validation support

### 🏗️ **Enterprise Features**
- **Remote Configuration**: Support for etcd/Consul (not shown in demo)
- **Security**: Password masking and sensitive data handling
- **Monitoring**: Configuration change notifications and logging

## Common Use Cases

1. **Microservice Configuration**: Environment-specific settings
2. **Feature Flags**: Runtime feature toggling
3. **Database Connections**: Multi-environment database configs
4. **API Configuration**: Timeouts, rate limits, CORS settings
5. **Security Settings**: JWT secrets, encryption keys
6. **Logging Configuration**: Levels, formats, outputs

## Best Practices Demonstrated

1. **Structured Configuration**: Use nested structs for organization
2. **Environment Prefixes**: Avoid naming conflicts
3. **Default Values**: Provide sensible defaults
4. **Validation**: Validate configuration at startup
5. **Sensitive Data**: Mask passwords and secrets in logs
6. **Documentation**: Clear mapping of environment variables

## Resources

- **Viper GitHub**: https://github.com/spf13/viper
- **Official Documentation**: https://pkg.go.dev/github.com/spf13/viper
- **Configuration Best Practices**: https://12factor.net/config

---

This demo showcases Viper as a production-ready configuration management solution that scales from simple applications to complex enterprise systems.
