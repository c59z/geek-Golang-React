import { makeAutoObservable } from "mobx";
import { http } from "@/utils";

class UserStore {
  userInfo = {};

  constructor() {
    makeAutoObservable(this);
  }

  async getUserInfo() {
    const res = await http.get("/auth/profile");
    // console.log(res);

    this.userInfo = res.data;
  }
}

export default UserStore;
