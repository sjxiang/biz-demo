

# 接口说明


## 1. 基础接口

抖音最基础的功能实现，支持所有用户刷抖音视频，同时允许用户注册账户，发布自己拍摄的视频，发布后能够被其他人刷到。


### /douyin/feed/ - 视频流接口

不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多 30 个。

接口类型 GET

接口定义


### /douyin/user/register/ - 用户注册接口



### /douyin/user/login/ - 用户登录接口



### /douyin/user/ - 用户信息接口



### /douyin/publish/action/ - 视频投稿



### /douyin/publish/list/ - 视频发布列表




## 2. 互动接口

每个登录用户支持点赞，同时维护用户自己的点赞视频列表，在个人信息页中查看。

所有用户能够查看视频的评论列表，但是只有登录用户能够对视频进行评论。


### /douyin/favorite/action/ - 赞操作

### /douyin/favorite/list/ - 喜欢列表

### /douyin/comment/action/ - 评论操作

### /douyin/comment/action/ - 视频评论列表


## 3. 社交接口

实现用户之间的关注关系维护，登录用户能够关注或取关其他用户，同时自己能够看到自己关注过的所有用户列表，以及所有关注自己的用户列表。

### /douyin/relation/action/ - 关系操作

### /douyin/relation/follow/list/ - 用户关注列表

### /douyin/relation/follower/list/ - 用户粉丝列表

### /douyin/relation/friend/list/ - 用户好友列表


## 4. 消息接口

客户端通过定时轮询服务端接口查询消息记录

### /douyin/message/chat/ - 聊天记录

### /douyin/message/action/ - 消息操作


