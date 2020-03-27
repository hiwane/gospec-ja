#!/bin/sh


RET=0
# 表記ユレチェック
for kwd in '　' 'ポインタ[^ー]' 'シグネチャ[^ー]' 'パラメータ[^ー]' '2 *項' 'インタフェース'
do
	if grep -H -n "${kwd}" README.md; then
		:
	else
		RET=1
	fi
done

# リンクチェック


exit ${RET}