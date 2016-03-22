CREATE TABLE notes (
    id integer PRIMARY KEY NOT NULL,
    title text NOT NULL,
    content text,
    priority integer
    created_at timestamp without time zone DEFAULT now() NOT NULL,
);