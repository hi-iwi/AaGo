#!/bin/bash

# Go 必须要开启支持 vendor 功能


readonly AaGoReporitories="AaGo"

initProject() {
    dir="$1"
    if [ ! -d "$dir" ]; then
        echo "Project directory $dir is not found"
        exit 1
    fi
    mkdir -p $dir'/app/aservice'

    mkdir -p $dir'/app/biz_test/cache'
    mkdir -p $dir'/app/biz_test/entity'
    mkdir -p $dir'/app/biz_test/module'
    mkdir -p $dir'/app/biz_test/mservice'

    mkdir -p $dir'/app/cache/proxycache'
    mkdir -p $dir'/app/router/middleware'

    mkdir -p $dir'/app/service/mq'
    mkdir -p $dir'/app/controller'


    mkdir -p $dir'/bootstrap/register'
    mkdir -p $dir'/conf'
    mkdir -p $dir'/deploy/config'
    mkdir -p $dir'/deploy/public'
    mkdir -p $dir'/deploy/view_src'

    mkdir -p $dir'/dic'
    mkdir -p $dir'/docs'
    mkdir -p $dir'/driver'
    mkdir -p $dir'/enum'
    mkdir -p $dir'/helper'

    # mkdir -p $dir'/app/register'
    mkdir -p $dir'/sdk'
    mkdir -p $dir'/storage/logs'
    mkdir -p $dir'/tests'   # 其他测试，都是直接在源文件下  _test.go 后缀即可，不需要重新建目录。这里是系统测试
}

goGetIwiReporitories() {
    update=0
    if [ "$1" == "update" ]; then
        update=1
    fi
    l="${GOPATH}/src/github.com/hi-iwi"
    cd $l
    for repo in $AaGoReporitories; do
        if [ $update -eq 1 ]; then
            go get -u -v "github.com/hi-iwi/${repo}"
        else
            go get -v "github.com/hi-iwi/${repo}"
        fi
    done
}

goGet() {
    p=$(pwd)
    gopath=$GOPATH
    repo="$1"
    update=0
    if [ "$2" == "update" ]; then
        update=1
    fi
    if [ ! -d "${p}/vendor" ]; then
        echo "vendor not found"
        exit 1
    fi

    # 优先在GOPATH下找，若没有，就根据情况下载到 vendor 下
    isInGoPath=0
    if [ -d "${GOPATH}/src/${repo}/.git" ]; then
        isInGoPath=1
    fi

    #  的一律放到$GOPATH下共用
    if [ $isInGoPath -eq 0 -a "${repo:0:17}" != "github.com/hi-iwi/" ]; then
        GOPATH="${p}/vendor"
    fi

    if [ -d "${p}/vendor/src" ]; then
        mv ${p}/vendor/src/* ${p}/vendor/
        rm -rf ${p}/vendor/src
    fi
    rm -rf "${p}/vendor/pkg"
    mkdir -p "${p}/.vendor/src"
    mv ${p}/vendor/* "${p}/.vendor/src"
    rm -rf "${p}/vendor"
    mv "${p}/.vendor" "${p}/vendor"

    src="${p}/vendor/src"

    if [ $update -eq 1 ]; then
        echo "go get -u -v $repo"
        go get -u -v $repo
    else
        echo "go get -v $repo"
        go get -v $repo
    fi

    mv ${p}/vendor/src/* ${p}/vendor/
    rm -rf ${p}/vendor/src
    rm -rf ${p}/vendor/pkg
    
    GOPATH=$gopath
    # 使用 go get 可以下载相关库
}

# 临时方案不用包管理，防止一些网络条件差情况下，dep 包管理总是超时问题；
depEnsureAdd() {
    update="$1"
    repo="$2"
    if [ "$repo" == "luexu" ]; then
        goGetIwiReporitories $update
    else
        goGet $repo $update
    fi
}

while getopts ':p:d:u:' opt
do
    case $opt in
    d)
        depEnsureAdd get ${OPTARG}
    ;;
    p)
        initProject ${OPTARG}
    ;;
    h)
        cat << EOF
Usage: AaGo.sh [\$options]
    -p <dir>   : create directories for new project
EOF
        exit 0
    ;;
    u)
        depEnsureAdd update ${OPTARG}
    ;;
    ?)
        echo "未知参数"
        exit 1
    ;;
    esac
done

shift $((OPTIND-1))