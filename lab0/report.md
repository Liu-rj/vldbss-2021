# Lab0 实验报告

## 实验结果

### 1. 完成 Map-Reduce 框架

`make test_example` 的实验截图：

![image-20211231142303574](E:\vldbss-2021\lab0\image-20211231142303574.png)

![image-20211231142359281](E:\vldbss-2021\lab0\image-20211231142359281.png)

### 2. 基于 Map-Reduce 框架编写 Map-Reduce 函数

`make test_homework` 的实验截图：

![](E:\vldbss-2021\lab0\image-20211231225827759.png)

![](E:\vldbss-2021\lab0\image-20211231225845543.png)

## 实验总结

实验过程中遇到的困难：

* 学习go并看懂提供的代码框架，这部分主要难点在于将框架的具体实现和Map-Reduce的理论框架结合起来，理解每段代码的含义和作用。
* 进行Map和Reduce函数优化的部分，首先要对Map-Reduce的原理有一个比较清楚认识，才能去发现原来函数中可以可以改进的地方，在弄懂原理之后，其实优化也就很直接简单了，原本的两层Map-Reduce其实可以直接删成一层，每一个Map Worker统计自己处理的Map File中url的出现次数，然后汇总给一个Reduce Worker进行统计和求Top 10就完成了整个过程。

对Map-Reduce计算框架的理解：

* Map-Reduce是Google提出的“三驾马车”之一，用于分布式系统的框架的大数据集并行计算。
* Map-Reduce的核心思想其实就是分而治之，将一个大问题拆成很多子问题，每个子问题独立的并行执行，最后再将结果汇总在一起得到最后的答案。
* Map-Reduce的具体实现思想是并行计算（本次lab使用到了go中的协程coroutine--轻量级线程），通过将不同任务动态规划给大规模计算集群中的计算单元，每个计算单元互不干扰的进行计算来实现整体的并行计算。
* Map-Reduce的本质是键值对的转换，Map-Reduce中一个很重要的概念叫做key-value pair，即我们所说的键值对，Map-Reduce通过Map和Reduce操作将键值对域进行了变换，最后依据键值对得到我们想要的结果，理解键值对是理解Map-Reduce的核心。

