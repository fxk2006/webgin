#!/bin/bash
# -*- coding:utf-8 -*-
#===============================================================
#		AUTHOR: YangYang root@opdays.com
#		CREATER: 2017-11-02 18:37:24
#		FILENAME: 1.sh
#		DESCRIPTION: 
#===============================================================

echo "任务开始"
    for x in $(seq 15);do
    echo -e "\033[31m${x}\033[0m"
    sleep 1
    done
echo "任务结束"
#tail -f log.log

