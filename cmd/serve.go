package cmd

import (
	"context"

	"entgo.io/ent/dialect"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/versionx"
	"go.uber.org/zap"

	"go.infratographer.com/example-api/internal/api"
	"go.infratographer.com/example-api/internal/config"
	ent "go.infratographer.com/example-api/internal/ent/generated"
)

const (
	defaultListenAddr = ":17608"
)

var (
	enablePlayground bool
	serveDevMode     bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the todo example Graph API",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	echox.MustViperFlags(viper.GetViper(), serveCmd.Flags(), defaultListenAddr)
	echojwtx.MustViperFlags(viper.GetViper(), serveCmd.Flags())

	// only available as a CLI arg because it shouldn't be something that could accidentially end up in a config file or env var
	serveCmd.Flags().BoolVar(&serveDevMode, "dev", false, "dev mode: enables playground, disables all auth checks, sets CORS to allow all, pretty logging, etc.")
	serveCmd.Flags().BoolVar(&enablePlayground, "playground", false, "enable the graph playground")
}

func serve(ctx context.Context) error {
	if serveDevMode {
		enablePlayground = true
		config.AppConfig.Logging.Debug = true
		config.AppConfig.Logging.Pretty = true
		config.AppConfig.Server.WithMiddleware(middleware.CORS())
	}

	cOpts := []ent.Option{}

	if config.AppConfig.Logging.Debug {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1", cOpts...)
	if err != nil {
		logger.Error("failed opening connection to sqlite", zap.Error(err))
		return err
	}
	defer client.Close()

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		logger.Errorf("failed creating schema resources", zap.Error(err))
		return err
	}

	// jwt auth middleware
	if !serveDevMode {
		if authconfig, err := echojwtx.AuthConfigFromViper(viper.GetViper()); err != nil {
			logger.Fatal("failed to initialize jwt authentication", zap.Error(err))
		} else if authconfig != nil {
			config.AppConfig.AuthConfig = *authconfig
			config.AppConfig.AuthConfig.JWTConfig.Skipper = echox.SkipDefaultEndpoints

			auth, err := echojwtx.NewAuth(ctx, config.AppConfig.AuthConfig)
			if err != nil {
				logger.Fatal("failed to initialize jwt authentication", zap.Error(err))
			}

			config.AppConfig.Server = config.AppConfig.Server.WithMiddleware(auth.Middleware())
		}
	}

	srv, err := echox.NewServer(logger.Desugar(), config.AppConfig.Server, versionx.BuildDetails())
	if err != nil {
		logger.Error("failed to create server", zap.Error(err))
	}

	r := api.NewResolver(client, logger.Named("resolvers"))
	handler := r.Handler(enablePlayground)

	srv.AddHandler(handler)

	if err := srv.RunWithContext(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return err
}
