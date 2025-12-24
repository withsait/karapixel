#!/bin/bash
set -e

echo "╔═══════════════════════════════════════════╗"
echo "║      KaraPanel Daemon Installer           ║"
echo "╚═══════════════════════════════════════════╝"

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "Please run as root"
    exit 1
fi

# Variables
INSTALL_DIR="/opt/karapanel"
SERVICE_NAME="karapanel-daemon"
BINARY_NAME="karapanel-daemon"

# Create directories
echo "[1/5] Creating directories..."
mkdir -p "$INSTALL_DIR/configs"

# Build binary (requires Go)
if command -v go &> /dev/null; then
    echo "[2/5] Building from source..."
    cd "$(dirname "$0")/../daemon"
    GOOS=linux GOARCH=amd64 go build -o "$INSTALL_DIR/$BINARY_NAME" .
else
    echo "[2/5] Go not found. Please copy pre-built binary to $INSTALL_DIR/$BINARY_NAME"
fi

# Copy config if not exists
if [ ! -f "$INSTALL_DIR/configs/config.yml" ]; then
    echo "[3/5] Copying config..."
    cp "$(dirname "$0")/../daemon/configs/config.yml" "$INSTALL_DIR/configs/"

    # Generate secure auth secret
    AUTH_SECRET=$(openssl rand -base64 32)
    sed -i "s/CHANGE_THIS_SECRET_IN_PRODUCTION/$AUTH_SECRET/" "$INSTALL_DIR/configs/config.yml"

    # Generate admin password hash
    echo "Enter admin password for KaraPanel:"
    read -s ADMIN_PASSWORD
    ADMIN_HASH=$(htpasswd -bnBC 10 "" "$ADMIN_PASSWORD" | tr -d ':\n' | sed 's/$2y/$2a/')
    sed -i "s|\$2a\$10\$example.hash.here|$ADMIN_HASH|" "$INSTALL_DIR/configs/config.yml"

    echo "Config created with new credentials"
else
    echo "[3/5] Config already exists, skipping..."
fi

# Install systemd service
echo "[4/5] Installing systemd service..."
cp "$(dirname "$0")/karapanel-daemon.service" /etc/systemd/system/
systemctl daemon-reload
systemctl enable "$SERVICE_NAME"

# Start service
echo "[5/5] Starting service..."
systemctl start "$SERVICE_NAME"

echo ""
echo "╔═══════════════════════════════════════════╗"
echo "║      Installation Complete!               ║"
echo "╠═══════════════════════════════════════════╣"
echo "║  Service: systemctl status $SERVICE_NAME  ║"
echo "║  Logs: journalctl -u $SERVICE_NAME -f     ║"
echo "║  API: http://localhost:8080               ║"
echo "╚═══════════════════════════════════════════╝"
