set names utf8;
DROP TABLE IF EXISTS `max_seq`;
CREATE TABLE `max_seq` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `secid` int(10) unsigned NOT NULL,
  `maxSeq` int(64) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
