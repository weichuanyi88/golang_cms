/*
Navicat MySQL Data Transfer

Source Server         : 本机mysql
Source Server Version : 50713
Source Host           : localhost:3306
Source Database       : d_go_cms

Target Server Type    : MYSQL
Target Server Version : 50713
File Encoding         : 65001

Date: 2016-08-22 16:26:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for t_article
-- ----------------------------
DROP TABLE IF EXISTS `t_article`;
CREATE TABLE `t_article` (
  `id` int(11) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `content` longtext,
  `addtime` int(11) DEFAULT NULL,
  `typeId` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_article
-- ----------------------------
INSERT INTO `t_article` VALUES ('1', '企业简介', '介绍公司。。。。。', null, '1');
INSERT INTO `t_article` VALUES ('2', '团队介绍', '团队介绍。。。', null, '1');
INSERT INTO `t_article` VALUES ('3', '企业理念', '企业理念。。。', null, '1');

-- ----------------------------
-- Table structure for t_article_type
-- ----------------------------
DROP TABLE IF EXISTS `t_article_type`;
CREATE TABLE `t_article_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `type` int(11) DEFAULT NULL COMMENT '1:关于我们,2:业务合作',
  `orderIndex` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_article_type
-- ----------------------------
INSERT INTO `t_article_type` VALUES ('1', '公司介绍', '0', '1');
INSERT INTO `t_article_type` VALUES ('2', '业务范围', '0', '2');
INSERT INTO `t_article_type` VALUES ('3', '服务流程', '0', '3');
INSERT INTO `t_article_type` VALUES ('4', '联系我们', '0', '4');
INSERT INTO `t_article_type` VALUES ('5', '线切割', '1', '1');
INSERT INTO `t_article_type` VALUES ('6', '五金冲压', '1', '2');
INSERT INTO `t_article_type` VALUES ('7', '专业设计', '1', '3');
INSERT INTO `t_article_type` VALUES ('8', '其他业务', '1', '4');

-- ----------------------------
-- Table structure for t_news
-- ----------------------------
DROP TABLE IF EXISTS `t_news`;
CREATE TABLE `t_news` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `content` longtext NOT NULL,
  `typeId` int(11) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `addTime` int(11) DEFAULT NULL,
  `parentTypeId` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_news
-- ----------------------------
INSERT INTO `t_news` VALUES ('1', '新闻标题一', '1content', '1', 'a.jpg', '1469416271', '0');
INSERT INTO `t_news` VALUES ('2', '新闻标题二', '2content', '1', 'a.jpg', '1469416271', '0');
INSERT INTO `t_news` VALUES ('3', '1469416271', '1469416271', '1', '1469416271', '1469416271', '0');
INSERT INTO `t_news` VALUES ('4', '达到司法等发达省份阿萨德', '打发的萨芬啊', '2', null, '1469416271', '0');
INSERT INTO `t_news` VALUES ('5', '的发生大幅ad司法的阿萨德', '发大幅爱的', '2', null, '1469416271', '0');
INSERT INTO `t_news` VALUES ('6', '的发射点发的发射点发地方', '的发射的地方', '2', null, '1469416271', '0');
INSERT INTO `t_news` VALUES ('7', '温热亲热亲热弃我而去', '去诶日起而且', '3', null, '1469416271', '0');
INSERT INTO `t_news` VALUES ('8', '的发射的发射的发射', '大师法第三方啊', '3', null, '1469416271', '0');
INSERT INTO `t_news` VALUES ('9', '的发射的发射的发射的发射的发射的发射点发', '打发打发', '3', null, '1469416271', '0');
INSERT INTO `t_news` VALUES ('10', '大富大沙发的', '打法的', '4', null, '1469416271', '0');
INSERT INTO `t_news` VALUES ('11', '萨的发射的发射点发', '新闻内容内容内容', '4', null, '1469416271', '0');
INSERT INTO `t_news` VALUES ('12', '大富大沙发大沙发', '打发打发', '4', null, '1469416271', '0');

-- ----------------------------
-- Table structure for t_news_type
-- ----------------------------
DROP TABLE IF EXISTS `t_news_type`;
CREATE TABLE `t_news_type` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `orderIndex` varchar(255) DEFAULT NULL,
  `parentId` int(5) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_news_type
-- ----------------------------
INSERT INTO `t_news_type` VALUES ('1', '公司新闻', '1', '0');
INSERT INTO `t_news_type` VALUES ('2', '客户公告', '2', '0');
INSERT INTO `t_news_type` VALUES ('3', '行业新闻', '3', '0');
INSERT INTO `t_news_type` VALUES ('4', '技术资讯', '4', '0');
INSERT INTO `t_news_type` VALUES ('5', '设备资讯', '1', '4');
INSERT INTO `t_news_type` VALUES ('6', '技术资讯', '2', '4');

-- ----------------------------
-- Table structure for t_product
-- ----------------------------
DROP TABLE IF EXISTS `t_product`;
CREATE TABLE `t_product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `typeId` int(11) NOT NULL COMMENT '产品类别',
  `image` varchar(255) DEFAULT NULL,
  `title` varchar(255) NOT NULL,
  `bianHao` varchar(50) DEFAULT NULL COMMENT '编号',
  `guiGe` varchar(50) DEFAULT NULL COMMENT '规格',
  `content` longtext,
  `addTime` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_product
-- ----------------------------
INSERT INTO `t_product` VALUES ('1', '1', '../static/images/11.jpg', '快走丝案例一', 'A-1001', '200*300', '案例详情', null);
INSERT INTO `t_product` VALUES ('2', '1', '../static/images/10.jpg', '快走丝案例二', 'A-1002', '400*500', '按理详情', null);
INSERT INTO `t_product` VALUES ('3', '1', '../static/images/10.jpg', '快走丝按理三', 'A-1003', '230*230', 'KJDLASKJDLA', null);
INSERT INTO `t_product` VALUES ('4', '1', '../static/images/8.jpg', '快走丝', 'B-1004', '12*12', 'AKDJLFKAJSDLK', null);
INSERT INTO `t_product` VALUES ('5', '1', '../static/images/8.jpg', 'dafsdadafdad', 'ad', 'aa', 'dadfas', null);
INSERT INTO `t_product` VALUES ('6', '1', '../static/images/8.jpg', 'adsa', 'ad', 'ad', 'sadf', null);
INSERT INTO `t_product` VALUES ('7', '1', '../static/images/8.jpg', 'wqerq', 'as', 'as', 'sa', null);
INSERT INTO `t_product` VALUES ('8', '1', '../static/images/8.jpg', 'fasd', 'as', 'da', 'sa', null);
INSERT INTO `t_product` VALUES ('9', '1', '../static/images/8.jpg', 'asdf', 'asd', 'fasd', 'asdf', null);
INSERT INTO `t_product` VALUES ('10', '1', '../static/images/8.jpg', 'fads', 'dad', 'dsasd', 'dfa', null);
INSERT INTO `t_product` VALUES ('11', '1', '../static/images/8.jpg', '测试数据', 'aa', 'aa', 'aa', null);
INSERT INTO `t_product` VALUES ('12', '1', '../static/images/8.jpg', '测试2', 'aa', 'aa', 'aa', null);

-- ----------------------------
-- Table structure for t_product_type
-- ----------------------------
DROP TABLE IF EXISTS `t_product_type`;
CREATE TABLE `t_product_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `orderIndex` int(11) DEFAULT NULL,
  `type` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_product_type
-- ----------------------------
INSERT INTO `t_product_type` VALUES ('1', '快走丝', '1', '1');
INSERT INTO `t_product_type` VALUES ('2', '中走丝', '2', '1');
INSERT INTO `t_product_type` VALUES ('3', '慢走丝', '3', '1');

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(255) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `nickname` varchar(255) NOT NULL COMMENT '昵称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of t_user
-- ----------------------------
INSERT INTO `t_user` VALUES ('1', 'weichuanyi', '123456', '韦传毅');
INSERT INTO `t_user` VALUES ('2', 'wcy02username', '1234567', 'wcy02');
INSERT INTO `t_user` VALUES ('3', 'wcy02username', '1234567', 'wcy03');
INSERT INTO `t_user` VALUES ('4', 'wcy02username', '1234567', 'wcy04');
INSERT INTO `t_user` VALUES ('5', 'wcy02username', '1234567', 'wcy05');
