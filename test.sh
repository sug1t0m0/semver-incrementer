#!/bin/bash

set -e

input_version="${1}"
input_method="${2}"


# Goアプリケーションの実行と出力を変数に保存する
command_output=$(./main -v="$input_version" -m="$input_method")

# 出力を行ごとに処理する
while IFS='=' read -r key value; do
    export "$key"="$value"
done <<< "$command_output"

# 確認
echo "Value of aaa: $version"
echo "Value of bbb: $method"