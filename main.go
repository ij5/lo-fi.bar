package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/cli"
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
	"github.com/maxence-charriere/go-app/v9/pkg/logs"
)

const (
	backgroundColor = "#000000"
)

type options struct {
	Port int `env:"PORT" help:"The port used to listen connections."`
}

type githubOptions struct {
	Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
	for _, l := range getLiveRadios() {
		app.Route("/"+l.Slug, newRadio())
	}
	app.Route("/", newRadio())
	app.RunWhenOnBrowser()

	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	h := app.Handler{
		Author:          "jh",
		BackgroundColor: backgroundColor,
		Description:     "Lofi music player",
		Icon: app.Icon{
			Default: "https://filedn.com/ljcU5RLfox2Sbr04VfCcCX5/lo-fi.png",
		},
		Keywords: []string{
			"lofi",
			"lo-fi",
			"music",
			"lofimusic",
			"chill",
			"chilled",
			"beats",
			"relax",
			"study",
			"sleep",
			"hiphop",
			"app",
			"pwa",
		},
		LoadingLabel: "Lofi music player to work, study and relax.",
		Name:         "Lo-Fi.bar",
		Image:        "https://filedn.com/ljcU5RLfox2Sbr04VfCcCX5/lo-fi.png",
		RawHeaders: []string{
			`<script>
			var isOnYouTubeIframeAPIReady = false;
			function onYouTubeIframeAPIReady() {
				isOnYouTubeIframeAPIReady = true;
			}
			</script>`,
		},
		Styles: []string{
			"https://fonts.googleapis.com/css2?family=Montserrat:wght@400;500&display=swap",
			"/web/app.css",
		},

		ThemeColor: backgroundColor,
		Title:      "Lo-Fi.bar",
	}

	opts := options{Port: 4000}
	cli.Register("local").
		Help(`Launches a server that serves the documentation app in a local environment.`).
		Options(&opts)

	githubOpts := githubOptions{}
	cli.Register("github").
		Help(`Generates the required resources to run Lofimusic app on GitHub Pages.`).
		Options(&githubOpts)
	cli.Load()

	switch cli.Load() {
	case "local":
		runLocal(ctx, &h, opts)

	case "github":
		generateGitHubPages(ctx, &h, githubOpts)
	}

}

func runLocal(ctx context.Context, h http.Handler, opts options) {
	app.Logf("%s", logs.New("starting lofimusic app server").
		Tag("port", opts.Port),
	)

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", opts.Port),
		Handler: h,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func generateGitHubPages(ctx context.Context, h *app.Handler, opts githubOptions) {
	radios := getLiveRadios()
	slugs := make([]string, len(radios))
	for i, r := range radios {
		slugs[i] = r.Slug
	}

	if err := app.GenerateStaticWebsite(opts.Output, h, slugs...); err != nil {
		panic(err)
	}
}

func exit() {
	err := recover()
	if err != nil {
		app.Logf("command failed: %s", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
