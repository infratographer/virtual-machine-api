package cmd

import (
	"context"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/crdbx"
	"go.infratographer.com/x/echojwtx"
	"go.infratographer.com/x/echox"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/otelx"
	"go.infratographer.com/x/versionx"
	"go.uber.org/zap"

	"go.infratographer.com/virtual-machine-api/internal/api"
	"go.infratographer.com/virtual-machine-api/internal/config"
	ent "go.infratographer.com/virtual-machine-api/internal/ent/generated"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/eventhooks"
)

const (
	defaultListenAddr = ":7911"
)

var (
	enablePlayground bool
	serveDevMode     bool
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the virtual machine API",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	echox.MustViperFlags(viper.GetViper(), serveCmd.Flags(), defaultListenAddr)
	echojwtx.MustViperFlags(viper.GetViper(), serveCmd.Flags())
	events.MustViperFlags(viper.GetViper(), serveCmd.Flags(), appName)
	permissions.MustViperFlags(viper.GetViper(), serveCmd.Flags())

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

	events, err := events.NewConnection(config.AppConfig.Events, events.WithLogger(logger))
	if err != nil {
		logger.Fatalw("failed to initialize events", "error", err)
	}

	err = otelx.InitTracer(config.AppConfig.Tracing, appName, logger)
	if err != nil {
		logger.Fatalw("failed to initialize tracer", "error", err)
	}

	db, err := crdbx.NewDB(config.AppConfig.CRDB, config.AppConfig.Tracing.Enabled)
	if err != nil {
		logger.Fatalw("failed to connect to database", "error", err)
	}

	defer db.Close()

	entDB := entsql.OpenDB(dialect.Postgres, db)

	cOpts := []ent.Option{ent.Driver(entDB), ent.EventsPublisher(events)}

	if config.AppConfig.Logging.Debug {
		cOpts = append(cOpts,
			ent.Log(logger.Named("ent").Debugln),
			ent.Debug(),
		)
	}

	client := ent.NewClient(cOpts...)
	defer client.Close()

	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		logger.Errorf("failed creating schema resources", zap.Error(err))
		return err
	}

	eventhooks.EventHooks(client)

	// var middleware []echo.MiddlewareFunc

	// jwt auth middleware
	if viper.GetBool("oidc.enabled") {
		// auth, err := echojwtx.NewAuth(ctx, config.AppConfig.OIDC)
		if err != nil {
			logger.Fatalw("failed to initialize jwt authentication", zap.Error(err))
		}
	}
	// middleware = append(middleware, auth.Middleware())

	srv, err := echox.NewServer(
		logger.Desugar(),
		config.AppConfig.Server,
		versionx.BuildDetails(),
		echox.WithLoggingSkipper(echox.SkipDefaultEndpoints),
	)
	if err != nil {
		logger.Error("failed to create server", zap.Error(err))
	}

	/*perms, err := permissions.New(config.AppConfig.Permissions,
		permissions.WithLogger(logger),
		permissions.WithDefaultChecker(permissions.DefaultAllowChecker),
	)*/
	// middleware = append(middleware, perms.Middleware())
	if err != nil {
		logger.Fatal("failed to initialize permissions", zap.Error(err))
	}

	r := api.NewResolver(client, logger.Named("resolvers"))
	handler := r.Handler(enablePlayground)

	srv.AddHandler(handler)

	if err := srv.RunWithContext(ctx); err != nil {
		logger.Error("failed to run server", zap.Error(err))
	}

	return err
}
