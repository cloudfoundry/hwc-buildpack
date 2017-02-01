Param(
  [Parameter(Mandatory=$True,Position=1)]
    [string]$BuildDir,

  [Parameter(Mandatory=$True,Position=2)]
    [string]$CacheDir
)

$ErrorActionPreference = "Stop";
trap { $host.SetShouldExit(1) }

$goInstallDir = "$Env:USERPROFILE\tmp\"
$buildpackRoot = "$PSScriptRoot\.."
$dependencies = "$buildpackRoot\dependencies"
$goVersion = "1.7.5"
$goZipHash = "f75db843f92eb873e8076f0d1ab9c0b7"

if(Test-Path -Path $dependencies){
  $archive = Get-Item("$dependencies\*go$goVersion*.zip")
} else {
  $archive = "$goInstallDir\go$goVersion.windows-amd64.zip"

  echo "-----> Downloading go $goVersion"
  (New-Object System.Net.WebClient).DownloadFile(
    "https://storage.googleapis.com/golang/go$goVersion.windows-amd64.zip",
    $archive)
  echo "       Done"

  $downloadHash = (Get-FileHash $archive -Algorithm MD5).Hash

  if ($downloadHash -ne $goZipHash) {
    echo "MD5 mismatch: expected $goZipHash, got $downloadHash"
    exit 1
  }
}

echo "-----> Extracting go to $goInstallDir\go..."
Add-Type -assembly "system.io.compression.filesystem"
[io.compression.zipfile]::ExtractToDirectory($archive, $goInstallDir)
echo "       Done"

$env:Path += ";$goInstallDir\go\bin\;$buildpackRoot\bin\"
$env:GOROOT = "$goInstallDir\go"
$env:GOPATH = $buildpackRoot

echo "-----> Compiling src/compile/compile.go"
go build -o "$buildpackRoot\bin\compile.exe" "compile"
echo "       Done"

echo "-----> Running bin/compile.exe"
compile.exe $BuildDir $CacheDir
echo "       Done"
