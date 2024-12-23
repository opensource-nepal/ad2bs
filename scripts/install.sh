#!/usr/bin/env bash

# TODO: support for windows

set -e

echo "Initializing download..."

# arguments
VERSION_ARG=$1

# Constants
REPO_URL="https://github.com/opensource-nepal/ad2bs"
RELEASES_URL="${REPO_URL}/releases"
FILE_BASENAME="ad2bs"
BINDIR=/usr/local/bin
BINARIES="ad2bs bs2ad"

github_release() {
    version=$1
    test -z "$version" && version="latest"
    json=$(curl -sL -H "Accept:application/json" "${RELEASES_URL}/${version}")

    # Check if JSON contains an error
    error=$(echo "$json" | grep -o '"error":"[^"]*' | sed 's/"error":"//')
    if [ "$error" = "Not Found" ]; then
        return 1
    fi

    # Extract version from JSON
    version=$(echo "$json" | tr -s '\n' ' ' | sed 's/.*"tag_name":"//' | sed 's/".*//')
    test -z "$version" && return 1
    echo "$version"
}

VERSION=$(github_release $VERSION_ARG) && true
VERSION_NUMBER=${VERSION#v}
if test -z "$VERSION"; then
    echo "Unable to find version '${VERSION_ARG}' - use 'latest' or see ${RELEASES_URL} for details" >&2
    exit 1
fi

# Generating tar filename
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH="$(uname -m)"
test "$ARCH" = "aarch64" && ARCH="arm64"
TAR_FILE="${FILE_BASENAME}_${VERSION_NUMBER}_${OS}_${ARCH}.tar.gz"

TMP_DIR="$(mktemp -d)"
cd "$TMP_DIR"
echo "Downloading ad2bs $VERSION"
curl -sfLO "$RELEASES_URL/download/$VERSION/$TAR_FILE" || {
    echo "Download failed. Version: $VERSION, OS: $OS $ARCH" >&2
    echo "Please raise an issue in ${REPO_URL}/issues" >&2
    exit 1
}

CHECKSUM_FILE="${FILE_BASENAME}_${VERSION_NUMBER}_checksums.txt"
curl -sfLO "$RELEASES_URL/download/$VERSION/$CHECKSUM_FILE" || {
    echo "Failed to download checksum. Version: $VERSION" >&2
    exit 1
}
# TODO: checksum check
echo "Download complete!"

echo "Installing ad2bs $VERSION"
tar -xf "$TMP_DIR/$TAR_FILE" -C "$TMP_DIR"
rm -rf "$TMP_DIR/$TAR_FILE"

# TEMP: for temp bs2ad
cp "${TMP_DIR}/ad2bs" "${TMP_DIR}/bs2ad"

# Installing binaries
for binexe in $BINARIES; do
    if [ "$OS" = "windows" ]; then
        binexe="${binexe}.exe"
    fi

    sudo install "${TMP_DIR}/$binexe" "${BINDIR}/"
done

rm -rf $TMP_DIR

echo "ad2bs $VERSION installed successfully!"
