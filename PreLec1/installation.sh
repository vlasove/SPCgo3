#!/bin/bash

# Первый шаг. Обновим конфиги
sudo apt-get update

# Установка git
sudo apt install git 

# Установка vscode (опционально)
sudo snap install --classic code 

# Установка утилиты make
sudo apt-get install make 

# Обновим конфиг перед установкой docker
sudo apt-get update

# Установка docker
# 1. Удалим старые зависимости
sudo apt-get remove docker docker-engine docker.io containerd runc


# 2. Установка зависимостей и подготовка репозитория
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

# 3. Сверка GPG
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# 4. Fingerprint
sudo apt-key fingerprint 0EBFCD88

# 5. Stable from 
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"


# 6. Update
sudo apt-get update
apt-cache policy docker-ce

# 7. Установка высокоуровневого интерфейса
sudo apt-get install docker-ce docker-ce-cli containerd.io

# 8. Проверим, что все работает
sudo docker run hello-world

