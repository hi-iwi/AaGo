#!/bin/bash

# 更新 Go mod

readonly root='/data/Aa/proj/go/src/github.com/hi-iwi'
comment="NO_COMMENT"
upgrade=0

for arg in "$@"; do
  case "$arg" in
    -u)
      upgrade=1
      ;;
    *)
      comment="$arg"
      ;;
  esac
done

pushAndUpgradeMod(){
  cd "$root/$1" || exit

  if [ $upgrade -eq 1 ]; then
    echo ">>> UPGRADING go.mod..."
    rm -f go.mod
    go mod init
    # 私有库问题
    env GIT_TERMINAL_PROMPT=1 go get -insecure github.com/hi-iwi/AaGo
    go get -u ./...
  fi
  echo ">>> git comment: $comment"
  git add -A .
  git commit -m "$comment"
  git push origin master
}

pushAndUpgradeMod 'AaGo'



