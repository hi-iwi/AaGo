#!/bin/bash

# 更新 Go mod

readonly root='/data/Aa/proj/go/src/github.com/hi-iwi'
readonly comment="${1:-'NO_COMMENT'}"
pushAndUpgradeMod(){
  echo "upgrading $1 ..."
  cd "$root/$1" || exit

#  sed -i '/github\.com\/hi-iwi\/aenum/d' go.mod
#  sed -i '/github\.com\/hi-iwi\/aenum/d' go.sum
#  sed -i '/github\.com\/hi-iwi\/code/d' go.mod
#  sed -i '/github\.com\/hi-iwi\/code/d' go.sum
#  sed -i '/github\.com\/hi-iwi\/dtype/d' go.mod
#  sed -i '/github\.com\/hi-iwi\/dtype/d' go.sum
#  sed -i '/github\.com\/hi-iwi\/AaGo/d' go.mod
#  sed -i '/github\.com\/hi-iwi\/AaGo/d' go.sum
#  sed -i '/github\.com\/hi-iwi\/aorm/d' go.mod
#  sed -i '/github\.com\/hi-iwi\/aorm/d' go.sum


  rm -f go.mod
  go mod init

  # 私有库问题
  env GIT_TERMINAL_PROMPT=1 go get -insecure github.com/hi-iwi/aenum
  env GIT_TERMINAL_PROMPT=1 go get -insecure github.com/hi-iwi/code
  env GIT_TERMINAL_PROMPT=1 go get -insecure github.com/hi-iwi/dtype
  env GIT_TERMINAL_PROMPT=1 go get -insecure github.com/hi-iwi/AaGo
  env GIT_TERMINAL_PROMPT=1 go get -insecure github.com/hi-iwi/aorm

  go get -u ./...
  git add -A .
  git commit -m "$comment"
  git push origin master
}


pushAndUpgradeMod 'aenum'
pushAndUpgradeMod 'code'

pushAndUpgradeMod 'dtype'
pushAndUpgradeMod 'AaGo'  # 依赖 dtype
pushAndUpgradeMod 'aorm'  # 依赖 AaGo


