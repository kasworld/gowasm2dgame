#!/usr/bin/env bash

################################################################################
echo "genlog -leveldatafile ./w2dlog/w2dlog.data -packagename w2dlog "
cd lib
genlog -leveldatafile ./w2dlog/w2dlog.data -packagename w2dlog 
cd ..

################################################################################
ProtocolW2DFiles="protocol_w2d/*.enum protocol_w2d/w2d_obj/protocol_*.go"
PROTOCOL_W2D_VERSION=`makesha256sum ${ProtocolW2DFiles}`
echo "Protocol W2D Version: ${PROTOCOL_W2D_VERSION}"

genprotocol -ver=${PROTOCOL_W2D_VERSION} -basedir=protocol_w2d -prefix=w2d -statstype=int
cd protocol_w2d
goimports -w .
cd ..

################################################################################
# generate enum
genenum -typename=TeamType -packagename=teamtype -basedir=enum -vectortype=int
genenum -typename=ActType -packagename=acttype -basedir=enum -vectortype=int
genenum -typename=GameObjType -packagename=gameobjtype -basedir=enum -vectortype=int
genenum -typename=EffectType -packagename=effecttype -basedir=enum -vectortype=int

cd enum 
goimports -w .
cd .. 

GameDataFiles="config/gameconst/*.go config/gamedata/*.go enum/*.enum"
Data_VERSION=`makesha256sum ${GameDataFiles}`
echo "Data Version: ${Data_VERSION}"
mkdir -p config/dataversion
echo "package dataversion
const DataVersion = \"${Data_VERSION}\"
" > config/dataversion/dataversion_gen.go 

################################################################################
# build bin

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

DATESTR=`date -Iseconds`
GITSTR=`git rev-parse HEAD`
BUILD_VER=${DATESTR}_${GITSTR}_release_linux
echo "Build version ${BUILD_VER}"

BIN_DIR="bin"
SRC_DIR="rundriver"

mkdir -p ${BIN_DIR}
echo ${BUILD_VER} > ${BIN_DIR}/BUILD_linux

BuildBin ${SRC_DIR}/server.go ${BIN_DIR} server
BuildBin ${SRC_DIR}/multiclient.go ${BIN_DIR} multiclient

cd rundriver
./genwasmclient.sh ${BUILD_VER}
cd ..

# copy data to bin folder
echo cp -r rundriver/serverdata ${BIN_DIR}
cp -r rundriver/serverdata ${BIN_DIR}
echo cp -r rundriver/clientdata ${BIN_DIR}
cp -r rundriver/clientdata ${BIN_DIR}

