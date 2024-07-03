# lee-activity-framework
活动框架

## 一、项目背景
* 活动大部分只是换皮，服务端逻辑极其相似，开发活动框架可以节省大部分重复工作量
* 降低使用成本，可以直接配置活动并观察活动数据
* 减少开发人员需要重复编写代码导致的BUG
* 提高维护效率，避免创建太多冗余的数据库表和配置

## 二、项目特点
* 方便维护
* 减少重复开发工作量
* 减少BUG
* ...

## 三、模块介绍
* **lee-activity-api**：通用api，负责存放公共功能、通用组件等，如http通用请求响应、自定义错误码、自定义logger等
* **lee-activity-config**：活动服务，负责活动相关配置
* **lee-activity-lottery**：抽奖服务，负责抽取奖励发放
* **lee-activity-mq-producer**：mq生产者服务，负责模拟发送一些用户行为事件，如送礼、发消息等
* **lee-activity-prop**：道具服务，负责道具管理和发放
* **lee-activity-rank**：榜单服务，负责监听数据源并统计榜单值
* **lee-activity-task**：任务服务，负责任务管理和监听累计任务完成进度

## 四、TODO LIST
### 1、功能项
- [x] 自定义logger(2024.03.07)
- [ ] 任务监听完成
- [ ] 补充新活动从0到1的配置流程
- [ ] 管理后台
- [ ] 榜单支持多数据源、支持分数相同按照时间排序
- [ ] 抽奖奖励发放
- [ ] 任务奖励自动&手动领取
- [ ] 灰度名单配置
- [ ] ...

### 2、优化项
- [ ] 分布式锁
- [x] Kafka幂等