# для обучения
git --version
git config
git config --global user.name Gogolev_vadim
git config --global user.email Gogolev_vadim92@gmail.com
git init
определяем что добавить
git add main.go
git add css/
git add .
git status
всё в ожидании
// удаление  фаилов
git rm --cached ____
git status

// необходим комит
добавить в локальное хранилище

git commit -m""
проверить статус
git status

изменили фаил
git status
git add

// помсотрим что добавлять можно
git add
добавить все фаилы расширения
git add *.html
все фаилы добавить из папки
git add css/*
git add css/**/*.html
git add !css/style.css - все фаилы кроме стайл
git add !index.html
git add !*.html

// вся цепочка изменений
git log
git log --oneline

//Игнор фаилов

make.txt
itproject/
css/
js/
bin/
admin/
*.php

git status

// Отмена действий и возращение к старым действием

проверить проект на момент комита
git checkout идентификатор

чтобы вернуться обратно нам
git checkout master


// команда позволяет отменить комит

git revert
:wq выйти из редактора q
git log --oneline

удалить все комиты до опреденного

git reset ____ до какого-то определённого но фаилы не меняются

git reset --hard - все фаилы меняются


// создание ветки

git branch forum
git checkout ветка
git branch -a показывает вертки

// возвращение на ветку
git checkout master

// создание ветки и переход
git checkout -b admin


// объеденить ветка
git merge forum

удалить вертку
git branch -D _____

//чтобы подключиться к удалённому репозиторию
git remote add origin https://github.com/f1rerabbit/git-hub-tutorial.git


// добавить фаилы github и авторизоваться
git push -u origin main

Посмотреть комиты /commits/master

создание фаила README.md

//как клонировать репозиторий
Можно просто скачать через zip

cd ../
git clone https://github.com/f1rerabbit/git-hub-tutorial.git
cd git-hub-tutorial
