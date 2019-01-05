# Postgres BiDirectional Replication 
Postgres 9.4 
BDR 1.0.7 - https://wiki.postgresql.org/wiki/BDR_User_Guide
## Requirements

VirtualBox - https://www.virtualbox.org/wiki/Downloads

Vagrant - https://www.vagrantup.com/downloads.html

Ansible - https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-the-control-machine

## Create VM
Checkout this repository, and run vagrant up

```
git clone git@github.com:ekryukov/pg_2active_replica.git pgvm
cd pgvm
vagrant up
```

## Run Ansible playbook
```
cd pgvm/provisioning
# edit invetory file - setup hosts
ansible-playbook -i inventory playbook.yml
```
