这个项目是基于开源框架pitaya进行二次开发的游戏服务器项目，pitaya连接: https://github.com/topfreegames/pitaya.git

在pitaya框架里，component作为对外提供服务的组件，在本项目中，game_service作为对component的包装，srvice_mgr.go实现对所有game_service的发现和注册，并能根据配置摒弃掉哪些不需要启动的对外服务。

在pitaya框架里，module作为系统的组件进行开发，在本项目中，game_modules作为对module的系统化实现，每个组件之间相互独立不影响，组件可以处理逻辑，也可以返回逻辑
返回逻辑时是需要指明立即返回和延迟返回，立即返回策略会将任务提升到module执行队列的前列，延迟返回则会处于低优先级，不需要返回的任务会排在队列的最末端。
