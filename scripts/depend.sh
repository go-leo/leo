#!/bin/sh

root=`pwd`

gomods=()

function recursive_list_dir(){
    for file_or_dir in `ls $1`
    do
        file=$1"/"$file_or_dir
        if [ -d $file ]
        then
            recursive_list_dir $file
        else
            if [ $file_or_dir = 'go.mod' ]
            then
                gomods[${#gomods[@]}]=$1
            fi
        fi
    done
}

recursive_list_dir $root

for gomod in ${gomods[@]}
do
    cd $gomod
    go get -u golang.org/x/...
    go get -u github.com/derekparker/trie
    go get -u github.com/stretchr/testify
    go mod tidy
    cd $root
done