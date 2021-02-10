#!/bin/bash

# 更新 Go mod

readonly root='/data/Aa/proj/go/src/github.com/hi-iwi'
readonly comment="${1:-'NO_COMMENT'}"
pushAndUpgradeMod(){
  cd $root'/'$1 || exit
  go get -u
  git add -A .
  git commit -m "$comment"
  git push origin master
}


pushAndUpgradeMod 'aenum'
pushAndUpgradeMod 'code'
pushAndUpgradeMod 'dtype'
pushAndUpgradeMod 'AaGo'  # 依赖 dtype
pushAndUpgradeMod 'aorm'  # 依赖 AaGo


