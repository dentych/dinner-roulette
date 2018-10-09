CREATE TABLE public.user (
  id       SERIAL      NOT NULL PRIMARY KEY,
  username VARCHAR(25) NOT NULL,
  email    TEXT        NOT NULL,
  hash     TEXT        NOT NULL,
  salt     TEXT        NOT NULL
);

CREATE TABLE public.recipe (
  id          SERIAL  NOT NULL PRIMARY KEY,
  name        TEXT    NOT NULL,
  description TEXT,
  url         TEXT,
  userid      INTEGER NOT NULL
    CONSTRAINT recipe_user_id_fk
    REFERENCES public.user
    ON DELETE CASCADE
);

CREATE TABLE public.ingredient (
  id   SERIAL NOT NULL PRIMARY KEY,
  name TEXT   NOT NULL,
  unit TEXT   NOT NULL
);

CREATE TABLE public.recipe_ingredient (
  recipeId     INTEGER
    CONSTRAINT fk_recipe_id
    REFERENCES public.recipe
    ON DELETE CASCADE,
  ingredientId INTEGER
    CONSTRAINT fk_ingredient_id
    REFERENCES public.ingredient
    ON DELETE CASCADE,
  amount       INTEGER
);