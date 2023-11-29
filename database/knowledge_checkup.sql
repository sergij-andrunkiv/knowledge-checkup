-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Nov 29, 2023 at 11:26 PM
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
(5, 'Andrunkiv', 'Sergiy', 'Romanovich', '2003', 'sergiy_2000', 'sergij.andrunkiv@gmail.com', 'qwerty', 1, 'Ч', 'dfdfdf', 1),
(6, 'test123', 'test123', 'test123', '2000', 'test_123', 'test123', '', 1, 'N/A', 'N/A', 0),
(10, 'Андруньків', 'Сергій', 'Романович', '2003', 'sergiy_andr', 'sergiy.andrunkiv@gmail.com', 'test123456', 1, 'Ч', '123456', 1),
(11, 'test123', 'test123', 'test123', '2000', 'test123', 'copini2940@eachart.com', 'test12345', 1, 'N/A', 'N/A', 0);

-- --------------------------------------------------------

--
-- Table structure for table `answers`
--

CREATE TABLE `answers` (
  `id_a` int(10) UNSIGNED NOT NULL,
  `id_q` int(10) UNSIGNED NOT NULL,
  `text` varchar(255) NOT NULL,
  `is_correct` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `answers`
--

INSERT INTO `answers` (`id_a`, `id_q`, `text`, `is_correct`) VALUES
(133, 144, '1 1', 1),
(134, 144, '1 2', 0),
(135, 144, '1 3', 0),
(136, 145, '2 1', 1),
(137, 145, '2 2', 0),
(138, 145, '2 3', 1),
(139, 146, '3 1', 0),
(140, 146, '3 2', 0),
(141, 146, '3 3', 1),
(142, 146, '3 4', 0),
(144, 147, '4 1', 0),
(145, 147, '4 2', 1),
(146, 148, '1 1', 0),
(147, 148, '1 2', 1),
(148, 149, '2 1', 0),
(149, 149, '2 2', 1),
(156, 154, 'ans1', 0),
(157, 154, 'ans2', 1),
(158, 155, '1', 1),
(159, 155, '2', 0),
(160, 155, '3', 1),
(161, 156, 'Ans1_Q1', 0),
(162, 156, 'Ans2_Q1', 1),
(163, 156, 'Ans3_Q1', 0),
(164, 156, 'Ans4_Q1', 0),
(165, 157, 'Ans1_Q2', 1),
(166, 157, 'Ans2_Q2', 0),
(167, 157, 'Ans3_Q2', 1),
(168, 157, 'Ans4_Q2', 0);

-- --------------------------------------------------------

--
-- Table structure for table `marks`
--

CREATE TABLE `marks` (
  `id_m` int(10) UNSIGNED NOT NULL,
  `user` int(10) UNSIGNED NOT NULL,
  `test` int(10) UNSIGNED NOT NULL,
  `mark` int(10) UNSIGNED NOT NULL,
  `count_of_correct_answers` int(10) UNSIGNED NOT NULL,
  `time_taken_s` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf32 COLLATE=utf32_general_ci;

--
-- Dumping data for table `marks`
--

INSERT INTO `marks` (`id_m`, `user`, `test`, `mark`, `count_of_correct_answers`, `time_taken_s`) VALUES
(8, 5, 20, 10, 2, 5),
(9, 5, 20, 10, 2, 5),
(10, 5, 21, 50, 2, 12),
(11, 1, 20, 20, 4, 44),
(12, 10, 20, 5, 1, 7),
(13, 10, 26, 10, 1, 10),
(14, 10, 26, 20, 2, 7),
(15, 6, 27, 15, 2, 4),
(16, 6, 27, 0, 0, 3),
(17, 5, 27, 15, 2, 14706),
(18, 5, 20, 10, 2, 22),
(19, 5, 20, 2, 0, 2),
(20, 5, 27, 10, 1, 2);

-- --------------------------------------------------------

--
-- Table structure for table `questions`
--

CREATE TABLE `questions` (
  `id_q` int(10) UNSIGNED NOT NULL,
  `text` varchar(255) NOT NULL,
  `id_creator` int(10) UNSIGNED NOT NULL,
  `type` enum('single','multiple') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `questions`
--

INSERT INTO `questions` (`id_q`, `text`, `id_creator`, `type`) VALUES
(144, 'First', 5, 'single'),
(145, 'Second', 5, 'multiple'),
(146, 'Thirdth', 5, 'single'),
(147, 'Fourth', 5, 'single'),
(148, '1', 5, 'multiple'),
(149, '2', 5, 'single'),
(154, 'ques1', 10, 'single'),
(155, 'ques2', 10, 'multiple'),
(156, 'Question_1', 10, 'single'),
(157, 'Question_2', 10, 'multiple');

-- --------------------------------------------------------

--
-- Table structure for table `tests`
--

CREATE TABLE `tests` (
  `id_t` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `title` varchar(255) NOT NULL,
  `count_of_questions` int(10) UNSIGNED NOT NULL,
  `max_mark` int(10) UNSIGNED NOT NULL,
  `tags` varchar(255) NOT NULL,
  `creator` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf32 COLLATE=utf32_general_ci;

--
-- Dumping data for table `tests`
--

INSERT INTO `tests` (`id_t`, `created_at`, `updated_at`, `title`, `count_of_questions`, `max_mark`, `tags`, `creator`) VALUES
(20, '2023-11-15 12:28:20', '2023-11-15 12:29:26', 'First Test', 4, 20, '#tag1, #tag2, #tag3', 5),
(21, '2023-11-15 12:34:08', '2023-11-15 12:59:27', 'Test 2', 2, 50, '1', 5),
(26, '2023-11-16 16:21:40', '2023-11-16 16:21:40', 'Sample Test', 2, 20, '#sampletest', 10),
(27, '2023-11-22 19:45:16', '2023-11-22 19:45:16', 'My test', 2, 20, '#mytest', 10);

--
-- Triggers `tests`
--
DELIMITER $$
CREATE TRIGGER `updated_at` BEFORE UPDATE ON `tests` FOR EACH ROW SET NEW.updated_at = CURRENT_TIMESTAMP()
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `tests_structure`
--

CREATE TABLE `tests_structure` (
  `id_t` int(10) UNSIGNED NOT NULL,
  `id_q` int(10) UNSIGNED NOT NULL,
  `id_a` int(10) UNSIGNED NOT NULL,
  `id_creator` int(10) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf16 COLLATE=utf16_general_ci;

--
-- Dumping data for table `tests_structure`
--

INSERT INTO `tests_structure` (`id_t`, `id_q`, `id_a`, `id_creator`) VALUES
(20, 144, 133, 5),
(20, 144, 134, 5),
(20, 144, 135, 5),
(20, 145, 136, 5),
(20, 145, 137, 5),
(20, 145, 138, 5),
(20, 146, 139, 5),
(20, 146, 140, 5),
(20, 146, 141, 5),
(20, 146, 142, 5),
(20, 147, 144, 5),
(20, 147, 145, 5),
(21, 148, 146, 5),
(21, 148, 147, 5),
(21, 149, 148, 5),
(21, 149, 149, 5),
(26, 154, 156, 10),
(26, 154, 157, 10),
(26, 155, 158, 10),
(26, 155, 159, 10),
(26, 155, 160, 10),
(27, 156, 161, 10),
(27, 156, 162, 10),
(27, 156, 163, 10),
(27, 156, 164, 10),
(27, 157, 165, 10),
(27, 157, 166, 10),
(27, 157, 167, 10),
(27, 157, 168, 10);

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
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `answers`
--
ALTER TABLE `answers`
  MODIFY `id_a` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=169;

--
-- AUTO_INCREMENT for table `marks`
--
ALTER TABLE `marks`
  MODIFY `id_m` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=21;

--
-- AUTO_INCREMENT for table `questions`
--
ALTER TABLE `questions`
  MODIFY `id_q` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=158;

--
-- AUTO_INCREMENT for table `tests`
--
ALTER TABLE `tests`
  MODIFY `id_t` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

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
