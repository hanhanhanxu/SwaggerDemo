
1、go get -u github.com/swaggo/swag/cmd/swag

2、main()添加注解，对外http接口添加注解

3、当前目录执行swag init命令,生成docs文件夹及文件(每次添加swagger相关注释注解时，都要执行一下)

4、main()所在文件import导入docs包：_ "SwaggerDemo/docs"

5、执行，访问 http://localhost:8000/swagger/index.html


总结：非常不友好，还得自己写许多注释，和java的一个注解搞定简直没法比，不是一个东西都。另外一直不明白为什么go里面的注释还能当成代码呢？（注释也有用处，也要给机器看，注释就是注释，代码就是代码，注释为什么要有代码意义呢？这样的注释还算注释吗）