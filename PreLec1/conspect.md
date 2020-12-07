## Преподготовка и настройка рабочей области

Хотим натсроить рабочую область для удобного напиания кода на данном курсе. Для этого нам понадобится:
* Стандартизированный путь до проектного расположения
* На курсе нам понадобятся ```git```, ```docker```
* Для эффективной работы с утилитами понадобятся аккаунты на ```github.com``` и ```docker-hub.com```
* Уметь пользоваться ```makefile```'ми
* Установочный ```bash```- скрипт

## Шаг 1. Создание и описание ***bash*** скрипта
Создадим ```bash``` скрипт с названием ```installation.sh```. Добавим следующий заголовок:
```
#!/bin/bash
```
***Перед описание начинки файла*** добавим скрипту возможность как выполнению:
```chmod +x installation.sh```

А теперь опишем все инстуркции, какие нам нужны:
* ***Тестовая инструкция*** : обновим конфигурацию (Выполним файл , чтобы проверить, что все работает ```./installation.sh```)
* Установка ```git```, ```code```, ```make``` + обновление конфигов после загрузки
```
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
```
* Выполним данный скрипт ```./installation.sh```

* Теперь установим ```docker```
```
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


```

* Выполним ```./installation.sh``` повторно