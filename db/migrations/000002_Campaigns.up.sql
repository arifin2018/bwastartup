CREATE TABLE IF NOT EXISTS campaigns (
  id int NOT NULL AUTO_INCREMENT,
  user_id int UNSIGNED NOT NULL,
  name VARCHAR(255) NOT NULL,
  short_description VARCHAR(255) NOT NULL,
  goal_amount int UNSIGNED NOT NULL,
  current_amount int UNSIGNED NOT NULL,
  perks text,
  backer_count int UNSIGNED,
  slug VARCHAR(255),
  created_at DATETIME NOT NULL,
  updated_at DATETIME,
  PRIMARY KEY (id),
  UNIQUE (name)
);