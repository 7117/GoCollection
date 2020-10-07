/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : test1

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2020-10-07 13:41:58
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for movie_info
-- ----------------------------
DROP TABLE IF EXISTS `movie_info`;
CREATE TABLE `movie_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `movie_id` int(11) unsigned NOT NULL COMMENT '电影id',
  `movie_name` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影名称',
  `movie_pic` varchar(200) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影图片',
  `movie_director` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影导演',
  `movie_writer` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影编剧',
  `movie_country` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影产地',
  `movie_language` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影语言',
  `movie_main_character` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影主演',
  `movie_type` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影类型',
  `movie_on_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '电影上映时间',
  `movie_span` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影时长',
  `movie_grade` varchar(5) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '电影评分',
  `remark` varchar(500) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '备注',
  `_create_time` timestamp NULL DEFAULT NULL,
  `_modify_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `_status` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='电影信息表';

-- ----------------------------
-- Records of movie_info
-- ----------------------------
