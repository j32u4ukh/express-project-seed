# express-project-seed
初始化 Express 所需專案與資源，快速開啟一個新的 Express 專案

## 使用 Seed 的初始化流程

### 1. 下載專案

```
git clone https://github.com/j32u4ukh/express-project-seed.git
```

### 2. 修改專案資訊

#### 1. 資料夾名稱
#### 2. package.json

將下圖紅框內的資訊，修改成自己專案的資訊

![package.json](/repo/images/package_json.png)

#### 3. package-lock.json

將下圖紅框內的資訊，修改成自己專案的資訊

![package-lock.json](/repo/images/package_lock_json.png)

### 3. 安裝依賴套件

```
npm install
```

### 4. 移除原始 git 檔案

移除下列檔案

* .git
* LICENSE
* README.md (當前檔案，也可不刪除，修改內容即可)
* repo 資料夾

---

## 原始專案初始化流程

### 1. npm init

建立 package.json 檔

```
npm init -y
```

#### 1.1 安裝 nodemon：非必須，但方便開發

```
npm install nodemon
```

透過 npm install 之後，資料夾中還多了一支 package-lock.json 的檔案。這支檔案會詳細記錄每一次我們使用 npm 安裝的檔案，主要是讓 npm 在執行時參考用的，多數時候你並不需要留意這支檔案。

### 2. 修改 scripts

修改 package.json，分別設定關鍵字 `start` 和 `dev` 對應的指令

```
"scripts": {  
    "start": "node index.js",
    "dev": "nodemon index.js",
    "test": "echo \"Error: no test specified\" && exit 1"
  },
```

使用 `$ npm run dev`，將執行 `dev` 對應的指令。

### 3. 安裝 Express

npm install: 安裝所需套件

```
npm i express
```

#### 3.1 安裝 express-handlebars 作為模板引擎

```
npm install express-handlebars
```

### 4. 載入 booststrap 資源

TODO: 補上檔案下載網址

> /css/bootstrap.css

> /css/bootstrap.css.map

> /js/bootstrap.js

> /js/bootstrap.js.map


#### 使用 `<link>` 導入的方式

下方示範引入 Bootstrap 第 5 版

在 `<head></head>` 中引入下方內容

```html
<link
    href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css"
    rel="stylesheet"
    integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN"
    crossorigin="anonymous"
/>
```

在 `<body></body>` 最末端引入下方內容，在 Bootstrap Bundle 的 JS 裡，包含了 Popper.js、和 Bootstrap.js ，因此不需要另外載入 Popper.js

```html
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js"
        integrity="sha384-C6RzsynM9kWDrMNeT87bh95OGNyZPhcTNXj1NW7RuBCsyN/o0jlpcV8Qyq46cDfL"
        crossorigin="anonymous"></script>
```


### 5. 載入 JS 套件

#### 導入 axios 套件

在 `<body></body>` 最末端引入下方內容

```html
<script src="https://cdn.jsdelivr.net/npm/axios@1.1.2/dist/axios.min.js"></script>
```

### 6. 建立 UI 模板

> /views/index.hbs

> /views/layouts/main.hbs

當中會載入通用的 CSS 和 JS 腳本