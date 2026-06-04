# FileMover

Um daemon que monitora sua pasta de downloads e move arquivos novos para pastas separadas por tipo (Vídeos, Músicas, Imagens, Documentos).

## Pré-requisitos

- Go 1.25 ou superior ([download](https://go.dev/dl/))

## Instalação

```bash
git clone https://github.com/jaimecabrito01/file-mover.git
cd file-mover
chmod +x install.sh
./install.sh
```

O script compila o binário, instala como serviço do systemd e inicia o daemon.

Para desinstalar:

```bash
chmod +x uninstall.sh
./uninstall.sh
```

## Uso

```bash
filemover --setup    # Configura os caminhos (interativo)
filemover --now      # Organiza arquivos já existentes na pasta de origem
filemover            # Inicia o daemon (monitora novos arquivos)
```

Na primeira execução, rode `--setup` para definir a pasta de origem (ex: Downloads) e as pastas de destino para cada tipo de arquivo.

Use `--now` para uma organização única dos arquivos que já estão na pasta de origem — útil após a configuração inicial.

Sem flags, o programa inicia o daemon em background e passa a monitorar a pasta configurada.

## Daemon (systemd)

```bash
sudo systemctl status filemover   # status do serviço
sudo journalctl -u filemover -f   # logs em tempo real
sudo systemctl stop filemover     # parar o daemon
sudo systemctl disable filemover  # remover da inicialização
```

## Configuração

O arquivo de configuração fica em `~/.config/filemover/config.json`:

```json
{
  "download": "/home/user/Downloads",
  "images": "/home/user/Pictures",
  "videos": "/home/user/Videos",
  "musics": "/home/user/Music",
  "documents": "/home/user/Documents"
}
```

## Licença

MIT
