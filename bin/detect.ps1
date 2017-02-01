Param(
  [Parameter(Mandatory=$True,Position=1)]
    [string]$BuildDir
)

if ((Get-Item("$BuildDir\web.config"))) {
  $version = Get-Content "$PSScriptRoot\..\VERSION"
  echo "hwc $version"
} else {
  exit 1
}
