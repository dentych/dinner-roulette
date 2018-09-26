CREATE TABLE public.user (
  id SERIAL NOT NULL CONSTRAINT user_pkey PRIMARY KEY,
  username VARCHAR(25) NOT NULL,
  email TEXT NOT NULL,
  hash TEXT NOT NULL,
  salt TEXT NOT NULL
);

CREATE TABLE public.recipe (
  id SERIAL NOT NULL CONSTRAINT recipe_pkey PRIMARY KEY,
  name TEXT NOT NULL,
  url TEXT,
  userid INTEGER NOT NULL
    CONSTRAINT recipe_user_id_fk
      REFERENCES public.user
      ON DELETE CASCADE
);

