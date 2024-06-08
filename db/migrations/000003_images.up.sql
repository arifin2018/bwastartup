CREATE TABLE IF NOT EXISTS images (
  id int NOT NULL AUTO_INCREMENT,
  campaign_id int UNSIGNED NOT NULL,
  file_name VARCHAR(255) NOT NULL,
  is_primary boolean NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME,
  PRIMARY KEY (id),
  UNIQUE (file_name)
);