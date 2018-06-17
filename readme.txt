简介：DevOps几乎是和云计算同一时间出现在IT行业，没有云平台的弹性和按需资源分配，DevOps就没有办法落地。
二者的出现彻底将瀑布式开发甚至Agile敏捷扫入历史的废墟，以AWS为首的西方公司已经可以做到每天上千次的迭代和变更，
这些都是技术发展的结果。本课程从一个简单的网站开发部署上线，变更，监控，自恢复等等方面简单介绍利用阿里云平台如何践行简单的DevOps。
1章 课程介绍
先介绍云计算基本概念以及什么是私有云，公有云，什么是Iaas，Paas，Saas等科普知识。
然后介绍DevOps的历史以及DevOps是什么，让大家认识云计算和DevOps

2章 阿里云平台的简单介绍
以阿里云平台为例进一步讲述公有云的基础知识，并且对ECS弹性计算服务和云监控服务作一个概要讲解，帮后面的课程打好基础

3章 网站上云
通过一个完整的项目开发，构建，上线，监控和简单运维过程全面讲述网站上云和DevOps理念


3-1案例和go简介
    案例为什么用go?
        1.go简单,开发效率高 2.全平台可以运行
        几个有用的命令:go install (编译打包)   go fmt(格式化) go test,  go get
3-2本地创建运行项目
        github新建一个空项目然后克隆到本地 goWebForDevOps
        新建目录webserver
        cd webserver
        新建main.go
            http.ListenAndServe绑定监听端口
            http.HandleFunc指定路由及处理函数
        在当前目录执行命令:go install
            那么在go的工作目录的bin目录下就生成了本项目的启动名称webserver,在终端直接输入项目名webserver就可以启动

3-3将web server部署到阿里云
    交叉编辑.go源文件:
        目标是编译可以在我的云服务器上运行的go项目,
        那么首先登录云服务查看linux的位数
            uname -a
            输出:Linux VM_0_2_centos 3.10.0-514.26.2.el7.x86_64 #1 SMP Tue Jul 4 15:04:05 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux
            看到x86_64,是64位的
        然后在本地,进入go项目下的目录webserver,执行编译命令:
        CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

              其中，CGO_ENABLED=0表示不使用cgo，
              GOOS：目标平台的操作系统（darwin、freebsd、linux、windows）
              GOARCH：目标平台的体系架构（386、amd64、arm）
              执行后,项目目录中出现了webserver,就是编译生成的目标文件
    将目标文件也传到github上面
    然后将github上的项目克隆到云服务器
            登录到云服务器,git clone github上的项目
    在云服务器运行编译后的文件
            cd到克隆的go项目
            找到编译的文件webserver
            运行:
                 ./webserver
    然后就可以在网页上输入"http://ip:端口"  进行浏览了
3-4 deployserver开发
    上面步骤比较繁琐
    git pull
    git push->git pull
    deploy
    能不能有自动化的方式,云服务虚拟机自动获取git的更新?
    在项目中新建目录deployserver
    vim main.go
        //重启服务
        func reLaunch() {
        	cmd := exec.Command("sh", "./deploy.sh")
        	err := cmd.Start()
        	if err != nil {
        		log.Fatal(err)
        	}
        	err = cmd.Wait()
        }
        func firstPage(w http.ResponseWriter, r *http.Request) {
        	io.WriteString(w, "<h1> Hello,this is my  deploy server</h1>")
        	reLaunch()
        }

        func main() {
        	http.HandleFunc("/", firstPage)
        	http.ListenAndServe(":5000", nil)
        }
    在项目中新建文件deploy.sh,内容如下
            #! /bin/sh

            kill -9 $(pgrep webserver)
            cd /data/gowork/goWebForDevOps/
            git pull https://github.com/CaryLy/goWebForDevOps.git
            cd webserver/
            ./webserver &
    同理上传到github,然后在云服务虚拟机拉下来,执行deployserver,就可以浏览器访问5000端口了

3.5完成自动化

    在github的当前项目goWebForDevOps,点击settings,点击左边菜单有webhooks,点击Add webhook
    Payload URL:http://ip:5000/
    which event:  选择Just the push event.
        那么当有push的时候,就会调用url http://ip:5000/,这个页面会执行脚本deploy.sh自动的拉取github上的代码



