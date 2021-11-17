#!/bin/bash

# 更新 Go mod

readonly root='/data/Aa/proj/go/src/github.com/hi-iwi'
comment="NO_COMMENT"
upgrade=0
incrTag=0

for arg in "$@"; do
  case "$arg" in
    -u)
      upgrade=1
      ;;
    -t)
      incrTag=1
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
    #env GIT_TERMINAL_PROMPT=1 go get -insecure github.com/hi-iwi/AaGo
  fi
  echo ">>> git comment: $comment"
  go build
  go get -u ./...
  go mod tidy
  git add -A .
  git commit -m "$comment"
  git push origin master
  # 自增tag
  if [ $incrTag -eq 1 ]; then
    echo "adding tag..."
    git fetch --tags
    latestTag=$(git describe --tags "$(git rev-list --tags --max-count=1)")
    if [ "$latestTag"  != "" ]; then
      tag=${latestTag%.*}
      id=${latestTag##*.}
      id=$((id+1))
      newTag=$tag'.'$id
      git tag -d "$latestTag"
      git push origin --delete tag "$latestTag"
      git tag "$newTag"
      git push origin --tags
    fi
  fi
}

pushAndUpgradeMod 'AaGo'



