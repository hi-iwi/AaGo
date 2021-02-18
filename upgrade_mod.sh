#!/bin/bash

# 更新 Go mod

readonly root='/data/Aa/proj/go/src/github.com/hi-iwi'
readonly comment="${1:-'NO_COMMENT'}"
upgrade=0


while getopts "u" opt
do
  # shellcheck disable=SC2220
  case $opt in
  u)
    echo "upgrading go.mod..."
    upgrade=1
  ;;
  esac
done


pushAndUpgradeMod(){
  echo "upgrading $1 ..."
  cd "$root/$1" || exit

  if [ $upgrade -eq 1 ]; then
    rm -f go.mod
    go mod init
      # 私有库问题
    env GIT_TERMINAL_PROMPT=1 go get -insecure github.com/hi-iwi/AaGo
    go get -u ./...
  fi
  git add -A .
  git commit -m "$comment"
  git push origin master
}


pushAndUpgradeMod 'AaGo'



