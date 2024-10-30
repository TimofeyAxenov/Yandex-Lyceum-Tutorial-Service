ALTER TABLE public.position
  ADD updatedAt date DEFAULT now() not NULL;
