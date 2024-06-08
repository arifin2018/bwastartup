CREATE TABLE IF NOT EXISTS transactions (
  id int NOT NULL AUTO_INCREMENT,
  campaigns_int int UNSIGNED NOT NULL,
  user_id int UNSIGNED NOT NULL,
  amount int UNSIGNED NOT NULL,
  status enum('package','pending','updated','done') NOT NULL,
  code int UNSIGNED NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME,
  PRIMARY KEY (id)
);