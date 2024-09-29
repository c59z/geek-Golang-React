import { Link } from "react-router-dom";
import { Table, Tag, Space } from "antd";
import { EditOutlined, DeleteOutlined } from "@ant-design/icons";

import {
  Card,
  Breadcrumb,
  Form,
  Button,
  Radio,
  DatePicker,
  Select,
} from "antd";
import "moment/locale/zh-cn";
import locale from "antd/es/date-picker/locale/zh_CN";
import "./index.scss";
import { useEffect, useState } from "react";
import { http } from "@/utils";
import { useNavigate } from "react-router-dom";
import { useStore } from "@/store";
import { observer } from "mobx-react-lite";

const { Option } = Select;
const { RangePicker } = DatePicker;
const data = [
  {
    id: "8218",
    comment_count: 0,
    cover: {
      images: ["http://geek.itheima.net/resources/images/15.jpg"],
    },
    like_count: 0,
    pubdate: "2019-03-11 09:00:00",
    read_count: 2,
    status: 2,
    title: "wkwebview离线化加载h5资源解决方案",
  },
];

const Article = () => {
  const { ChannelStore } = useStore();

  // const [channelList, setChannelList] = useState([]);

  // const loadChannelList = async () => {
  //   http.get("/channel/get").then((res) => {
  //     setChannelList(res.data.channels);
  //   });
  // };

  // useEffect(() => {
  //   loadChannelList();
  // }, []);

  // 文章列表数据管理
  const [article, setArticleList] = useState({
    list: [],
    count: 0,
  });

  // 参数管理
  const [params, setParams] = useState({
    page: 1,
    per_page: 10,
  });

  // 发送接口请求
  useEffect(() => {
    async function fetchArticleList() {
      const res = await http.get("/article/get", { params });
      const { results, total_count } = res.data;
      setArticleList({
        list: results,
        count: total_count,
      });
    }
    fetchArticleList();
  }, [params]);

  const onChange = (page) => {
    setParams({
      ...params,
      page,
    });
  };

  const delArticle = async (data) => {
    await http.delete(`/article/del/${data.ID}`);
    setParams({
      ...params,
      page: 1,
    });
  };

  const navigate = useNavigate();
  const goPublish = (data) => {
    navigate(`/publish?id=${data.ID}`);
  };

  const columns = [
    // {
    //   title: "封面",
    //   dataIndex: "cover",
    //   width: 120,
    //   render: (cover) => {
    //     // return <img src={cover || img404} width={80} height={60} alt="" />;
    //   },
    // },
    {
      title: "标题",
      dataIndex: "title",
      width: 220,
    },
    {
      title: "状态",
      dataIndex: "status",
      render: (data) => <Tag color="green">审核通过</Tag>,
    },
    {
      title: "发布时间",
      dataIndex: "CreatedAt",
    },
    // {
    //   title: "阅读数",
    //   dataIndex: "read_count",
    // },
    // {
    //   title: "评论数",
    //   dataIndex: "comment_count",
    // },
    // {
    //   title: "点赞数",
    //   dataIndex: "like_count",
    // },
    {
      title: "操作",
      render: (data) => {
        return (
          <Space size="middle">
            <Button
              type="primary"
              shape="circle"
              onClick={() => goPublish(data)}
              icon={<EditOutlined />}
            />
            <Button
              type="primary"
              danger
              shape="circle"
              onClick={() => delArticle(data)}
              icon={<DeleteOutlined />}
            />
          </Space>
        );
      },
    },
  ];

  const onFinish = (values) => {
    console.log(values);
  };

  return (
    <div>
      <Card
        title={
          <Breadcrumb separator=">">
            <Breadcrumb.Item>
              <Link to="/home">首页</Link>
            </Breadcrumb.Item>
            <Breadcrumb.Item>内容管理</Breadcrumb.Item>
          </Breadcrumb>
        }
        style={{ marginBottom: 20 }}
      >
        <Form onFinish={onFinish} initialValues={{ status: null }}>
          <Form.Item label="状态" name="status">
            <Radio.Group>
              <Radio value={-1}>全部</Radio>
              <Radio value={0}>草稿</Radio>
              <Radio value={1}>待审核</Radio>
              <Radio value={2}>审核通过</Radio>
              <Radio value={3}>审核失败</Radio>
            </Radio.Group>
          </Form.Item>

          <Form.Item label="频道" name="channel_id">
            <Select placeholder="请选择文章频道" style={{ width: 120 }}>
              {ChannelStore.channelList.map((item) => (
                <Option value={item.ID}>{item.Name}</Option>
              ))}
            </Select>
          </Form.Item>

          <Form.Item label="日期" name="date">
            {/* 传入locale属性 控制中文显示*/}
            <RangePicker locale={locale}></RangePicker>
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ marginLeft: 80 }}>
              筛选
            </Button>
          </Form.Item>
        </Form>
      </Card>

      <Card title={`根据筛选条件共查询到 ${article.count} 条结果：`}>
        <Table
          rowKey="id"
          columns={columns}
          dataSource={article.list}
          pagination={{
            pageSize: params.per_page,
            total: article.count,
            onChange: onChange,
          }}
        />
      </Card>
    </div>
  );
};

export default Article;
