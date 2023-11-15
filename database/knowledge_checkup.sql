-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:3306
-- Время создания: Ноя 15 2023 г., 15:39
-- Версия сервера: 8.0.24
-- Версия PHP: 7.4.27

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- База данных: `knowledge_checkup`
--

-- --------------------------------------------------------

--
-- Структура таблицы `accounts`
--

CREATE TABLE `accounts` (
  `id` int UNSIGNED NOT NULL,
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `accounts`
--

INSERT INTO `accounts` (`id`, `last_name`, `first_name`, `middle_name`, `year_of_birth`, `nickname`, `email`, `password`, `approved`, `gender`, `educational_institution`, `teacher_status`) VALUES
(1, 'test', 'test', 'test', 'test', 'test', 'test', 'test', 1, 'Ч', 'test', 0),
(2, '123', '123', '123', '123', '123', '123', '123', 1, 'Ч', ' ', 0),
(3, '1', '1', '1', '1', '1', '1', '1', 1, NULL, '', 0),
(4, '2000', '2000', '2000', '2000', '2000', '2000', '2000', 1, NULL, NULL, 0),
(5, 'Andrunkiv', 'Sergiy', 'Romanovich', '2003', 'sergiy_2000', 'sergij.andrunkiv@gmail.com', 'qwerty123', 1, 'N/A', 'N/A', 1),
(6, 'test123', 'test123', 'test123', '2000', 'testik', 'test123', 'test123', 1, 'N/A', 'N/A', 0),
(8, 'Lastname', 'Firstname', 'Middlename', '1888', 'Nickname', 'email@gmail.com', '12345', 1, 'N/A', 'N/A', 0),
(9, 'Surname', 'Name', 'Middlenam', '2222', 'nickname', 'test@gmail.com', 'password', 1, 'N/A', 'N/A', 0);

-- --------------------------------------------------------

--
-- Структура таблицы `answers`
--

CREATE TABLE `answers` (
  `id_a` int UNSIGNED NOT NULL,
  `id_q` int UNSIGNED NOT NULL,
  `text` varchar(255) NOT NULL,
  `is_correct` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `answers`
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
(149, 149, '2 2', 1);

-- --------------------------------------------------------

--
-- Структура таблицы `marks`
--

CREATE TABLE `marks` (
  `id_m` int UNSIGNED NOT NULL,
  `user` int UNSIGNED NOT NULL,
  `test` int UNSIGNED NOT NULL,
  `mark` int UNSIGNED NOT NULL,
  `count_of_correct_answers` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf32;

-- --------------------------------------------------------

--
-- Структура таблицы `questions`
--

CREATE TABLE `questions` (
  `id_q` int UNSIGNED NOT NULL,
  `text` varchar(255) NOT NULL,
  `id_creator` int UNSIGNED NOT NULL,
  `type` enum('single','multiple') NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--
-- Дамп данных таблицы `questions`
--

INSERT INTO `questions` (`id_q`, `text`, `id_creator`, `type`) VALUES
(144, 'First', 5, 'single'),
(145, 'Second', 5, 'multiple'),
(146, 'Thirdth', 5, 'single'),
(147, 'Fourth', 5, 'single'),
(148, '1', 5, 'single'),
(149, '2', 5, 'single');

-- --------------------------------------------------------

--
-- Структура таблицы `tests`
--

CREATE TABLE `tests` (
  `id_t` int UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `title` varchar(255) NOT NULL,
  `count_of_questions` int UNSIGNED NOT NULL,
  `max_mark` int UNSIGNED NOT NULL,
  `tags` varchar(255) NOT NULL,
  `creator` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf32;

--
-- Дамп данных таблицы `tests`
--

INSERT INTO `tests` (`id_t`, `created_at`, `updated_at`, `title`, `count_of_questions`, `max_mark`, `tags`, `creator`) VALUES
(20, '2023-11-15 12:28:20', '2023-11-15 12:29:26', 'First Test', 4, 20, '#tag1, #tag2, #tag3', 5),
(21, '2023-11-15 12:34:08', '2023-11-15 12:34:08', 'Test 2', 2, 50, '1', 5);

--
-- Триггеры `tests`
--
DELIMITER $$
CREATE TRIGGER `updated_at` BEFORE UPDATE ON `tests` FOR EACH ROW SET NEW.updated_at = CURRENT_TIMESTAMP()
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Структура таблицы `tests_structure`
--

CREATE TABLE `tests_structure` (
  `id_t` int UNSIGNED NOT NULL,
  `id_q` int UNSIGNED NOT NULL,
  `id_a` int UNSIGNED NOT NULL,
  `id_creator` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf16;

--
-- Дамп данных таблицы `tests_structure`
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
(21, 149, 149, 5);

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `accounts`
--
ALTER TABLE `accounts`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `answers`
--
ALTER TABLE `answers`
  ADD PRIMARY KEY (`id_a`),
  ADD KEY `Test_2` (`id_q`);

--
-- Индексы таблицы `marks`
--
ALTER TABLE `marks`
  ADD PRIMARY KEY (`id_m`),
  ADD KEY `Test_8` (`user`),
  ADD KEY `Test_9` (`test`);

--
-- Индексы таблицы `questions`
--
ALTER TABLE `questions`
  ADD PRIMARY KEY (`id_q`),
  ADD KEY `Test` (`id_creator`);

--
-- Индексы таблицы `tests`
--
ALTER TABLE `tests`
  ADD PRIMARY KEY (`id_t`),
  ADD KEY `Test_7` (`creator`);

--
-- Индексы таблицы `tests_structure`
--
ALTER TABLE `tests_structure`
  ADD PRIMARY KEY (`id_t`,`id_q`,`id_a`,`id_creator`),
  ADD KEY `Test_4` (`id_q`),
  ADD KEY `Test_5` (`id_a`),
  ADD KEY `Test_6` (`id_creator`);

--
-- AUTO_INCREMENT для сохранённых таблиц
--

--
-- AUTO_INCREMENT для таблицы `accounts`
--
ALTER TABLE `accounts`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT для таблицы `answers`
--
ALTER TABLE `answers`
  MODIFY `id_a` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=156;

--
-- AUTO_INCREMENT для таблицы `marks`
--
ALTER TABLE `marks`
  MODIFY `id_m` int UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT для таблицы `questions`
--
ALTER TABLE `questions`
  MODIFY `id_q` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=154;

--
-- AUTO_INCREMENT для таблицы `tests`
--
ALTER TABLE `tests`
  MODIFY `id_t` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `answers`
--
ALTER TABLE `answers`
  ADD CONSTRAINT `Test_2` FOREIGN KEY (`id_q`) REFERENCES `questions` (`id_q`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ограничения внешнего ключа таблицы `marks`
--
ALTER TABLE `marks`
  ADD CONSTRAINT `Test_8` FOREIGN KEY (`user`) REFERENCES `accounts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Test_9` FOREIGN KEY (`test`) REFERENCES `tests` (`id_t`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ограничения внешнего ключа таблицы `questions`
--
ALTER TABLE `questions`
  ADD CONSTRAINT `Test` FOREIGN KEY (`id_creator`) REFERENCES `accounts` (`id`);

--
-- Ограничения внешнего ключа таблицы `tests`
--
ALTER TABLE `tests`
  ADD CONSTRAINT `Test_7` FOREIGN KEY (`creator`) REFERENCES `accounts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Ограничения внешнего ключа таблицы `tests_structure`
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
