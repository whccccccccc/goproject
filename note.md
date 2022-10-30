# 变量与常量
## 声明方式
### 变量声明方式
``` golang
var name type = value
var name =value
name := value

```
### 常量声明方式

``` golang
const name type = value
const (
	c1 = 2
	c2 //=2
	c3//=2
	c4 = 6
	c5//=6
)
```
## 命名规则
1. 大/小写字母开头 控制是否可包外访问
## 可多重赋值
``` golang
var  a,b int =10,5
a,b=b,a// 交换ab值
```

## 匿名变量
> 例如函数返回两个返回值  可以用 _ 符号排除
```golang

getUserInfo()(name string,age int){
    return "tom",18
}
_,age=getUserInfo()//只获取age
```
# 数据类型
## 基本数据类型
1. 布尔类型：bool
2. 整型：int8、byte、int16、int、uint、uintptr 等
3. 浮点类型：float32、float64
4. 复数类型：complex64、complex128
5. 字符串：string
6. 字符类型：rune
7. 错误类型：error

## 复合数据类型
1. 指针（pointer）
2. 数组（array）
3. 切片（slice）
4. 字典（map）
5. 通道（chan）
6. 结构体（struct）
7. 接口（interface）

## 布尔类型
> 只能是true 或false 不可以 是 0~1的整型转换

## 整型及运算符

### 整形
> 注: int和int64是两种不同类型  ，编译器也不会帮你自动做类型转换  可以 类型(变量) 做强制转换
> 和其他编程语言一样，可以通过增加前缀 0 来表示八进制数（如：077），增加前缀 0x 来表示十六进制数（如：0xFF），以及使用 E 来表示 10 的连乘（如：1E3 = 1000）。
#### 有符号位
1. int8
2. int16
3. int32
4. int64
#### 无符号位
1. uint8
2. uint16
3. uint32
4. uint64

### 运算符

#### 算术运算符
> + - * / %  (%只能用于整数)
> 不同类型的整型值不能直接进行算术运算
> 在 Go 语言中，也支持自增/自减运算符，即 ++/--，但是只能作为语句，不能作为表达式，且只能用作后缀，不能放到变量前面：
> 支持 +=、-=、*=、/=、%= 这种快捷写法

#### 比较运算符
> Go语言支持一下集中常见的比较运算符: > < == >= <= 和!= 比较运算符运行的结果是布尔值
> 各种类型的整形变量都可以直接与字面常量进行比较 intValue==8
#### 逻辑运算符
> 逻辑运算符计算的结果也是布尔值，通常我们可以组合使用逻辑运算符和比较运算符：
| 运算符 | 含义                | 结果                                                   |
| :----- | :------------------ | :----------------------------------------------------- |
| x && y | 逻辑与运算符(AND)   | 如果 x 和 y 都是 true，则结果为 true，否则结果为 false |
| x      |                     | y                                                      | 逻辑或运算符（OR） | 如果 x 或 y 是 true，则结果为 true，否则结果为 false |
| !x     | 逻辑非运算符（NOT） | 如果 x 为 true，则结果为 false，否则结果为 true        |
#### 位运算符
> 位运算符以二进制的方式对数值进行运算，效率更高，性能更好，Go 语言支持以下这几种位运算符

| 运算符 | 含义     | 结果                                           |
| :----- | :------- | :--------------------------------------------- |
| x & y  | 按位与   | 把 x 和 y 都为 1 的位设为 1                    |
| x      | y        | 按位或                                         | 把 x 或 y 为 1 的位设为 1 |
| x ^ y  | 按位异或 | 把 x 和 y 一个为 1 一个为 0 的位设为 1         |
| ^x     | 按位取反 | 把 x 中为 0 的位设为 1，为 1 的位设为 0        |
| x << y | 左移     | 把 x 中的位向左移动 y 次，每次移动相当于乘以 2 |
| x >> y | 右移     | 把 x 中的位向右移动 y 次，每次移动相当于除以 2 |

#### 运算符优先级
> 上面介绍的 Go 语言运算符优先级如下所示（由上到下表示优先级从高到低，或者数字越大，优先级越高）：
> ++ 或 -- 只能出现在语句中，不能用于表达式，故不参与优先级判断。
1. 6      ^（按位取反） !
2. 5      *  /  %  <<  >>  &  &^
3. 4      +  -  |  ^（按位异或）
4. 3      ==  !=  <  <=  >  >=
5. 2      &&
6. 1      ||

## 浮点型与复数类型

### 浮点型
> Go 语言中的浮点数采用IEEE-754 标准的表达方式，定义了两个类型：float32 和 float64，其中 float32 是单精度浮点数，可以精确到小数点后 7 位（类似 PHP、Java 等语言的 float 类型），float64 是双精度浮点数，可以精确到小数点后 15 位（类似 PHP、Java 等语言的 double 类型）。
> 浮点数支持通过算术运算符进行四则运算，也支持通过比较运算符进行比较（前提是运算符两边的操作数类型一致），但是涉及到相等的比较除外，因为我们上面提到，看起来相等的两个十进制浮点数，在底层转化为二进制时会丢失精度
> 如果一定要判断相等，下面是一种替代的解决方案：
> 通过一个可以接受的最小误差值 p，约定如果两个浮点数的差值在此精度的误差范围之内，则判定这两个浮点数相等

