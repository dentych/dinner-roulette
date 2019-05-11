CREATE TABLE public.ingredient (
    name varchar(50)  NOT NULL,
    id serial  NOT NULL,
    kcal int  NOT NULL,
    CONSTRAINT ingredient_pk PRIMARY KEY (id)
);


-- Table: recipe_ingredient
CREATE TABLE public.recipe_ingredient (
    id serial  NOT NULL,
    ingredient_id int  NOT NULL,
    recipe_id int  NOT NULL,
    amount int NOT NULL,
    CONSTRAINT recipe_ingredient_pk PRIMARY KEY (id)
);
ALTER TABLE public.recipe_ingredient ADD CONSTRAINT recipe_ingredient_ingredient
    FOREIGN KEY (ingredient_id)
    REFERENCES ingredient (id)
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;


ALTER TABLE public.recipe_ingredient ADD CONSTRAINT recipe_ingredient_recipe
    FOREIGN KEY (recipe_id)
    REFERENCES recipe (id)
    NOT DEFERRABLE
    INITIALLY IMMEDIATE
;
