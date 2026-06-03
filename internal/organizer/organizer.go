package organizer

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/jaimecabrito01/separador-arquivos-service/entities"
)

func LoadConfig() (string, string, string, string, string, error) {
	configHome, _ := os.UserConfigDir()
	path := filepath.Join(configHome, "filemover", "config.json")

	data, err := os.ReadFile(path)
	if err != nil {
		return "", "", "", "", "", err
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return "", "", "", "", "", err
	}

	downloads, _ := m["download"].(string)
	musics, _ := m["musics"].(string)
	videos, _ := m["videos"].(string)
	documents, _ := m["documents"].(string)
	images, _ := m["images"].(string)

	return downloads, musics, videos, documents, images, nil
}

func Move(srcPath string, destDir string) {
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		os.MkdirAll(destDir, 0o755)
	}

	fileName := filepath.Base(srcPath)
	destPath := filepath.Join(destDir, fileName)

	srcFile, err := os.Open(srcPath)
	if err != nil {
		fmt.Printf("Erro ao abrir origem %s: %v\n", srcPath, err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		fmt.Printf("Erro ao criar destino %s: %v\n", destPath, err)
		return
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Printf("Erro ao copiar para %s: %v\n", destPath, err)
		os.Remove(destPath)
		return
	}

	srcFile.Close()
	destFile.Close()

	err = os.Remove(srcPath)
	if err != nil {
		fmt.Printf("Erro ao remover origem %s: %v\n", srcPath, err)
		return
	}

	fmt.Println("Movido para:", destPath)
}

func OrganizeExisting(downloads, musics, videos, documents, images string) {
	entries, err := os.ReadDir(downloads)
	if err != nil {
		fmt.Printf("Erro ao ler diretório %s: %v\n", downloads, err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if strings.HasSuffix(name, ".crdownload") ||
			strings.HasSuffix(name, ".part") {
			continue
		}

		srcPath := filepath.Join(downloads, name)
		ext := filepath.Ext(name)

		for _, e := range entities.VideoExtensions {
			if ext == string(e) {
				Move(srcPath, videos)
			}
		}
		for _, e := range entities.DocumentExtensions {
			if ext == string(e) {
				Move(srcPath, documents)
			}
		}
		for _, e := range entities.MusicExtensions {
			if ext == string(e) {
				Move(srcPath, musics)
			}
		}
		for _, e := range entities.ImageExtensions {
			if ext == string(e) {
				Move(srcPath, images)
			}
		}
	}
}
