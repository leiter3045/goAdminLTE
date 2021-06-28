/*
Navicat MySQL Data Transfer

Source Server         : go本地数据库
Source Server Version : 50650
Source Host           : 192.168.33.20:3306
Source Database       : gohoutai

Target Server Type    : MYSQL
Target Server Version : 50650
File Encoding         : 65001

Date: 2021-06-28 11:09:27
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for tp_access
-- ----------------------------
DROP TABLE IF EXISTS `tp_access`;
CREATE TABLE `tp_access` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` smallint(6) unsigned NOT NULL,
  `menu_id` smallint(6) NOT NULL COMMENT '菜单ID',
  `url` varchar(65) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_user_id` (`role_id`,`menu_id`) USING BTREE,
  KEY `role_id` (`role_id`) USING BTREE,
  KEY `node_id` (`url`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4 COMMENT='角色权限';

-- ----------------------------
-- Records of tp_access
-- ----------------------------
INSERT INTO `tp_access` VALUES ('42', '2', '3', '/role');
INSERT INTO `tp_access` VALUES ('43', '2', '9', '/menu');
INSERT INTO `tp_access` VALUES ('44', '2', '11', '/menu/edit');
INSERT INTO `tp_access` VALUES ('45', '2', '8', '/menu/del');
INSERT INTO `tp_access` VALUES ('46', '2', '7', '/menu/add');
INSERT INTO `tp_access` VALUES ('47', '2', '6', '/userlog');
INSERT INTO `tp_access` VALUES ('48', '2', '5', '/user');
INSERT INTO `tp_access` VALUES ('49', '2', '4', '/role');
INSERT INTO `tp_access` VALUES ('50', '1', '1', '/');
INSERT INTO `tp_access` VALUES ('51', '1', '13', '/config');
INSERT INTO `tp_access` VALUES ('52', '1', '14', '/config');

-- ----------------------------
-- Table structure for tp_menu
-- ----------------------------
DROP TABLE IF EXISTS `tp_menu`;
CREATE TABLE `tp_menu` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `pid` smallint(5) unsigned NOT NULL DEFAULT '0',
  `name` varchar(32) NOT NULL,
  `url` varchar(65) NOT NULL DEFAULT '' COMMENT '菜单URL',
  `icon` varchar(32) NOT NULL DEFAULT '',
  `sort` smallint(5) unsigned NOT NULL DEFAULT '50',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='管理菜单';

-- ----------------------------
-- Records of tp_menu
-- ----------------------------
INSERT INTO `tp_menu` VALUES ('1', '0', '系统首页', '/', 'fa fa-home', '50', '1');
INSERT INTO `tp_menu` VALUES ('3', '0', '权限管理', '/role', 'fa fa-group fa-fw', '50', '1');
INSERT INTO `tp_menu` VALUES ('4', '3', '角色管理', '/role', 'fa fa-circle-o', '50', '1');
INSERT INTO `tp_menu` VALUES ('5', '3', '员工管理', '/user', 'fa fa-circle-o', '50', '1');
INSERT INTO `tp_menu` VALUES ('6', '3', '操作日志', '/userlog', 'fa fa-circle-o', '50', '1');
INSERT INTO `tp_menu` VALUES ('7', '9', '添加', '/menu/add', 'fa fa-circle-o', '50', '1');
INSERT INTO `tp_menu` VALUES ('8', '9', '删除', '/menu/del', 'fa fa-circle-o', '50', '1');
INSERT INTO `tp_menu` VALUES ('9', '3', '菜单规则', '/menu', 'fa fa-circle-o', '50', '1');
INSERT INTO `tp_menu` VALUES ('11', '9', '编辑', '/menu/edit', 'fa fa-circle-o', '50', '1');
INSERT INTO `tp_menu` VALUES ('13', '0', '系统设置', '/config', 'fa fa-cog', '50', '1');
INSERT INTO `tp_menu` VALUES ('14', '13', '系统配置', '/config', 'fa fa-circle-o', '50', '1');
INSERT INTO `tp_menu` VALUES ('15', '5', '添加', '/user/add', 'fa fa-circle-o', '50', '1');

-- ----------------------------
-- Table structure for tp_role
-- ----------------------------
DROP TABLE IF EXISTS `tp_role`;
CREATE TABLE `tp_role` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `state` tinyint(1) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `status` (`state`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='管理角色';

-- ----------------------------
-- Records of tp_role
-- ----------------------------
INSERT INTO `tp_role` VALUES ('1', '审核组', '1');
INSERT INTO `tp_role` VALUES ('2', ' 维护组', '1');

-- ----------------------------
-- Table structure for tp_user
-- ----------------------------
DROP TABLE IF EXISTS `tp_user`;
CREATE TABLE `tp_user` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `role_id` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '员工角色',
  `username` varchar(32) NOT NULL,
  `userpass` char(32) NOT NULL,
  `last_login_ip` char(15) NOT NULL,
  `last_login_time` int(10) unsigned NOT NULL,
  `login_times` int(10) unsigned NOT NULL DEFAULT '0',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '0未禁止，1未正常',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='管理员';

-- ----------------------------
-- Records of tp_user
-- ----------------------------
INSERT INTO `tp_user` VALUES ('1', '1', 'admin', '1dec08486afbe30ab5d954b13a7abcf4', '', '0', '0', '1', '0');
INSERT INTO `tp_user` VALUES ('4', '2', 'ceshi003', '1dec08486afbe30ab5d954b13a7abcf4', '', '0', '0', '1', '1624505076');
INSERT INTO `tp_user` VALUES ('5', '1', 'ceshi004', '1dec08486afbe30ab5d954b13a7abcf4', '', '0', '0', '1', '0');

-- ----------------------------
-- Table structure for tp_user_log
-- ----------------------------
DROP TABLE IF EXISTS `tp_user_log`;
CREATE TABLE `tp_user_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT '管理员ID',
  `event_id` tinyint(3) unsigned NOT NULL COMMENT '操作事件ID',
  `module` varchar(30) NOT NULL DEFAULT '' COMMENT '所属模型',
  `relation_id` text COMMENT '关联信息ID，可能多个',
  `desc` varchar(127) NOT NULL COMMENT '操作描述',
  `ip` char(15) NOT NULL COMMENT '操作用户IP',
  `browser` varchar(500) NOT NULL DEFAULT '' COMMENT '浏览器',
  `cookie` varchar(32) NOT NULL DEFAULT '' COMMENT '电脑识别码',
  `add_time` int(10) unsigned NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`,`add_time`),
  KEY `event_id` (`event_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `mudule` (`module`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='管理员日志';

-- ----------------------------
-- Records of tp_user_log
-- ----------------------------
INSERT INTO `tp_user_log` VALUES ('1', '1', '2', '/login/ajaxout', '', '尝试员工退出登录成功', '192.168.33.46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36', 'd842d262b4895624fdcda933d42bedbb', '1624685323');
INSERT INTO `tp_user_log` VALUES ('2', '1', '2', '/login/ajaxout', '', '尝试员工退出登录成功', '192.168.33.46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36', 'd842d262b4895624fdcda933d42bedbb', '1624685420');
INSERT INTO `tp_user_log` VALUES ('3', '1', '2', '/login/ajaxout', '', '尝试员工退出登录成功', '192.168.33.46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36', 'd842d262b4895624fdcda933d42bedbb', '1624685478');
INSERT INTO `tp_user_log` VALUES ('4', '1', '2', '/login/ajaxout', '', '尝试员工退出登录成功', '192.168.33.46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36', 'd842d262b4895624fdcda933d42bedbb', '1624685522');
INSERT INTO `tp_user_log` VALUES ('5', '1', '2', '/login/ajaxout', '', '尝试员工退出登录成功', '192.168.33.46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36', 'd842d262b4895624fdcda933d42bedbb', '1624685612');
INSERT INTO `tp_user_log` VALUES ('6', '1', '2', '/login/ajaxout', '', '尝试员工退出登录成功', '192.168.33.46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36', 'd842d262b4895624fdcda933d42bedbb', '1624685664');
INSERT INTO `tp_user_log` VALUES ('7', '1', '2', '/login/ajaxout', '', '尝试员工退出登录成功', '192.168.33.46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36', 'd842d262b4895624fdcda933d42bedbb', '1624685692');
INSERT INTO `tp_user_log` VALUES ('8', '1', '4', '/menu/add', '15', '尝试添加菜单成功', '192.168.33.46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36', 'd842d262b4895624fdcda933d42bedbb', '1624692379');
INSERT INTO `tp_user_log` VALUES ('9', '1', '7', '/user/access', '5', '尝试编辑员工权限成功', '192.168.33.46', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36', 'd842d262b4895624fdcda933d42bedbb', '1624694763');
