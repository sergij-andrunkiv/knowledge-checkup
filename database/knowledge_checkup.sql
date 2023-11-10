-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Nov 10, 2023 at 10:06 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `knowledge_checkup`
--

-- --------------------------------------------------------

--
-- Table structure for table `accounts`
--

CREATE TABLE `accounts` (
  `id` int(10) UNSIGNED NOT NULL,
  `last_name` varchar(50) NOT NULL,
  `first_name` varchar(50) NOT NULL,
  `middle_name` varchar(50) NOT NULL,
  `year_of_birth` varchar(4) NOT NULL,
  `nickname` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `approved` tinyint(1) NOT NULL,
  `gender` enum('Ч','Ж','N/A') DEFAULT NULL,
  `educational_institution` varchar(70) DEFAULT NULL,
  `teacher_status` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `accounts`
--

INSERT INTO `accounts` (`id`, `last_name`, `first_name`, `middle_name`, `year_of_birth`, `nickname`, `email`, `password`, `approved`, `gender`, `educational_institution`, `teacher_status`) VALUES
(1, 'test', 'test', 'test', 'test', 'test', 'test', 'test', 1, 'Ч', 'test', 0),
(2, '123', '123', '123', '123', '123', '123', '123', 1, 'Ч', ' ', 0),
(3, '1', '1', '1', '1', '1', '1', '1', 1, NULL, '', 0),
(4, '2000', '2000', '2000', '2000', '2000', '2000', '2000', 1, NULL, NULL, 0),
(5, 'Andrunkiv', 'Sergiy', 'Romanovich', '2003', 'sergiy_2000', 'sergij.andrunkiv@gmail.com', 'qwerty123', 1, 'N/A', 'N/A', 1),
(6, 'test123', 'test123', 'test123', '2000', 'testik', 'test123', 'test123', 1, 'N/A', 'N/A', 0);

-- --------------------------------------------------------

--
-- Table structure for table `answers`
--

CREATE TABLE `answers` (
  `id_a` int(5) UNSIGNED NOT NULL,
  `id_q` int(5) UNSIGNED NOT NULL,
  `text` varchar(255) NOT NULL,
  `is_correct` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `answers`
--

INSERT INTO `answers` (`id_a`, `id_q`, `text`, `is_correct`) VALUES
(1, 54, '1', 0),
(2, 54, '2', 1),
(3, 54, '3', 0),
(4, 55, '4', 1),
(5, 55, '5', 1),
(6, 55, '6', 0),
(7, 55, '7', 1),
(8, 56, 'x', 0),
(9, 56, 'z', 1),
(10, 57, '1', 0),
(11, 57, '2', 0),
(12, 57, '3', 1),
(13, 58, '1', 0),
(14, 58, '2', 0),
(15, 58, '3', 1),
(16, 59, '4', 1),
(17, 59, '5', 0),
(18, 59, '6', 1),
(19, 60, '1111', 0),
(20, 60, '2222', 0),
(21, 60, '33333', 1),
(22, 61, '444444', 1),
(23, 61, '55555', 0),
(24, 61, '66666', 1),
(25, 62, '1', 0),
(26, 62, '2', 0),
(27, 62, '3', 1),
(28, 63, '1', 0),
(29, 63, '2', 0),
(30, 63, '3', 1),
(31, 64, '1', 0),
(32, 64, '2', 0),
(33, 64, '3', 1),
(34, 65, '1', 1),
(35, 106, '1', 0),
(36, 106, '2', 0),
(37, 106, '3', 1),
(38, 107, '4', 0),
(39, 107, '5', 1),
(40, 107, '6', 0),
(41, 107, '7', 1),
(42, 108, '1', 0),
(43, 108, '2', 1),
(44, 108, '3', 0),
(45, 109, '1', 0),
(46, 109, '2', 1),
(47, 109, '3', 0),
(48, 110, '1', 0),
(49, 110, '2', 1),
(50, 110, '3', 0),
(51, 111, '1', 0),
(52, 112, '1', 0),
(53, 113, '1', 0),
(54, 114, '1', 0),
(55, 114, '2', 1),
(56, 114, '3', 0),
(57, 115, '7', 1),
(58, 115, '8', 0),
(59, 115, '9', 0),
(60, 115, '0', 1),
(61, 116, '1', 0),
(62, 116, '2', 1),
(63, 116, '3', 0),
(64, 117, '7', 1),
(65, 117, '8', 0),
(66, 117, '9', 0),
(67, 117, '0', 1),
(68, 118, '1', 0),
(69, 118, '3', 1),
(70, 119, 'ddd', 0),
(71, 119, 'dd', 1),
(72, 119, 'd', 0),
(73, 120, '1', 1),
(74, 120, '2', 0),
(75, 121, '1', 0),
(76, 121, '2', 1),
(77, 121, '3', 0),
(78, 123, 'Ans1-q1', 0),
(79, 123, 'Ans2-q1', 0),
(80, 123, 'Ans3-q1', 1),
(81, 124, 'Ans1-q2', 0),
(82, 124, 'Ans2-q2', 1),
(83, 125, 'Ans1-q3', 1),
(84, 125, 'Ans2-q3', 0),
(85, 125, 'Ans3-q3', 1),
(86, 125, 'Ans4-q3', 0);

-- --------------------------------------------------------

--
-- Table structure for table `marks`
--

CREATE TABLE `marks` (
  `id_m` int(5) UNSIGNED NOT NULL,
  `user` int(5) UNSIGNED NOT NULL,
  `test` int(5) UNSIGNED NOT NULL,
  `mark` int(5) UNSIGNED NOT NULL,
  `count_of_correct_answers` int(5) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf32 COLLATE=utf32_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `questions`
--

CREATE TABLE `questions` (
  `id_q` int(5) UNSIGNED NOT NULL,
  `text` varchar(255) NOT NULL,
  `id_creator` int(11) UNSIGNED NOT NULL,
  `type` enum('single','multiple') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `questions`
--

INSERT INTO `questions` (`id_q`, `text`, `id_creator`, `type`) VALUES
(1, '123', 5, 'single'),
(2, '1234', 5, 'multiple'),
(3, '123', 5, 'single'),
(4, '1234', 5, 'multiple'),
(5, '123', 5, 'single'),
(6, '1234', 5, 'multiple'),
(7, '123', 5, 'single'),
(8, '1234', 5, 'multiple'),
(9, '123', 5, 'single'),
(10, '1234', 5, 'multiple'),
(11, '123', 5, 'single'),
(12, '1234', 5, 'multiple'),
(13, '123', 5, 'single'),
(14, '1234', 5, 'multiple'),
(15, '123', 5, 'single'),
(16, '1234', 5, 'multiple'),
(17, '123', 5, 'single'),
(18, '1234', 5, 'multiple'),
(19, '123', 5, 'single'),
(20, '1234', 5, 'multiple'),
(21, '123', 5, 'single'),
(22, '1234', 5, 'multiple'),
(23, '123', 5, 'single'),
(24, '1234', 5, 'multiple'),
(25, '123', 5, 'single'),
(26, '1234', 5, 'multiple'),
(27, '123', 5, 'single'),
(28, '1234', 5, 'multiple'),
(29, '123213131232123', 5, 'single'),
(50, '123', 5, 'single'),
(51, '123', 5, 'single'),
(52, '1234', 5, 'multiple'),
(53, '123213131232123', 5, 'single'),
(54, '123', 5, 'single'),
(55, '1234', 5, 'multiple'),
(56, '123213131232123', 5, 'single'),
(57, '123', 5, 'single'),
(58, '123', 5, 'single'),
(59, '456', 5, 'multiple'),
(60, '123', 5, 'single'),
(61, '456', 5, 'multiple'),
(62, '123', 5, 'single'),
(63, '123', 5, 'single'),
(64, '123', 5, 'single'),
(65, '123', 5, 'single'),
(66, '', 5, 'single'),
(67, '', 5, 'single'),
(68, '', 5, 'single'),
(69, '', 5, 'single'),
(70, '', 5, 'single'),
(71, '', 5, 'single'),
(72, '', 5, 'single'),
(73, '', 5, 'single'),
(74, '', 5, 'single'),
(75, '', 5, 'single'),
(76, '', 5, 'single'),
(77, '', 5, 'single'),
(78, '', 5, 'single'),
(79, '', 5, 'single'),
(80, '', 5, 'single'),
(81, '', 5, 'single'),
(82, '', 5, 'single'),
(83, '', 5, 'single'),
(84, '', 5, 'single'),
(85, '', 5, 'single'),
(86, '', 5, 'single'),
(87, '', 5, 'single'),
(88, '', 5, 'single'),
(89, '', 5, 'single'),
(90, '', 5, 'single'),
(91, '', 5, 'single'),
(92, '', 5, 'single'),
(93, '', 5, 'single'),
(94, '', 5, 'single'),
(95, '', 5, 'single'),
(96, '', 5, 'single'),
(97, '', 5, 'single'),
(98, '', 5, 'single'),
(99, '', 5, 'single'),
(100, '', 5, 'single'),
(101, '', 5, 'single'),
(102, '', 5, 'single'),
(103, '', 5, 'single'),
(104, '', 5, 'single'),
(105, '', 5, 'single'),
(106, 'Питання 1????', 5, 'single'),
(107, 'Питання 2????', 5, 'multiple'),
(108, '123', 5, 'single'),
(109, '123', 5, 'single'),
(110, '123', 5, 'single'),
(111, '123???', 5, 'single'),
(112, '123???', 5, 'single'),
(113, '123???', 5, 'single'),
(114, '123???', 5, 'single'),
(115, '1234????', 5, 'multiple'),
(116, '123???', 5, 'single'),
(117, '1234????', 5, 'multiple'),
(118, '123', 5, 'single'),
(119, 'gggg', 5, 'single'),
(120, '123', 5, 'single'),
(121, '123', 5, 'single'),
(122, '', 5, 'single'),
(123, 'Ques1', 5, 'single'),
(124, 'Ques2', 5, 'single'),
(125, 'Ques3', 5, 'multiple');

-- --------------------------------------------------------

--
-- Table structure for table `tests`
--

CREATE TABLE `tests` (
  `id_t` int(5) UNSIGNED NOT NULL,
  `title` varchar(255) NOT NULL,
  `count_of_questions` int(5) UNSIGNED NOT NULL,
  `max_mark` int(5) UNSIGNED NOT NULL,
  `tags` varchar(255) NOT NULL,
  `creator` int(5) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf32 COLLATE=utf32_general_ci;

--
-- Dumping data for table `tests`
--

INSERT INTO `tests` (`id_t`, `title`, `count_of_questions`, `max_mark`, `tags`, `creator`) VALUES
(1, 'My test', 2, 20, '#mytest', 5),
(2, '123', 1, 20, 'tag1', 5),
(3, '123', 1, 20, 'tag1', 5),
(4, 'My_Test', 2, 20, '#mytest', 5),
(5, 'My_Test', 2, 20, '#mytest', 5),
(6, 'My_Test', 2, 20, '#mytest', 5),
(7, 'My_Test', 2, 20, '#mytest', 5),
(8, 'My_Test', 2, 20, '#mytest', 5),
(9, '123', 1, 20, '123213123123', 5),
(10, ',fdsmfmsdfnksdnflndsf', 1, 15, 'dsfsdfsdfsdf', 5),
(11, '123', 1, 20, 'tag', 5),
(12, '123', 2, 0, '', 5),
(13, 'My New Test', 3, 30, '#mynewtest', 5);

-- --------------------------------------------------------

--
-- Table structure for table `tests_structure`
--

CREATE TABLE `tests_structure` (
  `id_t` int(5) UNSIGNED NOT NULL,
  `id_q` int(5) UNSIGNED NOT NULL,
  `id_a` int(5) UNSIGNED NOT NULL,
  `id_creator` int(5) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf16 COLLATE=utf16_general_ci;

--
-- Dumping data for table `tests_structure`
--

INSERT INTO `tests_structure` (`id_t`, `id_q`, `id_a`, `id_creator`) VALUES
(7, 114, 54, 5),
(7, 114, 55, 5),
(7, 114, 56, 5),
(7, 115, 57, 5),
(7, 115, 58, 5),
(7, 115, 59, 5),
(7, 115, 60, 5),
(8, 116, 61, 5),
(8, 116, 62, 5),
(8, 116, 63, 5),
(8, 117, 64, 5),
(8, 117, 65, 5),
(8, 117, 66, 5),
(8, 117, 67, 5),
(9, 118, 68, 5),
(9, 118, 69, 5),
(10, 119, 70, 5),
(10, 119, 71, 5),
(10, 119, 72, 5),
(11, 120, 73, 5),
(11, 120, 74, 5),
(12, 121, 75, 5),
(12, 121, 76, 5),
(12, 121, 77, 5),
(13, 123, 78, 5),
(13, 123, 79, 5),
(13, 123, 80, 5),
(13, 124, 81, 5),
(13, 124, 82, 5),
(13, 125, 83, 5),
(13, 125, 84, 5),
(13, 125, 85, 5),
(13, 125, 86, 5);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `accounts`
--
ALTER TABLE `accounts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `answers`
--
ALTER TABLE `answers`
  ADD PRIMARY KEY (`id_a`),
  ADD KEY `Test_2` (`id_q`);

--
-- Indexes for table `marks`
--
ALTER TABLE `marks`
  ADD PRIMARY KEY (`id_m`),
  ADD KEY `Test_8` (`user`),
  ADD KEY `Test_9` (`test`);

--
-- Indexes for table `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`id_q`),
  ADD KEY `Test` (`id_creator`);

--
-- Indexes for table `tests`
--
ALTER TABLE `tests`
  ADD PRIMARY KEY (`id_t`),
  ADD KEY `Test_7` (`creator`);

--
-- Indexes for table `tests_structure`
--
ALTER TABLE `tests_structure`
  ADD PRIMARY KEY (`id_t`,`id_q`,`id_a`,`id_creator`),
  ADD KEY `Test_4` (`id_q`),
  ADD KEY `Test_5` (`id_a`),
  ADD KEY `Test_6` (`id_creator`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `accounts`
--
ALTER TABLE `accounts`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `answers`
--
ALTER TABLE `answers`
  MODIFY `id_a` int(5) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=87;

--
-- AUTO_INCREMENT for table `marks`
--
ALTER TABLE `marks`
  MODIFY `id_m` int(5) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `questions`
--
ALTER TABLE `questions`
  MODIFY `id_q` int(5) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=126;

--
-- AUTO_INCREMENT for table `tests`
--
ALTER TABLE `tests`
  MODIFY `id_t` int(5) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `answers`
--
ALTER TABLE `answers`
  ADD CONSTRAINT `Test_2` FOREIGN KEY (`id_q`) REFERENCES `questions` (`id_q`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `marks`
--
ALTER TABLE `marks`
  ADD CONSTRAINT `Test_8` FOREIGN KEY (`user`) REFERENCES `accounts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Test_9` FOREIGN KEY (`test`) REFERENCES `tests` (`id_t`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `questions`
--
ALTER TABLE `questions`
  ADD CONSTRAINT `Test` FOREIGN KEY (`id_creator`) REFERENCES `accounts` (`id`);

--
-- Constraints for table `tests`
--
ALTER TABLE `tests`
  ADD CONSTRAINT `Test_7` FOREIGN KEY (`creator`) REFERENCES `accounts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `tests_structure`
--
ALTER TABLE `tests_structure`
  ADD CONSTRAINT `Test_3` FOREIGN KEY (`id_t`) REFERENCES `tests` (`id_t`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Test_4` FOREIGN KEY (`id_q`) REFERENCES `questions` (`id_q`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Test_5` FOREIGN KEY (`id_a`) REFERENCES `answers` (`id_a`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Test_6` FOREIGN KEY (`id_creator`) REFERENCES `accounts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
