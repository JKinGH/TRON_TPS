# TRON_TPS
---

## 官方资料
[私链搭建方法教程](https://github.com/tronprotocol/documentation/blob/master/TRX_CN/Tron-doc.md#463-%E6%90%AD%E5%BB%BA%E7%A7%81%E6%9C%89%E7%BD%91%E7%BB%9C)

**Note**
官方提供的私链搭建方法有问题,在此基础此处对[private_net_config.conf](https://github.com/tronprotocol/tron-deployment/blob/master/private_net_config.conf)进行修改

详情请见[出块节点supernode配置文件](./private_net/supernode/private_net_config.conf) 和
[全节点fullnode配置文件](./private_net/fullnode/private_net_config.conf)

[官方HTTP API](https://github.com/tronprotocol/documentation/blob/master/TRX_CN/Tron-http.md)

## 测试环境
![](https://github.com/JKinGH/TRON_TPS/blob/master/image/environment.jpg)

## 测试步骤
### 1.通过HTTP接口构造转账交易
### 2.多线程模拟多用户调用

## 注意事项
### 1.节点本身对HTTP请求做了限制，线程太多时会出现`connection reset`的情况
### 2.故增加`fullnode`模拟多节点私链，交易通过多个`fullnode`发送

## 测试结果
### 1.经测试，当存在一个 `supernode`和一个`fullnode`时TPS测试结果最佳
### 2.如下展示几组TPS在700左右的数据
```
 GetTransactionCountByBlocknum ,height=5635, in range[    >2000],contain transaction num= 2115 ,TPS=705 
 GetTransactionCountByBlocknum ,height=7082, in range[    >2000],contain transaction num= 2188 ,TPS=729 
 GetTransactionCountByBlocknum ,height=7152, in range[    >2000],contain transaction num= 2013 ,TPS=671 
 GetTransactionCountByBlocknum ,height=7165, in range[    >2000],contain transaction num= 2082 ,TPS=694 
 GetTransactionCountByBlocknum ,height=7174, in range[    >2000],contain transaction num= 2104 ,TPS=701 
 GetTransactionCountByBlocknum ,height=7208, in range[    >2000],contain transaction num= 2086 ,TPS=695 
 GetTransactionCountByBlocknum ,height=7253, in range[    >2000],contain transaction num= 2054 ,TPS=684 
 GetTransactionCountByBlocknum ,height=7265, in range[    >2000],contain transaction num= 2071 ,TPS=690 
 GetTransactionCountByBlocknum ,height=7279, in range[    >2000],contain transaction num= 2045 ,TPS=681 
```




