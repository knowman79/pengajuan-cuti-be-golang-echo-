/*
SQLyog Enterprise - MySQL GUI v7.02 
MySQL - 5.5.5-10.1.19-MariaDB : Database - db_makanan
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/`db_makanan` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `db_makanan`;

/*Table structure for table `tb_makanan` */

DROP TABLE IF EXISTS `tb_makanan`;

CREATE TABLE `tb_makanan` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `kode_makanan` char(10) DEFAULT NULL,
  `nama_makanan` varchar(200) DEFAULT NULL,
  `harga` double DEFAULT NULL,
  `stok` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `expired_at` date DEFAULT NULL,
  `update_by` int(11) DEFAULT NULL,
  `create_by` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `tb_makanan` */

insert  into `tb_makanan`(`id`,`kode_makanan`,`nama_makanan`,`harga`,`stok`,`created_at`,`updated_at`,`expired_at`,`update_by`,`create_by`) values (1,'K001','Kue Gajol',800,100,NULL,NULL,NULL,1,1),(2,'K002','Kue Lapis',1000,50,NULL,NULL,NULL,1,1);

/*Table structure for table `tb_role` */

DROP TABLE IF EXISTS `tb_role`;

CREATE TABLE `tb_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `tb_role` */

insert  into `tb_role`(`id`,`role`) values (1,'kasir'),(2,'admin');

/*Table structure for table `tb_transaksi` */

DROP TABLE IF EXISTS `tb_transaksi`;

CREATE TABLE `tb_transaksi` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_user` int(11) DEFAULT NULL,
  `kode_makanan` char(20) DEFAULT NULL,
  `jumlah_makanan` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `jumlah_uang` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

/*Data for the table `tb_transaksi` */

/*Table structure for table `tb_user` */

DROP TABLE IF EXISTS `tb_user`;

CREATE TABLE `tb_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nama_user` varchar(200) DEFAULT NULL,
  `username` varchar(200) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  `id_role` int(11) DEFAULT NULL,
  `phone` varchar(15) DEFAULT NULL,
  `email` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `tb_user` */

insert  into `tb_user`(`id`,`nama_user`,`username`,`password`,`id_role`,`phone`,`email`) values (1,'Nurikhsan','ikhsanhikari','ikhsanhikari29',2,'089656541471','ikhsanhikari29@gmail.com'),(2,'Dyah Nuraeni','Dyah09','dyah-potter',1,'088222500997','dyah_nuraeni@gmail.com');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
