#!/bin/sh

set -e

get_arch() {
  a=$(uname -m)
  case ${a} in
  "x86_64" | "amd64")
    echo "x86_64"
    ;;
  "i386")
    echo "i386"
    ;;
  "aarch64" | "arm64" | "arm")
    echo "arm64"
    ;;
  *)
    echo "${NIL}"
    ;;
  esac
}

get_os() {
  uname -s | awk '{print tolower($0)}'
}

owner="RossyWhite"
repo="tfsubst"
exe_name="tfsubst"
version=""

# parse flag
for i in "$@"; do
  case $i in
  -v=* | --version=*)
    version="${i#*=}"
    shift # past argument=value
    ;;
  *)
    # unknown option
    ;;
  esac
done

downloadFolder="$(mktemp -d)"
os=$(get_os)
arch=$(get_arch)
file_name="${exe_name}_${os}_${arch}.tar.gz"
downloaded_file="${downloadFolder}/${file_name}"
executable_folder="/usr/local/bin"

# if version is empty
if [ -z "$version" ]; then
  version="latest"
fi

asset_id=$(curl \
  -H "Accept: application/vnd.github+json" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  -sSf https://api.github.com/repos/"${owner}"/"${repo}"/releases/"${version}" \
  | jq -r ".assets[] | select(.name == \"${file_name}\") | .id")

echo "[1/3] Downloading ${file_name}"
rm -f "${downloaded_file}"
curl -fsSL -H "Accept: application/octet-stream" \
  -H "X-GitHub-Api-Version: 2022-11-28" \
  https://api.github.com/repos/"${owner}"/"${repo}"/releases/assets/"${asset_id}" \
  -o "${downloaded_file}"

echo "[2/3] Install ${exe_name} to the ${executable_folder}"
tar -xz -f "${downloaded_file}" -C ${executable_folder}
exe=${executable_folder}/${exe_name}
chmod +x "${exe}"

echo "[3/3] Set environment variables"
echo "${exe_name} was installed successfully to ${exe}"
if command -v "$exe_name" --help >/dev/null; then
  echo "Run '$exe_name --help' to get started"
else
  echo "Manually add the directory to your \$HOME/.bash_profile (or similar)"
  echo "  export PATH=${executable_folder}:\$PATH"
  echo "Run '$exe_name --help' to get started"
fi

exit 0
