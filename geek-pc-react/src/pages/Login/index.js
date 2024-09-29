import { Card, Form, Button, Input, Checkbox, message } from "antd";
import { replace, useNavigate } from "react-router-dom";
import logo from "@/assets/login.png";
import "./index.scss";

import { useStore } from "@/store";

const Login = () => {
  const { loginStore } = useStore();
  const navigate = useNavigate();
  const [messageApi] = message.useMessage();

  function onFinish(values) {
    console.log(values);
    try {
      loginStore.getToken({
        username: values.username,
        password: values.password,
      });
      // 登录成功后跳转到Layout页面
      messageApi.success("登录成功");
      navigate("/", { replace: true });
    } catch (e) {
      messageApi.error("登录失败");
    }
  }

  return (
    <div className="login">
      <Card className="login-container">
        <img className="login-logo" src={logo} alt="" />
        {/* 登录表单 */}
        <Form validateTrigger={["onBlur", "onChange"]} onFinish={onFinish}>
          <Form.Item
            name="username"
            rules={[
              {
                required: true,
                message: "Please input your username!",
              },
            ]}
          >
            <Input size="large" placeholder="请输入用户名" />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[
              {
                required: true,
                message: "Please input your password!",
              },
              {
                min: 6,
                max: 18,
                message: "密码长度不能小于6且不能大于18",
                validateTrigger: "onBlur",
              },
            ]}
          >
            <Input type="password" size="large" placeholder="请输入密码" />
          </Form.Item>
          <Form.Item name="remember" valuePropName="checked">
            <Checkbox className="login-checkbox-label">记住密码</Checkbox>
          </Form.Item>

          <Form.Item>
            {/* <!-- 渲染Button组件为submit按钮 --> */}
            <Button type="primary" htmlType="submit" size="large" block>
              登录
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};
export default Login;
