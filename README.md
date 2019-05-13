# TRON_TPS
---
## 背景
TRON波场采用DPOS共识机制，27个超级节点，3秒出一个块（1/3 BPS）,现对其进行TPS性能测试
---
## 官方资料
[私链搭建方法教程](https://github.com/tronprotocol/documentation/blob/master/TRX_CN/Tron-doc.md#463-%E6%90%AD%E5%BB%BA%E7%A7%81%E6%9C%89%E7%BD%91%E7%BB%9C)

**Note:**
官方提供的私链搭建方法有问题,在此基础此处对[private_net_config.conf](https://github.com/tronprotocol/tron-deployment/blob/master/private_net_config.conf)进行修改

详情请见[出块节点supernode配置文件](./private_net/supernode/private_net_config.conf) 和
[全节点fullnode配置文件](./private_net/fullnode/private_net_config.conf)

[官方HTTP API](https://github.com/tronprotocol/documentation/blob/master/TRX_CN/Tron-http.md)

---
## 测试环境
![](https://github.com/JKinGH/TRON_TPS/blob/master/image/environment.jpg)

---
## 测试步骤
### 1.通过HTTP接口`easytransferbyprivate`构造转账交易
### 2.多线程脚本模拟多用户调用
### 3.通过HTTP接口`gettransactioncountbyblocknum`扫描每个测试区块中包含的交易笔数

---
## 注意事项
### 1.节点本身对HTTP请求做了限制，线程太多时会出现`connection reset`的情况
### 2.故增加`fullnode`模拟多节点私链，交易通过多个`fullnode`发送

---
## 测试结果
### 1.经测试，当存在一个 `supernode`和一个`fullnode`时TPS测试结果最佳
### 2.测试结果数据
#### 测试程序运行20 mins后log记录如下，可见TPS在稳定在450左右，计算出平均值为441
```
 GetTransactionCountByBlocknum ,height=5356, in range[1000~2000],contain transaction num= 1152 ,TPS=384 
 GetTransactionCountByBlocknum ,height=5357, in range[1000~2000],contain transaction num= 1535 ,TPS=511 
 GetTransactionCountByBlocknum ,height=5358, in range[1000~2000],contain transaction num= 1371 ,TPS=457 
 GetTransactionCountByBlocknum ,height=5359, in range[1000~2000],contain transaction num= 1501 ,TPS=500 
 GetTransactionCountByBlocknum ,height=5360, in range[1000~2000],contain transaction num= 1219 ,TPS=406 
 GetTransactionCountByBlocknum ,height=5361, in range[1000~2000],contain transaction num= 1316 ,TPS=438 
 GetTransactionCountByBlocknum ,height=5362, in range[1000~2000],contain transaction num= 1380 ,TPS=460 
 GetTransactionCountByBlocknum ,height=5363, in range[1000~2000],contain transaction num= 1284 ,TPS=428 
 GetTransactionCountByBlocknum ,height=5364, in range[1000~2000],contain transaction num= 1310 ,TPS=436 
 GetTransactionCountByBlocknum ,height=5365, in range[1000~2000],contain transaction num= 1523 ,TPS=507 
 GetTransactionCountByBlocknum ,height=5366, in range[1000~2000],contain transaction num= 1190 ,TPS=396 
 GetTransactionCountByBlocknum ,height=5367, in range[1000~2000],contain transaction num= 1292 ,TPS=430 
 GetTransactionCountByBlocknum ,height=5368, in range[1000~2000],contain transaction num= 1372 ,TPS=457 
 GetTransactionCountByBlocknum ,height=5369, in range[1000~2000],contain transaction num= 1319 ,TPS=439 
 GetTransactionCountByBlocknum ,height=5370, in range[1000~2000],contain transaction num= 1394 ,TPS=464 
 GetTransactionCountByBlocknum ,height=5371, in range[1000~2000],contain transaction num= 1294 ,TPS=431 
 GetTransactionCountByBlocknum ,height=5372, in range[1000~2000],contain transaction num= 1354 ,TPS=451 
 GetTransactionCountByBlocknum ,height=5373, in range[1000~2000],contain transaction num= 1382 ,TPS=460 
 GetTransactionCountByBlocknum ,height=5374, in range[1000~2000],contain transaction num= 1460 ,TPS=486 
 GetTransactionCountByBlocknum ,height=5375, in range[1000~2000],contain transaction num= 1058 ,TPS=352 
 GetTransactionCountByBlocknum ,height=5376, in range[1000~2000],contain transaction num= 1421 ,TPS=473 
 GetTransactionCountByBlocknum ,height=5377, in range[1000~2000],contain transaction num= 1269 ,TPS=423 
 GetTransactionCountByBlocknum ,height=5378, in range[1000~2000],contain transaction num= 1303 ,TPS=434 
 GetTransactionCountByBlocknum ,height=5416, in range[1000~2000],contain transaction num= 1495 ,TPS=498 
 GetTransactionCountByBlocknum ,height=5417, in range[1000~2000],contain transaction num= 1441 ,TPS=480 
 GetTransactionCountByBlocknum ,height=5418, in range[1000~2000],contain transaction num= 1385 ,TPS=461 
 GetTransactionCountByBlocknum ,height=5419, in range[1000~2000],contain transaction num= 1575 ,TPS=525 
 GetTransactionCountByBlocknum ,height=5420, in range[1000~2000],contain transaction num= 1196 ,TPS=398 
 GetTransactionCountByBlocknum ,height=5421, in range[1000~2000],contain transaction num= 1446 ,TPS=482 
 GetTransactionCountByBlocknum ,height=5422, in range[1000~2000],contain transaction num= 1215 ,TPS=405 
 GetTransactionCountByBlocknum ,height=5423, in range[1000~2000],contain transaction num= 1481 ,TPS=493 
 GetTransactionCountByBlocknum ,height=5424, in range[1000~2000],contain transaction num= 1477 ,TPS=492 
 GetTransactionCountByBlocknum ,height=5425, in range[1000~2000],contain transaction num= 1392 ,TPS=464 
 GetTransactionCountByBlocknum ,height=5426, in range[1000~2000],contain transaction num= 1561 ,TPS=520 
 GetTransactionCountByBlocknum ,height=5427, in range[1000~2000],contain transaction num= 1664 ,TPS=554 
 GetTransactionCountByBlocknum ,height=5428, in range[1000~2000],contain transaction num= 1298 ,TPS=432 
 GetTransactionCountByBlocknum ,height=5429, in range[1000~2000],contain transaction num= 1405 ,TPS=468 
 GetTransactionCountByBlocknum ,height=5430, in range[1000~2000],contain transaction num= 1350 ,TPS=450 
 GetTransactionCountByBlocknum ,height=5431, in range[1000~2000],contain transaction num= 1549 ,TPS=516 
 GetTransactionCountByBlocknum ,height=5432, in range[1000~2000],contain transaction num= 1276 ,TPS=425 
 GetTransactionCountByBlocknum ,height=5433, in range[1000~2000],contain transaction num= 1649 ,TPS=549 
 GetTransactionCountByBlocknum ,height=5435, in range[1000~2000],contain transaction num= 1552 ,TPS=517 
 GetTransactionCountByBlocknum ,height=5437, in range[1000~2000],contain transaction num= 1590 ,TPS=530 
 GetTransactionCountByBlocknum ,height=5438, in range[1000~2000],contain transaction num= 1450 ,TPS=483 
 GetTransactionCountByBlocknum ,height=5439, in range[1000~2000],contain transaction num= 1394 ,TPS=464 
 GetTransactionCountByBlocknum ,height=5440, in range[1000~2000],contain transaction num= 1403 ,TPS=467 
 GetTransactionCountByBlocknum ,height=5441, in range[1000~2000],contain transaction num= 1511 ,TPS=503 
 GetTransactionCountByBlocknum ,height=5442, in range[1000~2000],contain transaction num= 1423 ,TPS=474 
 GetTransactionCountByBlocknum ,height=5443, in range[1000~2000],contain transaction num= 1436 ,TPS=478 
 GetTransactionCountByBlocknum ,height=5444, in range[1000~2000],contain transaction num= 1621 ,TPS=540 
 GetTransactionCountByBlocknum ,height=5445, in range[1000~2000],contain transaction num= 1255 ,TPS=418 
 GetTransactionCountByBlocknum ,height=5446, in range[1000~2000],contain transaction num= 1387 ,TPS=462 
 GetTransactionCountByBlocknum ,height=5447, in range[1000~2000],contain transaction num= 1465 ,TPS=488 
 GetTransactionCountByBlocknum ,height=5448, in range[1000~2000],contain transaction num= 1313 ,TPS=437 
 GetTransactionCountByBlocknum ,height=5449, in range[1000~2000],contain transaction num= 1512 ,TPS=504 
 GetTransactionCountByBlocknum ,height=5450, in range[1000~2000],contain transaction num= 1765 ,TPS=588 
 GetTransactionCountByBlocknum ,height=5451, in range[1000~2000],contain transaction num= 1229 ,TPS=409 
 GetTransactionCountByBlocknum ,height=5452, in range[1000~2000],contain transaction num= 1386 ,TPS=462 
 GetTransactionCountByBlocknum ,height=5453, in range[1000~2000],contain transaction num= 1162 ,TPS=387 
 GetTransactionCountByBlocknum ,height=5454, in range[1000~2000],contain transaction num= 1576 ,TPS=525 
 GetTransactionCountByBlocknum ,height=5455, in range[1000~2000],contain transaction num= 1423 ,TPS=474 
 GetTransactionCountByBlocknum ,height=5456, in range[1000~2000],contain transaction num= 1462 ,TPS=487 
 GetTransactionCountByBlocknum ,height=5457, in range[1000~2000],contain transaction num= 1579 ,TPS=526 
 GetTransactionCountByBlocknum ,height=5458, in range[1000~2000],contain transaction num= 1420 ,TPS=473 
 GetTransactionCountByBlocknum ,height=5459, in range[1000~2000],contain transaction num= 1436 ,TPS=478 
 GetTransactionCountByBlocknum ,height=5460, in range[1000~2000],contain transaction num= 1231 ,TPS=410 
 GetTransactionCountByBlocknum ,height=5461, in range[1000~2000],contain transaction num= 1238 ,TPS=412 
 GetTransactionCountByBlocknum ,height=5511, in range[1000~2000],contain transaction num= 1461 ,TPS=487 
 GetTransactionCountByBlocknum ,height=5512, in range[1000~2000],contain transaction num= 1404 ,TPS=468 
 GetTransactionCountByBlocknum ,height=5513, in range[1000~2000],contain transaction num= 1726 ,TPS=575 
 GetTransactionCountByBlocknum ,height=5514, in range[1000~2000],contain transaction num= 1161 ,TPS=387 
 GetTransactionCountByBlocknum ,height=5515, in range[1000~2000],contain transaction num= 1604 ,TPS=534 
 GetTransactionCountByBlocknum ,height=5516, in range[1000~2000],contain transaction num= 1265 ,TPS=421 
 GetTransactionCountByBlocknum ,height=5517, in range[1000~2000],contain transaction num= 1391 ,TPS=463 
 GetTransactionCountByBlocknum ,height=5518, in range[1000~2000],contain transaction num= 1541 ,TPS=513 
 GetTransactionCountByBlocknum ,height=5519, in range[1000~2000],contain transaction num= 1303 ,TPS=434 
 GetTransactionCountByBlocknum ,height=5520, in range[1000~2000],contain transaction num= 1357 ,TPS=452 
 GetTransactionCountByBlocknum ,height=5534, in range[1000~2000],contain transaction num= 1200 ,TPS=400 
 GetTransactionCountByBlocknum ,height=5535, in range[1000~2000],contain transaction num= 1316 ,TPS=438 
 GetTransactionCountByBlocknum ,height=5536, in range[1000~2000],contain transaction num= 1304 ,TPS=434 
 GetTransactionCountByBlocknum ,height=5537, in range[1000~2000],contain transaction num= 1450 ,TPS=483 
 GetTransactionCountByBlocknum ,height=5538, in range[1000~2000],contain transaction num= 1194 ,TPS=398 
 GetTransactionCountByBlocknum ,height=5539, in range[1000~2000],contain transaction num= 1717 ,TPS=572 
 GetTransactionCountByBlocknum ,height=5540, in range[1000~2000],contain transaction num= 1344 ,TPS=448 
 GetTransactionCountByBlocknum ,height=5541, in range[1000~2000],contain transaction num= 1386 ,TPS=462 
 GetTransactionCountByBlocknum ,height=5543, in range[1000~2000],contain transaction num= 1710 ,TPS=570 
 GetTransactionCountByBlocknum ,height=5584, in range[1000~2000],contain transaction num= 1818 ,TPS=606 
 GetTransactionCountByBlocknum ,height=5585, in range[1000~2000],contain transaction num= 1579 ,TPS=526 
 GetTransactionCountByBlocknum ,height=5586, in range[1000~2000],contain transaction num= 1672 ,TPS=557 
 GetTransactionCountByBlocknum ,height=5587, in range[1000~2000],contain transaction num= 1500 ,TPS=500 
 GetTransactionCountByBlocknum ,height=5588, in range[1000~2000],contain transaction num= 1732 ,TPS=577 
 GetTransactionCountByBlocknum ,height=5589, in range[1000~2000],contain transaction num= 1250 ,TPS=416 
 GetTransactionCountByBlocknum ,height=5590, in range[1000~2000],contain transaction num= 1527 ,TPS=509 
 GetTransactionCountByBlocknum ,height=5591, in range[1000~2000],contain transaction num= 1169 ,TPS=389 
 GetTransactionCountByBlocknum ,height=5592, in range[1000~2000],contain transaction num= 1447 ,TPS=482 
 GetTransactionCountByBlocknum ,height=5593, in range[1000~2000],contain transaction num= 1523 ,TPS=507 
 GetTransactionCountByBlocknum ,height=5594, in range[1000~2000],contain transaction num= 1472 ,TPS=490 
 GetTransactionCountByBlocknum ,height=5595, in range[1000~2000],contain transaction num= 1471 ,TPS=490 
 GetTransactionCountByBlocknum ,height=5596, in range[1000~2000],contain transaction num= 1349 ,TPS=449 
 GetTransactionCountByBlocknum ,height=5597, in range[1000~2000],contain transaction num= 1321 ,TPS=440 
 GetTransactionCountByBlocknum ,height=5598, in range[1000~2000],contain transaction num= 1285 ,TPS=428 
 GetTransactionCountByBlocknum ,height=5599, in range[1000~2000],contain transaction num= 1447 ,TPS=482 
 GetTransactionCountByBlocknum ,height=5600, in range[1000~2000],contain transaction num= 1518 ,TPS=506 
 GetTransactionCountByBlocknum ,height=5601, in range[1000~2000],contain transaction num= 1446 ,TPS=482 
 GetTransactionCountByBlocknum ,height=5602, in range[1000~2000],contain transaction num= 1702 ,TPS=567 
 GetTransactionCountByBlocknum ,height=5603, in range[1000~2000],contain transaction num= 1301 ,TPS=433 
 GetTransactionCountByBlocknum ,height=5604, in range[1000~2000],contain transaction num= 1447 ,TPS=482 
 GetTransactionCountByBlocknum ,height=5605, in range[1000~2000],contain transaction num= 1260 ,TPS=420 
 GetTransactionCountByBlocknum ,height=5606, in range[1000~2000],contain transaction num= 1303 ,TPS=434 
 GetTransactionCountByBlocknum ,height=5607, in range[1000~2000],contain transaction num= 1458 ,TPS=486 
 GetTransactionCountByBlocknum ,height=5608, in range[1000~2000],contain transaction num= 1365 ,TPS=455 
 GetTransactionCountByBlocknum ,height=5609, in range[1000~2000],contain transaction num= 1540 ,TPS=513 
 GetTransactionCountByBlocknum ,height=5610, in range[1000~2000],contain transaction num= 1273 ,TPS=424 
 GetTransactionCountByBlocknum ,height=5611, in range[1000~2000],contain transaction num= 1499 ,TPS=499 
 GetTransactionCountByBlocknum ,height=5612, in range[1000~2000],contain transaction num= 1088 ,TPS=362 
 GetTransactionCountByBlocknum ,height=5613, in range[1000~2000],contain transaction num= 1666 ,TPS=555 
 GetTransactionCountByBlocknum ,height=5614, in range[1000~2000],contain transaction num= 1191 ,TPS=397 
 GetTransactionCountByBlocknum ,height=5615, in range[1000~2000],contain transaction num= 1526 ,TPS=508 
 GetTransactionCountByBlocknum ,height=5616, in range[1000~2000],contain transaction num= 1408 ,TPS=469 
 GetTransactionCountByBlocknum ,height=5617, in range[1000~2000],contain transaction num= 1490 ,TPS=496 
 GetTransactionCountByBlocknum ,height=5618, in range[1000~2000],contain transaction num= 1422 ,TPS=474 
 GetTransactionCountByBlocknum ,height=5619, in range[1000~2000],contain transaction num= 1606 ,TPS=535 
 GetTransactionCountByBlocknum ,height=5620, in range[1000~2000],contain transaction num= 1340 ,TPS=446 
 GetTransactionCountByBlocknum ,height=5621, in range[1000~2000],contain transaction num= 1385 ,TPS=461 
 GetTransactionCountByBlocknum ,height=5622, in range[1000~2000],contain transaction num= 1328 ,TPS=442 
 GetTransactionCountByBlocknum ,height=5623, in range[1000~2000],contain transaction num= 1432 ,TPS=477 
 GetTransactionCountByBlocknum ,height=5624, in range[1000~2000],contain transaction num= 1408 ,TPS=469 
 GetTransactionCountByBlocknum ,height=5625, in range[1000~2000],contain transaction num= 1464 ,TPS=488 
 GetTransactionCountByBlocknum ,height=5626, in range[1000~2000],contain transaction num= 1457 ,TPS=485 
 GetTransactionCountByBlocknum ,height=5627, in range[1000~2000],contain transaction num= 1439 ,TPS=479 
 GetTransactionCountByBlocknum ,height=5628, in range[1000~2000],contain transaction num= 1665 ,TPS=555 
 GetTransactionCountByBlocknum ,height=5629, in range[1000~2000],contain transaction num= 1178 ,TPS=392 
 GetTransactionCountByBlocknum ,height=5630, in range[1000~2000],contain transaction num= 1476 ,TPS=492 
 GetTransactionCountByBlocknum ,height=5631, in range[1000~2000],contain transaction num= 1655 ,TPS=551 
 GetTransactionCountByBlocknum ,height=5632, in range[1000~2000],contain transaction num= 1436 ,TPS=478 
 GetTransactionCountByBlocknum ,height=5633, in range[1000~2000],contain transaction num= 1252 ,TPS=417 
 GetTransactionCountByBlocknum ,height=5635, in range[    >2000],contain transaction num= 2115 ,TPS=705 
 GetTransactionCountByBlocknum ,height=5636, in range[1000~2000],contain transaction num= 1291 ,TPS=430 
 GetTransactionCountByBlocknum ,height=5898, in range[1000~2000],contain transaction num= 1079 ,TPS=359 
 GetTransactionCountByBlocknum ,height=5901, in range[1000~2000],contain transaction num= 1108 ,TPS=369 
 GetTransactionCountByBlocknum ,height=5904, in range[1000~2000],contain transaction num= 1050 ,TPS=350 
 GetTransactionCountByBlocknum ,height=5905, in range[1000~2000],contain transaction num= 1076 ,TPS=358 
 GetTransactionCountByBlocknum ,height=5909, in range[1000~2000],contain transaction num= 1225 ,TPS=408 
 GetTransactionCountByBlocknum ,height=5911, in range[1000~2000],contain transaction num= 1201 ,TPS=400 
 GetTransactionCountByBlocknum ,height=5917, in range[1000~2000],contain transaction num= 1196 ,TPS=398 
 GetTransactionCountByBlocknum ,height=5919, in range[1000~2000],contain transaction num= 1205 ,TPS=401 
 GetTransactionCountByBlocknum ,height=5921, in range[1000~2000],contain transaction num= 1257 ,TPS=419 
 GetTransactionCountByBlocknum ,height=5923, in range[1000~2000],contain transaction num= 1105 ,TPS=368 
 GetTransactionCountByBlocknum ,height=5927, in range[1000~2000],contain transaction num= 1346 ,TPS=448 
 GetTransactionCountByBlocknum ,height=5928, in range[1000~2000],contain transaction num= 1111 ,TPS=370 
 GetTransactionCountByBlocknum ,height=5931, in range[1000~2000],contain transaction num= 1032 ,TPS=344 
 GetTransactionCountByBlocknum ,height=5932, in range[1000~2000],contain transaction num= 1606 ,TPS=535 
 GetTransactionCountByBlocknum ,height=5934, in range[1000~2000],contain transaction num= 1124 ,TPS=374 
 GetTransactionCountByBlocknum ,height=5935, in range[1000~2000],contain transaction num= 1141 ,TPS=380 
 GetTransactionCountByBlocknum ,height=5937, in range[1000~2000],contain transaction num= 1331 ,TPS=443 
 GetTransactionCountByBlocknum ,height=5938, in range[1000~2000],contain transaction num= 1034 ,TPS=344 
 GetTransactionCountByBlocknum ,height=5939, in range[1000~2000],contain transaction num= 1123 ,TPS=374 
 GetTransactionCountByBlocknum ,height=5941, in range[1000~2000],contain transaction num= 1139 ,TPS=379 
 GetTransactionCountByBlocknum ,height=5944, in range[1000~2000],contain transaction num= 1192 ,TPS=397 
 GetTransactionCountByBlocknum ,height=5946, in range[1000~2000],contain transaction num= 1036 ,TPS=345 
 GetTransactionCountByBlocknum ,height=5947, in range[1000~2000],contain transaction num= 1014 ,TPS=338 
 GetTransactionCountByBlocknum ,height=5948, in range[1000~2000],contain transaction num= 1144 ,TPS=381 
 GetTransactionCountByBlocknum ,height=5949, in range[1000~2000],contain transaction num= 1329 ,TPS=443 
 GetTransactionCountByBlocknum ,height=5951, in range[1000~2000],contain transaction num= 1007 ,TPS=335 
 GetTransactionCountByBlocknum ,height=5952, in range[1000~2000],contain transaction num= 1295 ,TPS=431 
 GetTransactionCountByBlocknum ,height=5954, in range[1000~2000],contain transaction num= 1209 ,TPS=403 
 GetTransactionCountByBlocknum ,height=5955, in range[1000~2000],contain transaction num= 1176 ,TPS=392 
 GetTransactionCountByBlocknum ,height=5957, in range[1000~2000],contain transaction num= 1208 ,TPS=402 
 GetTransactionCountByBlocknum ,height=5959, in range[1000~2000],contain transaction num= 1027 ,TPS=342 
 GetTransactionCountByBlocknum ,height=5960, in range[1000~2000],contain transaction num= 1012 ,TPS=337 
 GetTransactionCountByBlocknum ,height=5961, in range[1000~2000],contain transaction num= 1018 ,TPS=339 
 GetTransactionCountByBlocknum ,height=5962, in range[1000~2000],contain transaction num= 1221 ,TPS=407 
 GetTransactionCountByBlocknum ,height=5963, in range[1000~2000],contain transaction num= 1035 ,TPS=345 
 GetTransactionCountByBlocknum ,height=5965, in range[1000~2000],contain transaction num= 1084 ,TPS=361 
 GetTransactionCountByBlocknum ,height=5966, in range[1000~2000],contain transaction num= 1198 ,TPS=399 
 GetTransactionCountByBlocknum ,height=5967, in range[1000~2000],contain transaction num= 1188 ,TPS=396 
 GetTransactionCountByBlocknum ,height=5969, in range[1000~2000],contain transaction num= 1064 ,TPS=354 
 GetTransactionCountByBlocknum ,height=5970, in range[1000~2000],contain transaction num= 1111 ,TPS=370 
 GetTransactionCountByBlocknum ,height=5971, in range[1000~2000],contain transaction num= 1263 ,TPS=421 
 GetTransactionCountByBlocknum ,height=5974, in range[1000~2000],contain transaction num= 1200 ,TPS=400 
 GetTransactionCountByBlocknum ,height=5976, in range[1000~2000],contain transaction num= 1166 ,TPS=388 
 GetTransactionCountByBlocknum ,height=5981, in range[1000~2000],contain transaction num= 1042 ,TPS=347 
 GetTransactionCountByBlocknum ,height=5983, in range[1000~2000],contain transaction num= 1040 ,TPS=346 
 GetTransactionCountByBlocknum ,height=5984, in range[1000~2000],contain transaction num= 1224 ,TPS=408 
 GetTransactionCountByBlocknum ,height=5986, in range[1000~2000],contain transaction num= 1148 ,TPS=382 
 GetTransactionCountByBlocknum ,height=5988, in range[1000~2000],contain transaction num= 1005 ,TPS=335 
 GetTransactionCountByBlocknum ,height=5991, in range[1000~2000],contain transaction num= 1859 ,TPS=619 
 GetTransactionCountByBlocknum ,height=5994, in range[1000~2000],contain transaction num= 1440 ,TPS=480 
 GetTransactionCountByBlocknum ,height=5995, in range[1000~2000],contain transaction num= 1007 ,TPS=335 
 GetTransactionCountByBlocknum ,height=5996, in range[1000~2000],contain transaction num= 1012 ,TPS=337 
 GetTransactionCountByBlocknum ,height=5997, in range[1000~2000],contain transaction num= 1309 ,TPS=436 
 GetTransactionCountByBlocknum ,height=5999, in range[1000~2000],contain transaction num= 1044 ,TPS=348 
 GetTransactionCountByBlocknum ,height=6000, in range[1000~2000],contain transaction num= 1163 ,TPS=387 
 GetTransactionCountByBlocknum ,height=6002, in range[1000~2000],contain transaction num= 1132 ,TPS=377 
 GetTransactionCountByBlocknum ,height=6003, in range[1000~2000],contain transaction num= 1001 ,TPS=333 
 GetTransactionCountByBlocknum ,height=6006, in range[1000~2000],contain transaction num= 1142 ,TPS=380 
 GetTransactionCountByBlocknum ,height=6009, in range[1000~2000],contain transaction num= 1331 ,TPS=443 
 GetTransactionCountByBlocknum ,height=6011, in range[1000~2000],contain transaction num= 1106 ,TPS=368 
 GetTransactionCountByBlocknum ,height=6013, in range[1000~2000],contain transaction num= 1050 ,TPS=350 
 GetTransactionCountByBlocknum ,height=6014, in range[1000~2000],contain transaction num= 1220 ,TPS=406 
 GetTransactionCountByBlocknum ,height=6015, in range[1000~2000],contain transaction num= 1203 ,TPS=401 
 GetTransactionCountByBlocknum ,height=6016, in range[1000~2000],contain transaction num= 1005 ,TPS=335 
 GetTransactionCountByBlocknum ,height=6928, in range[1000~2000],contain transaction num= 1219 ,TPS=406 
 GetTransactionCountByBlocknum ,height=6929, in range[1000~2000],contain transaction num= 1233 ,TPS=411 
 GetTransactionCountByBlocknum ,height=6930, in range[1000~2000],contain transaction num= 1473 ,TPS=491 
 GetTransactionCountByBlocknum ,height=6931, in range[1000~2000],contain transaction num= 1283 ,TPS=427 
 GetTransactionCountByBlocknum ,height=6932, in range[1000~2000],contain transaction num= 1013 ,TPS=337 
 GetTransactionCountByBlocknum ,height=6933, in range[1000~2000],contain transaction num= 1311 ,TPS=437 
 GetTransactionCountByBlocknum ,height=6937, in range[1000~2000],contain transaction num= 1573 ,TPS=524 
 GetTransactionCountByBlocknum ,height=6938, in range[1000~2000],contain transaction num= 1182 ,TPS=394 
 GetTransactionCountByBlocknum ,height=6939, in range[1000~2000],contain transaction num= 1166 ,TPS=388 
 GetTransactionCountByBlocknum ,height=6940, in range[1000~2000],contain transaction num= 1168 ,TPS=389 
 GetTransactionCountByBlocknum ,height=6941, in range[1000~2000],contain transaction num= 1022 ,TPS=340 
 GetTransactionCountByBlocknum ,height=6942, in range[1000~2000],contain transaction num= 1233 ,TPS=411 
 GetTransactionCountByBlocknum ,height=6943, in range[1000~2000],contain transaction num= 1304 ,TPS=434 
 GetTransactionCountByBlocknum ,height=6944, in range[1000~2000],contain transaction num= 1005 ,TPS=335 
 GetTransactionCountByBlocknum ,height=6945, in range[1000~2000],contain transaction num= 1249 ,TPS=416 
 GetTransactionCountByBlocknum ,height=6947, in range[1000~2000],contain transaction num= 1208 ,TPS=402 
 GetTransactionCountByBlocknum ,height=6948, in range[1000~2000],contain transaction num= 1157 ,TPS=385 
 GetTransactionCountByBlocknum ,height=6949, in range[1000~2000],contain transaction num= 1104 ,TPS=368 
 GetTransactionCountByBlocknum ,height=6950, in range[1000~2000],contain transaction num= 1432 ,TPS=477 
 GetTransactionCountByBlocknum ,height=6952, in range[1000~2000],contain transaction num= 1213 ,TPS=404 
 GetTransactionCountByBlocknum ,height=6953, in range[1000~2000],contain transaction num= 1277 ,TPS=425 
 GetTransactionCountByBlocknum ,height=6954, in range[1000~2000],contain transaction num= 1216 ,TPS=405 
 GetTransactionCountByBlocknum ,height=6955, in range[1000~2000],contain transaction num= 1343 ,TPS=447 
 GetTransactionCountByBlocknum ,height=6956, in range[1000~2000],contain transaction num= 1243 ,TPS=414 
 GetTransactionCountByBlocknum ,height=6957, in range[1000~2000],contain transaction num= 1103 ,TPS=367 
 GetTransactionCountByBlocknum ,height=6958, in range[1000~2000],contain transaction num= 1220 ,TPS=406 
 GetTransactionCountByBlocknum ,height=6959, in range[1000~2000],contain transaction num= 1086 ,TPS=362 
 GetTransactionCountByBlocknum ,height=6960, in range[1000~2000],contain transaction num= 1076 ,TPS=358 
 GetTransactionCountByBlocknum ,height=6961, in range[1000~2000],contain transaction num= 1164 ,TPS=388 
 GetTransactionCountByBlocknum ,height=6962, in range[1000~2000],contain transaction num= 1489 ,TPS=496 
 GetTransactionCountByBlocknum ,height=6964, in range[1000~2000],contain transaction num= 1388 ,TPS=462 
 GetTransactionCountByBlocknum ,height=6965, in range[1000~2000],contain transaction num= 1122 ,TPS=374 
 GetTransactionCountByBlocknum ,height=6966, in range[1000~2000],contain transaction num= 1406 ,TPS=468 
 GetTransactionCountByBlocknum ,height=6967, in range[1000~2000],contain transaction num= 1362 ,TPS=454 
 GetTransactionCountByBlocknum ,height=6968, in range[1000~2000],contain transaction num= 1360 ,TPS=453 
 GetTransactionCountByBlocknum ,height=6969, in range[1000~2000],contain transaction num= 1361 ,TPS=453 
 GetTransactionCountByBlocknum ,height=6970, in range[1000~2000],contain transaction num= 1505 ,TPS=501 
 GetTransactionCountByBlocknum ,height=6971, in range[1000~2000],contain transaction num= 1246 ,TPS=415 
 GetTransactionCountByBlocknum ,height=6972, in range[1000~2000],contain transaction num= 1212 ,TPS=404 
 GetTransactionCountByBlocknum ,height=6973, in range[1000~2000],contain transaction num= 1440 ,TPS=480 
 GetTransactionCountByBlocknum ,height=6974, in range[1000~2000],contain transaction num= 1158 ,TPS=386 
 GetTransactionCountByBlocknum ,height=6975, in range[1000~2000],contain transaction num= 1386 ,TPS=462 
 GetTransactionCountByBlocknum ,height=6976, in range[1000~2000],contain transaction num= 1634 ,TPS=544 
 GetTransactionCountByBlocknum ,height=6977, in range[1000~2000],contain transaction num= 1011 ,TPS=337 
 GetTransactionCountByBlocknum ,height=6978, in range[1000~2000],contain transaction num= 1529 ,TPS=509 
 GetTransactionCountByBlocknum ,height=6980, in range[1000~2000],contain transaction num= 1359 ,TPS=453 
 GetTransactionCountByBlocknum ,height=6981, in range[1000~2000],contain transaction num= 1189 ,TPS=396 
 GetTransactionCountByBlocknum ,height=6982, in range[1000~2000],contain transaction num= 1468 ,TPS=489 
 GetTransactionCountByBlocknum ,height=6983, in range[1000~2000],contain transaction num= 1180 ,TPS=393 
 GetTransactionCountByBlocknum ,height=6984, in range[1000~2000],contain transaction num= 1405 ,TPS=468 
 GetTransactionCountByBlocknum ,height=6985, in range[1000~2000],contain transaction num= 1339 ,TPS=446 
 GetTransactionCountByBlocknum ,height=6986, in range[1000~2000],contain transaction num= 1304 ,TPS=434 
 GetTransactionCountByBlocknum ,height=6988, in range[1000~2000],contain transaction num= 1343 ,TPS=447 
 GetTransactionCountByBlocknum ,height=6989, in range[1000~2000],contain transaction num= 1144 ,TPS=381 
 GetTransactionCountByBlocknum ,height=6990, in range[1000~2000],contain transaction num= 1506 ,TPS=502 
 GetTransactionCountByBlocknum ,height=6991, in range[1000~2000],contain transaction num= 1009 ,TPS=336 
 GetTransactionCountByBlocknum ,height=6992, in range[1000~2000],contain transaction num= 1228 ,TPS=409 
 GetTransactionCountByBlocknum ,height=6993, in range[1000~2000],contain transaction num= 1125 ,TPS=375 
 GetTransactionCountByBlocknum ,height=6994, in range[1000~2000],contain transaction num= 1244 ,TPS=414 
 GetTransactionCountByBlocknum ,height=6995, in range[1000~2000],contain transaction num= 1152 ,TPS=384 
 GetTransactionCountByBlocknum ,height=6996, in range[1000~2000],contain transaction num= 1201 ,TPS=400 
 GetTransactionCountByBlocknum ,height=6997, in range[1000~2000],contain transaction num= 1267 ,TPS=422 
 GetTransactionCountByBlocknum ,height=6998, in range[1000~2000],contain transaction num= 1275 ,TPS=425 
 GetTransactionCountByBlocknum ,height=6999, in range[1000~2000],contain transaction num= 1048 ,TPS=349 
 GetTransactionCountByBlocknum ,height=7039, in range[1000~2000],contain transaction num= 1160 ,TPS=386 
 GetTransactionCountByBlocknum ,height=7040, in range[1000~2000],contain transaction num= 1292 ,TPS=430 
 GetTransactionCountByBlocknum ,height=7042, in range[1000~2000],contain transaction num= 1477 ,TPS=492 
 GetTransactionCountByBlocknum ,height=7043, in range[1000~2000],contain transaction num= 1519 ,TPS=506 
 GetTransactionCountByBlocknum ,height=7044, in range[1000~2000],contain transaction num= 1457 ,TPS=485 
 GetTransactionCountByBlocknum ,height=7079, in range[1000~2000],contain transaction num= 1210 ,TPS=403 
 GetTransactionCountByBlocknum ,height=7082, in range[    >2000],contain transaction num= 2188 ,TPS=729 
 GetTransactionCountByBlocknum ,height=7090, in range[1000~2000],contain transaction num= 1568 ,TPS=522 
 GetTransactionCountByBlocknum ,height=7099, in range[1000~2000],contain transaction num= 1476 ,TPS=492 
 GetTransactionCountByBlocknum ,height=7128, in range[1000~2000],contain transaction num= 1295 ,TPS=431 
 GetTransactionCountByBlocknum ,height=7129, in range[1000~2000],contain transaction num= 1081 ,TPS=360 
 GetTransactionCountByBlocknum ,height=7130, in range[1000~2000],contain transaction num= 1874 ,TPS=624 
 GetTransactionCountByBlocknum ,height=7131, in range[1000~2000],contain transaction num= 1074 ,TPS=358 
 GetTransactionCountByBlocknum ,height=7132, in range[1000~2000],contain transaction num= 1604 ,TPS=534 
 GetTransactionCountByBlocknum ,height=7133, in range[1000~2000],contain transaction num= 1502 ,TPS=500 
 GetTransactionCountByBlocknum ,height=7135, in range[1000~2000],contain transaction num= 1951 ,TPS=650 
 GetTransactionCountByBlocknum ,height=7139, in range[1000~2000],contain transaction num= 1056 ,TPS=352 
 GetTransactionCountByBlocknum ,height=7146, in range[1000~2000],contain transaction num= 1322 ,TPS=440 
 GetTransactionCountByBlocknum ,height=7147, in range[1000~2000],contain transaction num= 1402 ,TPS=467 
 GetTransactionCountByBlocknum ,height=7149, in range[1000~2000],contain transaction num= 1150 ,TPS=383 
 GetTransactionCountByBlocknum ,height=7152, in range[    >2000],contain transaction num= 2013 ,TPS=671 
 GetTransactionCountByBlocknum ,height=7154, in range[1000~2000],contain transaction num= 1202 ,TPS=400 
 GetTransactionCountByBlocknum ,height=7155, in range[1000~2000],contain transaction num= 1150 ,TPS=383 
 GetTransactionCountByBlocknum ,height=7158, in range[1000~2000],contain transaction num= 1094 ,TPS=364 
 GetTransactionCountByBlocknum ,height=7159, in range[1000~2000],contain transaction num= 1305 ,TPS=435 
 GetTransactionCountByBlocknum ,height=7161, in range[1000~2000],contain transaction num= 1767 ,TPS=589 
 GetTransactionCountByBlocknum ,height=7162, in range[1000~2000],contain transaction num= 1166 ,TPS=388 
 GetTransactionCountByBlocknum ,height=7163, in range[1000~2000],contain transaction num= 1318 ,TPS=439 
 GetTransactionCountByBlocknum ,height=7165, in range[    >2000],contain transaction num= 2082 ,TPS=694 
 GetTransactionCountByBlocknum ,height=7167, in range[1000~2000],contain transaction num= 1070 ,TPS=356 
 GetTransactionCountByBlocknum ,height=7169, in range[1000~2000],contain transaction num= 1089 ,TPS=363 
 GetTransactionCountByBlocknum ,height=7170, in range[1000~2000],contain transaction num= 1242 ,TPS=414 
 GetTransactionCountByBlocknum ,height=7171, in range[1000~2000],contain transaction num= 1183 ,TPS=394 
 GetTransactionCountByBlocknum ,height=7174, in range[    >2000],contain transaction num= 2104 ,TPS=701 
 GetTransactionCountByBlocknum ,height=7176, in range[1000~2000],contain transaction num= 1116 ,TPS=372 
 GetTransactionCountByBlocknum ,height=7177, in range[1000~2000],contain transaction num= 1299 ,TPS=433 
 GetTransactionCountByBlocknum ,height=7178, in range[1000~2000],contain transaction num= 1113 ,TPS=371 
 GetTransactionCountByBlocknum ,height=7179, in range[1000~2000],contain transaction num= 1285 ,TPS=428 
 GetTransactionCountByBlocknum ,height=7180, in range[1000~2000],contain transaction num= 1375 ,TPS=458 
 GetTransactionCountByBlocknum ,height=7181, in range[1000~2000],contain transaction num= 1037 ,TPS=345 
 GetTransactionCountByBlocknum ,height=7182, in range[1000~2000],contain transaction num= 1183 ,TPS=394 
 GetTransactionCountByBlocknum ,height=7183, in range[1000~2000],contain transaction num= 1314 ,TPS=438 
 GetTransactionCountByBlocknum ,height=7184, in range[1000~2000],contain transaction num= 1266 ,TPS=422 
 GetTransactionCountByBlocknum ,height=7185, in range[1000~2000],contain transaction num= 1232 ,TPS=410 
 GetTransactionCountByBlocknum ,height=7186, in range[1000~2000],contain transaction num= 1184 ,TPS=394 
 GetTransactionCountByBlocknum ,height=7188, in range[1000~2000],contain transaction num= 1083 ,TPS=361 
 GetTransactionCountByBlocknum ,height=7190, in range[1000~2000],contain transaction num= 1324 ,TPS=441 
 GetTransactionCountByBlocknum ,height=7191, in range[1000~2000],contain transaction num= 1169 ,TPS=389 
 GetTransactionCountByBlocknum ,height=7192, in range[1000~2000],contain transaction num= 1232 ,TPS=410 
 GetTransactionCountByBlocknum ,height=7193, in range[1000~2000],contain transaction num= 1438 ,TPS=479 
 GetTransactionCountByBlocknum ,height=7195, in range[1000~2000],contain transaction num= 1319 ,TPS=439 
 GetTransactionCountByBlocknum ,height=7197, in range[1000~2000],contain transaction num= 1713 ,TPS=571 
 GetTransactionCountByBlocknum ,height=7198, in range[1000~2000],contain transaction num= 1285 ,TPS=428 
 GetTransactionCountByBlocknum ,height=7199, in range[1000~2000],contain transaction num= 1281 ,TPS=427 
 GetTransactionCountByBlocknum ,height=7200, in range[1000~2000],contain transaction num= 1142 ,TPS=380 
 GetTransactionCountByBlocknum ,height=7201, in range[1000~2000],contain transaction num= 1145 ,TPS=381 
 GetTransactionCountByBlocknum ,height=7202, in range[1000~2000],contain transaction num= 1422 ,TPS=474 
 GetTransactionCountByBlocknum ,height=7204, in range[1000~2000],contain transaction num= 1908 ,TPS=636 
 GetTransactionCountByBlocknum ,height=7205, in range[1000~2000],contain transaction num= 1324 ,TPS=441 
 GetTransactionCountByBlocknum ,height=7208, in range[    >2000],contain transaction num= 2086 ,TPS=695 
 GetTransactionCountByBlocknum ,height=7212, in range[1000~2000],contain transaction num= 1482 ,TPS=494 
 GetTransactionCountByBlocknum ,height=7213, in range[1000~2000],contain transaction num= 1437 ,TPS=479 
 GetTransactionCountByBlocknum ,height=7214, in range[1000~2000],contain transaction num= 1396 ,TPS=465 
 GetTransactionCountByBlocknum ,height=7215, in range[1000~2000],contain transaction num= 1106 ,TPS=368 
 GetTransactionCountByBlocknum ,height=7216, in range[1000~2000],contain transaction num= 1470 ,TPS=490 
 GetTransactionCountByBlocknum ,height=7217, in range[1000~2000],contain transaction num= 1071 ,TPS=357 
 GetTransactionCountByBlocknum ,height=7218, in range[1000~2000],contain transaction num= 1158 ,TPS=386 
 GetTransactionCountByBlocknum ,height=7219, in range[1000~2000],contain transaction num= 1578 ,TPS=526 
 GetTransactionCountByBlocknum ,height=7221, in range[1000~2000],contain transaction num= 1500 ,TPS=500 
 GetTransactionCountByBlocknum ,height=7222, in range[1000~2000],contain transaction num= 1573 ,TPS=524 
 GetTransactionCountByBlocknum ,height=7224, in range[1000~2000],contain transaction num= 1327 ,TPS=442 
 GetTransactionCountByBlocknum ,height=7225, in range[1000~2000],contain transaction num= 1038 ,TPS=346 
 GetTransactionCountByBlocknum ,height=7226, in range[1000~2000],contain transaction num= 1256 ,TPS=418 
 GetTransactionCountByBlocknum ,height=7227, in range[1000~2000],contain transaction num= 1270 ,TPS=423 
 GetTransactionCountByBlocknum ,height=7228, in range[1000~2000],contain transaction num= 1420 ,TPS=473 
 GetTransactionCountByBlocknum ,height=7229, in range[1000~2000],contain transaction num= 1274 ,TPS=424 
 GetTransactionCountByBlocknum ,height=7231, in range[1000~2000],contain transaction num= 1226 ,TPS=408 
 GetTransactionCountByBlocknum ,height=7232, in range[1000~2000],contain transaction num= 1094 ,TPS=364 
 GetTransactionCountByBlocknum ,height=7233, in range[1000~2000],contain transaction num= 1281 ,TPS=427 
 GetTransactionCountByBlocknum ,height=7234, in range[1000~2000],contain transaction num= 1094 ,TPS=364 
 GetTransactionCountByBlocknum ,height=7235, in range[1000~2000],contain transaction num= 1383 ,TPS=461 
 GetTransactionCountByBlocknum ,height=7236, in range[1000~2000],contain transaction num= 1141 ,TPS=380 
 GetTransactionCountByBlocknum ,height=7239, in range[1000~2000],contain transaction num= 1016 ,TPS=338 
 GetTransactionCountByBlocknum ,height=7240, in range[1000~2000],contain transaction num= 1238 ,TPS=412 
 GetTransactionCountByBlocknum ,height=7241, in range[1000~2000],contain transaction num= 1224 ,TPS=408 
 GetTransactionCountByBlocknum ,height=7242, in range[1000~2000],contain transaction num= 1636 ,TPS=545 
 GetTransactionCountByBlocknum ,height=7243, in range[1000~2000],contain transaction num= 1159 ,TPS=386 
 GetTransactionCountByBlocknum ,height=7244, in range[1000~2000],contain transaction num= 1358 ,TPS=452 
 GetTransactionCountByBlocknum ,height=7245, in range[1000~2000],contain transaction num= 1100 ,TPS=366 
 GetTransactionCountByBlocknum ,height=7246, in range[1000~2000],contain transaction num= 1214 ,TPS=404 
 GetTransactionCountByBlocknum ,height=7247, in range[1000~2000],contain transaction num= 1402 ,TPS=467 
 GetTransactionCountByBlocknum ,height=7248, in range[1000~2000],contain transaction num= 1118 ,TPS=372 
 GetTransactionCountByBlocknum ,height=7249, in range[1000~2000],contain transaction num= 1179 ,TPS=393 
 GetTransactionCountByBlocknum ,height=7250, in range[1000~2000],contain transaction num= 1031 ,TPS=343 
 GetTransactionCountByBlocknum ,height=7251, in range[1000~2000],contain transaction num= 1152 ,TPS=384 
 GetTransactionCountByBlocknum ,height=7253, in range[    >2000],contain transaction num= 2054 ,TPS=684 
 GetTransactionCountByBlocknum ,height=7255, in range[1000~2000],contain transaction num= 1421 ,TPS=473 
 GetTransactionCountByBlocknum ,height=7256, in range[1000~2000],contain transaction num= 1308 ,TPS=436 
 GetTransactionCountByBlocknum ,height=7258, in range[1000~2000],contain transaction num= 1332 ,TPS=444 
 GetTransactionCountByBlocknum ,height=7260, in range[1000~2000],contain transaction num= 1310 ,TPS=436 
 GetTransactionCountByBlocknum ,height=7261, in range[1000~2000],contain transaction num= 1398 ,TPS=466 
 GetTransactionCountByBlocknum ,height=7263, in range[1000~2000],contain transaction num= 1260 ,TPS=420 
 GetTransactionCountByBlocknum ,height=7265, in range[    >2000],contain transaction num= 2071 ,TPS=690 
 GetTransactionCountByBlocknum ,height=7266, in range[1000~2000],contain transaction num= 1139 ,TPS=379 
 GetTransactionCountByBlocknum ,height=7267, in range[1000~2000],contain transaction num= 1193 ,TPS=397 
 GetTransactionCountByBlocknum ,height=7268, in range[1000~2000],contain transaction num= 1025 ,TPS=341 
 GetTransactionCountByBlocknum ,height=7269, in range[1000~2000],contain transaction num= 1343 ,TPS=447 
 GetTransactionCountByBlocknum ,height=7270, in range[1000~2000],contain transaction num= 1130 ,TPS=376 
 GetTransactionCountByBlocknum ,height=7271, in range[1000~2000],contain transaction num= 1097 ,TPS=365 
 GetTransactionCountByBlocknum ,height=7273, in range[1000~2000],contain transaction num= 1319 ,TPS=439 
 GetTransactionCountByBlocknum ,height=7275, in range[1000~2000],contain transaction num= 1132 ,TPS=377 
 GetTransactionCountByBlocknum ,height=7279, in range[    >2000],contain transaction num= 2045 ,TPS=681 
 GetTransactionCountByBlocknum ,height=7281, in range[1000~2000],contain transaction num= 1094 ,TPS=364 
```
#### 筛选出700左右的TPS,平均值=694（峰值）
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