### 复数类型
> 除了整型和浮点型之外，Go 语言还支持复数类型，与复数相对，我们可以把整型和浮点型这种日常比较常见的数字称为实数，复数是实数的延伸，可以通过两个实数（在计算机中用浮点数表示）构成，一个表示实部（real），一个表示虚部（imag），常见的表达形式如下：
> z = a + bi
> 其中 a、b 均为实数，i 称为虚数单位，当 b = 0 时，z 就是常见的实数，当 a = 0 而 b ≠ 0 时，将 z 称之为纯虚数，如果你理解数学概念中的复数概念，这些都很好理解，下面我们来看下复数在 Go 语言中的表示和使用。
> 在 Go 语言中，复数支持两种类型：complex64（32 位实部和虚部） 和 complex128（64 位实部与虚部），对应的表示示例如下，和数学概念中的复数表示形式一致：

## 字符串及底层字符类型
### 字符串
> 在 Go 语言中，字符串是一种基本类型，默认是通过 UTF-8 编码的字符序列，当字符为 ASCII 码时则占用 1 个字节，其它字符根据需要占用 2-4 个字节，比如中文编码通常需要 3 个字节。

#### 声明和初始化
``` golang
var str string         // 声明字符串变量
str = "Hello World"    // 变量初始化
str2 := "你好，学院君"   // 也可以同时进行声明和初始化
```
#### 格式化输出
> 还可以通过 Go 语言内置的 len() 函数获取指定字符串的长度，以及通过 fmt 包提供的 Printf 进行字符串格式化输出：
``` golang
fmt.Printf("The length of \"%s\" is %d \n", str, len(str)) 
fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)
```
#### 转义字符
Go 语言的字符串不支持单引号，只能通过双引号定义字符串字面值，如果要对特定字符进行转义，可以通过 \ 实现，就像我们上面在字符串中转义双引号和换行符那样，常见的需要转义的字符如下所示：
1. \n ：换行符
2. \r ：回车符
3. \t ：tab 键
4. \u 或 \U ：Unicode 字符
5. \\\ ：反斜杠自身
> 也可以使用 `` 的方式

