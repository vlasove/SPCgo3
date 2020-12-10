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
var path = require('path');
module.exports = 
{
    resolve:
    {
        extensions: ['.js', '.jsx']
    },
    mode: 'development',
    entry: '/.app/main.js',
    cache: true,
    output:
    {
        path: __dirname,
        filename: './assets/script.js'
    },
    module:
    {
        rules:
        [
            {
                test: path.join(__dirname, '.'),
                exclude: /(node_modules)/,
                loader : 'babel-loader',
                query:
                {
                    cacheDirectory: true,
                    presets: ['es2015', 'react']
                }
            }
        ]
    }
};
```

### Шаг 4. Сборка ReactApp и необходимые компоненты
* Создадим файлы ```app/components/react-app.jsx```
* ```app/components/employee-app.jxs```
* ```app/components/employee.jsx```
* ```app/components/employee-list.jsx```
* ```app/components/add-employee.jsx```

### Шаг 5. Сборка билда
* npm install
* npm run build 

### Шаг 6. Запуск собранного проекта
* ```go run server.go```
