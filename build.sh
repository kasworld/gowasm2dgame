#!/usr/bin/env bash


DATESTR=`date -Iseconds`
GITSTR=`git rev-parse HEAD`
BUILD_VER=${DATESTR}_${GITSTR}_release
echo "Build" ${BUILD_VER}


BuildBin() {
    local srcfile=${1}
    local dstdir=${2}
    local dstfile=${3}
    local args="-X main.Ver=${BUILD_VER}"

    echo "[BuildBin] go build -i -o ${dstdir}/${dstfile} -ldflags "${args}" ${srcfile}"

    mkdir -p ${dstdir}
    go build -i -o ${dstdir}/${dstfile} -ldflags "${args}" ${srcfile}

    if [ ! -f "${dstdir}/${dstfile}" ]; then
        echo "${dstdir}/${dstfile} build fail, build file: ${srcfile}"
        # exit 1
    fi
    strip "${dstdir}/${dstfile}"
}


cd lib
genlog -leveldatafile ./w2dlog/w2dlog.data -packagename w2dlog 
cd ..


ProtocolW2DFiles="protocol_w2d/w2d_gendata/command.data \
protocol_w2d/w2d_gendata/error.data \
protocol_w2d/w2d_gendata/noti.data \
"

PROTOCOL_W2D_VERSION=`cat ${ProtocolW2DFiles}| sha256sum | awk '{print $1}'`
echo "Protocol W2D Version:" ${PROTOCOL_W2D_VERSION}

cd protocol_w2d
genprotocol -ver=${PROTOCOL_W2D_VERSION} \
    -basedir=. \
    -prefix=w2d -statstype=int

goimports -w .

cd ..

genenum -typename=TeamType -packagename=teamtype -basedir=enums -statstype=int
goimports -w enums/teamtype/teamtype_gen.go
goimports -w enums/teamtype_stats/teamtype_stats_gen.go

genenum -typename=ActType -packagename=acttype -basedir=enums -statstype=int
goimports -w enums/acttype/acttype_gen.go
goimports -w enums/acttype_stats/acttype_stats_gen.go

genenum -typename=GameObjType -packagename=gameobjtype -basedir=enums -statstype=int
goimports -w enums/gameobjtype/gameobjtype_gen.go
goimports -w enums/gameobjtype_stats/gameobjtype_stats_gen.go

genenum -typename=EffectType -packagename=effecttype -basedir=enums -statstype=int
goimports -w enums/effecttype/effecttype_gen.go
goimports -w enums/effecttype_stats/effecttype_stats_gen.go

GameDataFiles="
game/gameconst/gameconst.go \
game/gameconst/serviceconst.go \
game/gamedata/*.go \
enums/*.enum \
"
Data_VERSION=`cat ${GameDataFiles}| sha256sum | awk '{print $1}'`
echo "Data Version:" ${Data_VERSION}

echo "
package gameconst

const DataVersion = \"${Data_VERSION}\"
" > game/gameconst/dataversion_gen.go 

# build bin

BIN_DIR="bin"
SRC_DIR="rundriver"

echo ${BUILD_VER} > ${BIN_DIR}/BUILD

BuildBin ${SRC_DIR}/server.go ${BIN_DIR} server
BuildBin ${SRC_DIR}/multiclient.go ${BIN_DIR} multiclient

cd rundriver
echo "build wasm client"
GOOS=js GOARCH=wasm go build -o www/wasmclient.wasm wasmclient.go
cd ..