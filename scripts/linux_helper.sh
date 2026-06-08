mkcd() {
  local path
  path=$(mkcd-bin "$@") || return 1
  cd "$path" || return 1
}