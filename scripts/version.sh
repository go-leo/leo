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

tags=()

for gomod in ${gomods[@]}
do
    if [ $gomod = $root ]
    then
       tags[${#tags[@]}]=$1
#       echo $1
    else
       dir=${gomod: ${#root}+1}
       tags[${#tags[@]}]=$dir"/"$1
#       echo $dir"/"$1
    fi
done

for tag in ${tags[@]}
do
    git tag $tag
done
