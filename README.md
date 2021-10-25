## 1 说明

jenkinstools是一个命令行小工具，目前可实现同一个jenkins下的功能有：

1. 拷贝一个Folder下的所有job至另一个Folder（Folder需要事先存在），参数跟上替换字符串可自动替换job参数
2. 创建Folder，目前支持直接创建Folder、指定parentFolder下创建子Folder、指定View下的ParentFolder下创建子Folder（待完善的功能项：指定View下的Folder）
3. 创建View，目前只支持两种类型的View创建：ListView和MyView；其它类型的View也可以实现，没写

## 2 用法

以Windows下举例演示，Linux类似。
修改config.ini文件，选择要连接的jenkins url和账号密码

    [jenkins]
    JenkinsUrl="http://jenkins.rsq.local"
    UserName="rensongqi"
    Password="123456"

### 2.1 拷贝jobs from folder(-c)

    jenkinstools.exe -c rsq_test1 rsq_test2 test1 test2

参数说明：
    
    rsq_test1: 代表src Folder
    rsq_test2: 代表dst Folder
    
    test1: 哪个字符串需要替换
    test2: 替换为哪个字符串

### 2.2 创建目录folder(-f)

1. 指定在RSQ的View下的rsq_tes1目录下创建rsq_folder目录
 
    jenkinstools.exe -f rsq_folder RSQ rsq_test1
   
2. 指定在rsq_test1的Folder下创建rsq_test1 folder

    jenkinstools.exe -f rsq_folder "" rsq_test1
    
3. 直接在ALL View视图下创建Folder(只需要把不需要的参数为空即可)

    jenkinstools.exe  -f rsq_folder "" ""

### 2.3 创建View(-v)

目前只支持两种ViewType：ListView和MyView

1. 指定ViewType创建View(创建名为RSQ的类型为MyView的View)

    jenkinstools.exe -v RSQ MyView
    
2. 如果不指定ViewType，默认会创建ListView的View

    jenkinstools.exe -v RSQ ""
    
    
## 3 Attention

最好标准化项目的命名规范，这样copy jobs的功能可以最大化实现。

**新项目的命名最好都用小写**，这样替换字符串的时候可以全部替换，不用担心大小写的问题。


