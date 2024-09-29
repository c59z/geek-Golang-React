// 把所有的模块做统一处理
// 导出一个统一的方法 useStore

import LoginStore from "./login.Store";
import UserStore from "./userStore";
import ChannelStore from "./channel.Store";
import ArticleStore from "./article.Store";
import React from "react";

class RootStore {
  constructor() {
    this.loginStore = new LoginStore();
    this.userStore = new UserStore();
    this.ChannelStore = new ChannelStore();
    this.ArticleStore = new ArticleStore();
  }
}

// 实例化根
// 导出useStore context
const rootStore = new RootStore();
const context = React.createContext(rootStore);

const useStore = () => React.useContext(context);

export { useStore };
