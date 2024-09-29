/*
 Navicat Premium Data Transfer

 Source Server         : 本机mysql
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : geekdemo

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 29/09/2024 22:01:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `author` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `channel_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of articles
-- ----------------------------
INSERT INTO `articles` VALUES (1, '测试文章1', '内容啦啦啦啦啦', '1', '2024-09-26 19:24:28', '2024-09-27 22:37:30', NULL, 2);
INSERT INTO `articles` VALUES (2, '测试文章2', '啦啦啦啦啦', '1', '2024-09-26 19:24:41', '2024-09-27 22:37:31', NULL, 2);
INSERT INTO `articles` VALUES (3, '测试文章3', '啦啦啦啦啦', '1', '2024-09-26 19:24:52', '2024-09-27 22:37:33', NULL, 2);
INSERT INTO `articles` VALUES (4, '测试文章4', '4', '1', '2024-09-26 19:25:06', '2024-09-27 22:37:35', NULL, 2);
INSERT INTO `articles` VALUES (5, '测试文章5', '5', '1', '2024-09-26 19:25:15', '2024-09-27 22:37:36', '2024-09-26 21:04:07', 2);
INSERT INTO `articles` VALUES (6, '测试文章React001', '<p>123qwe</p>', '1', '2024-09-27 22:37:59', '2024-09-27 22:37:59', NULL, 2);
INSERT INTO `articles` VALUES (7, '测试文章React001', '<p>123qwe</p>', '1', '2024-09-27 22:38:42', '2024-09-27 22:38:42', NULL, 2);
INSERT INTO `articles` VALUES (8, '测试文章React003', '<p>啦啦啦啦啦啦啦</p>', '1', '2024-09-27 22:42:39', '2024-09-27 22:42:39', NULL, 2);
INSERT INTO `articles` VALUES (9, '测试文章React003', '<p>啦啦啦啦啦啦啦</p>', '1', '2024-09-27 22:42:48', '2024-09-27 22:42:48', NULL, 2);
INSERT INTO `articles` VALUES (10, 'messageTest001', '<p>123123</p>', '1', '2024-09-27 22:47:04', '2024-09-27 22:47:04', NULL, 2);
INSERT INTO `articles` VALUES (11, 'POSTMAN测试', 'POSTMAN测试', '1', '2024-09-28 21:10:54', '2024-09-28 21:12:59', NULL, 1);

-- ----------------------------
-- Table structure for channels
-- ----------------------------
DROP TABLE IF EXISTS `channels`;
CREATE TABLE `channels`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of channels
-- ----------------------------
INSERT INTO `channels` VALUES (1, '推荐');
INSERT INTO `channels` VALUES (2, 'IT');
INSERT INTO `channels` VALUES (3, '日语');

-- ----------------------------
-- Table structure for userinfo
-- ----------------------------
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `photo` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of userinfo
-- ----------------------------
INSERT INTO `userinfo` VALUES (1, 'Yuki', '$2a$10$dg9dI7E5IdAjs3AwzD1LK.esKbZQzPi59S2uKo8xoQ0hI0Z5jHbEu', 'bzh666@gmail.com', '2024-09-19 16:42:23', '2024-09-20 16:58:57', NULL, 'https://s2.loli.net/2023/09/11/XRI6tAPH7TxQKJr.jpg');

SET FOREIGN_KEY_CHECKS = 1;
