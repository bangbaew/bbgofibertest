cd "$(dirname "$BASH_SOURCE")" || {
    echo "Error getting script directory" >&2
    exit 1
}
mutagen-compose -f docker-compose.mutagen.yml up -d