#!/bin/sh

# 默认值
DRY_RUN=0
VERSION=""
# 1. 处理参数
while [[ $# -gt 0 ]]; do
    case "$1" in --dry-run=*)
            DRY_RUN="${1#*=}"
            ;;
        *)
            echo "Unknown option: $1"
            exit 1
            ;;
    esac
    shift
done
echo "DRY_RUN: $DRY_RUN"



# 2. 递归查找文件夹中的 go.mod
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

gomods=()
root=`pwd`
recursive_list_dir $root


# 3. 确认版本号
version=""
if [ "$DRY_RUN" -eq 1 ]; then
    # 不执行
    latest_tag=$(semantic-release --dry-run)
    echo $latest_tag
    if [[ $latest_tag =~ "The next release version is "([[:digit:]\.]+).*$ ]]; then
        last_segment=${BASH_REMATCH[1]}
    fi

    version="v"$last_segment
else
    # 执行
    semantic-release
    git checkout develop
    git merge master
    git push origin develop
    git checkout master
    git merge develop
    git push origin master

    # 获取最新 tag
    latest_tag=$(git describe --tags $(git rev-list --tags --max-count=1))
    # 使用正则表达式匹配版本号的最后一段
    if [[ $latest_tag =~ \/v([[:digit:]\.]+)$ ]]; then
        last_segment=${BASH_REMATCH[1]}
    fi

    version="v"$last_segment
fi

echo "The next release version is "$version


# 初始化 tags 数组
tags=()

# 4. 确定各子包 tag 名称
for gomod in ${gomods[@]}
do
    if [ $gomod != $root ]; then
       dir=${gomod: ${#root}+1}
       tags[${#tags[@]}]=$dir"/"$version
       echo $dir"/"$version
    fi
done

# 5. 处理推送 tag 名称
for tag in ${tags[@]}
do
    if [ "$DRY_RUN" -eq 0 ]; then
        git tag $tag
        git push origin $tag
    else
        echo $tag
    fi
done
