BEGIN;

CREATE TABLE modules
(
    id        BIGSERIAL PRIMARY KEY,
    domain    VARCHAR(255) NOT NULL,
    category  VARCHAR(255) NOT NULL,
    purpose   VARCHAR(255) NOT NULL,
    structure JSON DEFAULT NULL
);

CREATE TABLE sections
(
    id        BIGSERIAL PRIMARY KEY,
    module_id INT          NOT NULL,
    name      VARCHAR(255) NOT NULL,
    position     INT DEFAULT NULL,
    FOREIGN KEY (module_id) REFERENCES modules (id)
);

CREATE TABLE items
(
    id          BIGSERIAL PRIMARY KEY,
    section_id  INT          NOT NULL,
    name        VARCHAR(255) NOT NULL,
    type        VARCHAR(255) DEFAULT NULL,
    duration    INT          DEFAULT NULL,
    difficulty  VARCHAR(255) DEFAULT NULL,
    description TEXT         DEFAULT NULL,
    FOREIGN KEY (section_id) REFERENCES sections (id)
);

CREATE TABLE variations
(
    id          BIGSERIAL PRIMARY KEY,
    module_id   INT          NOT NULL,
    name        VARCHAR(255) NOT NULL,
    description TEXT DEFAULT NULL,
    FOREIGN KEY (module_id) REFERENCES modules (id)
);

CREATE TABLE variations_sections
(
    id           BIGSERIAL PRIMARY KEY,
    variation_id INT NOT NULL,
    section_id   INT NOT NULL,
    position        INT DEFAULT NULL,
    FOREIGN KEY (variation_id) REFERENCES variations (id),
    FOREIGN KEY (section_id) REFERENCES sections (id)
);

CREATE TABLE variations_items
(
    id           BIGSERIAL PRIMARY KEY,
    variation_id INT NOT NULL,
    item_id      INT NOT NULL,
    position        INT DEFAULT NULL,
    FOREIGN KEY (variation_id) REFERENCES variations (id),
    FOREIGN KEY (item_id) REFERENCES items (id)
);

END;
