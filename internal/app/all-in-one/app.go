package all_in_one

import (
	"context"
	"github.com/leinodev/go-microservice-template/config"
	"github.com/leinodev/go-microservice-template/internal/core/example"
	"github.com/leinodev/go-microservice-template/internal/infrastructure/controllers/http"
	in_memory "github.com/leinodev/go-microservice-template/internal/infrastructure/repositories/in-memory"
	"github.com/leinodev/go-microservice-template/internal/server"
	"os/signal"
	"syscall"
)

// App - структура приложения all-in-one, здесь происходит конфигурация и запуск основных компонентов.
type App struct {
	config   *config.Config
	stopFunc context.CancelFunc
}

func NewApp(
	config *config.Config,
) *App {
	// create instance of application
	return &App{
		config: config,
	}
}

func (app *App) Configure(ctx context.Context) error {
	// initialize all components of application:
	// - controllers
	// - repositories
	// - cores
	// - consumer etc.

	return nil
}

func (app *App) Run(ctx context.Context) error {
	ctx, app.stopFunc = signal.NotifyContext(ctx, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM) // signals to graceful shutdown
	defer app.stopFunc()

	s := server.BuildServer(app.config.Service)

	repo := in_memory.NewExampleRepository()
	core := example.NewGetExampleByIDCore(repo)
	http.Register(s, core)

	errors := make(chan error, 1)
	go func() {
		errors <- s.Listen(app.config.Service.Address)
	}()

	shutdown := func(ctx context.Context) error {
		if err := s.ShutdownWithTimeout(app.config.Service.ShutdownTimeout); err != nil {
			return err
		}
		return nil
	}

	select {
	case err := <-errors:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		defer app.stopFunc()
		err := shutdown(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *App) Stop() {
	app.stopFunc()
}
