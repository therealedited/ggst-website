/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE IF NOT EXISTS `ggst` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `ggst`;

CREATE TABLE IF NOT EXISTS `character` (
  `idChar` int(11) NOT NULL DEFAULT '-1',
  `name` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`idChar`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `character_move` (
  `idChar` int(11) DEFAULT NULL,
  `idMove` int(11) DEFAULT NULL,
  KEY `FK_idChar` (`idChar`),
  KEY `FK_idMove` (`idMove`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `move` (
  `idMove` int(11) NOT NULL DEFAULT '-1',
  `name` varchar(50) DEFAULT NULL,
  `input` varchar(50) DEFAULT NULL,
  `damage` varchar(50) DEFAULT NULL,
  `guard` varchar(50) DEFAULT NULL,
  `startup` varchar(50) DEFAULT NULL,
  `active` varchar(50) DEFAULT NULL,
  `recovery` varchar(50) DEFAULT NULL,
  `onBlock` varchar(50) DEFAULT NULL,
  `onHit` varchar(50) DEFAULT NULL,
  `invuln` varchar(50) DEFAULT NULL,
  `type` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`idMove`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;