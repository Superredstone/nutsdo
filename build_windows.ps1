if([System.IO.File]::Exists(".\build")){
    mkdir build
}

$env:GOARCH = "amd64"
$env:GOOS = "linux"
go build -o build/nutsdo-linux-amd64 .
$env:GOOS = "darwin"
go build -o build/nutsdo-darwin-amd64 .
$env:GOOS = "windows"
go build -o build/nutsdo-windows-amd64.exe .