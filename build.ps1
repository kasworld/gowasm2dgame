
# del generated code 
# Get-ChildItem .\enum\ -Recurse -Include *_gen.go | Remove-Item
# Get-ChildItem .\protocol_w2d\ -Recurse -Include *_gen.go | Remove-Item
# Remove-Item lib\w2dlog\log_gen.go
# Remove-Item config/dataversion/dataversion_gen.go 

################################################################################
Set-Location lib
Write-Output "genlog -leveldatafile ./w2dlog/w2dlog.data -packagename w2dlog "
genlog -leveldatafile ./w2dlog/w2dlog.data -packagename w2dlog 
Set-Location ..

################################################################################
$PROTOCOL_W2D_VERSION=makesha256sum protocol_w2d/*.enum protocol_w2d/w2d_obj/protocol_*.go
Write-Output "Protocol W2D Version: ${PROTOCOL_W2D_VERSION}"
Write-Output "genprotocol -ver=${PROTOCOL_W2D_VERSION} -basedir=protocol_w2d -prefix=w2d -statstype=int"
genprotocol -ver="${PROTOCOL_W2D_VERSION}" -basedir=protocol_w2d -prefix=w2d -statstype=int
Set-Location protocol_w2d
goimports -w .
Set-Location ..

################################################################################
# generate enum
Write-Output "generate enums"
genenum -typename=TeamType -packagename=teamtype -basedir=enum -vectortype=int
genenum -typename=ActType -packagename=acttype -basedir=enum -vectortype=int
genenum -typename=GameObjType -packagename=gameobjtype -basedir=enum -vectortype=int
genenum -typename=EffectType -packagename=effecttype -basedir=enum -vectortype=int

Set-Location enum
goimports -w .
Set-Location ..

$Data_VERSION=makesha256sum config/gameconst/*.go config/gamedata/*.go enum/*.enum
Write-Output "Data Version: ${Data_VERSION}"
mkdir -ErrorAction SilentlyContinue config/dataversion
Write-Output "package dataversion
const DataVersion = `"${Data_VERSION}`" 
" > config/dataversion/dataversion_gen.go 


################################################################################
$DATESTR=Get-Date -UFormat '+%Y-%m-%dT%H:%M:%S%Z:00'
$GITSTR=git rev-parse HEAD
$BUILD_VER="${DATESTR}_${GITSTR}_release_windows"
Write-Output "Build Version: ${BUILD_VER}"

################################################################################
# build bin

$BIN_DIR="bin"
$SRC_DIR="rundriver"

mkdir -ErrorAction SilentlyContinue "${BIN_DIR}"
Write-Output ${BUILD_VER} > ${BIN_DIR}/BUILD_windows

# build bin here
go build -o "${BIN_DIR}\server.exe" -ldflags "-X main.Ver=${BUILD_VER}" "${SRC_DIR}\serverwin.go"
go build -o "${BIN_DIR}\multiclient.exe" -ldflags "-X main.Ver=${BUILD_VER}" "${SRC_DIR}\multiclient.go"

Set-Location rundriver
./genwasmclient.ps1 ${BUILD_VER}
Set-Location ..

Write-Output "cp -r rundriver/serverdata ${BIN_DIR}"
Copy-Item -Force -r rundriver/serverdata ${BIN_DIR}
Write-Output "cp -r rundriver/clientdata ${BIN_DIR}"
Copy-Item -Force -r rundriver/clientdata ${BIN_DIR}

