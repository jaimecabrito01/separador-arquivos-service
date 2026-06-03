#!/bin/bash

SERVICE_NAME="filemover"
BINARY_NAME="filemover"
INSTALL_DIR="/usr/local/bin"
SERVICE_FILE="/etc/systemd/system/$SERVICE_NAME.service"

echo ">>> Instalando $SERVICE_NAME ..."

echo ">>> Compilando binário..."
go build -o $BINARY_NAME ./cmd/daemon/

echo ">>> Movendo binário para $INSTALL_DIR..."
sudo mv $BINARY_NAME $INSTALL_DIR/

echo ">>> Criando service systemd..."
sudo bash -c "cat > $SERVICE_FILE" <<EOF
[Unit]
Description=Separador de arquivos automático
After=network.target

[Service]
ExecStart=$INSTALL_DIR/$BINARY_NAME
Restart=always
User=$USER
StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
EOF

echo ">>> Recarregando systemd..."
sudo systemctl daemon-reload

echo ">>> Ativando serviço..."
sudo systemctl enable $SERVICE_NAME

echo ">>> Executando configuração inicial..."
$INSTALL_DIR/$BINARY_NAME --setup

echo ">>> Iniciando $SERVICE_NAME..."
sudo systemctl start $SERVICE_NAME

echo ">>> Instalação concluída!"
echo "Status do serviço:"
systemctl status $SERVICE_NAME --no-pager
