package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jaimecabrito01/separador-arquivos-service/internal/cli"
	"github.com/jaimecabrito01/separador-arquivos-service/internal/organizer"
	"github.com/jaimecabrito01/separador-arquivos-service/internal/watcher"
)

func main() {
	setup := flag.Bool("setup", false, "configure paths")
	now := flag.Bool("now", false, "organize existing files now")
	flag.Parse()

	if *setup {
		cli.Init()
		return
	}

	if *now {
		downloads, musics, videos, documents, images, err := organizer.LoadConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao carregar config: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(">>> Organizando arquivos existentes...")
		organizer.OrganizeExisting(downloads, musics, videos, documents, images)
		fmt.Println(">>> Organização concluída!")
		return
	}

	watcher.Watcher()
}
