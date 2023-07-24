![2](https://github.com/liyatai/Vue3-Gin--rememberWord/assets/102398022/7fdc7a80-4603-473f-a8cd-79cb49b79363)# Vue3-Gin--rememberWord
vue3+ts+vite+gin开发的记单词app
## 

## Word-vue3-gin词汇软件使用文档

![2ac79e23-831a-4eb6-880f-ff74ebd7e4ea](https://github.com/liyatai/Vue3-Gin--rememberWord/assets/102398022/782f14a1-47c3-4dc9-898f-2c1c64d64cc9)

+ client文件夹是前端源代码文件夹
+ server文件夹是后端源代码文件夹
+ package文件夹是编译后的可执行文件的文件夹

## 可执行文件package目录

```bash
|--- dist
|    |--- index.html
|    └--- assets
|--- json
|--- client.exe
|--- server.exe
|--- 结束.exe
└--- 开始.exe
```

+ 其中dist存放的是静态文件（html、js和css）。

+ json存放的是json格式文件，这些文件用来存放一些数据。

+ client.exe和server.exe是golang的可执行文件，单独打开可以运行前端和后端项目

+ 开始.exe是用来启动项目，一键启动前后端项目。

+ 结束.exe用来结束项目，结束所有进程。
  
  

## 使用教程

### 启动项目

&emsp;&emsp;单击`开始.exe`，会启动前后端项目并且打开指定的前端页面。

![e074e5c5-bd8c-496b-95bd-4598e620a475](https://github.com/liyatai/Vue3-Gin--rememberWord/assets/102398022/e62c4d46-406b-40a1-b10c-7ca3f88f4940)


### 使用记单词功能

&emsp;&emsp;开启项目后会自动进入记单词的模块，在记单词模块，可以选择不同的单词组（可以自己管理单词组），选择单词组以后的功能类似扇贝记单词，会出现英文，如果认识就删除这个单词，如果不认识，查看中文解释再单击切换。同时还有读音功能。

### 管理词汇

#### 管理词汇表

&emsp;&emsp;可以**新建**、**删除**、**查询**和**管理**词汇表。可以根据自己的习惯取表名，目前不支持在前端页面修改表名（忘了考虑），可以在json文件夹下修改表名，单击**管理**词汇表可以对词汇表的单词进行管理。

![1](https://github.com/liyatai/Vue3-Gin--rememberWord/assets/102398022/48805fe2-475c-4ddd-a7e3-00dd8a166d72)



#### 管理表中的词汇

&emsp;&emsp;可以对表中的词汇进行**增删改查**的操作。


## 使用场景

&emsp;&emsp;适合自己搭配词汇书学习，把不会的单词录入到系统中，或者将自己的做过卷子中的生词录入到系统中，方便记忆，不论是考研还是四六级考试，攻克自己的生词将会大大提高做题的能力！
