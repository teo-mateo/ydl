CREATE SEQUENCE public.queue_id_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 167
  CACHE 1;
ALTER TABLE public.queue_id_seq
  OWNER TO postgres;


-------------------

CREATE TABLE public.yqueue
(
  id integer NOT NULL DEFAULT nextval('queue_id_seq'::regclass),
  yturl text NOT NULL,
  status integer NOT NULL,
  file text,
  lastupdate timestamp with time zone,
  who text,
  CONSTRAINT queue_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.yqueue
  OWNER TO postgres;


------------------------

CREATE TABLE public.ystatus
(
  id integer NOT NULL,
  name text,
  CONSTRAINT "YStatus_pkey" PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.ystatus
  OWNER TO postgres;


----------

alter table public.yqueue add column ytkey text