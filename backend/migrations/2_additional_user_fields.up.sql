ALTER TABLE public.user ADD COLUMN firstName VARCHAR(20);
ALTER TABLE public.user ADD COLUMN lastName VARCHAR(30);
ALTER TABLE public.user ADD COLUMN createdTimestamp TIMESTAMP DEFAULT NOW();

ALTER TABLE public.recipe ADD COLUMN description TEXT;