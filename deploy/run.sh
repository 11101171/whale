#!/bin/bash

#function that prints out usage syntax
syntax () {
    echo " "
    echo "./app-xxx.sh [options]"
    echo " "
    echo "  [options] is one of create | backup | rm | read "
    echo "    start             - start concurrent application "
    echo " "
    echo "    stop             - stop concurrent application"
    echo " "
    echo "    restart                 - restart concurren application"
    echo " "
    echo "    bstart               - build and start concurrent application"
    echo " "
    echo "    brestart               - build and restart concurrent application"
    echo " "
    echo " "
}


project='whale'
port=''$2
runmode=''$3
application='app-'${runmode}'-'${port}
cd ./../../ # project path
runpath=`pwd`'/'${project}'/deploy/'${application}

if echo "$port" | grep -q '^[0-9]\+$'; then
    echo "["$1"] application: $application ,port: $port."
else
    echo "$0 {run:start|stop|restart} {port:number} {runmode:dev|test|prod}"
    exit 4
fi

if [ $runmode != "dev" ] && [ $runmode != "test" ] && [ $runmode != "prod" ]; then
    echo "$0 {run:start|stop|restart} {port:number} {runmode:dev|test|prod}"
    exit 4
fi

function build(){
     if [ -d '${runpath}' ];then
        rm -rf ${runpath}
     fi
     mkdir ${runpath}
     touch ${runpath}/run.log
     sed -i '' "s/^http.port.*/http.port = ${port}/g" ./${project}/conf/app.conf
     if [ -f '${project}.tar.gz' ];then 
        rm ${project}.tar.gz
     fi
     GOOS=linux GOARCH=amd64 revel package whale
     # GOOS=darwin GOARCH=amd64 revel package ${project}
     tar -zxf whale.tar.gz -C ${runpath}
     echo ${runpath}" ${1}模式 ${runmode}"
     
}

function start() {
    nohup "${runpath}/${project}" -importPath ${project} -srcPath "${runpath}/src" -runMode prod 2>&1 >> ${runpath}/run.log 2>&1 /dev/null &
    echo $! > ${runpath}/run_pid
    echo ${runpath}" ${1}模式 ${runmode}"
}

function kill() {
    kill -9 `cat ${runpath}/run_pid`
    # killall ${application}
    echo ${runpath}" ${1}"
}

case $1 in
    start)
        start " 服务已重启..."
    ;;
    restart)
        kill " 服务已停止..."
        start " 服务已重启..."
    ;;
    bstart)
        build " 服务已构建..."
        start " 服务已启动..."
    ;;
    stop)
        kill " 服务已停止..."  
    ;;
    brestart)
        kill " 服务已停止..."
        build "服务已构建..."
        start " 服务已重启..."
    ;;
    *)
        echo "$0 {start|stop|restart}"
        exit 4
    ;;
esac
