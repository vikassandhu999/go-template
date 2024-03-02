BEGIN;

CREATE TABLE contents(
    id BIGSERIAL PRIMARY KEY,
    page_id INT NOT NULL,
    html TEXT,
    css TEXT,
    json_structure jsonb,
    publised_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (page_id) REFERENCES contents(id)
);

CREATE TABLE pages(
    id BIGSERIAL PRIMARY KEY,
    kind INT NOT NULL,
    title TEXT,
    slug TEXT,
    working_content_id INT NOT NULL,
    published_content_id INT,
    FOREIGN KEY (working_content_id) REFERENCES contents(id),
    FOREIGN KEY (published_content_id) REFERENCES contents(id)
);

CREATE TABLE coaching_pages(
    coaching_id INT NOT NULL,
    page_id INT NOT NULL,
    PRIMARY KEY(coaching_id,page_id),
    FOREIGN KEY (page_id) REFERENCES contents(id)
);

CREATE TABLE meeting_pages(
    meeting_id INT NOT NULL,
    page_id INT NOT NULL,
    PRIMARY KEY(meeting_id,page_id),
    FOREIGN KEY (page_id) REFERENCES contents(id)
);

END;
