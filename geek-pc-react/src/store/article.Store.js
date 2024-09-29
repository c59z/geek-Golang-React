import { makeAutoObservable } from "mobx";
import { http } from "@/utils";

class ArticleStore {
  article = {};
  constructor() {
    makeAutoObservable(this);
  }

  getArticleById = async (id) => {
    const res = await http.get("/article/findOne/" + id);
    // console.log("获取文章信息：");
    // console.log(res.data.article);
    this.article = res.data.article;
  };
}

export default ArticleStore;
