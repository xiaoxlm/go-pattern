# 装饰器模式
在不破环原来对象的情况下，扩展对象的信息

## 用例
如测试代码中的例子，之前我们的People对象只需要年龄和名字. 但现在我们需要在People有描述， 在不破坏NewPeople方法的情况下， 我们又能对NewPeople进行扩展， 于是就装饰后的NewPeople方法**NewPeopleFNCWithDesc**