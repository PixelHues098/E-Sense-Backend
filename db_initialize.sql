CREATE TABLE users(
    user_id  SERIAL PRIMARY KEY,
	first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
	project_id INT NULL,
	username VARCHAR(255) NOT NULL,
	password_hash bytea NOT NULL,
	salt bytea NULL,
	email VARCHAR(255) NOT NULL,
	date_created DATE NOT NULL default current_timestamp

);

CREATE TABLE projects(
	project_id SERIAL PRIMARY KEY,
	owner_id INT NOT NULL,
	date_created DATE NOT NULL
	
);

CREATE TABLE sprints(
	sprint_id SERIAL PRIMARY KEY,
	project_id INT NOT NULL,
	sprint_name VARCHAR(255) NOT NULL,
	date_created DATE NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	FOREIGN KEY (project_id) REFERENCES projects (project_id)
);

CREATE TABLE swimlanes(
	swimlane_id SERIAL PRIMARY KEY,
	project_id INT NOT NULL,
	swimlane_name VARCHAR(255) NOT NULL,
	swimlane_position SMALLINT NOT NULL,
	FOREIGN KEY (project_id) REFERENCES projects (project_id)
);

CREATE TABLE issues(
	issues SERIAL PRIMARY KEY,
	project_id INT NOT NULL,
	reporter_id INT NOT NULL,
	assignee_id INT NOT NULL,
	sprint_id INT NOT NULL,
	date_created DATE NOT NULL,
	issue_type SMALLINT NOT NULL,
	priority_type SMALLINT NOT NULL,
	
	FOREIGN KEY (project_id)  REFERENCES projects (project_id),
	FOREIGN KEY (reporter_id) REFERENCES users    (user_id),
	FOREIGN KEY (assignee_id)  REFERENCES users    (user_id),
	FOREIGN KEY (sprint_id)   REFERENCES sprints  (sprint_id)
);