## Лекция 11. Добавление REACT клиента
***Задача*** : создать клиент с использованием ```frontend``` фреймворка ```REACT-JS```

### Шаг 1. Определение шаблона
Создадим директорию ```assets``` и добавим туда стандартный шаблон ```index.html```:
```
<html>
    <head lang="en">
        <meta charset="UTF-8">
        <title>Our React client</title>
    </head>
    <body>
        <div id="react"></div>
        <script src="/script.js"></script>
    </body>
</html>
```

### Шаг 2. Инициализация package.json
Теперь инициализируем сторонние зависимости проекта (клиента):
* Выполним команду ```npm init``` (аналог ```maven/gradle``` скрипта)
В ```package.json``` : 
```
{
  "name": "reactjs-client",
  "version": "1.0.0",
  "description": "ReactJs Client",
  "keywords": [
    "react"
  ],
  "author": "Evgen Vlasov",
  "dependencies": {
    "axios": "^0.18.0",
    "lodash": "^4.17.5",
    "react": "^16.2.0",
    "react-dom": "^16.2.0",
    "react-router-dom": "^4.2.2",
    "webpack": "^4.2.0",
    "webpack-cli": "^4.2.0"
  },
  "scripts": {
    "build": "webpack",
    "watch": "webpack --watch -d"
  },
  "devDependencies": {
    "babel-core": "^6.18.2",
    "babel-loader": "^7.1.4",
    "babel-polyfill": "^6.16.0",
    "babel-preset-es2015": "^6.18.0",
    "babel-preset-react": "^6.16.0"
  }
}
```

### Шаг 3. Инициализация сборщика проекта
В качестве сборщика проекта используется ```webpack```. Для его описания создадим файл ```webpack.config.js```:
```

```
