import {
  Card,
  Breadcrumb,
  Form,
  Button,
  Radio,
  Input,
  Upload,
  Space,
  Select,
  message,
} from "antd";
import { Link } from "react-router-dom";
import { useStore } from "@/store";
import "./index.scss";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";
import { useLocation } from "react-router-dom";
import { http } from "@/utils";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const { Option } = Select;

const Publish = () => {
  const [form] = Form.useForm(); // 创建Form实例
  const { ChannelStore, userStore, ArticleStore } = useStore();
  const location = useLocation();
  const searchParams = new URLSearchParams(location.search);
  const [articleId, setArticleId] = useState();
  const id = searchParams.get("id");
  const navigate = useNavigate();

  useEffect(() => {
    if (id === undefined || id === null) {
      form.setFieldsValue({
        title: "",
        channel_id: undefined,
        content: "",
      });
    }
  }, [id]);

  useEffect(() => {
    setArticleId(id);
    console.log(articleId);
    if (articleId !== undefined && articleId !== null) {
      ArticleStore.getArticleById(articleId).then(() => {
        // 获取原来的文章内容
        const articleInfo = ArticleStore.article;
        console.log("已获取文章信息:");
        console.log(articleInfo);

        if (articleInfo) {
          console.log("设置表单");

          form.setFieldsValue({
            title: articleInfo.title,
            channel_id: articleInfo.ChannelID,
            content: articleInfo.content,
          });
        }
      });
    } else {
      form.setFieldsValue({
        title: "",
        channel_id: undefined,
        content: "",
      });
    }
  }, [articleId, ArticleStore]);

  async function onFinish(data) {
    if (articleId === undefined) {
      // 添加文章
      const res = await http.post("/article/add", {
        channel_id: data.channel_id,
        content: data.content,
        title: data.title,
        author: String(userStore.userInfo.ID),
      });

      if (res.data.code === 200) {
        message.success("插入成功");
        navigate("/", { replace: true });
      } else {
        message.error("插入失败");
      }
    } else {
      // 修改文章
      const res = await http.post(`/article/add`, {
        ID: articleId,
        channel_id: data.channel_id,
        content: data.content,
        title: data.title,
        author: String(userStore.userInfo.ID),
      });

      if (res.data.code === 200) {
        message.success("修改成功");
        navigate("/", { replace: true });
      } else {
        message.error("修改失败");
      }
    }
  }

  return (
    <div className="publish">
      <Card
        title={
          <Breadcrumb separator=">">
            <Breadcrumb.Item>
              <Link to="/">首页</Link>
            </Breadcrumb.Item>
            <Breadcrumb.Item>发布文章</Breadcrumb.Item>
          </Breadcrumb>
        }
      >
        <Form
          form={form} // 绑定表单实例
          labelCol={{ span: 4 }}
          wrapperCol={{ span: 16 }}
          initialValues={{ type: 1 }}
          onFinish={onFinish}
        >
          <Form.Item
            label="标题"
            name="title"
            rules={[{ required: true, message: "请输入文章标题" }]}
          >
            <Input placeholder="请输入文章标题" style={{ width: 400 }} />
          </Form.Item>
          <Form.Item
            label="频道"
            name="channel_id"
            rules={[{ required: true, message: "请选择文章频道" }]}
          >
            <Select placeholder="请选择文章频道" style={{ width: 400 }}>
              {ChannelStore.channelList.map((item) => (
                <Option value={item.ID}>{item.Name}</Option>
              ))}
            </Select>
          </Form.Item>

          {/* <Form.Item label="封面">
            <Form.Item name="type">
              <Radio.Group>
                <Radio value={1}>单图</Radio>
                <Radio value={3}>三图</Radio>
                <Radio value={0}>无图</Radio>
              </Radio.Group>
            </Form.Item>
            <Upload
              name="image"
              listType="picture-card"
              className="avatar-uploader"
              showUploadList
            >
              <div style={{ marginTop: 8 }}>
                <PlusOutlined />
              </div>
            </Upload>
          </Form.Item> */}
          <Form.Item
            label="内容"
            name="content"
            rules={[{ required: true, message: "请输入文章内容" }]}
          >
            <ReactQuill
              className="publish-quill"
              theme="snow"
              placeholder="请输入文章内容"
            />
          </Form.Item>

          <Form.Item wrapperCol={{ offset: 4 }}>
            <Space>
              <Button size="large" type="primary" htmlType="submit">
                发布文章
              </Button>
            </Space>
          </Form.Item>
        </Form>
      </Card>
    </div>
  );
};

export default Publish;
