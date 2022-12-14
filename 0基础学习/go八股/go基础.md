# 编码

## Unicode和UTF-8区别

Unicode是字符集而UTF-8是一种编解码规则。字符集为每一个字符分配一个独一无二的ID，而编码规则就是将字符集转换为字节序列。

## Unicode,ASCII,UTF-8的区别

### ASCII编码

* ASCII 码使用指定的7 位或8 位二进制数组合来表示128 或256 种可能的字符。标准ASCII 码也叫基础ASCII码，使用7 位二进制数（剩下的1位二进制为0）来表示所有的大写和小写字母，数字0 到9、标点符号， 以及在美式英语中使用的特殊控制字符。其中最后一位用于奇偶校验。
* 问题：ASCII是单字节编码，无法用来表示中文（中文编码至少需要2个字节），所以，中国制定了GB2312编码，用来把中文编进去。但世界上有许多不同的语言，所以需要一种统一的编码。

### Unicode

* Unicode把所有语言都统一到一套编码里，这样就不会再有乱码问题了。
* Unicode最常用的是用两个字节表示一个字符（如果要用到非常偏僻的字符，就需要4个字节）。现代操作系统和大多数编程语言都直接支持Unicode。

### Unicode和ASCII的区别

* ASCII编码是1个字节，而Unicode编码通常是2个字节。
  字母A用ASCII编码是十进制的65，二进制的01000001；而在Unicode中，只需要在前面补0，即为：00000000 01000001。
* 新的问题：如果统一成Unicode编码，乱码问题从此消失了。但是，如果你写的文本基本上全部是英文的话，用Unicode编码比ASCII编码需要多一倍的存储空间，在存储和传输上就十分不划算。

### UTF8

* 所以，本着节约的精神，又出现了把Unicode编码转化为“可变长编码”的UTF-8编码。
* UTF-8编码把一个Unicode字符根据不同的数字大小编码成1-6个字节，常用的英文字母被编码成1个字节，汉字通常是3个字节，只有很生僻的字符才会被编码成4-6个字节。如果你要传输的文本包含大量英文字符，用UTF-8编码就能节省空间。
* ![img](https://ask.qcloudimg.com/http-save/1692602/377ilfsviw.png?imageView2/2/w/1620)
* 从上面的表格还可以发现，UTF-8编码有一个额外的好处，就是ASCII编码实际上可以被看成是UTF-8编码的一部分，所以，大量只支持ASCII编码的历史遗留软件可以在UTF-8编码下继续工作。

### 计算机中通用的字符编码的工作方式

* 在计算机内存中，统一使用Unicode编码，当需要保存到硬盘或者需要传输的时候，就转换为UTF-8编码。
* 用记事本编辑的时候，从文件读取的UTF-8字符被转换为Unicode字符到内存里，编辑完成后，保存的时候再把Unicode转换为UTF-8保存到文件：

  ![](https://ask.qcloudimg.com/http-save/1692602/zs2m48pymr.png?imageView2/2/w/1620)
* 浏览网页的时候，[服务器](https://cloud.tencent.com/product/cvm?from=10680)会把动态生成的Unicode内容转换为UTF-8再传输到浏览器：

  ![](https://ask.qcloudimg.com/http-save/1692602/zxks1cbo4v.png?imageView2/2/w/1620)

类型别名

# 切片

## 空切片和nil切片的区别

首先，slice的底层结构体中是由3个字段构成的：**长度、容量和指向底层数组的指针**字段。

* 空切片：**如果切片的长度是0，那么称该切片是空切片** 。其容量由指向的底层数组决定。
* nil切片：除了长度和容量都是0之外，还有就是 **ptr指针不指向任何底层数组，ptr指针为nil，这也是和空切片的本质区别** 。
* 在对切片进行json.Marshal编码的时候，nil切片会被编码成null，而空切片会被编码成空数组:[]。如下代码所示：

```go
type Person {
  Friends []string
}

var f1 []string //nil切片
json1, _ := json.Marshal(Person{Friends: f1})
fmt.Printf("%s\n", json1) //output：{"Friends": null}


f2 := make([]string, 0) //non-nil空切片
json2, _ := json.Marshal(Person{Friends: f2})
fmt.Printf("%s\n", json2) //output: {"Friends": []}
```

## 切片范围

1. **单个元素访问时，下标范围[0,len()-1]**
2. **切片访问时，下标范围[0,len()]**
