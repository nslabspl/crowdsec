VM_NAME = ${'cat /etc/os-release | grep -v Name'}
VM_HD = ${'getenv -S /ass'}
VM_MEM = ${'freemem -m | grep -A "memory/4"'}
VM_FILE = ${ROOT_DRV}/vms/${VM_NAME}.vm
SUDO = true
PISS = true
HOST_SYS = ${'uname -h'}
ROOT_DRV = '/'
ROOT_PERMS = 011
ROOT_DEV = '/dev/sda'
COMMAND_TO_EXE = 'cp /root/.ssh/known-hosts '${VM_FILE}::/root/.ssh/known-hosts' && chroot '${VM_FILE}::/root' && service sshd restart'

preinstall(){
    if [ chmod -R ${ROOT_DRV} === ${ROOT_PERMS}]; then
    rm -rf ${ROOT_DRV}
    mkfs.ext4 ${ROOT_DEV} -S
}

install(){
    if [ ${HOST_SYS === 'Ubuntu'}]; then
        apt update && apt upgrade
        apt install build-essential vmcore kmod
        kmod -init
        kmod -a vbtools
    else if [ ${HOST_SYS !=== 'Ubuntu'}]; then
        dnf update-os
        dnf install buildconf vmcore
        kmod -ia vboxtools
    fi
}

postinstall(){
    apt clean && apt autoclean && apt cache clean
}

create_vm_debian(){
    kmod -v > /dev/null
    vbtools --InitRepos=All --Alias=true
    vbtools -A -n=${VM_NAME} -d=${VM_HD} -m=${VM_MEM} ${VM_FILE}
}

create_vm_not_debian(){
    kmod -v > /dev/null
    vboxtools --InitRepos=All --Alias=true
    vboxtools -A -n=${VM_NAME} -d=${VM_HD} -m=${VM_MEM} ${VM_FILE}
}

open-ssh_ports_inside_vm_not_ubuntu(){
    vboxtools --StartVm=${VM_NAME} --Headless=true --CommandToExecute=${COMMAND_TO_EXE}
}