#### 多行字符串
> 对于多行字符串，也可以通过 ` 构建：  如果使用+号链接的话 +号要留在上一行

#### 不可变值类型
> 虽然可以通过数组下标方式访问字符串中的字符但是和数组不同，在 Go 语言中，字符串是一种不可变值类型，一旦初始化之后，它的内容不能被修改，比如看下面这个例子：
``` golang
str := "Hello world"
str[0] = 'X' // 编译错误
```
### 字符串操作
#### 字符串连接
> Go 内置提供了丰富的字符串函数，常见的操作包含连接、获取长度和指定字符，获取长度和指定字符前面已经介绍过，字符串连接只需要通过 + 连接符即可
> 如果字符串长度较长，需要换行，则 + 连接符必须出现在上一行的末尾，否则会报错：
#### 字符串切片
> 在 Go 语言中，可以通过字符串切片实现获取子串的功能：
> Go 切片区间可以对比数学中的区间概念来理解，它是一个左闭右开的区间，比如上述 str[0:5] 对应到字符串元素的区间是 [0,5)，str[:5] 对应的区间是 [0,5)（数组索引从 0 开始），str[7:] 对应的区间是 [7:len(str)]（这是闭区间，是个例外，因为没有指定区间结尾）
> :index 代表index 0~index 不含index
> index:代表 index~结束 包含index
> 包头不包尾
> str[:] 会打印出完整的字符串来。
>  Go 字符串也支持字符串比较、是否包含指定字符/子串、获取指定子串索引位置、字符串替换、大小写转换、trim 等操作，更多操作 API 参考标准库 strings 包，

#### 字符串遍历
> Go 语言支持两种方式遍历字符串。
1. 以字节数组的方式遍历
2. 以 Unicode 字符遍历

### 底层字符类型
> Go 语言对字符串中的单个字符进行了单独的类型支持，在 Go 语言中支持两种字符类型
> 1. byte，代表 UTF-8 编码中单个字节的值（它也是 uint8 类型的别名，两者是等价的，因为正好占据 1 个字节的内存空间
> 2. rune，代表单个 Unicode 字符（它也是 int32 类型的别名，因为正好占据 4 个字节的内存空间。关于 rune 相关的操作，可查阅 Go 标准库的 unicode 包）。

#### UTF-8 和 Unicode 的区别
> Unicode 是一种字符集，囊括了目前世界上所有语言的所有字符，与之类似的术语还有 ASCII 字符集（仅包含 256 个字符）、ISO 8859-1 字符集等（包含所有西方拉丁字母），广义的 Unicode 既包含了字符集，也包含了编码规则，比如 UTF-8、UTF-16、UTF8MB4、GBK 等。
> 因此 UTF-8 是 Unicode 字符集的实现方式之一，它会将 Unicode 字符以某种方式进行编码。在具体实现时，UTF-8 是一种变长的编码规则，从 1~4 个字节不等，比如英文字符是 1 个字节，中文字符是 3 个字节。通过 UTF-8 编码的 Unicode 字符以最大长度 4 个字节作为单个字符固定占据的内存空间，在 Go 语言中可以通过 unicode/utf8 包进行 UTF-8 和 Unicode 之间的转换。
> 所以如果从 Unicode 字符集的视角看，字符串的每个字符都是一个字符的独立单元，但如果从 UTF-8 编码的视角看，一个字符可能是由多个字节编码而来的。
> 我们通过 len 函数获取到的是字符串的字节长度，再据此通过字符数组的方式遍历字符串时，是以 UTF-8 编码的角度切入的；而当我们通过 range 关键字遍历字符串时，又是从 Unicode 字符集的角度切入的，如此一来就得到了不同的结果。
> 出于简化语言的考虑，Go 语言的多数 API 都假设字符串为 UTF-8 编码。

#### 将 Unicode 编码转化为可打印字符
> 如果你想要将 Unicode 字符编码转化为对应的字符，可以使用 string 函数进行转化

## 基本数据类型之间的转化
>  Go 语言中的基本数据类型，分别是布尔类型、整型、浮点型、复数类型、字符串和字符类型
### 数值类型之间的转化
> 像这样直接用要转换的类型包裹即可
> 需要注意在有符号与无符号以及高位数字向低位数字转化时，需要注意数字的溢出和截断。
``` golang
uint()
int64()
```
#### 整型之间的转化


##### 整型与浮点型之间的转化
> 浮点型转化为整型时，小数位被丢弃：
> 整型转化为浮点型时，直接调用对应的函数即可：

##### 数值和布尔类型之间的转化
>  Go 语言不支持将数值类型转化为布尔型，你需要自己根据需求去实现类似的转化。

### 字符串和其他基本类型之间的转化
#### 将整型转化为字符串
> 整型数据可以通过 Unicode 字符集转化为对应的 UTF-8 编码的字符串
> 还可以将 byte 数组或者 rune 数组转化为字符串，因为字符串底层就是通过这两个基本字符类型构建的

#### strconv 包
> Go 语言默认不支持将字符串类型强制转化为数值类型，即使字符串中包含数字也不行。
> 如果要实现更强大的基本数据类型与字符串之间的转化，可以使用 Go 官方 strconv 包提供的函数

##  数组使用入门及其不足

### 数组的声明和初始化
> 数组是所有语言编程中最常用的数据结构之一，Go 语言也不例外，与 PHP、JavaScript 等弱类型动态语言不同，在 Go 语言中，数组是固定长度的、同一类型的数据集合。数组中包含的每个数据项被称为数组元素，一个数组包含的元素个数被称为数组的长度。 和java一样
> 在 Go 语言中，你可以通过 [] 来标识数组类型，但需要指定长度和元素类型。以下是一些常见的数组声明方法：
``` golang
var a [8]byte // 长度为8的数组，每个元素为一个字节
var b [3][3]int // 二维数组（9宫格）
var c [3][3][3]float64 // 三维数组（立体的9宫格）
var d = [3]int{1, 2, 3}  // 声明时初始化
var e = new([3]string)   // 通过 new 初始化
//数组的格式定义如下所示:
//[capacity]data_type{element_values}
//还可以通过这种语法糖省略数组长度的声明：
a := [...]int{1, 2, 3}
//数组在初始化的时候，如果没有填满，则空位会通过对应的元素类型零值填充：
//我们还可以初始化指定下标位置的元素值，未设置的位置也会以对应元素类型的零值填充：
a := [5]int{1: 3, 3: 7}
//[0 3 0 7 0]
```
## 切片使用入门与数据共享问题处理
> 数组的一个特点：数组的长度在定义之后无法修改，数组长度是数组类型本身的一部分，是数组的一个内置常量，因此我们无法在数组上做动态的元素增删操作。
> Go语言提供了切片（slice）来弥补数组的不足，切片一个最强大的功能就是支持对元素做动态增删操作

### 切片的定义
> 在 Go 语言中，切片是一个新的数据类型，与数组最大的不同在于，切片的类型字面量中只有元素的类型，没有长度：
> 因此它是一个可变长度的、同一类型元素集合，切片的长度可以随着元素数量的增长而增长（但不会随着元素数量的减少而减少），不过切片从底层管理上来看依然使用数组来管理元素，可以看作是对数组做了一层简单的封装。基于数组，切片添加了一系列管理功能，可以随时动态扩充存储空间  
> 类似于Java 集合中的 List

### 创建切片
> 创建切片的方法主要有三种 —— 基于数组、切片和直接创建
#### 基于数组