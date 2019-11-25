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

goimports -w w2d_version/version_gen.go
goimports -w w2d_idcmd/command_gen.go
goimports -w w2d_idnoti/noti_gen.go
goimports -w w2d_error/error_gen.go
goimports -w w2d_const/const_gen.go
goimports -w w2d_packet/packet_gen.go
goimports -w w2d_obj/objtemplate_gen.go
goimports -w w2d_msgp/serialize_gen.go
goimports -w w2d_json/serialize_gen.go
goimports -w w2d_gob/serialize_gen.go
goimports -w w2d_handlersp/fnobjtemplate_gen.go
goimports -w w2d_handlersp/fnbytestemplate_gen.go
goimports -w w2d_handlereq/fnobjtemplate_gen.go
goimports -w w2d_handlereq/fnbytestemplate_gen.go
goimports -w w2d_handlenoti/fnobjtemplate_gen.go
goimports -w w2d_handlenoti/fnbytestemplate_gen.go
goimports -w w2d_serveconnbyte/serveconnbyte_gen.go
goimports -w w2d_conntcp/conntcp_gen.go
goimports -w w2d_connwasm/connwasm_gen.go
goimports -w w2d_connwsgorilla/connwsgorilla_gen.go
goimports -w w2d_loopwsgorilla/loopwsgorilla_gen.go
goimports -w w2d_looptcp/looptcp_gen.go
goimports -w w2d_pid2rspfn/pid2rspfn_gen.go
goimports -w w2d_statnoti/statnoti_gen.go
goimports -w w2d_statcallapi/statcallapi_gen.go
goimports -w w2d_statserveapi/statserveapi_gen.go
goimports -w w2d_statapierror/statapierror_gen.go
goimports -w w2d_authorize/authorize_gen.go
goimports -w w2d_error_stats/w2d_error_stats_gen.go
goimports -w w2d_idcmd_stats/w2d_idcmd_stats_gen.go
goimports -w w2d_idnoti_stats/w2d_idnoti_stats_gen.go

cd ..

genenum -typename=ActType -packagename=acttype -basedir=enums -statstype=int
goimports -w enums/acttype/acttype_gen.go
goimports -w enums/acttype_stats/acttype_stats_gen.go

genenum -typename=GameObjType -packagename=gameobjtype -basedir=enums -statstype=int
goimports -w enums/gameobjtype/gameobjtype_gen.go
goimports -w enums/gameobjtype_stats/gameobjtype_stats_gen.go

genenum -typename=TeamType -packagename=teamtype -basedir=enums -statstype=int
goimports -w enums/teamtype/teamtype_gen.go
goimports -w enums/teamtype_stats/teamtype_stats_gen.go


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
BuildBin ${SRC_DIR}/client.go ${BIN_DIR} client

cd rundriver
echo "build wasm client"
GOOS=js GOARCH=wasm go build -o www/wasmclient.wasm wasmclient.go
cd ..