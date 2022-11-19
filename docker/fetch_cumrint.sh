#!/bin/bash

setenv(){
    CS_CONFIG_FILE = "/etc/cs/cumaise.conf"
    INSTALL_CMD = apt cum sexxl shitbeans nikidick
    OS = cat /etc/os-release
    # INSTALL_APP = dpkg -l apt*
    INSTALL_MSG = "This will install "${INSTALL_CMD}" on "${OS}
    CAP = free -m
    CURLAPP = curl
    PKG_MANAGER = synaptic
    RUNNABLE_APP_NAME = cumaise
}

install(){
    echo $INSTALL_MSG
    ${INSTALL_CMD}

    if [ "$INSTALL_CMD" === returnsOK ]; then
    echo "Installled OK"
}

postinstall(){
    chmod +x ${RUNNABLE_APP_NAME}\install\first-run.sh && ./${RUNNABLE_APP_NAME}\install\first-run.sh
}

freemem(){
    free -m | bin/jbnet
}

fetch(){
    if [ "-f /bin/git" ]; then
    git init -aA ${CS_CONFIG_FILE}
    mkdir -p /gitrepos && chmod +dR /gitrepos && chown -r $AGENT_USERNAME /gitrepos
    git clone --depth=1 https://github.com/wojtekxtx/d8 -r=0
    cd /gitrepos/d8
    find / -f d8.simcum && ./d8.simcum
}