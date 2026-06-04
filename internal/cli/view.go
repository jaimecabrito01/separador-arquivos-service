package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/jaimecabrito01/separador-arquivos-service/entities"
)

func getUserInput(prompt string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s (Padrão: %s): ", prompt, defaultValue)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("\nErro de leitura: %v. Usando o valor padrão.\n", err)
		return defaultValue
	}

	input = strings.TrimSpace(input)

	if input == "" {
		fmt.Printf("=> Usando o valor padrão: %s\n", defaultValue)
		return defaultValue
	}

	absPath, err := filepath.Abs(input)
	if err != nil {
		fmt.Println("Aviso: Não foi possível resolver o caminho absoluto, usando a entrada como está.")
		return input
	}
	return absPath
}

func Init() {
	fmt.Println("--- Configuração de Rotas ---")

	defaultDownloadsPath := xdg.UserDirs.Download
	defaultDocs := xdg.UserDirs.Documents
	defaultVideos := xdg.UserDirs.Videos
	defaultImages := xdg.UserDirs.Pictures
	defaultMusics := xdg.UserDirs.Music
	paths := make(map[string]string)

	fmt.Println("De onde voce quer organizar os arquivos?Digite o a rota:\n")
	fmt.Printf("Se você não digitar nada, será usado: [%s] (Padrão)\n", defaultDownloadsPath)
	paths["Downloads"] = getUserInput("Downloads: ", defaultDownloadsPath)

	fmt.Println("\n--- Rotas de Destino (Onde os arquivos serão classificados) ---")

	paths["Documents"] = getUserInput("Documentos:", defaultDocs)
	paths["Videos"] = getUserInput("Videos: ", defaultVideos)
	paths["Musics"] = getUserInput("Musicas: ", defaultMusics)
	paths["Images"] = getUserInput("Imagens: ", defaultImages)

	fmt.Println("\n=================================")
	fmt.Println("Configurações de Rota Aplicadas:")
	fmt.Println("=================================")
	for category, path := range paths {
		fmt.Printf("%-30s: %s\n", category, path)
	}
	fmt.Println("=================================")

	pat := entities.Paths{
		Downloads: paths["Downloads"],
		Images:    paths["Images"],
		Videos:    paths["Videos"],
		Musics:    paths["Musics"],
		Documents: paths["Documents"],
	}
	entities.NewConfig(&pat)
}
