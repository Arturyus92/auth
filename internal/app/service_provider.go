package app

import (
	"context"
	"log"

	"github.com/Arturyus92/auth/internal/api/access"
	"github.com/Arturyus92/auth/internal/api/auth"
	"github.com/Arturyus92/auth/internal/api/user"
	"github.com/Arturyus92/auth/internal/config"
	"github.com/Arturyus92/auth/internal/config/env"
	"github.com/Arturyus92/auth/internal/repository"
	logRepository "github.com/Arturyus92/auth/internal/repository/log"
	permRepository "github.com/Arturyus92/auth/internal/repository/permission"
	secretRepository "github.com/Arturyus92/auth/internal/repository/secret"
	userRepository "github.com/Arturyus92/auth/internal/repository/user"
	"github.com/Arturyus92/auth/internal/service"
	accessService "github.com/Arturyus92/auth/internal/service/access"
	authService "github.com/Arturyus92/auth/internal/service/auth"
	userService "github.com/Arturyus92/auth/internal/service/user"
	"github.com/Arturyus92/platform_common/pkg/closer"
	"github.com/Arturyus92/platform_common/pkg/db"
	"github.com/Arturyus92/platform_common/pkg/db/pg"
	"github.com/Arturyus92/platform_common/pkg/db/transaction"
)

type serviceProvider struct {
	pgConfig         config.PGConfig
	grpcConfig       config.GRPCConfig
	httpConfig       config.HTTPConfig
	swaggerConfig    config.SwaggerConfig
	loggerConfig     config.LoggerConfig
	prometheusConfig config.PrometheusConfig

	dbClient         db.Client
	txManager        db.TxManager
	userRepository   repository.UserRepository
	logRepository    repository.LogRepository
	permRepository   repository.PermRepository
	secretRepository repository.SecretRepository

	userService   service.UserService
	authService   service.AuthService
	accessService service.AccessService

	userImpl   *user.Implementation
	authImpl   *auth.Implementation
	accessImpl *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// PrometheusConfig - ...
func (s *serviceProvider) PrometheusConfig() config.PrometheusConfig {
	if s.prometheusConfig == nil {
		cfg, err := env.NewPrometheusConfig()
		if err != nil {
			log.Fatalf("failed to get prometheus config: %s", err.Error())
		}

		s.prometheusConfig = cfg
	}

	return s.prometheusConfig
}

// LoggerConfig - ...
func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := env.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

// PGConfig - ...
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// GRPCConfig - ...
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// HTTPConfig - ...
func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

// SwaggerConfig - ...
func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

// DBClient - ...
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// TxManager - ...
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// UserRepository - ...
func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

// LogRepository - ...
func (s *serviceProvider) LogRepository(ctx context.Context) repository.LogRepository {
	if s.logRepository == nil {
		s.logRepository = logRepository.NewRepository(s.DBClient(ctx))
	}

	return s.logRepository
}

// PermRepository - ...
func (s *serviceProvider) PermRepository(ctx context.Context) repository.PermRepository {
	if s.permRepository == nil {
		s.permRepository = permRepository.NewRepository(s.DBClient(ctx))
	}

	return s.permRepository
}

// SecretRepository - ...
func (s *serviceProvider) SecretRepository(ctx context.Context) repository.SecretRepository {
	if s.secretRepository == nil {
		s.secretRepository = secretRepository.NewRepository(s.DBClient(ctx))
	}

	return s.secretRepository
}

// UserService - ...
func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
			s.LogRepository(ctx),
		)
	}

	return s.userService
}

// AccessService - ...
func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.PermRepository(ctx),
			s.SecretRepository(ctx),
		)
	}

	return s.accessService
}

// AuthService - ...
func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(s.UserRepository(ctx),
			s.SecretRepository(ctx),
		)
	}

	return s.authService
}

// UserImpl - ...
func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

// AuthImpl - ...
func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

// AccessImpl - ...
func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}
