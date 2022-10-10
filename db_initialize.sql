CREATE TABLE users(
	user_id INT NOT NULL,
	project_id INT NOT NULL,
	username VARCHAR(255) NOT NULL,
	password_hash bytea NOT NULL,
	salt bytea NOT NULL,
	email VARCHAR(255) NOT NULL,
	date_created DATE NOT NULL,
	verfication_code VARCHAR(6) NOT NULL,
	validation_date DATE NOT NULL,
	PRIMARY KEY (user_id)

);

CREATE TABLE projects(
	project_id INT NOT NULL,
	owner_id INT NOT NULL,
	date_created DATE NOT NULL,
	
	PRIMARY KEY (project_id)
);

CREATE TABLE sprints(
	sprint_id int NOT NULL,
	project_id INT NOT NULL,
	sprint_name VARCHAR(255) NOT NULL,
	created_date DATE NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	PRIMARY KEY (sprint_id),
	FOREIGN KEY (project_id) REFERENCES projects (project_id)
);

CREATE TABLE swimlanes(
	swimlane_id INT NOT NULL,
	project_id INT NOT NULL,
	swimlane_name VARCHAR(255) NOT NULL,
	swimlane_position SMALLINT NOT NULL,
	PRIMARY KEY (swimlane_id),
	FOREIGN KEY (project_id) REFERENCES projects (project_id)
);

CREATE TABLE issues(
	issues INT NOT NULL,
	project_id INT NOT NULL,
	reporter_id INT NOT NULL,
	asignee_id INT NOT NULL,
	sprint_id INT NOT NULL,
	date_created DATE NOT NULL,
	issue_type SMALLINT NOT NULL,
	priority_type SMALLINT NOT NULL,
	
	PRIMARY KEY (issues),
	FOREIGN KEY (project_id)  REFERENCES projects (project_id),
	FOREIGN KEY (reporter_id) REFERENCES users    (user_id),
	FOREIGN KEY (asignee_id)  REFERENCES users    (user_id),
	FOREIGN KEY (sprint_id)   REFERENCES sprints  (sprint_id)
);