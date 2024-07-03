简单的原理：

1. 客户需要支付20.05usdt
2. 服务器有一个hash表存储钱包地址对应的待支付金额 例如:address_1 : 20.05
3. 发起支付的时候，我们可以判定钱包address_1的20.05金额是否被占用，如果没有被占用那么可以直接返回这个钱包地址和金额给客户，告知客户需按规定金额20.05准确支付，少一分都不行。且将钱包地址和金额 address_1:20.05锁起来，有效期10分钟。
4. 如果订单并发下，又有一个20.05元需要支付，但是在第3步的时候上一个客户已经锁定了该金额，还在等待支付中...，那么我们将待支付金额加上0.0001，再次尝试判断address_1:20.0501金额是否被占用？如果没有则重复第三步，如果还是被占用就继续累加尝试，直到加了100次后都失败
5. 新开一个线程去监听所有钱包的USDT入账事件(类型为`Transfer`)，网上有公开的api或rpc节点。如果发现`订单创建时间 到 订单创建时间 + 订单有效时间`这个区间是否有入账金额与待支付的金额相等。则判断该笔订单支付成功(前提是该`transaction_id`在当前系统中不存在)！

技术栈思考

```
                                                    ---------->  1. 发送一条延迟队列用于处理订单超时关闭结束后台与TRC20轮询
                                                    |
生成订单数据 -> 保存订单到MySQL -> 将订单发送RabbitMQ     
                                                    |
                                                    ---------->  2. 发送一条消息到队列(Pub/Sub)用于处理轮询查询订单状态(未查到不ack重复继续消费)
                                                    
```
```
生成订单数据
    |
    
```