## Go类型

### 1. Go的类型
1. 基本类型 ： int,float(float32,float64),bool,string(切记，string字符串在Go中是基本类型)
2. 结构类型 ： struct , array , slice , map, channel
3. 只描述类型的行为 : interface

对于结构类型，当没有实际值时，默认是nil

### 类型转换
Go语言中不存在隐式类型转换(Nice...)，即所有的转换都必须显示声明的

### Go的命名规范
> 尽量使用驼峰命名法，不推荐使用小写单词+下划线的形式


### 1.布尔类型bool

var isMan = true

对于bool我们只需要记住以下几点即可：
* 布尔型只有两个值：true , false
* 使用 == 或者 != 进行比较获得bool值（必须类型相同才可以进行比较）
* 格式化输出时，**%t**用来输出布尔型值
* 对于布尔型变量的命名，以is 或者Is开头将提升代码的可读性。

重要的事情说三遍：

**只有类型相同的值才可以比较**

**只有类型相同的值才可以比较**

**只有类型相同的值才可以比较**

PS：如果值的类型是接口(interface)，它们也必须都实现相同的接口。


### 2.Go的数字类型
> Go语言支持整型和浮点型，且原生支持复数。

#### 2.1 整型int

有符号的：
* int8（-128 -> 127）
* int16（-32768 -> 32767）
* int32（-2,147,483,648 -> 2,147,483,647）
* int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）

无符号的：
* uint8（0 -> 255）
* uint16（0 -> 65,535）
* uint32（0 -> 4,294,967,295）
* uint64（0 -> 18,446,744,073,709,551,615）

* uintptr 的长度被设定为足够存放一个指针即可。


#### 2.2 浮点型(float)
> PS：Go语言中没有float类型 ，Just for  **float32** and **float64**

* float32（+- 1e-45 -> +- 3.4 * 1e38）
* float64（+- 5 * 1e-324 -> 107 * 1e308）

----------

In summary：

1. int型是计算最快的一种类型；
2. 整型的初始值为0，浮点型的初始值为0.0；
3. float32精确到小数点后7位，float64精确到小数点后15位；
4. 尽可能地使用float64，因为math包中的所有有关数字运算的函数都会要求这个类型；
5. 8进制表示法：前缀0   ；
6. 16进制表示法：前缀0x ；
7. 使用e来表示10的连乘(1e3 = 1000)
8. Go不允许不同类型之间的混合使用，应尽量避免；

#### 2.3 格式化说明符
在格式化字符串中：
1. %d 用于格式化整数
    1. %x , %X  格式化16进制的数字；
2. %g 用于格式化浮点型
    1. %f 输出浮点数
    2. %e 输出科学计数法
3. %0d 用于规定输出定长的整数，其中开头的数字0是必须的；
4. %n.mg用于表示数字n并精确到小数点后m位,除了使用g之外,还可以使用e或者f

#### 2.4 字符串
> 字符串是一种值类型，且值不可变，即创建某个文本后无法修改这个文本的内容。从本质上来讲，
字符串是**字节的定长数组**

Go支持两种形式的字符串：
1. 解释型字符串(双引号括起来，其相关的转义字符将被替换)：
2. 非解释型字符串(单引号括起来，支持换行)

PS：
1. 字符串的拼接符 +，常用但效率不高；
2. 字符串本质上是定长的字节数组;
3. str[i]拿到的是个字节，如a - 97
4. 获取字符串中的某个字节的地址是非法的

重要的事情说3遍：

* 非解释型字符串使用的是**反引号**

* 非解释型字符串使用的是**反引号**

* 非解释型字符串使用的是**反引号**

### 字符串操作

对于字符串，每种语言都有一些对于字符串的预定义处理函数。

在Go中，使用**strings包**来完成对字符串的主要操作。
常用的字符串操作函数：
1. strings.HasPrefix()
2. strings.HasSuffix()
3. 包含关系：
    * strings.Contains(s,subStr) bool
4. 字符索引(-1表示没有)：
    * strings.Index(s,str string) int
    * strings.LastIndex(s,str string) int
5. 字符串替换：
    * strings.Replace(str,old,new,n) string （n = -1表示全部替换）
6. 统计字符串出现的次数
    * strings.Count(s,str string) int
7. 重复字符串：
    * strings.Repeat(s,count int) string 
8. 字符串大小写切换：
    * strings.ToLower(s) string 小写
    * strings.ToUpper(s) string 大写
9. 修剪字符串：
    * strings.TrimSpace(s) string 剔除字符串开头和结尾的空白符号
    * strings.TrimLeft(s,str string) string 剔除字符串开头的str
    * strings.TrimRight(s,str string) string 剔除字符串结尾的str
    * strings.Trim(s,str string) string 剔除开头和结尾的str
10. 切割字符串
    * strings.Field(s) string 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块
    * strings.Split(s,sep string) string 用于自定义分割符号sep来对指定字符串进行分割
11. 拼接字符串
    * strings.Join(sl []string, sep string) string Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
  
  [更多关于strings的用法](https://golang.org/pkg/strings/)
    
    
> 与字符串相关的类型转换都是通过**strconv包**实现的。

[更多关于strconv包的用法](https://golang.org/pkg/strconv/)





































