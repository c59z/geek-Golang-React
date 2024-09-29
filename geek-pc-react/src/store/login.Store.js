import { makeAutoObservable } from "mobx";
import { http, setToken, getToken, removeToken } from "@/utils";

// login module
class LoginStore {
  token = getToken() || "";
  constructor() {
    makeAutoObservable(this); // 设置响应式
  }
  getToken = async ({ username, password }) => {
    const res = await http.post("/auth/login", {
      username,
      password,
    });
    console.log(res);
    this.token = res.data.token;
    // 存入ls
    setToken(this.token);
  };

  loginOut = () => {
    this.token = "";
    removeToken();
  };
}

export default LoginStore;
