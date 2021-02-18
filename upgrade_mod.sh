#!/bin/bash

# 更新 Go mod

readonly root='/data/Aa/proj/go/src/github.com/hi-iwi'
readonly comment="${1:-'NO_COMMENT'}"
pushAndUpgradeMod(){
  echo "upgrading $1 ..."
  cd "$root/$1" || exit


  # 暂时用不到了
  # 私有库问题
  #rm -f go.mod
  #go mod init
  #env GIT_TERMINAL_PROMPT=1 go get -insecure github.com/hi-iwi/AaGo
  #go get -u ./...

  git add -A .
  git commit -m "$comment"
  git push origin master
}


pushAndUpgradeMod 'AaGo'  # 依赖 dtype



