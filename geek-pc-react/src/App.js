import Login from "@/pages/Login";
import Layout from "@/pages/Layout";

import { BrowserRouter, Routes, Route } from "react-router-dom";
// import { AuthComponent } from "./components/AuthComponents";
import "antd/dist/antd";
import "./index.css";
import "./app.css";
import Article from "./pages/Article";
import Home from "./pages/Home";
import Publish from "./pages/Publish";
import { HistoryRouter, history } from "./utils/history";
import { observer } from "mobx-react-lite";

function App() {
  return (
    <HistoryRouter history={history}>
      <div className="App">
        <Routes>
          {/* <Route
            path="/"
            element={
              <AuthComponent>
                <Layout></Layout>
              </AuthComponent>
            }
          ></Route> */}
          <Route path="/" element={<Layout></Layout>}>
            <Route index element={<Home></Home>}></Route>
            <Route path="/article" element={<Article></Article>}></Route>
            <Route path="/publish" element={<Publish></Publish>}></Route>
          </Route>
          <Route path="/login" element={<Login></Login>}></Route>
        </Routes>
      </div>
    </HistoryRouter>
  );
}

export default observer(App);
