/* http://yensdesign.com/tutorials/forumsdatabase/schemaFull.jpg */

DROP TABLE IF EXISTS forum_pollsOptions;
DROP TABLE IF EXISTS forum_post;
DROP TABLE IF EXISTS forum_user;
DROP TABLE IF EXISTS forum_group;
DROP TABLE IF EXISTS forum_subCategory;
DROP TABLE IF EXISTS forum_category;

CREATE TABLE forum_category (
	category_id INT PRIMARY KEY AUTO_INCREMENT,
	category_title VARCHAR(64) NOT NULL UNIQUE,
	category_description TEXT,
	date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE forum_subCategory (
	subCategory_id INT PRIMARY KEY AUTO_INCREMENT,
	subCategory_idCategory INT NOT NULL,
	subCategory_title VARCHAR(45) NOT NULL UNIQUE,
	subCategory_description TEXT,
	subCategory_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT FOREIGN KEY(subCategory_idCategory) REFERENCES forum_category(category_id)	 
);

CREATE TABLE forum_group (
	group_id INT PRIMARY KEY AUTO_INCREMENT,
	group_name VARCHAR(45) NOT NULL UNIQUE
);

CREATE TABLE forum_user (
	user_id INT PRIMARY KEY AUTO_INCREMENT,
	user_pseudo VARCHAR(16) NOT NULL UNIQUE,
	user_password VARCHAR(256) NOT NULL,
	user_email VARCHAR(128) NOT NULL UNIQUE,
	user_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	user_date_lastConnection TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	user_groupId INT NOT NULL DEFAULT 1,
	user_karma FLOAT(2,1),
	user_ip VARCHAR(20) NOT NULL,
	CONSTRAINT FOREIGN KEY(user_groupId) REFERENCES forum_group(group_id)
);

CREATE TABLE forum_post (
	post_id INT PRIMARY KEY AUTO_INCREMENT,
	post_idSubcategory INT NOT NULL,
	post_idUser INT NOT NULL,
	post_idParentPost INT NOT NULL DEFAULT 0,
	post_title VARCHAR(90) NOT NULL UNIQUE,
	post_content TEXT NOT NULL,
	post_isPoll BOOLEAN NOT NULL DEFAULT 0,
	post_pollTitle VARCHAR(90) NOT NULL,
	post_isEdited BOOLEAN NOT NULL DEFAULT 0,
	post_createDate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	post_editDate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	post_ip VARCHAR(20) NOT NULL,
	CONSTRAINT FOREIGN KEY(post_idSubcategory) REFERENCES forum_subCategory(subCategory_id),
	CONSTRAINT FOREIGN KEY(post_idUser) REFERENCES forum_user(user_id),
	CONSTRAINT FOREIGN KEY(post_idParentPost) REFERENCES forum_post(post_id)
);

CREATE TABLE forum_pollsOptions (
	pollsOptions_id INT PRIMARY KEY AUTO_INCREMENT,
	pollsOptions_idPost INT NOT NULL,
	pollsOptions_title VARCHAR(64) NOT NULL,
	pollsOptions_value INT NOT NULL,
	date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT FOREIGN KEY(pollsOptions_idPost) REFERENCES forum_post(post_id)
);

INSERT INTO forum_group (`group_name`) VALUES ('Utilisateur'), ('Modérateur'), ('Administrateur');
