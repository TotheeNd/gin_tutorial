/*
 Navicat Premium Data Transfer

 Source Server         : 188.131.140.15
 Source Server Type    : MariaDB
 Source Server Version : 100244 (10.2.44-MariaDB)
 Source Host           : 188.131.140.15:3306
 Source Schema         : ranking

 Target Server Type    : MariaDB
 Target Server Version : 100244 (10.2.44-MariaDB)
 File Encoding         : 65001

 Date: 19/07/2023 16:29:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '活动名称',
  `add_time` int(11) NULL DEFAULT 0 COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of activity
-- ----------------------------
INSERT INTO `activity` VALUES (1, '活动1', 1686105642);

-- ----------------------------
-- Table structure for player
-- ----------------------------
DROP TABLE IF EXISTS `player`;
CREATE TABLE `player`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `aid` int(11) NULL DEFAULT 0 COMMENT '所属排名活动',
  `ref` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '编号',
  `nickname` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `declaration` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '宣言',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
  `score` int(11) NULL DEFAULT 0 COMMENT '分数',
  `add_time` int(11) NULL DEFAULT 0 COMMENT '添加时间',
  `update_time` int(11) NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of player
-- ----------------------------
INSERT INTO `player` VALUES (1, 1, '0001', '下雨le', '很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！', '/images/11.png', 2, 1686105642, 0);
INSERT INTO `player` VALUES (2, 1, '0002', '栀夏', '很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！', '/images/12.png', 3, 1686105642, 0);
INSERT INTO `player` VALUES (3, 1, '0003', '闷油瓶', '很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！', '/images/13.png', 3, 1686105642, 0);
INSERT INTO `player` VALUES (4, 1, '0004', '喵喵兽', '很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！', '/images/14.png', 0, 1686105642, 0);
INSERT INTO `player` VALUES (5, 1, '0005', '弃殇', '很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！', '/images/15.png', 3, 1686105642, 0);
INSERT INTO `player` VALUES (6, 1, '0006', '冷巷', '很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！', '/images/16.png', 0, 1686105642, 0);
INSERT INTO `player` VALUES (7, 1, '0007', '栀子香气', '很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！', '/images/17.png', 7, 1686105642, 0);
INSERT INTO `player` VALUES (8, 1, '0008', '失与倦', '很开心和大家见面，我喜欢书法，画画，希望大家能够喜欢我！', '/images/18.png', 0, 1686105642, 0);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '密码',
  `add_time` int(11) NULL DEFAULT 0 COMMENT '添加时间',
  `update_time` int(11) NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'test', '96e79218965eb72c92a549dd5a330112', 1686130840, 1686130840);
INSERT INTO `user` VALUES (2, 'test111', '96e79218965eb72c92a549dd5a330112', 1686212844, 1686212844);
INSERT INTO `user` VALUES (3, 'test2', 'e3ceb5881a0a1fdaad01296d7554868d', 1686558256, 1686558256);
INSERT INTO `user` VALUES (5, 'liubei', '96e79218965eb72c92a549dd5a330112', 1688747874, 1688747874);
INSERT INTO `user` VALUES (6, 'zhja', '96e79218965eb72c92a549dd5a330112', 1689171940, 1689171940);
INSERT INTO `user` VALUES (7, 'zhja111', '96e79218965eb72c92a549dd5a330112', 1689175012, 1689175012);

-- ----------------------------
-- Table structure for vote
-- ----------------------------
DROP TABLE IF EXISTS `vote`;
CREATE TABLE `vote`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT 0 COMMENT '投票用户ID',
  `player_id` int(11) NULL DEFAULT 0 COMMENT '选手ID',
  `add_time` int(11) NULL DEFAULT 0 COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vote
-- ----------------------------
INSERT INTO `vote` VALUES (1, 2, 3, 1686279370);
INSERT INTO `vote` VALUES (2, 1, 3, 1686305248);
INSERT INTO `vote` VALUES (3, 3, 2, 1686617334);
INSERT INTO `vote` VALUES (4, 3, 7, 1686617960);
INSERT INTO `vote` VALUES (5, 5, 3, 1688751615);
INSERT INTO `vote` VALUES (6, 1, 7, 1688827908);
INSERT INTO `vote` VALUES (7, 1, 2, 1688827991);
INSERT INTO `vote` VALUES (8, 6, 1, 1689172370);
INSERT INTO `vote` VALUES (9, 6, 2, 1689174889);
INSERT INTO `vote` VALUES (10, 7, 1, 1689175040);
INSERT INTO `vote` VALUES (11, 7, 7, 1689175057);

SET FOREIGN_KEY_CHECKS = 1;
