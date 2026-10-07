#!/usr/bin/env sh
set -eu

repo="Pointdexter37/Kin"
install_dir="${KIN_INSTALL_DIR:-$HOME/.local/bin}"
tmp_dir="$(mktemp -d)"

cleanup() {
    rm -rf "$tmp_dir"
}
trap cleanup EXIT

os="$(uname -s | tr '[:upper:]' '[:lower:]')"
case "$os" in
    linux) os="linux" ;;
    darwin) os="darwin" ;;
    *)
        printf '%s\n' "Unsupported operating system: $os" >&2
        exit 1
        ;;
esac

arch="$(uname -m)"
case "$arch" in
    x86_64|amd64) arch="amd64" ;;
    arm64|aarch64) arch="arm64" ;;
    *)
        printf '%s\n' "Unsupported architecture: $arch" >&2
        exit 1
        ;;
esac

release_json="$tmp_dir/release.json"
curl --fail --silent --show-error --location \
    -A "kin-installer" \
    "https://api.github.com/repos/$repo/releases/latest" > "$release_json"

asset_name="$(grep -o "\"name\": \"Kin_[^\"]*_${os}_${arch}\.tar\.gz\"" "$release_json" |
    sed -n 's/.*"name": "\([^"]*\)".*/\1/p' | head -n 1)"
if [ -z "$asset_name" ]; then
    printf '%s\n' "The latest release does not contain an archive for $os/$arch." >&2
    exit 1
fi

asset_url="https://github.com/$repo/releases/latest/download/$asset_name"
checksum_url="https://github.com/$repo/releases/latest/download/checksums.txt"
curl --fail --silent --show-error --location -A "kin-installer" "$asset_url" -o "$tmp_dir/archive.tar.gz"
curl --fail --silent --show-error --location -A "kin-installer" "$checksum_url" -o "$tmp_dir/checksums.txt"

expected_hash="$(grep " $asset_name$" "$tmp_dir/checksums.txt" | awk '{print $1}')"
if [ -z "$expected_hash" ]; then
    printf '%s\n' "No checksum was found for $asset_name." >&2
    exit 1
fi

if command -v sha256sum >/dev/null 2>&1; then
    actual_hash="$(sha256sum "$tmp_dir/archive.tar.gz" | awk '{print $1}')"
elif command -v shasum >/dev/null 2>&1; then
    actual_hash="$(shasum -a 256 "$tmp_dir/archive.tar.gz" | awk '{print $1}')"
else
    printf '%s\n' "A SHA-256 utility (sha256sum or shasum) is required." >&2
    exit 1
fi

if [ "$actual_hash" != "$expected_hash" ]; then
    printf '%s\n' "Checksum verification failed for $asset_name." >&2
    exit 1
fi

mkdir -p "$install_dir"
tar -xzf "$tmp_dir/archive.tar.gz" -C "$tmp_dir"
install -m 0755 "$tmp_dir/kin" "$install_dir/kin"

printf 'Installed Kin to %s\n' "$install_dir/kin"
printf '%s\n' "Run: kin init"
