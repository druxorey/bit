#!/bin/sh
set -e

# Bit installer script
# Usage: curl -sfL https://raw.githubusercontent.com/paulilaaso/bit/main/install.sh | sh

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
    linux)
        OS="Linux"
        ;;
    darwin)
        OS="Darwin"
        ;;
    *)
        echo "Unsupported operating system: $OS"
        exit 1
        ;;
esac

case "$ARCH" in
    x86_64)
        ARCH="x86_64"
        ;;
    amd64)
        ARCH="x86_64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Get latest release version
LATEST_VERSION=$(curl -s https://api.github.com/repos/paulilaaso/bit/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo "Failed to get latest version"
    exit 1
fi

echo "Installing Bit $LATEST_VERSION for $OS $ARCH..."

# Construct download URL
FILENAME="bit_${LATEST_VERSION#v}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/paulilaaso/bit/releases/download/$LATEST_VERSION/$FILENAME"

# Create temp directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

# Download and extract
echo "Downloading $URL..."
curl -sfL "$URL" -o "$FILENAME"

echo "Extracting..."
tar -xzf "$FILENAME"

# Install binary
INSTALL_DIR="/usr/local/bin"

if [ -w "$INSTALL_DIR" ]; then
    mv bit "$INSTALL_DIR/"
else
    echo "Installing to $INSTALL_DIR (requires sudo)..."
    sudo mv bit "$INSTALL_DIR/"
fi

# Cleanup
cd -
rm -rf "$TMP_DIR"

echo ""
echo "✓ Bit installed successfully!"
echo ""
echo "Usage:"
echo "  bit              - Start interactive UI"
echo "  bit -list        - List all available fonts"
echo "  bit \"Hello\"    - Quick render text"
echo "  bit -help        - Show all options"
