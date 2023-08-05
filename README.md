这个项目是基于开源框架pitaya进行二次开发的游戏服务器项目，pitaya连接: https://github.com/topfreegames/pitaya.git

<1> main是启动函数， 负责注册所有的game_module系统组件，并将这些组件交给pitaya管理。
<2> module_mgr是系统组件管理器，负责维护管理所有的系统组件。