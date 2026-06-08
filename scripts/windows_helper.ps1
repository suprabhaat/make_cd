function mkcd {
  $path = mkcd-bin @args
  if ($LASTEXITCODE -ne 0) { return }
  Set-Location $path
}